package lexer

import (
	"testing"

	"github.com/kenju/go-minruby/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `
	p(3)
	p(1 + 1)
	4 - 3
	3 * 4
	10 / 2
	10 % 3
	1 + 2 + 3 % 4 * 5 * 6 + 7 - 8 / 9
	-10
	x = 42
	true
	false
	if 7 > 5
		1
	else
		0
	end
	!true
	10 == 10
	10 != 9
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NEWLINE, "\n"},

		// p(3)
		{token.IDENT, "p"},
		{token.LPAREN, "("},
		{token.INT, "3"},
		{token.RPAREN, ")"},

		{token.NEWLINE, "\n"},

		// p(1 + 1)
		{token.IDENT, "p"},
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

		{token.NEWLINE, "\n"},

		// 10 % 3
		{token.INT, "10"},
		{token.PERCENTAGE, "%"},
		{token.INT, "3"},

		{token.NEWLINE, "\n"},

		// 1 + 2 + 3 % 4 * 5 * 6 + 7 - 8 / 9
		{token.INT, "1"},
		{token.PLUS, "+"},
		{token.INT, "2"},
		{token.PLUS, "+"},
		{token.INT, "3"},
		{token.PERCENTAGE, "%"},
		{token.INT, "4"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.ASTERISK, "*"},
		{token.INT, "6"},
		{token.PLUS, "+"},
		{token.INT, "7"},
		{token.MINUS, "-"},
		{token.INT, "8"},
		{token.SLASH, "/"},
		{token.INT, "9"},

		{token.NEWLINE, "\n"},

		// -10
		{token.MINUS, "-"},
		{token.INT, "10"},

		{token.NEWLINE, "\n"},

		// x = 42
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "42"},

		{token.NEWLINE, "\n"},

		// true
		{token.TRUE, "true"},

		{token.NEWLINE, "\n"},

		// false
		{token.FALSE, "false"},

		{token.NEWLINE, "\n"},

		// if 7 > 5
		//   1
		//  else
		//    0
		//  end
		{token.IF, "if"},
		{token.INT, "7"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.NEWLINE, "\n"},
		{token.INT, "1"},
		{token.NEWLINE, "\n"},
		{token.ELSE, "else"},
		{token.NEWLINE, "\n"},
		{token.INT, "0"},
		{token.NEWLINE, "\n"},
		{token.END, "end"},

		{token.NEWLINE, "\n"},

		// !true
		{ token.BANG, "!" },
		{ token.TRUE, "true" },

		{token.NEWLINE, "\n"},

		// 10 == 10
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},

		{token.NEWLINE, "\n"},

		// 10 != 9
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
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
