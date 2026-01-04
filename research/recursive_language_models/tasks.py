from __future__ import annotations

import random
import string
from dataclasses import dataclass
from typing import Tuple


@dataclass
class NeedleTask:
    """Represents a synthetic S-NIAH style example."""

    context: str
    question: str
    answer: str


def generate_s_niah(
    needle: str,
    total_chars: int = 20_000,
    seed: int = 7,
    surround: str = "S-NIAH NEEDLE",
) -> NeedleTask:
    """
    Build a single-needle-in-a-haystack prompt where the model must recover the needle.

    Args:
        needle: The string to embed in the haystack.
        total_chars: Total size of the generated context.
        seed: Seed for deterministic placement.
        surround: Phrase around the needle to make peeking easier to validate.
    """
    rng = random.Random(seed)
    filler = _random_text(max(0, total_chars - len(needle) - len(surround) * 2), rng)
    insert_at = rng.randint(0, len(filler))
    context = filler[:insert_at] + surround + needle + surround + filler[insert_at:]
    question = f"Return the exact needle string that appears between two '{surround}' markers."
    return NeedleTask(context=context, question=question, answer=needle)


def _random_text(length: int, rng: random.Random) -> str:
    alphabet = string.ascii_letters + string.digits + " "
    return "".join(rng.choice(alphabet) for _ in range(length))
