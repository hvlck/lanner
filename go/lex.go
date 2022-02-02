package main

import (
	"bufio"
	"io"
	"unicode"
)

type Token int

type Span struct {
	column int
	line   int
}

const (
	EOI = iota

	ADD
	SUBTRACT
	MULTIPLY
	DIVIDE

	POWER
	FN

	LPARAN
	RPARAN
	POINT
	WHITESPACE

	NUMBER
)

var tokens = []string{
	EOI: "EOI",

	ADD:      "+",
	SUBTRACT: "-",
	MULTIPLY: "*",
	DIVIDE:   "/",

	POWER: "^",
	FN:    "fn()",

	LPARAN:     "(",
	RPARAN:     ")",
	POINT:      ".",
	WHITESPACE: "WHITESPACE",

	NUMBER: "NUMBER",
}

type TokenList []Token

func (t Token) String() string {
	return tokens[t]
}

type TokenEntry struct {
	span  Span
	text  string
	token Token
}

type Lexer struct {
	tokens []TokenEntry
	reader *bufio.Reader
	span   Span
}

func Lex(src io.Reader) *Lexer {
	return &Lexer{
		reader: bufio.NewReader(src),
		span:   Span{line: 1, column: 0},
	}
}

func (l *Lexer) Lex(rn rune) (Span, Token, string) {
	switch rn {
	case '+':
		return l.span, ADD, "+"
	case '-':
		return l.span, SUBTRACT, "-"
	case '*':
		return l.span, MULTIPLY, "*"
	case '/':
		return l.span, DIVIDE, "/"
	default:
		{
			if unicode.IsDigit(rn) {
				n := l.lexNumber()
				return l.span, NUMBER, n
			} else if unicode.IsSpace(rn) {
				return l.span, WHITESPACE, string(rn)
			}
		}
	}
	return Span{}, EOI, ""
}

func (l *Lexer) lexNumber() string {
	l.unadvance(1)

	number := ""
	for {
		// fix err at some point
		r, _, _ := l.reader.ReadRune()

		l.span.column++

		if unicode.IsDigit(r) {
			number += string(r)
		} else {
			l.unadvance(1)
			break
		}
	}

	return number
}

func (l *Lexer) unadvance(n int) {
	for i := 0; i < n; i++ {
		if err := l.reader.UnreadRune(); err != nil {
			panic(err)
		}

		l.span.column--
	}
}

func (l *Lexer) removeWhitespace() {
	var n []TokenEntry
	for _, v := range l.tokens {
		if v.token != WHITESPACE {
			n = append(n, v)
		}
	}

	l.tokens = n
}

func (l *Lexer) Start() {
	for {
		rn, _, err := l.reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				// needs to append spans and such
				l.tokens = append(l.tokens, TokenEntry{token: EOI, span: l.span, text: "EOI"})
			}

			break
		}

		if rn == '\n' || rn == '\r' {
			l.advanceLine()
		}

		l.span.column++

		span, tok, txt := l.Lex(rn)

		l.tokens = append(l.tokens, TokenEntry{span: span, token: tok, text: txt})
	}
}

func (l *Lexer) advanceLine() {
	l.span.column = 0
	l.span.line++
}
