from __future__ import annotations

import argparse
import json
from dataclasses import dataclass, field
from pathlib import Path
from typing import Any, Dict, Iterable, List, Mapping, Optional, Sequence

from .llm_clients import LLMClient


Action = Dict[str, Any]
Message = Mapping[str, str]


@dataclass
class RLMConfig:
    """Configuration controlling recursion and peeking."""

    max_depth: int = 3
    max_steps: int = 12
    peek_chars: int = 800
    system_prompt: str = field(
        default_factory=lambda: (
            "You are a Recursive Language Model (RLM). "
            "The full document lives outside your context as the variable `context`. "
            "You may not quote it directly unless you peek at spans first. "
            "Respond with a JSON object describing exactly one action:\n"
            '- {"action": "peek", "start": int, "end": int, "rationale": str}\n'
            '- {"action": "recurse", "start": int, "end": int, "question": str, "rationale": str}\n'
            '- {"action": "answer", "response": str}\n'
            "Indices are zero-based and inclusive of `start`, exclusive of `end`. "
            "Use recursion to decompose long contexts. "
            "Do not hallucinate content; only rely on spans you have peeked."
        )
    )


@dataclass
class TraceEntry:
    """Records each step for inspection and debugging."""

    step: int
    action: Action
    observation: Optional[str] = None
    depth: int = 0


class RecursiveLanguageModel:
    """
    Minimal RLM runtime that exposes programmatic tools (peek, recurse) to an LLM client.
    """

    def __init__(
        self,
        context: str,
        client: LLMClient,
        config: Optional[RLMConfig] = None,
        start: int = 0,
        end: Optional[int] = None,
        depth: int = 0,
    ):
        self.context = context
        self.client = client
        self.config = config or RLMConfig()
        self.start = start
        self.end = len(context) if end is None else end
        self.depth = depth
        self.trace: List[TraceEntry] = []

    def run(self, question: str) -> Dict[str, Any]:
        """Execute the RLM loop until the model returns an answer."""
        for step in range(1, self.config.max_steps + 1):
            messages = self._build_messages(question)
            raw = self.client.complete(messages)
            action = self._parse_action(raw)
            entry = TraceEntry(step=step, action=action, depth=self.depth)

            if action["action"] == "peek":
                observation = self._peek(action["start"], action["end"])
                entry.observation = observation
                self.trace.append(entry)
                continue

            if action["action"] == "recurse":
                if self.depth + 1 >= self.config.max_depth:
                    raise RuntimeError("Max recursion depth exceeded.")
                span_context = self._peek(action["start"], action["end"])
                child = RecursiveLanguageModel(
                    context=span_context,
                    client=self.client,
                    config=self.config,
                    start=0,
                    end=len(span_context),
                    depth=self.depth + 1,
                )
                child_result = child.run(action["question"])
                entry.observation = child_result["answer"]
                self.trace.append(entry)
                combined_trace = [self._trace_as_dict(t) for t in self.trace]
                combined_trace.extend(child_result.get("trace", []))
                return {"answer": child_result["answer"], "trace": combined_trace}

            if action["action"] == "answer":
                entry.observation = action.get("response")
                self.trace.append(entry)
                materialized_trace = [self._trace_as_dict(t) for t in self.trace]
                return {"answer": action.get("response", ""), "trace": materialized_trace}

            raise RuntimeError(f"Unsupported action: {action}")

        raise RuntimeError("RLM exceeded max_steps without answering.")

    def _build_messages(self, question: str) -> List[Message]:
        summary = {
            "question": question,
            "scope": {"start": self.start, "end": self.end, "length": self.end - self.start},
            "trace": [self._trace_as_dict(t) for t in self.trace],
            "peek_limit": self.config.peek_chars,
            "depth": self.depth,
            "max_depth": self.config.max_depth,
        }
        return [
            {"role": "system", "content": self.config.system_prompt},
            {"role": "user", "content": json.dumps(summary)},
        ]

    def _peek(self, start: int, end: int) -> str:
        absolute_start = self.start + max(0, start)
        absolute_end = self.start + min(self.end - self.start, end)
        span = self.context[absolute_start:absolute_end]
        if len(span) > self.config.peek_chars:
            span = span[: self.config.peek_chars]
        return span

    @staticmethod
    def _parse_action(raw: str) -> Action:
        try:
            action = json.loads(raw)
        except json.JSONDecodeError as exc:
            raise RuntimeError(f"Model returned non-JSON: {raw}") from exc

        if "action" not in action:
            raise RuntimeError(f"Model response missing 'action': {action}")
        return action

    @staticmethod
    def _trace_as_dict(entry: TraceEntry) -> Dict[str, Any]:
        data = {
            "step": entry.step,
            "action": entry.action,
            "depth": entry.depth,
        }
        if entry.observation is not None:
            data["observation"] = entry.observation
        return data


def load_context_from_file(path: Path) -> str:
    return path.read_text(encoding="utf-8")


def main(argv: Optional[Sequence[str]] = None) -> None:
    parser = argparse.ArgumentParser(description="Run a Recursive Language Model demo.")
    parser.add_argument("--context-file", type=Path, required=True, help="Path to a UTF-8 text file.")
    parser.add_argument("--question", required=True, help="Question to ask about the context.")
    parser.add_argument("--model", default="gpt-4.1-mini", help="Base model name for the OpenAI client.")
    args = parser.parse_args(argv)

    from .llm_clients import OpenAIClient

    context = load_context_from_file(args.context_file)
    client = OpenAIClient(model=args.model)
    rlm = RecursiveLanguageModel(context=context, client=client)
    result = rlm.run(args.question)
    print(json.dumps(result, indent=2))


if __name__ == "__main__":
    main()
