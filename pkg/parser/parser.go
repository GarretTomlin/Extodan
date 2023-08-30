package parser

import (
	"extodan/pkg/ast"   // Import the ast package
	"extodan/pkg/lexer" // Import the lexer package
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
		statement := p.ParseStatement()
		if statement.Type != ast.NodeMissingExpression {
			programNode = ast.AddChild(programNode, statement)
		}
	}
	return programNode
}


func (p *Parser) ParseParameter() ast.Node {
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

func (p *Parser) ParseStatement() ast.Node {
	// Currently, only return statements are supported for demonstration purposes
	if p.curToken.Type == lexer.TokenKeyword && p.curToken.Value == "return" {
		returnStmt := ast.NewNode(ast.NodeReturnStatement, "")
		p.advanceToken()
		expression := p.ParseExpression()
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



func (p *Parser) ParseExpression() ast.Node {
	expression := p.parseAdditiveExpression()
	return expression
}

func (p *Parser) parseAdditiveExpression() ast.Node {
	leftOperand := p.parseMultiplicativeExpression()

	for p.curToken.Type == lexer.TokenOperator && (p.curToken.Value == "+" || p.curToken.Value == "-") {
		operator := p.curToken.Value
		p.advanceToken() // Consume the operator

		rightOperand := p.parseMultiplicativeExpression()

		expression := ast.NewNode(ast.NodeExpression, "")
		expression = ast.AddChild(expression, leftOperand)

		operatorNode := ast.NewNode(ast.NodeOperator, operator)
		expression = ast.AddChild(expression, operatorNode)

		expression = ast.AddChild(expression, rightOperand)

		leftOperand = expression
	}

	return leftOperand
}

func (p *Parser) parseMultiplicativeExpression() ast.Node {
	leftOperand := p.parsePrimaryExpression()

	for p.curToken.Type == lexer.TokenOperator && (p.curToken.Value == "*" || p.curToken.Value == "/") {
		operator := p.curToken.Value
		p.advanceToken() // Consume the operator

		rightOperand := p.parsePrimaryExpression()

		expression := ast.NewNode(ast.NodeExpression, "")
		expression = ast.AddChild(expression, leftOperand)

		operatorNode := ast.NewNode(ast.NodeOperator, operator)
		expression = ast.AddChild(expression, operatorNode)

		expression = ast.AddChild(expression, rightOperand)

		leftOperand = expression
	}

	return leftOperand
}

func (p *Parser) parsePrimaryExpression() ast.Node {
	if p.curToken.Type == lexer.TokenInteger {
		intNode := ast.NewNode(ast.NodeInteger, p.curToken.Value)
		p.advanceToken() // Consume the integer token
		return intNode
	}

	// Handle other types of primary expressions (e.g., parentheses, identifiers, etc.)
	// Return the appropriate AST node based on the parsed expression
	return ast.NewNode(ast.NodeMissingExpression, "Unsupported Primary Expression")
}



func (p *Parser) ParseFunctionDeclaration() ast.Node {
	if p.curToken.Type == lexer.TokenKeyword && p.curToken.Value == "func" {
		p.advanceToken() // Consume "func" keyword
		if p.curToken.Type == lexer.TokenIdentifier {
			funcName := p.curToken.Value
			p.advanceToken() // Consume function name

			// Parse function parameters
			parameters := p.ParseParameter()

			// Parse function body
			body := p.ParseFunctionBody()

			// Parse function expressiona
			expression := p.ParseExpression()

			//parse Statement

			statement := p.ParseStatement()

			// Create the function declaration node
			funcDeclNode := ast.NewNode(ast.NodeFunctionDeclaration, funcName)
			funcDeclNode = ast.AddChild(funcDeclNode, parameters)
			funcDeclNode = ast.AddChild(funcDeclNode, body)
			funcDeclNode = ast.AddChild(funcDeclNode, expression)
			funcDeclNode = ast.AddChild(funcDeclNode, statement)

			return funcDeclNode
		}
	}

	// Create a placeholder node for unsupported function declaration
	return ast.NewNode(ast.NodeFunctionDeclaration, "UnsupportedFunction")
}


func (p *Parser) ParseFunctionBody() ast.Node {
	if p.curToken.Type == lexer.TokenKeyword && p.curToken.Value == "do" {
		p.advanceToken() // Consume "do" keyword

		// Parse statements within the function body
		functionBody := ast.NewNode(ast.NodeFunctionBody, "")
		for p.curToken.Type != lexer.TokenKeyword || p.curToken.Value != "endfunc" {
			statement := p.ParseStatement()
			functionBody = ast.AddChild(functionBody, statement)
		}

		// Consume "endfunc" keyword
		p.advanceToken()

		return functionBody
	}

	// Create a placeholder node for unsupported function body
	return ast.NewNode(ast.NodeFunctionBody, "UnsupportedFunctionBody")
}

// Parse the program using the provided lexer.
func (p *Parser) Parse() ast.Node {
	return p.parseProgram()
}
