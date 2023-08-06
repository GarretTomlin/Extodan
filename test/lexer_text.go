package tests

import (
	"testing"

    "github.com/GarretTomlin/extodan/pkg/lexer" // Update the import path here
)

func TestLexer(t *testing.T) {
	// Your test cases for the lexer go here
	// ...

	t.Run("ExampleTest", func(t *testing.T) {
		sourceCode := `def add(a, b) do
			return a + b
		end`

		expectedTokens := []lexer.Token{
			{Type: lexer.TokenKeyword, Value: "def"},
			{Type: lexer.TokenIdentifier, Value: "add"},
			{Type: lexer.TokenPunctuation, Value: "("},
			{Type: lexer.TokenIdentifier, Value: "a"},
			{Type: lexer.TokenPunctuation, Value: ","},
			{Type: lexer.TokenIdentifier, Value: "b"},
			{Type: lexer.TokenPunctuation, Value: ")"},
			{Type: lexer.TokenKeyword, Value: "do"},
			{Type: lexer.TokenIdentifier, Value: "return"},
			{Type: lexer.TokenIdentifier, Value: "a"},
			{Type: lexer.TokenOperator, Value: "+"},
			{Type: lexer.TokenIdentifier, Value: "b"},
			{Type: lexer.TokenKeyword, Value: "end"},
			{Type: lexer.TokenEOF, Value: ""},
		}

		lexer := lexer.NewLexer(sourceCode)

		for _, expectedToken := range expectedTokens {
			token := lexer.GetNextToken()
			if token.Type != expectedToken.Type || token.Value != expectedToken.Value {
				t.Errorf("Expected %s: %s, got %s: %s", expectedToken.Type, expectedToken.Value, token.Type, token.Value)
			}
		}

		if !lexer.IsEOF() {
			t.Errorf("Expected end of file, got %s: %s", lexer.CurrentToken().Type, lexer.CurrentToken().Value)
		}
	})
}
