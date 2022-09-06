package parse

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hvlck/lanner/ast"
	lanner "github.com/hvlck/lanner/lex"
)

type parse_node struct {
	expectedType string
	value        interface{}
}

func expect_int(stmt ast.Expression, val interface{}) string {
	if num, ok := stmt.(*ast.NumberLiteral); ok {
		valNum := val.(*ast.NumberLiteral)
		if num.Value == valNum.Value {
			return ""
		}

		return fmt.Sprintf("expected int of value %v, got %v", valNum.Value, num.Value)
	}

	fmt.Printf("%T\n", stmt)
	return "not a number"
}

func expect_infix(stmt ast.Expression, val interface{}) string {
	if num, ok := stmt.(*ast.InfixExpression); ok {
		if valExpr, ok := val.(*ast.InfixExpression); ok {
			if strings.Contains(valExpr.Left.String(), "(") {
				fmt.Println(num.Left, num)
				r := expect_infix(num.Left, valExpr.Left)
				if len(r) != 0 {
					return r
				}
			}

			if strings.Contains(valExpr.Right.String(), "(") {
				r := expect_infix(num.Right, valExpr.Right)
				if len(r) != 0 {
					return r
				}
			}

			if valExpr.Left.String() == num.Left.String() && valExpr.Right.String() == num.Right.String() && valExpr.Op == num.Op {
				return ""
			}

			return "expressions are not equal"
		}
		return "value passed is not infix expression pointer"
	} else if num, ok := stmt.(*ast.NumberLiteral); ok {
		if _, ok := val.(*ast.NumberLiteral); !ok {
			fmt.Println(val)
			return fmt.Sprintf("value is not a number literal, got %T instead", val)
		}
		return expect_int(num, val)
	}

	return "statement is not infix expression"
}

func expect_parse(t *testing.T, input string, expected []parse_node) {
	lexer := lanner.New(input)
	parser := New(lexer)
	program := parser.Parse()
	if len(program.Errors) != 0 {
		t.Fail()
	}

	for i, v := range expected {
		e := ""
		if len(program.Expressions) <= i {
			t.Fail()
			return
		}

		switch v.expectedType {
		case "int":
			e = expect_int(program.Expressions[i], v.value)
		case "op":
			fmt.Println("parsing ", input)
			e = expect_infix(program.Expressions[0], v.value)
		}

		if len(e) != 0 {
			t.Fatalf("error: %v", e)
		}
	}
}

func TestInfix(t *testing.T) {
	input := []string{
		`10`,
		`(10+10)`,
		`(10+20)*10`,
	}
	expected := [][]parse_node{
		{{
			expectedType: "int",
			value: &ast.NumberLiteral{
				Value: 10,
			},
		}},
		{{
			expectedType: "op",
			value: &ast.InfixExpression{
				Left:  &ast.NumberLiteral{Value: 10},
				Right: &ast.NumberLiteral{Value: 10},
				Op:    "+",
			},
		}},
		{{
			expectedType: "op",
			value: &ast.InfixExpression{
				Left: &ast.InfixExpression{
					Left:  &ast.NumberLiteral{Value: 10},
					Right: &ast.NumberLiteral{Value: 20},
					Op:    "+",
				},
				Right: &ast.NumberLiteral{Value: 10},
				Op:    "*",
			},
		}},
	}

	for i, v := range input {
		expect_parse(t, v, expected[i])
	}
}
