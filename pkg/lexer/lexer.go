package lexer

import (
	"fmt"
	"regexp"
)


// Lexer holds the input source code and the current position.
type Lexer struct {
	input        string
	position     int
	currentToken Token
}

// NewLexer creates a new Lexer with the input source code.
func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

// isEOF checks if the lexer has reached the end of the input.
func (l *Lexer) isEOF() bool {
	return l.position >= len(l.input)
}

// readNextToken reads the next token from the input source code.
func (l *Lexer) readNextToken() {
	if l.isEOF() {
		l.currentToken = Token{Type: TokenEOF, Value: ""}
		return
	}

	// Regular expressions for different token types
	identifierRegex := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*`)
	integerRegex := regexp.MustCompile(`^\d+`)
	stringRegex := regexp.MustCompile(`^"(.*?)"`)
	keywordRegex := regexp.MustCompile(`^(def|do|end|if|else|return)`)
	operatorRegex := regexp.MustCompile(`^(\+|-|\*|\/|=|==|!=|<|<=|>|>=)`)
	punctuationRegex := regexp.MustCompile(`^(\(|\)|\{|\}|\[|\]|,|;)`)
	commentRegex := regexp.MustCompile(`^#.*$`)

	if matches := identifierRegex.FindString(l.input[l.position:]); matches != "" {
		l.currentToken = Token{Type: TokenIdentifier, Value: matches}
		l.position += len(matches)
		return
	} else if matches := integerRegex.FindString(l.input[l.position:]); matches != "" {
		l.currentToken = Token{Type: TokenInteger, Value: matches}
		l.position += len(matches)
		return
	} else if matches := stringRegex.FindStringSubmatch(l.input[l.position:]); len(matches) == 2 {
		l.currentToken = Token{Type: TokenString, Value: matches[1]}
		l.position += len(matches[0])
		return
	} else if matches := keywordRegex.FindString(l.input[l.position:]); matches != "" {
		l.currentToken = Token{Type: TokenKeyword, Value: matches}
		l.position += len(matches)
		return
	} else if matches := operatorRegex.FindString(l.input[l.position:]); matches != "" {
		l.currentToken = Token{Type: TokenOperator, Value: matches}
		l.position += len(matches)
		return
	} else if matches := punctuationRegex.FindString(l.input[l.position:]); matches != "" {
		l.currentToken = Token{Type: TokenPunctuation, Value: matches}
		l.position += len(matches)
		return
	} else if matches := commentRegex.FindString(l.input[l.position:]); matches != "" {
		l.currentToken = Token{Type: TokenComment, Value: matches}
		l.position += len(matches)
		return
	}

	// If no token pattern is matched, consider it an unrecognized character
	l.currentToken = Token{Type: "UNKNOWN", Value: string(l.input[l.position])}
	l.position++
}

// GetNextToken returns the next token from the input source code.
func (l *Lexer) GetNextToken() Token {
	l.readNextToken()
	return l.currentToken
}

func main() {
	// Example usage of the lexer
	sourceCode := `
def add(a, b) do
    return a + b
end

result = add(10, 20)
if result > 50 do
    print("Result is greater than 50!")
else
    print("Result is less than or equal to 50!")
end
`

	lexer := NewLexer(sourceCode)

	// Tokenize and print the tokens
	for !lexer.isEOF() {
		token := lexer.GetNextToken()
		fmt.Printf("%s: %s\n", token.Type, token.Value)
	}
}
