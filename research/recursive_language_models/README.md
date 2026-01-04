## Recursive Language Models (RLM) Playground

This folder contains a lightweight, testable reimplementation of the core ideas from **Recursive Language Models** (Zhang et al., 2025). The goal is to make the paper practical inside this repository by providing:

- A minimal runtime that lets a base LLM *peek* into long contexts, recurse on sub-spans, and keep a tool-usage trace.
- Synthetic tasks (e.g., S-NIAH / long haystack retrieval) to exercise the recursion loop without needing the full benchmarks from the paper.
- Tests that validate the control flow with a mock LLM so you can iterate without external API access.

### Files

- `rlm.py` – Core runtime with `RecursiveLanguageModel`, tool definitions, and a traceable execution loop.
- `tasks.py` – Synthetic generators for S-NIAH-style long-context prompts and questions.
- `tests/` – Pytest coverage using a mock LLM to validate peeking and recursion behavior.
- `llm_clients.py` – Simple client interface plus an optional OpenAI-compatible client for real-world runs.

### Quickstart

1. (Optional) Install runtime deps for live model runs:
   ```bash
   pip install openai
   ```
2. Run the built-in tests (no network required):
   ```bash
   python -m pytest research/recursive_language_models/tests
   ```
3. Try a manual demo with your own `OPENAI_API_KEY`:
   ```bash
   python -m research.recursive_language_models.rlm --question "Find the needle" --context-file /path/to/long.txt
   ```

### Design Notes

- **Environment-as-context:** The RLM never feeds the entire document to the model. Instead it exposes `peek(start, end)` slices and lets the model request spans programmatically.
- **Recursive calls:** The model can ask the runtime to spawn sub-queries over narrower spans, mirroring the paper’s recursive descent strategy.
- **Traceability:** Every action (peek, recurse, answer) is recorded so you can inspect tool usage and cost.
- **Pluggable LLMs:** Any client that implements the simple `LLMClient` protocol works; the included OpenAI client is optional.

### Extending

- Add new tools (e.g., regex search, semantic chunk selection) by extending `Toolset` in `rlm.py`.
- Swap in a different base model by implementing `LLMClient`.
- Plug in real benchmarks by generating contexts in `tasks.py` or loading your own corpora, then invoking `RecursiveLanguageModel.run`.
