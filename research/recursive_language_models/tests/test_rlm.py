from __future__ import annotations

import json
from pathlib import Path

from research.recursive_language_models.llm_clients import MockLLMClient
from research.recursive_language_models.rlm import RecursiveLanguageModel, RLMConfig
from research.recursive_language_models.tasks import generate_s_niah


def test_peek_then_answer():
    context = "prefix NEEDLE suffix"
    client = MockLLMClient(
        [
            {"action": "peek", "start": 7, "end": 13, "rationale": "Look near the center"},
            {"action": "answer", "response": "NEEDLE"},
        ]
    )
    rlm = RecursiveLanguageModel(context=context, client=client, config=RLMConfig(max_steps=4))
    result = rlm.run("What is the needle?")
    assert result["answer"] == "NEEDLE"
    assert len(result["trace"]) == 2
    assert result["trace"][0]["observation"] == "NEEDLE"


def test_recurses_into_span_and_returns_child_answer(tmp_path: Path):
    task = generate_s_niah("TARGET", total_chars=400, seed=3)
    child_script = [
        {"action": "peek", "start": 0, "end": 80, "rationale": "read span"},
        {"action": "answer", "response": "TARGET"},
    ]
    parent_script = [
        {"action": "recurse", "start": 0, "end": 120, "question": task.question, "rationale": "delegate"},
    ]

    client = MockLLMClient(parent_script + child_script)
    rlm = RecursiveLanguageModel(task.context, client=client, config=RLMConfig(max_steps=6))
    result = rlm.run(task.question)

    assert result["answer"] == "TARGET"
    # Parent recurse action plus two child steps.
    assert len(result["trace"]) == 3
    assert result["trace"][0]["action"]["action"] == "recurse"
    assert result["trace"][1]["action"]["action"] == "peek"
    assert result["trace"][2]["action"]["action"] == "answer"
