use crate::index::{PartialIndex, SymbolNode};
use crate::output::{format_context, format_results, Format, ResultItem};
use crate::parser::{LanguageId, Manager};
use crate::ripgrep::{Client as RgClient, SearchOptions};
use crate::util::walk_files;
use clap::{Arg, ArgAction, Command};
use std::collections::HashMap;
use std::path::{Path, PathBuf};
use std::time::Instant;

pub fn run() -> anyhow::Result<()> {
    let matches = Command::new("atse")
        .version(env!("CARGO_PKG_VERSION"))
        .about("Tree-sitter powered code analysis - fast, accurate, zero false positives")
        .subcommand_required(true)
        .arg_required_else_help(true)
        .arg(
            Arg::new("format")
                .long("format")
                .default_value("text")
                .help("Output format: text, json, locations")
                .global(true),
        )
        .arg(
            Arg::new("verbose")
                .long("verbose")
                .short('v')
                .action(ArgAction::SetTrue)
                .help("Verbose output")
                .global(true),
        )
        .arg(
            Arg::new("include")
                .long("include")
                .num_args(1)
                .action(ArgAction::Append)
                .help("Include glob pattern")
                .global(true),
        )
        .arg(
            Arg::new("exclude")
                .long("exclude")
                .num_args(1)
                .action(ArgAction::Append)
                .help("Exclude glob pattern")
                .global(true),
        )
        .arg(
            Arg::new("limit")
                .long("limit")
                .default_value("0")
                .help("Limit results (0 = unlimited)")
                .global(true),
        )
        .subcommand(
            Command::new("search")
                .about("Search for code symbols")
                .arg(Arg::new("query").required(true))
                .arg(Arg::new("path").default_value("."))
                .arg(
                    Arg::new("type")
                        .long("type")
                        .value_parser(["function", "class", "method", "type", "const"]),
                )
                .arg(Arg::new("fuzzy").long("fuzzy").action(ArgAction::SetTrue)),
        )
        .subcommand(
            Command::new("find-fn")
                .about("Find function calls")
                .arg(Arg::new("name").required(true))
                .arg(Arg::new("path").default_value(".")),
        )
        .subcommand(
            Command::new("context")
                .about("Get context for position")
                .arg(Arg::new("location").required(true)),
        )
        .get_matches();

    let format = Format::parse(matches.get_one::<String>("format").unwrap());
    let verbose = matches.get_flag("verbose");
    let include: Vec<String> = matches
        .get_many::<String>("include")
        .map(|v| v.map(|s| s.to_string()).collect())
        .unwrap_or_default();
    let exclude: Vec<String> = matches
        .get_many::<String>("exclude")
        .map(|v| v.map(|s| s.to_string()).collect())
        .unwrap_or_default();
    let mut limit: usize = matches
        .get_one::<String>("limit")
        .unwrap()
        .parse()
        .unwrap_or(0);

    match matches.subcommand() {
        Some(("search", sub)) => {
            let query = sub.get_one::<String>("query").unwrap();
            let path = sub.get_one::<String>("path").unwrap();
            let typ = sub.get_one::<String>("type").cloned();
            let fuzzy = sub.get_flag("fuzzy");
            run_search(
                query, path, typ, fuzzy, &include, &exclude, format, verbose, &mut limit,
            )
        }
        Some(("find-fn", sub)) => {
            let name = sub.get_one::<String>("name").unwrap();
            let path = sub.get_one::<String>("path").unwrap();
            run_find_fn(name, path, &include, &exclude, format, verbose, &mut limit)
        }
        Some(("context", sub)) => {
            let loc = sub.get_one::<String>("location").unwrap();
            run_context(loc, format)
        }
        _ => unreachable!(),
    }
}

