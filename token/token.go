package token

const (
	// Literals
	INT = "INT"

	// Operators
	PLUS = "+"

	// Symbols
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	FUNCTION = "FUNCTION"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
