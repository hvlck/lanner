package main

import (
	"errors"
	"fmt"
	"strconv"
)

// todo: replace with own implementation
func evalComparison(v UnaryOp) bool {
	switch v.op {
	case LT:
		return v.lhs < v.rhs
	case LTE:
		return v.lhs <= v.rhs
	case GT:
		return v.lhs > v.rhs
	case GTE:
		return v.lhs >= v.rhs
	default:
		return false
	}
}

// todo: replace with own implementation
func evalUnary(v UnaryOp) float64 {
	switch v.op {
	case ADD:
		return v.lhs + v.rhs
	case SUBTRACT:
		return v.lhs - v.rhs
	case MULTIPLY:
		return v.lhs * v.rhs
	case DIVIDE:
		return v.lhs / v.rhs
	default:
		return 0.0
	}
}

// converts any given operation type to a value
// return values: value, # of tokens consumed, error
func (l *Lexer) toValue(v interface{}) (interface{}, uint8, error) {
	fmt.Println(v)
	switch z := v.(type) {
	case UnaryOp:
		{
			if UNARY_OPERATORS[z.op] {
				return evalUnary(z), 3, nil
			} else if COMPARISON_OPS[z.op] {
				return evalComparison(z), 3, nil
			}
		}
	default:
		{
			break
		}
	}

	return 0.0, 0, nil
}

// converts a given array of (three) tokens to a unary operation
func toUnary(v []TokenEntry) (UnaryOp, error) {
	if len(v) < 3 {
		return UnaryOp{}, errors.New("not unary")
	}

	lhs, err := strconv.ParseFloat(v[0].text, 64)
	rhs, er := strconv.ParseFloat(v[2].text, 64)
	if er != nil || err != nil {
		return UnaryOp{}, err
	}

	return UnaryOp{
		lhs: lhs,
		rhs: rhs,
		op:  v[1].token,
	}, nil
}

type UnaryOp struct {
	lhs float64
	rhs float64
	op  Token
}

var UNARY_OPERATORS = map[Token]bool{
	ADD:      true,
	SUBTRACT: true,
	MULTIPLY: true,
	DIVIDE:   true,
}

var COMPARISON_OPS = map[Token]bool{
	LT:  true,
	LTE: true,
	GT:  true,
	GTE: true,
}

func (l *Lexer) groupOp(v ...TokenEntry) (interface{}, error) {
	op := v[1].token

	for a, r := range v {
		if r.token == EOI {
			return nil, errors.New("end of input")
		}
	}
	if (len(v) == 3) && (UNARY_OPERATORS[op] || COMPARISON_OPS[op]) {
		un, err := toUnary(v)

		if err != nil {
			return nil, err
		}

		return un, nil
	}

	return nil, nil
}

// evaluates tokens
func (l *Lexer) eval() (interface{}, error) {
	var value interface{}
	t := 0
	for t <= len(l.tokens)-1 {
		tok := l.tokens[t]

		if t+2 > len(l.tokens) {
			break
		}

		op, err := l.groupOp(l.tokens[t : t+3]...)
		fmt.Println(op)
		if op == nil {
			return value, nil
		}

		if err != nil {
			return 0.0, err
		}

		if tok.token == LNBREAK || tok.token == EOI {
			return value, nil
		}

		z, idx, err := l.toValue(op)
		if err != nil {
			return value, nil
		}
		t += int(idx)

		switch zt := z.(type) {
		case bool:
			return z, nil
		case float64:
			{
				switch r := value.(type) {
				case float64:
					value = r + zt
				}
			}
		}

		t++
	}

	// fmt.Print(value)

	return value, nil
}
