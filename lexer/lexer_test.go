package lexer

import (
	"testing"

	"github.com/kenju/go-minruby/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `
	p(1 + 1)
	4 - 3
	3 * 4
	10 / 2
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NEWLINE, "\n"},

		// p(1 + 1)
		{token.FUNCTION, "p"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.PLUS, "+"},
		{token.INT, "1"},
		{token.RPAREN, ")"},

		{token.NEWLINE, "\n"},

		// 4 - 3
		{token.INT, "4"},
		{token.MINUS, "-"},
		{token.INT, "3"},

		{token.NEWLINE, "\n"},

		// 3 * 4
		{token.INT, "3"},
		{token.ASTERISK, "*"},
		{token.INT, "4"},

		{token.NEWLINE, "\n"},

		// 10 / 2
		{token.INT, "10"},
		{token.SLASH, "/"},
		{token.INT, "2"},
	}

	lexer := New(input)
	for i, tt := range tests {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