fn run_search(
    query: &str,
    path: &str,
    typ: Option<String>,
    fuzzy: bool,
    include: &[String],
    exclude: &[String],
    format: Format,
    verbose: bool,
    limit: &mut usize,
) -> anyhow::Result<()> {
    let start = Instant::now();
    let rg = RgClient::default();
    if !rg.available() {
        anyhow::bail!("ripgrep (rg) not found in PATH. Please install ripgrep.");
    }
    let candidates = rg.search(
        query,
        path,
        &SearchOptions {
            includes: include.to_vec(),
            excludes: exclude.to_vec(),
        },
    )?;
    let mut unique_files = HashMap::new();
    for c in candidates {
        unique_files.insert(c.path.clone(), c);
    }

    let mgr = Manager::new();
    let index = PartialIndex::default();

    for (file_path, _) in &unique_files {
        let p = PathBuf::from(file_path);
        let tree = match mgr.parse_file(&p) {
            Ok(t) => t,
            Err(_) => continue,
        };
        let content = match mgr.get_content(&p) {
            Ok(c) => c,
            Err(_) => continue,
        };
        let lang = match mgr.infer_language(&p) {
            Ok(l) => l,
            Err(_) => continue,
        };
        let symbols = extract_symbols(&mgr, &tree, lang, &content, file_path);
        index.add_file(file_path.clone(), content, symbols);
    }

    let mut matches = Vec::new();
    for syms in index.symbols.read().unwrap().values() {
        for sym in syms {
            if matches_query(&sym.symbol, query, fuzzy) {
                if let Some(ref t) = typ {
                    if sym.symbol_type != *t {
                        continue;
                    }
                }
                let score = calculate_score(&sym.symbol, query);
                matches.push((score, sym.clone()));
            }
        }
    }

    matches.sort_by(|a, b| b.0.cmp(&a.0).then(a.1.symbol.len().cmp(&b.1.symbol.len())));
    let mut out_items = Vec::new();
    if *limit > 0 && matches.len() > *limit {
        matches.truncate(*limit);
    }
    for (_, sym) in matches {
        out_items.push(ResultItem {
            file: sym.file_path.clone(),
            line: sym.line,
            column: 0,
            name: Some(sym.symbol.clone()),
            text: None,
        });
    }
    println!("{}", format_results(&out_items, format, verbose));
    if verbose {
        eprintln!(
            "Processed {} files, {} symbols in {:?}",
            unique_files.len(),
            index.symbols.read().unwrap().len(),
            start.elapsed()
        );
    }
    Ok(())
}

fn extract_symbols(
    mgr: &Manager,
    tree: &tree_sitter::Tree,
    lang: LanguageId,
    content: &[u8],
    file_path: &str,
) -> Vec<SymbolNode> {
    let queries = build_symbol_queries(lang);
    let mut symbols = Vec::new();
    for (symbol_type, q) in queries {
        if let Ok(results) = mgr.query(tree, &q, lang, content) {
            for res in results {
                let name = res.text.trim();
                if name.is_empty() {
                    continue;
                }
                symbols.push(SymbolNode {
                    symbol: name.to_string(),
                    symbol_type: symbol_type.to_string(),
                    file_path: file_path.to_string(),
                    line: res.start_position.row as u32,
                });
            }
        }
    }
    symbols
}

fn build_symbol_queries(lang: LanguageId) -> HashMap<String, String> {
    let mut map = HashMap::new();
    match lang {
        LanguageId::JavaScript | LanguageId::Tsx | LanguageId::TypeScript => {
            map.insert(
                "function".into(),
                "(function_declaration name: (identifier) @name)".into(),
            );
            map.insert(
                "class".into(),
                "(class_declaration name: (type_identifier) @name)".into(),
            );
            map.insert(
                "method".into(),
                "(method_definition name: (property_identifier) @name)".into(),
            );
            map.insert(
                "const".into(),
                "(lexical_declaration (variable_declarator name: (identifier) @name)) @const"
                    .into(),
            );
        }
        LanguageId::Go => {
            map.insert(
                "function".into(),
                "(function_declaration name: (identifier) @name)".into(),
            );
            map.insert(
                "method".into(),
                "(method_declaration name: (field_identifier) @name)".into(),
            );
            map.insert(
                "type".into(),
                "(type_spec name: (type_identifier) @name)".into(),
            );
        }
        LanguageId::Python => {
            map.insert(
                "function".into(),
                "(function_definition name: (identifier) @name)".into(),
            );
            map.insert(
                "class".into(),
                "(class_definition name: (identifier) @name)".into(),
            );
        }
        LanguageId::Java => {
            map.insert(
                "class".into(),
                "(class_declaration name: (identifier) @name)".into(),
            );
            map.insert(
                "interface".into(),
                "(interface_declaration name: (identifier) @name)".into(),
            );
            map.insert(
                "enum".into(),
                "(enum_declaration name: (identifier) @name)".into(),
            );
            map.insert(
                "method".into(),
                "(method_declaration name: (identifier) @name)".into(),
            );
            map.insert(
                "constructor".into(),
                "(constructor_declaration name: (identifier) @name)".into(),
            );
        }
    }
    map
}

