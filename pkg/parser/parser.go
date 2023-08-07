package parser

import (
	"fmt"
	"strconv"
	"extodan/pkg/lexer"
)

// Parser holds the lexer and the current token.
type Parser struct {
	lexer        *lexer.Lexer
	currentToken lexer.Token
}

// NewParser creates a new Parser with the given lexer.
func NewParser(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer: lexer,
	}
	parser.readNextToken()
	return parser
}

// readNextToken reads the next token from the lexer.
func (p *Parser) readNextToken() {
	p.currentToken = p.lexer.GetNextToken()
}

// consumeToken checks if the current token type matches the expected type
// and moves to the next token if they match. Otherwise, it raises an error.
func (p *Parser) consumeToken(expectedType string) {
	if p.currentToken.Type == expectedType {
		p.readNextToken()
	} else {
		panic(fmt.Sprintf("Expected token type '%s', but got '%s'", expectedType, p.currentToken.Type))
	}
}

// parseInteger parses an INTEGER token and returns its integer value.
func (p *Parser) parseInteger() int {
	value, err := strconv.Atoi(p.currentToken.Value)
	if err != nil {
		panic("Invalid integer value")
	}
	p.consumeToken(lexer.TokenInteger)
	return value
}

// parseFactor parses a factor in the expression.
func (p *Parser) parseFactor() int {
	if p.currentToken.Type == lexer.TokenInteger {
		return p.parseInteger()
	} else if p.currentToken.Type == lexer.TokenPunctuation && p.currentToken.Value == "(" {
		p.consumeToken(lexer.TokenPunctuation)
		result := p.parseExpression()
		p.consumeToken(lexer.TokenPunctuation)
		return result
	} else {
		panic("Invalid factor")
	}
}

// parseTerm parses a term in the expression.
func (p *Parser) parseTerm() int {
	result := p.parseFactor()
	for p.currentToken.Type == lexer.TokenOperator && (p.currentToken.Value == "*" || p.currentToken.Value == "/") {
		operator := p.currentToken.Value
		p.consumeToken(lexer.TokenOperator)
		operand := p.parseFactor()
		if operator == "*" {
			result *= operand
		} else {
			result /= operand
		}
	}
	return result
}

// parseExpression parses an expression.
func (p *Parser) parseExpression() int {
	result := p.parseTerm()
	for p.currentToken.Type == lexer.TokenOperator && (p.currentToken.Value == "+" || p.currentToken.Value == "-") {
		operator := p.currentToken.Value
		p.consumeToken(lexer.TokenOperator)
		operand := p.parseTerm()
		if operator == "+" {
			result += operand
		} else {
			result -= operand
		}
	}
	return result
}

// ParseProgram starts parsing the input source code.
func (p *Parser) ParseProgram() int {
	return p.parseExpression()
}
