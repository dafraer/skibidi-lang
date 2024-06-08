package lexer

import (
	"fmt"
	"skibidilang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `skibidi five1 = 5;
skibidi ten_ = 10;
skibidi add = ohio(x, y) {
x + y;
};
skibidi result1 = add(five1, ten_);
!-/*5^;
5 < 10 > 5;

&|%[]
if (5 < 10) {
	goon alpha;
} else {
	goon beta;
}
10 == 10;
10 != 9;
>=
<=
++
--
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "skibidi"},
		{token.IDENT, "five1"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "skibidi"},
		{token.IDENT, "ten_"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "skibidi"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "ohio"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.ADD, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "skibidi"},
		{token.IDENT, "result1"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five1"},
		{token.COMMA, ","},
		{token.IDENT, "ten_"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.NOT, "!"},
		{token.SUB, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.POWER, "^"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.AND, "&"},
		{token.OR, "|"},
		{token.MOD, "%"},
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "goon"},
		{token.TRUE, "alpha"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "goon"},
		{token.FALSE, "beta"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.GEQ, ">="},
		{token.LEQ, "<="},
		{token.INC, "++"},
		{token.DEC, "--"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Printf("iteration %v, token type %v, token literal %v\n", i, tok.Type, tok.Literal)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
