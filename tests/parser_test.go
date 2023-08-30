
package tests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"extodan/pkg/lexer"
	"extodan/pkg/parser"
	"extodan/pkg/ast"
)

func TestParser_ParseStatement(t *testing.T) {
	sourceCode := `
		return 42
		return 10
	`
	sourceCode = strings.TrimSpace(sourceCode)

	l := lexer.NewLexer(sourceCode)
	p := parser.NewParser(l)
	program := p.Parse()

	assert.NotNil(t, program)
	assert.Equal(t, ast.NodeProgram, program.Type)
	assert.Len(t, program.Children, 2)

	assert.Equal(t, ast.NodeReturnStatement, program.Children[0].Type)
	assert.Equal(t, ast.NodeReturnStatement, program.Children[1].Type)

	assert.Equal(t, ast.NodeInteger, program.Children[0].Children[0].Type)
	assert.Equal(t, "42", program.Children[0].Children[0].Value)

	assert.Equal(t, ast.NodeInteger, program.Children[1].Children[0].Type)
	assert.Equal(t, "10", program.Children[1].Children[0].Value)
}

func TestParser_ParseFunctionDeclaration(t *testing.T) {
	sourceCode := `
		func add(a, b) do
			return a + b
		endfunc
	`
	sourceCode = strings.TrimSpace(sourceCode)

	l := lexer.NewLexer(sourceCode)
	p := parser.NewParser(l)
	program := p.Parse()

	assert.NotNil(t, program)
	assert.Equal(t, ast.NodeProgram, program.Type)
	assert.Len(t, program.Children, 1)

	assert.Equal(t, ast.NodeFunctionDeclaration, program.Children[0].Type)
	funcDeclNode := program.Children[0]

	assert.Len(t, funcDeclNode.Children, 4) // Function name, parameters, function body

	// Assert other details of the function declaration and its contents
	assert.Equal(t, "add", funcDeclNode.Value) // Check the function name

	// Check the parameters
	assert.Len(t, funcDeclNode.Children[0].Children, 2)
	assert.Equal(t, "a", funcDeclNode.Children[0].Children[0].Value)
	assert.Equal(t, "b", funcDeclNode.Children[0].Children[1].Value)

	// Check the function body
	assert.Equal(t, ast.NodeProgram, funcDeclNode.Children[1].Type)
	assert.Len(t, funcDeclNode.Children[1].Children, 1) // Return statement

	// Check the return statement
	assert.Equal(t, ast.NodeReturnStatement, funcDeclNode.Children[1].Children[0].Type)
	assert.NotNil(t, funcDeclNode.Children[1].Children[0].Children[0]) // Check that an expression is present
	assert.Equal(t, ast.NodeExpression, funcDeclNode.Children[1].Children[0].Children[0].Type)

	// Check the specific contents of the expression
	expression := funcDeclNode.Children[1].Children[0].Children[0]
	assert.Equal(t, "+", expression.Value) // Check the operator
	assert.NotNil(t, expression.Children[0]) // Check that left operand is present
	assert.NotNil(t, expression.Children[1]) // Check that right operand is present
	assert.Equal(t, ast.NodeIdentifier, expression.Children[0].Type) // Check the left operand type
	assert.Equal(t, "a", expression.Children[0].Value) // Check the left operand value
	assert.Equal(t, ast.NodeIdentifier, expression.Children[1].Type) // Check the right operand type
	assert.Equal(t, "b", expression.Children[1].Value) // Check the right operand value
}
