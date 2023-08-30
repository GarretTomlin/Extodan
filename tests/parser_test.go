package tests

import (
	"strings"
	"testing"

	"extodan/pkg/ast"
	"extodan/pkg/lexer"
	"extodan/pkg/parser"

	"github.com/stretchr/testify/assert"
)

func TestParser_ParseStatement(t *testing.T) {
	sourceCode := `
		return 42
		return 10
	`
	sourceCode = strings.TrimSpace(sourceCode)

	l := lexer.NewLexer(sourceCode)
	p := parser.NewParser(l)

	// Parse the first return statement
	statement1 := p.ParseStatement()
	assert.NotNil(t, statement1)
	assert.Equal(t, ast.NodeReturnStatement, statement1.Type)

	// Parse the second return statement
	statement2 := p.ParseStatement()
	assert.NotNil(t, statement2)
	assert.Equal(t, ast.NodeReturnStatement, statement2.Type)


}


func TestParser_ParseExpression(t *testing.T) {
	sourceCode := "4 + 2"
	l := lexer.NewLexer(sourceCode)
	p := parser.NewParser(l)
	expression := p.ParseExpression()

	// Assert the result
	assert.NotNil(t, expression)
	assert.Equal(t, ast.NodeExpression, expression.Type)
	assert.Equal(t, "+", expression.Children[1].Value)


	// Assert left and right operands
assert.Len(t, expression.Children, 3)
assert.Equal(t, ast.NodeInteger, expression.Children[0].Type)
assert.Equal(t, "4", expression.Children[0].Value)
assert.Equal(t, ast.NodeInteger, expression.Children[2].Type)
assert.Equal(t, "2", expression.Children[2].Value)

}


func TestParser_ParseParameter(t *testing.T) {
	sourceCode := "a"
	l := lexer.NewLexer(sourceCode)
	p := parser.NewParser(l)

	parameterNode := p.ParseParameter()

	assert.NotNil(t, parameterNode)
	assert.Equal(t, ast.NodeParameter, parameterNode.Type)
	assert.Equal(t, "a", parameterNode.Value)
}




