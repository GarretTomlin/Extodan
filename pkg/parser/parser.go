package parser

import (
	"extodan/pkg/lexer" // Import the lexer package
	"extodan/pkg/ast"   // Import the ast package
)

// Parser holds the lexer and the current token.
type Parser struct {
	lexer    *lexer.Lexer
	curToken lexer.Token
}

// NewParser creates a new parser using a lexer.
func NewParser(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	parser.advanceToken()
	return parser
}

// advanceToken advances to the next token.
func (p *Parser) advanceToken() {
	p.curToken = p.lexer.GetNextToken()
}

func (p *Parser) parseProgram() ast.Node {
	programNode := ast.NewNode(ast.NodeProgram, "")
	for !p.lexer.IsEOF() {
		statement := p.parseStatement()
		if statement.Type != ast.NodeMissingExpression {
			programNode = ast.AddChild(programNode, statement)
		}
	}
	return programNode
}


func (p *Parser) parseFunctionDeclaration() ast.Node {
	if p.curToken.Type == lexer.TokenKeyword && p.curToken.Value == "func" {
		p.advanceToken() // Consume 'func'

		// Parse function name and parameters
		funcName := p.curToken.Value
		p.advanceToken() // Consume function name

		p.advanceToken() // Consume '('
		parameters := []ast.Node{} // Store parameters
		for p.curToken.Type == lexer.TokenIdentifier {
			parameter := p.parseParameter()
			parameters = append(parameters, parameter)
			if p.curToken.Type == lexer.TokenPunctuation {
				p.advanceToken() // Consume ','
			} else {
				break
			}
		}
		p.advanceToken() // Consume ')'

		// Parse function body
		p.advanceToken() // Consume 'do'
		functionBody := p.parseProgram()
		p.advanceToken() // Consume 'endfunc'

		// Create the function declaration node
		funcDeclNode := ast.NewNode(ast.NodeFunctionDeclaration, funcName)
		for _, parameter := range parameters {
			funcDeclNode = ast.AddChild(funcDeclNode, parameter)
		}
		funcDeclNode = ast.AddChild(funcDeclNode, functionBody)

		return funcDeclNode
	}
	return ast.NewNode(ast.NodeFunctionDeclaration, "UnsupportedFunction")
}

func (p *Parser) parseParameter() ast.Node {
	if p.curToken.Type == lexer.TokenIdentifier {
		paramName := p.curToken.Value
		p.advanceToken() // Consume parameter name

		// Create the parameter node
		paramNode := ast.NewNode(ast.NodeParameter, paramName)
		return paramNode
	}

	// Create a placeholder node for unsupported parameter
	return ast.NewNode(ast.NodeParameter, "UnsupportedParameter")
}

func (p *Parser) parseStatement() ast.Node {
	// Currently, only return statements are supported for demonstration purposes
	if p.curToken.Type == lexer.TokenKeyword && p.curToken.Value == "return" {
		returnStmt := ast.NewNode(ast.NodeReturnStatement, "")
		p.advanceToken()
		expression := p.parseExpression()
		if expression.Type != ast.NodeMissingExpression {
			returnStmt = ast.AddChild(returnStmt, expression)
		} else {
			// Use the missing expression node to represent a missing expression
			returnStmt = ast.AddChild(returnStmt, expression)
		}
		return returnStmt
	}

	// Return a placeholder node for unsupported statements
	return ast.NewNode(ast.NodeReturnStatement, "Unsupported Statement")
}

func (p *Parser) parseExpression() ast.Node {
	// For demonstration purposes, let's assume we only parse integers
	if p.curToken.Type == lexer.TokenInteger {
		intNode := ast.NewNode(ast.NodeInteger, p.curToken.Value)
		p.advanceToken()
		return intNode
	}

	// Return a placeholder node for unsupported expressions
	return ast.NewNode(ast.NodeExpression, "Unsupported Expression")
}

// Parse the program using the provided lexer.
func (p *Parser) Parse() ast.Node {
	return p.parseProgram()
}
