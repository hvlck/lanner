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
	if r.numerator != 10 && r.denominator != 20 && r.positive != false {
		t.Fail()
	}
}

func TestSimplify(t *testing.T) {
	r := CreateRatio(-20, 30)

	r.Simplify()

	if r.numerator != 2 && r.denominator != 3 && r.positive != false {
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

func TestMultiply(t *testing.T) {
	r := CreateRatio(-1, 2)
	n := CreateRatio(2, 1)

	resultant := Multiply(r, n)

	if resultant.numerator != 2 && resultant.denominator != 2 {
		t.Fail()
	}

	if resultant.positive == true {
		t.Fail()
	}
}

func TestDivide(t *testing.T) {
	r := CreateRatio(-1, 2)
	n := CreateRatio(2, 1)

	resultant := Divide(r, n)

	if resultant.numerator != 1 && resultant.denominator != 4 {
		t.Fail()
	}

	if resultant.positive == true {
		t.Fail()
	}
}