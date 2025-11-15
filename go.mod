module github.com/NightTrek/atse

go 1.23

toolchain go1.24.10

require (
	github.com/smacker/go-tree-sitter v0.0.0-20240827094217-dd81d9e9be82
	github.com/spf13/cobra v1.10.1
	github.com/tiktoken-go/tokenizer v0.7.0
)

require (
	github.com/dlclark/regexp2 v1.11.5 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
)

replace github.com/tree-sitter/tree-sitter-go => github.com/tree-sitter/tree-sitter-go v0.25.0
