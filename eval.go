package main

import (
	"fmt"
	"strconv"
)

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
	}

	return 0.0
}

func (l *Lexer) toValue(v interface{}) (float64, error) {
	switch z := v.(type) {
	case UnaryOp:
		return evalUnary(z), nil
	default:
		{
			break
		}
	}

	return 0.0, nil
}

type UnaryOp struct {
	lhs float64
	rhs float64
	op  Token
}

func (l *Lexer) groupOp(v ...TokenEntry) (interface{}, error) {
	op := v[1].token

	if (len(v) == 3) && (op == ADD || op == SUBTRACT || op == DIVIDE || op == MULTIPLY) {
		lhs, err := strconv.ParseFloat(v[0].text, 64)
		rhs, er := strconv.ParseFloat(v[2].text, 64)
		if er != nil || err != nil {
			return nil, err
		}

		return UnaryOp{
			lhs: lhs,
			rhs: rhs,
			op:  op,
		}, nil
	}
	return nil, nil
}

func (l *Lexer) eval() (float64, error) {
	value := 0.0
	for t := 0; t <= len(l.tokens)-1; t++ {
		tok := l.tokens[t]
		if t+2 > len(l.tokens) {
			break
		}

		op, err := l.groupOp(l.tokens[t : t+3]...)
		fmt.Println(op)

		if err != nil {
			return 0.0, err
		}

		if tok.token == LNBREAK {
			return value, nil
		}

		z, _ := l.toValue(op)
		value = value + z
	}

	return value, nil
}
