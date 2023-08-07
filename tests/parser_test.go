
package tests 

import (

	"testing"

	"extodan/pkg/lexer"
	"extodan/pkg/parser"
)
func TestParser(t *testing.T) {
	input := "2 + 3 * (4 - 2)"
	expectedResult := 8

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	result := parser.ParseProgram()
	if result != expectedResult {
		t.Errorf("Expected result %d, but got %d", expectedResult, result)
	}
}
