package lexer

import (
	"fmt"
	"monkey-lang/token"
	"reflect"
	"testing"
)

func TestLexer_readChar(t *testing.T) {
	content := "stringasfa"
	var ch byte
	//var r int32
	ch = content[0]
	//r = content[0]
	fmt.Println(ch, string(ch))
	for v, h := range content {
		fmt.Println(v, string(h))

	}

	//fmt.Println(content[1])
}

func TestLexer_NextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// TODO: Add test cases.
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPARAN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPARAN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPARAN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPARAN, ")"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := l.NextToken()
			if !reflect.DeepEqual(tt.expectedType, got.Type) {
				t.Errorf("NextToken() = %v, want %v", got.Type, tt.expectedType)
			}
			if !reflect.DeepEqual(tt.expectedLiteral, got.Literal) {
				t.Errorf("NextToken() = %v, want %v", got.Literal, tt.expectedLiteral)
			}

		})
	}
}
