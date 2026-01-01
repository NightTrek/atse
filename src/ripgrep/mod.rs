use serde::Deserialize;
use std::process::{Command, Stdio};
use std::time::Duration;
use std::{io::BufRead, io::BufReader};

#[derive(Debug, Clone, Default)]
pub struct SearchOptions {
    pub includes: Vec<String>,
    pub excludes: Vec<String>,
}

#[derive(Debug, Clone)]
pub struct FileMatch {
    pub path: String,
    pub line: u32,
    pub column: u32,
    pub match_text: String,
}

#[derive(Debug, Clone)]
pub struct Client {
    pub executable: String,
    pub timeout: Duration,
}

impl Default for Client {
    fn default() -> Self {
        Self {
            executable: "rg".to_string(),
            timeout: Duration::from_secs(30),
        }
    }
}

impl Client {
    pub fn available(&self) -> bool {
        which::which(&self.executable).is_ok()
    }

    pub fn search(
        &self,
        query: &str,
        path: &str,
        opts: &SearchOptions,
    ) -> anyhow::Result<Vec<FileMatch>> {
        let mut args = vec!["--json".to_string()];
        // Respect the configured timeout to avoid runaway searches.
        args.push("--timeout".to_string());
        args.push(self.timeout.as_millis().to_string());
        for inc in &opts.includes {
            args.push("--glob".into());
            args.push(inc.clone());
        }
        for exc in &opts.excludes {
            args.push("--glob".into());
            args.push(format!("!{}", exc));
        }
        args.push("-e".into());
        args.push(query.to_string());
        args.push(path.to_string());

        let mut cmd = Command::new(&self.executable);
        cmd.args(&args)
            .stdout(Stdio::piped())
            .stderr(Stdio::piped());

        let mut child = cmd.spawn()?;
        let stdout = child.stdout.take().expect("stdout available");
        let reader = BufReader::new(stdout);

        let mut matches = Vec::new();
        for line in reader.lines() {
            let line = line?;
            if line.trim().is_empty() {
                continue;
            }
            if let Ok(event) = serde_json::from_str::<RipgrepEvent>(&line) {
                if let Some(mat) = event.to_match() {
                    matches.push(mat);
                }
            }
        }

        let status = child.wait()?;
        if status.success() || status.code() == Some(1) {
            Ok(matches)
        } else {
            let mut stderr = String::new();
            if let Some(mut err) = child.stderr.take() {
                use std::io::Read;
                err.read_to_string(&mut stderr)?;
            }
            anyhow::bail!("ripgrep failed: {}", stderr);
        }
    }
}

#[derive(Debug, Deserialize)]
struct RipgrepEvent {
    #[serde(rename = "type")]
    kind: String,
    data: Option<RipgrepData>,
}

#[derive(Debug, Deserialize)]
struct RipgrepData {
    path: Option<RipgrepPath>,
    lines: Option<RipgrepLines>,
    line_number: Option<u32>,
    submatches: Option<Vec<RipgrepSubmatch>>,
}

#[derive(Debug, Deserialize)]
struct RipgrepPath {
    text: String,
}

#[derive(Debug, Deserialize)]
struct RipgrepLines {
    text: String,
}

#[derive(Debug, Deserialize)]
struct RipgrepSubmatch {
    start: u32,
}

impl RipgrepEvent {
    fn to_match(self) -> Option<FileMatch> {
        if self.kind != "match" {
            return None;
        }
        let data = self.data?;
        let path = data.path?.text;
        let line = data.line_number.unwrap_or(0);
        let match_text = data.lines.map(|l| l.text).unwrap_or_default();
        let column = data
            .submatches
            .as_ref()
            .and_then(|s| s.first())
            .map(|s| s.start)
            .unwrap_or(0);
        Some(FileMatch {
            path,
            line,
            column,
            match_text,
        })
    }
}
