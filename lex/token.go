package lanner

import "fmt"

type Span struct {
	Line, Col int
}

type Token struct {
	Type       TokenType
	Literal    string
	Start, End Span
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %v, Literal: %v, Start: %v, End: %v}", t.Type, t.Literal, t.Start, t.End)
}

type TokenType string

const (
	EOI      = ""
	ILLEGAL  = "ILLEGAL"
	LN_BREAK = ""

	LPAREN = "("
	RPAREN = ")"
	ADD    = "+"
	SUB    = "-"
	MULT   = "*"
	DIV    = "/"
	EXP    = "^"
	BANG   = "!"

	EQUALITY     = "=="
	NOT_EQUALITY = "!="
	GT           = ">"
	LT           = "<"
	GTEQ         = ">="
	LTEQ         = "<="

	// variables, etc
	IDENT = "IDENT"

	INT   = "INTEGER"
	FLOAT = "FLOAT"

	TRUE  = "TRUE"
	FALSE = "FALSE"
)
