package lexer

import (
	"regexp"
	"strings"
)

// Lexer holds the input source code and the current position.
type Lexer struct {
	input        string
	position     int
	currentToken Token
}

// KeywordMap maps keywords to their corresponding token types.
var KeywordMap = map[string]string{
    "func":   TokenKeyword,
    "do":     TokenKeyword,
    "end":    TokenKeyword,
    "return": TokenKeyword,
    "fn":     TokenKeyword,
    "endfunc": TokenKeyword,
    "when":   TokenKeyword,
    "else":   TokenKeyword,
    "if":     TokenKeyword,
}


// NewLexer creates a new Lexer with the input source code.
func NewLexer(input string) *Lexer {
	// Remove any leading and trailing spaces from the input source code
	input = strings.TrimSpace(input)

	// Add an EOF token at the end of the input source code
	input = input + "\n"

	return &Lexer{
		input: input,
	}
}

// isEOF checks if the lexer has reached the end of the input.
func (l *Lexer) isEOF() bool {
	return l.position >= len(l.input)-1
}

// readNextToken reads the next token from the input source code.
func (l *Lexer) readNextToken() {
	// Regular expressions for different token types
	identifierRegex := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*`)
	integerRegex := regexp.MustCompile(`^\d+`)
	stringRegex := regexp.MustCompile(`^"(.*?)"`)
	operatorRegex := regexp.MustCompile(`^\+|-|\*|/`)
	punctuationRegex := regexp.MustCompile(`^(\(|\)|,|\{|}|\[|\]|;|\.)`)
	commentRegex := regexp.MustCompile(`^#.*$`)
	spaceRegex := regexp.MustCompile(`^\s+`)

	newlineRegex := regexp.MustCompile(`^\r?\n`)

	for !l.isEOF() {
		// Check for spaces and newlines first
		if matches := spaceRegex.FindString(l.input[l.position:]); matches != "" {
			l.position += len(matches)
		} else if matches := newlineRegex.FindString(l.input[l.position:]); matches != "" {
			l.position += len(matches)
		} else {
			// Check for different token patterns
			if matches := identifierRegex.FindString(l.input[l.position:]); matches != "" {
				// Check if the identifier is a keyword
				if tokenType, isKeyword := KeywordMap[strings.ToLower(matches)]; isKeyword {
					l.currentToken = Token{Type: tokenType, Value: matches}
				} else {
					l.currentToken = Token{Type: TokenIdentifier, Value: matches}
				}
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
			} else if matches := operatorRegex.FindString(l.input[l.position:]); matches != "" {
				l.currentToken = Token{Type: TokenOperator, Value: matches}
				l.position++
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
	}

	// At the end of the input, return the EOF token
	l.currentToken = Token{Type: TokenEOF, Value: ""}
}

// GetNextToken returns the next token from the input source code.
func (l *Lexer) GetNextToken() Token {
	l.readNextToken()
	return l.currentToken
}

// IsEOF checks if the lexer has reached the end of the input.
func (l *Lexer) IsEOF() bool {
	return l.isEOF()
}

// CurrentToken returns the current token from the lexer.
func (l *Lexer) CurrentToken() Token {
	return l.currentToken
}
