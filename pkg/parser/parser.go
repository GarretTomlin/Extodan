package parser

import (
	"extodan/pkg/lexer"
)

// Parser represents the parsing component that processes tokens from the lexer.
type Parser struct {
	lexer *lexer.Lexer // The lexer used to tokenize the source code.
}

// NewParser creates a new Parser instance with the provided lexer.
func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer: lexer,
	}
}
