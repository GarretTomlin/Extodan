package errors

import (
    "fmt"
)

// SourcePosition represents a position in the source code.
type SourcePosition struct {
    Line   int
    Column int
}

// ExtodanError represents a generic error in the Extodan language.
type ExtodanError struct {
    Message string
    Position SourcePosition
}

func (e ExtodanError) Error() string {
    return fmt.Sprintf("error at line %d, column %d: %s", e.Position.Line, e.Position.Column, e.Message)
}

// ParsingError represents an error that occurs during parsing.
type ParsingError struct {
    ExtodanError
}

// UnexpectedTokenError represents an error for unexpected tokens.
type UnexpectedTokenError struct {
    ExtodanError
    Token string
}

func (e UnexpectedTokenError) Error() string {
    return fmt.Sprintf("unexpected token '%s' at line %d, column %d: %s", e.Token, e.Position.Line, e.Position.Column, e.Message)
}

// SyntaxError represents a syntax error in the code.
type SyntaxError struct {
    ExtodanError
}

// SemanticError represents a semantic error in the code.
type SemanticError struct {
    ExtodanError
}

// IOValidationError represents an error related to I/O validation.
type IOValidationError struct {
    ExtodanError
    Path string
}


// NewSourcePosition creates a SourcePosition instance.
func NewSourcePosition(line, column int) SourcePosition {
    return SourcePosition{Line: line, Column: column}
}


// CreateParsingError creates a ParsingError instance with the provided message, line, and column.
func CreateParsingError(message string, line, column int) ParsingError {
    return ParsingError{
        ExtodanError: ExtodanError{Message: message, Position: NewSourcePosition(line, column)},
    }
}

// CreateUnexpectedTokenError creates an UnexpectedTokenError instance with the provided message, token, line, and column.
func CreateUnexpectedTokenError(message, token string, line, column int) UnexpectedTokenError {
    return UnexpectedTokenError{
        ExtodanError: ExtodanError{Message: message, Position: NewSourcePosition(line, column)},
        Token:        token,
    }
}

// CreateSyntaxError creates a SyntaxError instance with the provided message, line, and column.
func CreateSyntaxError(message string, line, column int) SyntaxError {
    return SyntaxError{
        ExtodanError: ExtodanError{Message: message, Position: NewSourcePosition(line, column)},
    }
}

// CreateSemanticError creates a SemanticError instance with the provided message, line, and column.
func CreateSemanticError(message string, line, column int) SemanticError {
    return SemanticError{
        ExtodanError: ExtodanError{Message: message, Position: NewSourcePosition(line, column)},
    }
}

// CreateIOValidationError creates an IOValidationError instance with the provided message, path, line, and column.
func CreateIOValidationError(message, path string, line, column int) IOValidationError {
    return IOValidationError{
        ExtodanError: ExtodanError{Message: message, Position: NewSourcePosition(line, column)},
        Path:         path,
    }
}
