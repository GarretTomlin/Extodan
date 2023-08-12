package tests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"extodan/pkg/lexer"
)

func TestLexer(t *testing.T) {
	sourceCode := `func multiple(a, b) do
	return a + b
endfunc`

	// Remove any leading and trailing spaces from the sourceCode
	sourceCode = strings.TrimSpace(sourceCode)

	expectedTokens := []lexer.Token{
		{Type: lexer.TokenKeyword, Value: "func"},
		{Type: lexer.TokenIdentifier, Value: "multiple"},
		{Type: lexer.TokenPunctuation, Value: "("},
		{Type: lexer.TokenIdentifier, Value: "a"},
		{Type: lexer.TokenPunctuation, Value: ","},
		{Type: lexer.TokenIdentifier, Value: "b"},
		{Type: lexer.TokenPunctuation, Value: ")"},
		{Type: lexer.TokenKeyword, Value: "do"},
		{Type: lexer.TokenKeyword, Value: "return"},
		{Type: lexer.TokenIdentifier, Value: "a"},
		{Type: lexer.TokenOperator, Value: "+"},
		{Type: lexer.TokenIdentifier, Value: "b"},
		{Type: lexer.TokenKeyword, Value: "endfunc"},
		{Type: lexer.TokenEOF, Value: ""},
	}

	l := lexer.NewLexer(sourceCode)

	for _, expectedToken := range expectedTokens {
		token := l.GetNextToken()
		assert.Equal(t, expectedToken, token)
	}

	assert.True(t, l.IsEOF())
}
