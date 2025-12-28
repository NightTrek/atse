use serde::Serialize;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Format {
    Text,
    Json,
    Locations,
}

impl Format {
    pub fn parse(s: &str) -> Self {
        match s.to_lowercase().as_str() {
            "json" => Format::Json,
            "locations" => Format::Locations,
            _ => Format::Text,
        }
    }
}

#[derive(Debug, Clone, Serialize)]
pub struct ResultItem {
    pub file: String,
    pub line: u32,
    pub column: u32,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub text: Option<String>,
}

pub fn format_results(results: &[ResultItem], format: Format, verbose: bool) -> String {
    match format {
        Format::Json => serde_json::to_string_pretty(results)
            .unwrap_or_else(|e| format!("{{\"error\":\"{}\"}}", e)),
        Format::Locations => results
            .iter()
            .map(|r| format!("{}:{}:{}", r.file, r.line + 1, r.column))
            .collect::<Vec<_>>()
            .join("\n"),
        Format::Text => {
            if results.is_empty() {
                return "No results found.".into();
            }
            let mut out = String::new();
            for r in results {
                out.push_str(&format!("{}:{}:{}", r.file, r.line + 1, r.column));
                if let Some(name) = &r.name {
                    out.push_str(&format!(" - {}", name));
                }
                out.push('\n');
                if verbose {
                    if let Some(text) = &r.text {
                        for line in text.lines() {
                            if !line.trim().is_empty() {
                                out.push_str("  ");
                                out.push_str(line);
                                out.push('\n');
                            }
                        }
                        out.push('\n');
                    }
                }
            }
            out
        }
    }
}

pub fn format_context(
    context: &str,
    file_path: &str,
    line: u32,
    col: u32,
    format: Format,
) -> String {
    match format {
        Format::Json => serde_json::to_string_pretty(&ResultItem {
            file: file_path.to_string(),
            line,
            column: col,
            name: None,
            text: Some(context.to_string()),
        })
        .unwrap_or_else(|e| format!("{{\"error\":\"{}\"}}", e)),
        _ => format!("{}:{}:{}\n{}\n", file_path, line + 1, col, context),
    }
}
