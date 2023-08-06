// func main() {
// 	// Example usage of the lexer
// 	sourceCode := `
// def add(a, b) do
//     return a + b
// end

// result = add(10, 20)
// if result > 50 do
//     print("Result is greater than 50!")
// else
//     print("Result is less than or equal to 50!")
// end
// `

// 	lexer := NewLexer(sourceCode)

// 	// Tokenize and print the tokens
// 	for !lexer.isEOF() {
// 		token := lexer.GetNextToken()
// 		fmt.Printf("%s: %s\n", token.Type, token.Value)
// 	}
// }