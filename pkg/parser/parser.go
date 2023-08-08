package parser

import (
	"extodan/pkg/lexer"
)

type Parser struct {
	lexer *lexer.Lexer
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer: lexer,
	}
}


