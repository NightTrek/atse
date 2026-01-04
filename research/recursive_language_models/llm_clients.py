from __future__ import annotations

import importlib.util
import json
from dataclasses import dataclass
from typing import Iterable, List, Mapping, Protocol


Message = Mapping[str, str]


class LLMClient(Protocol):
    """Minimal interface for chat-based models."""

    def complete(self, messages: Iterable[Message]) -> str:
        """Return the model's raw string response."""
        raise NotImplementedError


@dataclass
class OpenAIClient:
    """
    Thin wrapper around the OpenAI Chat Completions API.

    This client is intentionally lightweight to avoid pulling the dependency
    unless the caller opts in. It expects the `openai` package to be available
    and an `OPENAI_API_KEY` environment variable to be set.
    """

    model: str = "gpt-4.1-mini"

    def complete(self, messages: Iterable[Message]) -> str:
        if importlib.util.find_spec("openai") is None:  # pragma: no cover - import presence checked in integration
            raise RuntimeError("Install `openai` and set OPENAI_API_KEY to use OpenAIClient.")
        from openai import OpenAI  # type: ignore

        client = OpenAI()
        completion = client.chat.completions.create(
            model=self.model,
            messages=list(messages),
            response_format={"type": "json_object"},
            max_tokens=400,
        )
        return completion.choices[0].message.content or ""


class MockLLMClient:
    """
    Deterministic client for tests.

    Provide a list of JSON-serializable actions. Each `complete` call pops the
    next action and returns it as a JSON string, letting tests drive the runtime
    deterministically without a real model.
    """

    def __init__(self, scripted_actions: List[Mapping[str, object]]):
        self._actions = list(scripted_actions)

    def complete(self, messages: Iterable[Message]) -> str:  # pragma: no cover - exercised via tests
        if not self._actions:
            raise RuntimeError("MockLLMClient exhausted scripted actions.")
        action = self._actions.pop(0)
        return json.dumps(action)
