package parser

import (
	"extodan/pkg/lexer"
)

// Parser holds the lexer and the current token.
type Parser struct {
	lexer  *lexer.Lexer
	errors []string
}

// NewParser creates a new Parser with the provided Lexer.
func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{
		lexer:  l,
		errors: make([]string, 0),
	}
}

