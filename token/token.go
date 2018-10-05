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
	ASSIGN     = "="
	PLUS       = "+"
	MINUS      = "-"
	ASTERISK   = "*"
	SLASH      = "/"
	PERCENTAGE = "%"
	LT         = "<"
	GT         = ">"

	// Symbols
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	END      = "END"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
	"if":    IF,
	"else":  ELSE,
	"end":   END,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
