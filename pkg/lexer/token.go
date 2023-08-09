package lexer

// Token types
const (
	TokenEOF         = "EOF"         
	TokenIdentifier  = "IDENTIFIER"  
	TokenInteger     = "INTEGER"     
	TokenString      = "STRING"      
	TokenKeyword     = "KEYWORD"     
	TokenOperator    = "OPERATOR"    
	TokenPunctuation = "PUNCTUATION" 
	TokenComment     = "COMMENT"     
)

// Token represents a single token with its type, value, and position.
type Token struct {
	Type     string         
	Value    string         
	Position SourcePosition 
}
