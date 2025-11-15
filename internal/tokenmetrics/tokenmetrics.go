package tokenmetrics

import (
	"unicode/utf8"

	"github.com/tiktoken-go/tokenizer"
)

// DefaultModel is the model used when no explicit token-model is provided.
// It should match the model you most commonly use in benchmarks.
const DefaultModel = "gpt-4o"

// Metrics captures basic token statistics for a piece of text.
type Metrics struct {
	Model      string `json:"model"`
	CharCount  int    `json:"output_char_count"`
	TokenCount int    `json:"output_token_count"`
}

// Count computes token metrics for the given text using the specified model.
// If the model is unsupported, it falls back to a generic encoding (cl100k_base).
func Count(text, model string) (Metrics, error) {
	if model == "" {
		model = DefaultModel
	}

	codec, err := tokenizer.ForModel(tokenizer.Model(model))
	if err != nil {
		// Fallback: use a generic encoding that works for many chat models
		codec, err = tokenizer.Get(tokenizer.Cl100kBase)
		if err != nil {
			return Metrics{}, err
		}
	}

	tokens, err := codec.Count(text)
	if err != nil {
		return Metrics{}, err
	}

	return Metrics{
		Model:      codec.GetName(),
		CharCount:  utf8.RuneCountInString(text),
		TokenCount: tokens,
	}, nil
}
