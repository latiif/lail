package lexer

import (
	"testing"

	"github.com/latiif/lail/pkg/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x,y) {
  x+y;
};
let result = add(five,ten);
!-/*5;
5 < 10 > 5;
if (5 < 10) {
  return true;
} else {
  return false;
}
10 == 10;
10 != 9;
>= <=
"foobar"
"foo bar"
import "file"
`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Ident, "five"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "ten"},
		{token.Assign, "="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.Lparen, "("},
		{token.Ident, "x"},
		{token.Comma, ","},
		{token.Ident, "y"},
		{token.Rparen, ")"},
		{token.Lbrace, "{"},
		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "result"},
		{token.Assign, "="},
		{token.Ident, "add"},
		{token.Lparen, "("},
		{token.Ident, "five"},
		{token.Comma, ","},
		{token.Ident, "ten"},
		{token.Rparen, ")"},
		{token.Semicolon, ";"},
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Astersik, "*"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Int, "5"},
		{token.LT, "<"},
		{token.Int, "10"},
		{token.GT, ">"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.If, "if"},
		{token.Lparen, "("},
		{token.Int, "5"},
		{token.LT, "<"},
		{token.Int, "10"},
		{token.Rparen, ")"},
		{token.Lbrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Else, "else"},
		{token.Lbrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Int, "10"},
		{token.EQ, "=="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Int, "10"},
		{token.NEQ, "!="},
		{token.Int, "9"},
		{token.Semicolon, ";"},
		{token.GTE, ">="},
		{token.LTE, "<="},
		{token.String, "foobar"},
		{token.String, "foo bar"},
		{token.Import, "import"},
		{token.String, "file"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tc := range tests {
		tok := l.NextToken()

		if tok.Type != tc.expectedType {
			t.Fatalf("test[%d] - token.Type wrong. got: %q, want: %q", i, tok.Type, tc.expectedType)
		}

		if tok.Literal != tc.expectedLiteral {
			t.Fatalf("test[%d] - token.Literal wrong. got: %q, want: %q", i, tok.Literal, tc.expectedLiteral)
		}

		if tok.Col == 0 || tok.Line == 0 {
			t.Fatalf("test[%d] (%q) - token.Col or token.Line are not correctly set. Both are 0.", i, tc.expectedLiteral)
		}
	}
}

func TestTokenCoordinates(t *testing.T) {
	input :=
		`
let myvar = 5;
if "my long string"
`
	tests := []struct {
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		{"let", 1, 2},
		{"myvar", 5, 2},
		{"=", 11, 2},
		{"5", 13, 2},
		{";", 14, 2},
		{"if", 1, 3},
		{"my long string", 4, 3},
	}

	l := New(input)
	for i, tc := range tests {
		tok := l.NextToken()

		if tok.Literal != tc.expectedLiteral {
			t.Fatalf("test[%d] - token.Literal wrong. got: %q, want: %q", i, tok.Literal, tc.expectedLiteral)
		}

		if tok.Line != tc.expectedLine {
			t.Fatalf("test[%d] - token.Line wrong. got: %d, want: %d", i, tok.Line, tc.expectedLine)
		}

		if tok.Col != tc.expectedColumn {
			t.Fatalf("test[%d] - token.Line wrong. got: %d, want: %d", i, tok.Col, tc.expectedColumn)
		}
	}
}
