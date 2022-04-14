package main

import (
	"testing"
)

func TestGCD(t *testing.T) {
	four := gcd(8, 12)

	if four != 4 {
		t.Fail()
	}

	hundred := gcd(100, 1000)

	if hundred != 100 {
		t.Fail()
	}
}

func TestRatio(t *testing.T) {
	r := CreateRatio(10, -20)
	if r.numerator != 10 && r.denominator != 20 {
		t.Fail()
	}

	if r.Evaluate() > 0 {
		t.Fail()
	}
}

func TestSimplify(t *testing.T) {
	r := CreateRatio(-20, 30)

	r.Simplify()

	if r.numerator != 2 && r.denominator != 3 {
		t.Fail()
	}

	if r.Evaluate() > 0 {
		t.Fail()
	}
}

func TestEquality(t *testing.T) {
	r := CreateRatio(10, -20)
	n := CreateRatio(-20, 40)

	if Equals(r, n) != true {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	r := CreateRatio(-10, 20)
	n := CreateRatio(7, 3)

	// -1/2 + 7/3 = -3/6 + 14/6 = 11/3
	resultant := Add(r, n)

	if resultant.numerator != 11 && resultant.denominator != 6 {
		t.Fail()
	}

	if resultant.Evaluate() < 0 {
		t.Fail()
	}
}

func TestSubtract(t *testing.T) {
	r := CreateRatio(-10, 20)
	n := CreateRatio(7, 3)

	// -1/2 - 7/3 = -3/6 - 14/6 = -19/6
	resultant := Subtract(r, n)

	if resultant.numerator != 17 && resultant.denominator != 6 {
		t.Fail()
	}

	if resultant.Evaluate() > 0 {
		t.Fail()
	}
}
func TestMultiply(t *testing.T) {
	r := CreateRatio(-2, 3)
	n := CreateRatio(10, 4)

	resultant := Multiply(r, n)

	if resultant.numerator != 5 && resultant.denominator != 3 {
		t.Fail()
	}

	if resultant.Evaluate() > 0 {
		t.Fail()
	}
}

func TestDivide(t *testing.T) {
	r := CreateRatio(-2, 3)
	n := CreateRatio(10, 4)

	resultant := Divide(r, n)

	if resultant.numerator != 4 && resultant.denominator != 15 {
		t.Fail()
	}

	if resultant.Evaluate() > 0 {
		t.Fail()
	}
}
