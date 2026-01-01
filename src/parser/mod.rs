use anyhow::Context;
use std::collections::HashMap;
use std::path::Path;
use std::sync::RwLock;
use tree_sitter as ts;
use tree_sitter_go as ts_go;
use tree_sitter_java as ts_java;
use tree_sitter_javascript as ts_js;
use tree_sitter_python as ts_py;
use tree_sitter_typescript::{language_tsx, language_typescript};

#[derive(Debug, Clone)]
pub struct QueryMatch {
    pub name: String,
    pub text: String,
    pub start_position: ts::Point,
    pub end_position: ts::Point,
    pub start_byte: usize,
    pub end_byte: usize,
}

#[derive(Debug, Clone)]
pub struct NodeInfo {
    pub name: String,
    pub node_type: String,
    pub text: String,
    pub start_position: ts::Point,
    pub end_position: ts::Point,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub enum LanguageId {
    Go,
    JavaScript,
    TypeScript,
    Tsx,
    Python,
    Java,
}

fn language_for(ext: &str) -> Option<LanguageId> {
    match ext.to_lowercase().as_str() {
        "go" => Some(LanguageId::Go),
        "js" | "jsx" | "mjs" | "cjs" => Some(LanguageId::JavaScript),
        "ts" => Some(LanguageId::TypeScript),
        "tsx" => Some(LanguageId::Tsx),
        "py" => Some(LanguageId::Python),
        "java" => Some(LanguageId::Java),
        _ => None,
    }
}

pub struct Manager {
    languages: HashMap<LanguageId, ts::Language>,
    trees: RwLock<HashMap<String, ts::Tree>>,
    contents: RwLock<HashMap<String, Vec<u8>>>,
}

impl Manager {
    pub fn new() -> Self {
        let mut languages = HashMap::new();
        languages.insert(LanguageId::Go, ts_go::language());
        languages.insert(LanguageId::JavaScript, ts_js::language());
        languages.insert(LanguageId::TypeScript, language_typescript());
        languages.insert(LanguageId::Tsx, language_tsx());
        languages.insert(LanguageId::Python, ts_py::language());
        languages.insert(LanguageId::Java, ts_java::language());

        Self {
            languages,
            trees: RwLock::new(HashMap::new()),
            contents: RwLock::new(HashMap::new()),
        }
    }

    pub fn infer_language(&self, path: &Path) -> anyhow::Result<LanguageId> {
        let ext = path
            .extension()
            .and_then(|s| s.to_str())
            .ok_or_else(|| anyhow::anyhow!("unsupported file extension"))?;
        language_for(ext).ok_or_else(|| anyhow::anyhow!("unsupported file extension: {}", ext))
    }

    pub fn parse_file(&self, path: &Path) -> anyhow::Result<ts::Tree> {
        if let Some(tree) = self
            .trees
            .read()
            .unwrap()
            .get(path.to_string_lossy().as_ref())
        {
            return Ok(tree.clone());
        }

        let content = std::fs::read(path)
            .with_context(|| format!("failed to read file {}", path.display()))?;
        let lang = self.infer_language(path)?;
        let language = *self
            .languages
            .get(&lang)
            .ok_or_else(|| anyhow::anyhow!("language not loaded"))?;
        let mut parser = ts::Parser::new();
        parser.set_language(language)?;
        let tree = parser
            .parse(&content, None)
            .ok_or_else(|| anyhow::anyhow!("failed to parse tree"))?;

        self.trees
            .write()
            .unwrap()
            .insert(path.to_string_lossy().into(), tree.clone());
        self.contents
            .write()
            .unwrap()
            .insert(path.to_string_lossy().into(), content);
        Ok(tree)
    }

    pub fn get_content(&self, path: &Path) -> anyhow::Result<Vec<u8>> {
        if let Some(content) = self
            .contents
            .read()
            .unwrap()
            .get(path.to_string_lossy().as_ref())
        {
            return Ok(content.clone());
        }
        Ok(std::fs::read(path)?)
    }

    pub fn query(
        &self,
        tree: &ts::Tree,
        query: &str,
        lang: LanguageId,
        content: &[u8],
    ) -> anyhow::Result<Vec<QueryMatch>> {
        let language = *self
            .languages
            .get(&lang)
            .ok_or_else(|| anyhow::anyhow!("language not loaded"))?;
        let query = ts::Query::new(language, query)?;
        let mut cursor = ts::QueryCursor::new();
        let mut matches = Vec::new();
        for m in cursor.matches(&query, tree.root_node(), content) {
            for c in m.captures {
                let name = query.capture_names()[c.index as usize].clone();
                let node = c.node;
                let start = node.start_position();
                let end = node.end_position();
                let text = node.utf8_text(content).unwrap_or("").to_string();
                matches.push(QueryMatch {
                    name,
                    text,
                    start_position: start,
                    end_position: end,
                    start_byte: node.start_byte(),
                    end_byte: node.end_byte(),
                });
            }
        }
        Ok(matches)
    }

    pub fn list_nodes_by_type(
        &self,
        tree: &ts::Tree,
        node_type: &str,
        content: &[u8],
    ) -> Vec<NodeInfo> {
        let mut nodes = Vec::new();
        fn traverse<'a>(node: ts::Node<'a>, content: &[u8], target: &str, out: &mut Vec<NodeInfo>) {
            if node.kind() == target {
                let name_node = node.child_by_field_name("name");
                let name = name_node
                    .and_then(|n| n.utf8_text(content).ok())
                    .unwrap_or("(anonymous)")
                    .to_string();
                out.push(NodeInfo {
                    name,
                    node_type: node.kind().to_string(),
                    text: node.utf8_text(content).unwrap_or("").to_string(),
                    start_position: node.start_position(),
                    end_position: node.end_position(),
                });
            }
            for i in 0..node.child_count() {
                traverse(node.child(i).unwrap(), content, target, out);
            }
        }
        traverse(tree.root_node(), content, node_type, &mut nodes);
        nodes
    }

    pub fn get_context_at_position(
        &self,
        tree: &ts::Tree,
        row: u32,
        col: u32,
        content: &[u8],
    ) -> String {
        let point = ts::Point {
            row: row as usize,
            column: col as usize,
        };
        let node_opt = tree
            .root_node()
            .named_descendant_for_point_range(point, point);
        if node_opt.is_none() {
            return String::new();
        }
        let mut node = node_opt.unwrap();
        let block_types = [
            "function_declaration",
            "method_definition",
            "class_declaration",
            "arrow_function",
            "function_definition",
            "class_definition",
            "method_declaration",
            "type_declaration",
        ];
        loop {
            if block_types.contains(&node.kind()) {
                return node.utf8_text(content).unwrap_or("").to_string();
            }
            if let Some(parent) = node.parent() {
                node = parent;
            } else {
                break;
            }
        }
        tree.root_node()
            .utf8_text(content)
            .unwrap_or("")
            .to_string()
    }
}
