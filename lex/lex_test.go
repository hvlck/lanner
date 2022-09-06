package lanner

import (
	"reflect"
	"testing"
)

func expect_lex_deepeq(t *testing.T, input string, expected []Token) {
	l := New(input)
	for _, v := range expected {
		tok := l.Next()
		if !reflect.DeepEqual(tok, v) {
			t.Fatalf("expected %v to be %v", tok, v)
		}
	}
}

type lex_tok struct {
	TokenType
	string
}

func expect_lex_equal(t *testing.T, input string, expected []lex_tok) {
	l := New(input)
	for _, v := range expected {
		tok := l.Next()
		if tok.Literal != v.string || tok.Type != v.TokenType {
			t.Fatalf("expected %v to be %v", tok, v)
		}
	}
}

func TestComplex(t *testing.T) {
	input := `(10+10)*(20 / 10.0)`

	expected := []lex_tok{
		{LPAREN, "("},
		{INT, "10"},
		{ADD, "+"},
		{INT, "10"},
		{RPAREN, ")"},
		{MULT, "*"},
		{LPAREN, "("},
		{INT, "20"},
		{DIV, "/"},
		{FLOAT, "10.0"},
		{RPAREN, ")"},
		{EOI, ""},
	}

	expect_lex_equal(t, input, expected)
}

func TestInfix(t *testing.T) {
	input := `10+10`
	expected := []Token{
		{
			Type:    INT,
			Literal: "10",
			Start: Span{
				Col:  0,
				Line: 1,
			},
			End: Span{
				Col:  2,
				Line: 1,
			},
		},
		{
			Type:    ADD,
			Literal: "+",
			Start: Span{
				Col:  3,
				Line: 1,
			},
			End: Span{
				Col:  3,
				Line: 1,
			},
		},
		{
			Type:    INT,
			Literal: "10",
			Start: Span{
				Col:  4,
				Line: 1,
			},
			End: Span{
				Col:  5,
				Line: 1,
			},
		},
	}
	expect_lex_deepeq(t, input, expected)

}
