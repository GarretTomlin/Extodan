package tests

import (
	"fmt"
	"strings"
	"testing"

	"extodan/pkg/lexer"
)

func TestLexer(t *testing.T) {
	sourceCode := `func add(a, b) do
	return a + b
end`

	// Remove any leading and trailing spaces from the sourceCode
	sourceCode = strings.TrimSpace(sourceCode)

	expectedTokens := []lexer.Token{
		{Type: lexer.TokenKeyword, Value: "func"},
		{Type: lexer.TokenIdentifier, Value: "add"},
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
		{Type: lexer.TokenKeyword, Value: "end"},
		{Type: lexer.TokenEOF, Value: ""}, 
	}

	l := lexer.NewLexer(sourceCode)

	for i, expectedToken := range expectedTokens {
		token := l.GetNextToken()
		fmt.Printf("Expected: %s: %s, Got: %s: %s\n", expectedToken.Type, expectedToken.Value, token.Type, token.Value)
		if token.Type != expectedToken.Type || token.Value != expectedToken.Value {
			t.Errorf("Mismatch at index %d. Expected %s: %s, got %s: %s", i, expectedToken.Type, expectedToken.Value, token.Type, token.Value)
		}
	}

	if !l.IsEOF() {
		t.Errorf("Expected end of input, got %s: %s", l.CurrentToken().Type, l.CurrentToken().Value)
	}
}
