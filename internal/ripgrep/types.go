package ripgrep

// Result represents a single match from ripgrep
type Result struct {
	Type string       `json:"type"`
	Data SortableData `json:"data"`
}

type SortableData struct {
	Path           TextData   `json:"path"`
	Lines          TextData   `json:"lines"`
	LineNumber     int        `json:"line_number"`
	AbsoluteOffset int        `json:"absolute_offset"`
	Submatches     []Submatch `json:"submatches"`
}

type TextData struct {
	Text string `json:"text"`
}

type Submatch struct {
	Match string `json:"match"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

// FileMatch represents a file containing a match
type FileMatch struct {
	Path      string
	Line      int
	Column    int
	MatchText string
}
