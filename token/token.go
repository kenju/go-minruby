package token

const (
	EOF     = "EOF"
	NEWLINE = "NEWLINE"
	ILLEGAL = "ILLEGAL"

	// Identifiers
	IDENT = "IDENT"

	// Literals
	INT = "INT"

	// Operators
	PLUS       = "+"
	MINUS      = "-"
	ASTERISK   = "*"
	SLASH      = "/"
	PERCENTAGE = "%"

	// Symbols
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
