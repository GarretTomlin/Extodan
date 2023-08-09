package lexer

// Token types
const (
	TokenEOF         = "EOF"         // Represents the end of the input.
	TokenIdentifier  = "IDENTIFIER"  // Represents an identifier token.
	TokenInteger     = "INTEGER"     // Represents an integer token.
	TokenString      = "STRING"      // Represents a string token.
	TokenKeyword     = "KEYWORD"     // Represents a keyword token.
	TokenOperator    = "OPERATOR"    // Represents an operator token.
	TokenPunctuation = "PUNCTUATION" // Represents a punctuation token.
	TokenComment     = "COMMENT"     // Represents a comment token.
)

// Token represents a single token with its type, value, and position.
type Token struct {
	Type     string         // The type of the token.
	Value    string         // The value associated with the token.
	Position SourcePosition // The position of the token in the source code.
}
