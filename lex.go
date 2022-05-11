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

	POWER     // exponentiation
	FN        // function
	FACTORIAL // factorial (!)
	PIPE

	// comparison
	GT  // greater than
	LT  // less than
	LTE // less than equal to
	GTE // greater than equal to
	EQ  // equal (=)

	LPARAN
	RPARAN
	POINT
	WHITESPACE
	LNBREAK

	NUMBER
	FLOAT

	UNIT
)

var tokens = []string{
	EOI: "EOI",

	ADD:      "+",
	SUBTRACT: "-",
	MULTIPLY: "*",
	DIVIDE:   "/",

	POWER:     "^",
	FN:        "fn()",
	FACTORIAL: "!",
	PIPE:      "|",

	GT:  ">",
	LT:  "<",
	LTE: "<=",
	GTE: ">=",
	EQ:  "=",

	LPARAN:     "(",
	RPARAN:     ")",
	POINT:      ".",
	WHITESPACE: "WHITESPACE",
	LNBREAK:    "LNBREAK",

	NUMBER: "NUMBER",
	FLOAT:  "FLOAT",

	UNIT: "UNIT",
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
	case '^':
		return l.span, POWER, "^"
	case '(':
		return l.span, LPARAN, "("
	case ')':
		return l.span, RPARAN, ")"
	case '!':
		return l.span, FACTORIAL, "!"
	case '=':
		return l.span, EQ, "="
	default:
		{
			if unicode.IsDigit(rn) {
				n, fp := l.lexNumber()
				if fp {
					return l.span, FLOAT, n
				}

				return l.span, NUMBER, n
			} else if rn == '<' || rn == '>' {
				p := l.advanceIf('=')

				if p == true && rn == '<' {
					return l.span, LTE, "<="
				} else if p == true && rn == '>' {
					return l.span, GTE, ">="
				} else if p == false {
					if rn == '>' {
						return l.span, GT, ">"
					} else if rn == '<' {
						return l.span, LT, "<"
					}
				}
			} else if rn == '\n' || rn == '\r' {
				return l.span, LNBREAK, string(rn)
			} else if unicode.IsSpace(rn) {
				return l.span, WHITESPACE, string(rn)
			} else if unicode.IsLetter(rn) {
				fn := l.lexText()
				return l.span, FN, fn
			}
		}
	}
	return Span{}, EOI, ""
}

// list of functions
// (s) supported | (ns) not supported
// sine sin(x) (ns)
// cosine cos(x) (ns)
// tangent tan(x) (ns)
// arcsine arcsin(x) (ns)
const FUNCTIONS = "SIN COS TAN ARCSIN ARCCOS ARCTAN SEC CSC COT ARCSEC ARCCSC ARCCOT MIN MAX"

// list of US customary units
// in inch length 12in=1ft
// ft foot length 3ft=1yd
// yd yard length
// mi mile length 5280ft=1mi
// oz ounce weight 16oz=1lb
// lb pound weight 2000lb=1T
const US_CUSTOMARY_UNITS = "in ft yd mi oz lb T floz C pt qt F"

// list of SI unit bases
// base units:
// m meter length
// g gram mass
// s second time
// A ampere electric current
// K kelvin temperature
// mol mole amount of substance
// cd candela luminous intensity
// partial list of derived units:
// N (kg*m/s^2) Newton force
// J (N*m) Joule energy/work
// W (J/s) Watts power
// Hz (s^-1) Hertz frequency
// V (m^2·kg·s^-3·A^-1) volt electric potential difference
// ohm (m^2·kg·s^-3·A^-2) ohm electric resistance
const SI_UNITS = "m g s A K mol cd N J W Hz V ohm"
const SI_PREFIXES = "Y Z E P T G M k h da d c m micro n p f a z y"

func (l *Lexer) advanceIf(rn rune) bool {
	r, _, _ := l.reader.ReadRune()

	if r == rn {
		return true
	} else {
		l.unadvance(1)
		return false
	}
}

func (l *Lexer) lexText() string {
	l.unadvance(1)

	fn := ""
	for {
		r, _, _ := l.reader.ReadRune()

		l.span.column++

		if unicode.IsLetter(r) {
			fn += string(r)
		} else if r == '(' {
			// same behavior for now
			l.unadvance(1)
			break
		} else {
			l.unadvance(1)
			break
		}
	}

	return fn
}

// lexes a complete number
// returns `false` in the boolean return value if the number is not floating-point, `true` otherwise
func (l *Lexer) lexNumber() (string, bool) {
	l.unadvance(1)

	number := ""
	fp := true
	for {
		// fix err at some point
		r, _, _ := l.reader.ReadRune()

		l.span.column++

		if unicode.IsDigit(r) {
			number += string(r)
		} else if r == '.' {
			fp = true
			number += "."
		} else {
			l.unadvance(1)
			break
		}
	}

	return number, fp
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
