package main

import (
	"errors"
	"math/big"
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
func evalUnaryFloat(v UnaryOp) big.Float {
	lhs := big.NewFloat(v.lhs)
	rhs := big.NewFloat(v.rhs)
	switch v.op {
	case ADD:
		return *lhs.Add(lhs, rhs)
	case SUBTRACT:
		return *lhs.Sub(lhs, rhs)
	case MULTIPLY:
		return *lhs.Mul(lhs, rhs)
	case DIVIDE:
		return *lhs.Quo(lhs, rhs)
	default:
		return *big.NewFloat(0.0)
	}
}

// converts any given operation type to a value
// return values: value, # of tokens consumed, error
func (l *Lexer) toValue(v interface{}) (interface{}, uint8, error) {
	switch z := v.(type) {
	case UnaryOp:
		{
			if UNARY_OPERATORS[z.op] {
				return evalUnaryFloat(z), 3, nil
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

	for _, r := range v {
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

func IsInt(v interface{}) (*big.Int, bool) {
	switch r := v.(type) {
	case *big.Int:
		return r, true
	default:
		return new(big.Int), false
	}
}

func IsFloat(v interface{}) (*big.Float, bool) {
	switch r := v.(type) {
	case *big.Float:
		return r, true
	default:
		return new(big.Float), false
	}
}

// evaluates tokens
func (l *Lexer) eval() (interface{}, error) {
	value := big.NewFloat(0)

	for t := 0; t < len(l.tokens)-1; t++ {
		tok := l.tokens[t]

		op, err := l.groupOp(l.tokens[t : t+3]...)
		if op == nil {
			return value, nil
		}

		if err != nil {
			return nil, err
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
		case big.Float:
			{
				value = value.Add(value, &zt)
			}
		}
	}

	return value, nil
}
