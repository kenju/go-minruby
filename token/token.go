package token

const (
	EOF     = "EOF"
	NEWLINE = "NEWLINE"

	// Literals
	INT = "INT"

	// Operators
	PLUS  = "+"
	MINUS = "-"
	ASTERISK = "*"
	SLASH = "/"
	PERCENTAGE = "%"

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
