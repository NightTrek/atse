use std::collections::HashMap;
use std::sync::RwLock;

#[derive(Debug, Clone)]
pub struct SymbolNode {
    pub symbol: String,
    pub symbol_type: String,
    pub file_path: String,
    pub line: u32,
}

#[derive(Debug, Clone)]
pub struct FileIndex {
    pub path: String,
    pub content: Vec<u8>,
    pub symbols: Vec<SymbolNode>,
}

#[derive(Debug, Default)]
pub struct PartialIndex {
    pub parsed_files: RwLock<HashMap<String, FileIndex>>,
    pub symbols: RwLock<HashMap<String, Vec<SymbolNode>>>,
}

impl PartialIndex {
    pub fn add_file(&self, path: String, content: Vec<u8>, symbols: Vec<SymbolNode>) {
        {
            let mut pf = self.parsed_files.write().unwrap();
            pf.insert(
                path.clone(),
                FileIndex {
                    path: path.clone(),
                    content,
                    symbols: symbols.clone(),
                },
            );
        }
        let mut sym_map = self.symbols.write().unwrap();
        for sym in symbols {
            sym_map.entry(sym.symbol.clone()).or_default().push(sym);
        }
    }

    pub fn has_file(&self, path: &str) -> bool {
        self.parsed_files.read().unwrap().contains_key(path)
    }

    pub fn find_symbol(&self, name: &str) -> Vec<SymbolNode> {
        self.symbols
            .read()
            .unwrap()
            .get(name)
            .cloned()
            .unwrap_or_default()
    }

    pub fn count(&self) -> (usize, usize) {
        let pf = self.parsed_files.read().unwrap();
        let sym = self.symbols.read().unwrap();
        (pf.len(), sym.len())
    }
}

#[derive(Debug, Clone)]
pub struct TypeUsage {
    pub type_name: String,
    pub usage_type: String,
    pub file_path: String,
    pub line: u32,
}

#[derive(Debug, Clone)]
pub struct Import {
    pub import_path: String,
    pub file_path: String,
    pub line: u32,
}

#[derive(Debug, Clone)]
pub struct CallSite {
    pub caller_symbol: String,
    pub caller_file_path: String,
    pub caller_line: u32,
    pub called_symbol: String,
}

#[derive(Debug, Clone)]
pub struct EventUsage {
    pub event_name: String,
    pub event_type: String,
    pub file_path: String,
    pub line: u32,
}
