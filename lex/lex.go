package lanner

import (
	"unicode"
)

type Lexer struct {
	input   string
	current int

	line int
	col  int
}

func (l *Lexer) unread() {
	l.current--
}

func (l *Lexer) nextCharacter() (rune, bool) {
	if len(l.input) > l.current+1 {
		l.current++
		rn := rune(l.input[l.current])

		if rn == '\n' {
			l.line++
			l.col = 0
		}

		l.col++
		return rn, false
	} else {
		return 0, true
	}
}

func makeToken(tokType TokenType, literal string, start Span) Token {
	return Token{
		Type:    tokType,
		Literal: literal,
		Start:   start,
	}
}

func (l *Lexer) Next() Token {
	var token Token

	nxt, eoi := l.nextCharacter()

	starting := Span{Col: l.col, Line: l.line}
	if eoi {
		token = makeToken(EOI, "", starting)
		token.End = starting

		return token
	}

	switch nxt {
	case '(':
		token = makeToken(LPAREN, "(", starting)
	case ')':
		token = makeToken(RPAREN, ")", starting)
	case '+':
		token = makeToken(ADD, "+", starting)
	case '-':
		token = makeToken(SUB, "-", starting)
	case '*':
		token = makeToken(MULT, "*", starting)
	case '/':
		token = makeToken(DIV, "/", starting)
	case '^':
		token = makeToken(EXP, "^", starting)
	default:
		if unicode.IsNumber(nxt) {
			n, tok := l.lexNumber()
			token = makeToken(tok, n, starting)
		} else if unicode.IsSpace(nxt) {
			if nxt == '\n' {
				token = makeToken(LN_BREAK, "", starting)
			} else {
				return l.Next()
			}
		} else {
			token = makeToken(ILLEGAL, string(nxt), starting)
		}
	}

	token.End = Span{
		Line: l.line,
		Col:  l.col,
	}

	return token
}

func (l *Lexer) lexNumber() (string, TokenType) {
	starting := l.current
	isFloat := false
	for v, eoi := rune(l.input[l.current]), false; unicode.IsNumber(v) || v == '.'; v, eoi = l.nextCharacter() {
		if eoi || l.current == len(l.input)-1 {
			l.current = len(l.input)
			break
		}

		if v == '.' {
			isFloat = true
		}
	}

	defer l.unread()

	num := l.input[starting:l.current]
	if isFloat {
		return num, FLOAT
	}

	return num, INT
}

func New(src string) *Lexer {
	return &Lexer{
		input:   src,
		current: -1,
		line:    1,
		col:     -1,
	}
}
