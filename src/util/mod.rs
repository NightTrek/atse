use globset::{Glob, GlobSet, GlobSetBuilder};
use ignore::DirEntry;
use std::path::{Path, PathBuf};
use walkdir::{DirEntry as WalkEntry, WalkDir};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum FileCategory {
    Production,
    Test,
    Generated,
    Config,
}

#[derive(Debug, Clone)]
pub struct FileMatch {
    pub path: PathBuf,
    pub category: FileCategory,
}

pub const DEFAULT_EXCLUDES: &[&str] = &[
    "*.test.ts",
    "*.test.js",
    "*.test.tsx",
    "*.test.jsx",
    "*.spec.ts",
    "*.spec.js",
    "*.spec.tsx",
    "*.spec.jsx",
    "*Test.java",
    "*Tests.java",
    "*IT.java",
    "*_test.go",
    "**/*_test.go",
    "__tests__/**",
    "**/__tests__/**",
    "**/test/**",
    "**/tests/**",
    "**/*.generated.*",
    "**/generated/**",
    "**/*.pb.go",
    "**/*.pb.ts",
    "**/*.pb.js",
    "**/*.d.ts",
    "**/node_modules/**",
    "**/vendor/**",
    "**/target/**",
    "**/__mocks__/**",
    "**/__fixtures__/**",
    "**/dist/**",
    "**/build/**",
    "**/.next/**",
];

fn build_globset(patterns: &[String]) -> Result<GlobSet, globset::Error> {
    let mut builder = GlobSetBuilder::new();
    for pat in patterns {
        builder.add(Glob::new(pat)?);
    }
    builder.build()
}

fn is_supported(path: &Path) -> bool {
    match path
        .extension()
        .and_then(|s| s.to_str())
        .map(|s| s.to_lowercase())
    {
        Some(ext) => matches!(
            ext.as_str(),
            "go" | "js" | "jsx" | "mjs" | "cjs" | "ts" | "tsx" | "py" | "java"
        ),
        None => false,
    }
}

fn classify(path: &Path) -> FileCategory {
    let lower = path.to_string_lossy().to_lowercase();
    let base = path.file_name().and_then(|s| s.to_str()).unwrap_or("");

    for pat in [
        ".test.",
        ".spec.",
        "__test",
        "__tests__",
        "/test/",
        "/tests/",
        "__mocks__",
        "__fixtures__",
    ] {
        if lower.contains(pat) {
            return FileCategory::Test;
        }
    }

    for pat in [
        ".generated.",
        "/generated/",
        ".pb.",
        "_pb.",
        ".d.ts",
        "/dist/",
        "/build/",
        "/.next/",
    ] {
        if lower.contains(pat) {
            return FileCategory::Generated;
        }
    }

    for ext in ["json", "yaml", "yml", "toml", "ini", "config"] {
        if base.ends_with(ext) {
            return FileCategory::Config;
        }
    }

    FileCategory::Production
}

pub fn walk_files(
    root: &Path,
    recursive: bool,
    include: &[String],
    exclude: &[String],
    apply_defaults: bool,
) -> anyhow::Result<Vec<FileMatch>> {
    let mut effective_excludes: Vec<String> = exclude.to_vec();
    if apply_defaults {
        effective_excludes.extend(DEFAULT_EXCLUDES.iter().map(|s| s.to_string()));
    }

    let include_globs = if include.is_empty() {
        None
    } else {
        Some(build_globset(include)?)
    };
    let exclude_globs = build_globset(&effective_excludes)?;

    let walker = WalkDir::new(root).into_iter().filter_entry(|e| {
        if !recursive && e.depth() > 0 && e.file_type().is_dir() {
            return false;
        }
        true
    });

    let mut out = Vec::new();
    for entry in walker {
        let entry: WalkEntry = match entry {
            Ok(e) => e,
            Err(err) => return Err(err.into()),
        };
        if entry.file_type().is_dir() {
            continue;
        }
        let path = entry.path();
        if !is_supported(path) {
            continue;
        }
        if exclude_globs.is_match(path) {
            continue;
        }
        if let Some(include_globs) = &include_globs {
            if !include_globs.is_match(path) {
                continue;
            }
        }
        out.push(FileMatch {
            path: path.to_path_buf(),
            category: classify(path),
        });
    }

    Ok(out)
}

pub fn dir_entry_to_path(entry: &DirEntry) -> Option<PathBuf> {
    if entry.file_type().map(|ft| ft.is_file()).unwrap_or(false) {
        Some(entry.path().to_path_buf())
    } else {
        None
    }
}