fn matches_query(name: &str, query: &str, fuzzy: bool) -> bool {
    let name_lower = name.to_lowercase();
    let query_lower = query.to_lowercase();
    if name_lower == query_lower
        || name_lower.starts_with(&query_lower)
        || name_lower.contains(&query_lower)
    {
        return true;
    }
    if fuzzy {
        fuzzy_match(&name_lower, &query_lower)
    } else {
        false
    }
}

fn fuzzy_match(name: &str, query: &str) -> bool {
    let mut idx = 0;
    for ch in query.chars() {
        if let Some(pos) = name[idx..].find(ch) {
            idx += pos + 1;
        } else {
            return false;
        }
    }
    true
}

fn calculate_score(name: &str, query: &str) -> i32 {
    let n = name.to_lowercase();
    let q = query.to_lowercase();
    if n == q {
        1000
    } else if n.starts_with(&q) {
        500
    } else if n.contains(&q) {
        250
    } else {
        100
    }
}

fn run_find_fn(
    name: &str,
    path: &str,
    include: &[String],
    exclude: &[String],
    format: Format,
    verbose: bool,
    limit: &mut usize,
) -> anyhow::Result<()> {
    let mgr = Manager::new();
    let files = walk_files(Path::new(path), true, include, exclude, true)?;
    if files.is_empty() {
        println!("No files found matching criteria.");
        return Ok(());
    }
    let mut results = Vec::new();
    for f in &files {
        let tree = match mgr.parse_file(&f.path) {
            Ok(t) => t,
            Err(err) => {
                if verbose {
                    eprintln!("Warning: failed to parse {}: {}", f.path.display(), err);
                }
                continue;
            }
        };
        let content = mgr.get_content(&f.path)?;
        let lang = mgr.infer_language(&f.path)?;
        let query = build_function_call_query(lang, name);
        if query.is_empty() {
            continue;
        }
        let matches = mgr.query(&tree, &query, lang, &content)?;
        for m in matches {
            results.push(ResultItem {
                file: f.path.to_string_lossy().into(),
                line: m.start_position.row as u32,
                column: m.start_position.column as u32,
                name: Some(name.to_string()),
                text: Some(m.text.clone()),
            });
        }
    }
    if *limit == 0 {
        *limit = 50;
    }
    if *limit > 0 && results.len() > *limit {
        results.truncate(*limit);
        if verbose {
            eprintln!(
                "Warning: Found more than limit; showing first {} (use --limit 0 for all)",
                *limit
            );
        }
    }
    println!("{}", format_results(&results, format, verbose));
    Ok(())
}

fn build_function_call_query(lang: LanguageId, func_name: &str) -> String {
    match lang {
        LanguageId::JavaScript | LanguageId::TypeScript | LanguageId::Tsx => format!(
            "(call_expression\n  function: (identifier) @fn\n  (#eq? @fn \"{}\"))",
            func_name
        ),
        LanguageId::Go => format!(
            "(call_expression\n  function: (identifier) @fn\n  (#eq? @fn \"{}\"))",
            func_name
        ),
        LanguageId::Python => format!(
            "(call\n  function: (identifier) @fn\n  (#eq? @fn \"{}\"))",
            func_name
        ),
        LanguageId::Java => format!(
            "(method_invocation\n  name: (identifier) @fn\n  (#eq? @fn \"{}\"))",
            func_name
        ),
    }
}

fn run_context(location: &str, format: Format) -> anyhow::Result<()> {
    let parts: Vec<&str> = location.split(':').collect();
    if parts.len() != 3 {
        anyhow::bail!("invalid format: expected file:line:column");
    }
    let file = PathBuf::from(parts[0]);
    let line: u32 = parts[1].parse()?;
    let col: u32 = parts[2].parse()?;
    let mgr = Manager::new();
    let tree = mgr.parse_file(&file)?;
    let content = mgr.get_content(&file)?;
    let context = mgr.get_context_at_position(&tree, line - 1, col, &content);
    println!(
        "{}",
        format_context(&context, parts[0], line - 1, col, format)
    );
    Ok(())
}
