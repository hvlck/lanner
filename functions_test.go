package main

import (
	"testing"
)

func TestAbsolute(t *testing.T) {
	pos := Absolute(10.22)
	neg := Absolute(-10.22)

	if pos != neg || neg < 0 {
		t.Log("absolute() returning negative absolute value")
		t.Fail()
	}
}

func TestFactorial(t *testing.T) {
	two := Factorial(2)
	fifteen := Factorial(15)

	if two != 2 || fifteen != 1307674368000 {
		t.Log("factorials incorrect")
		t.Fail()
	}
}
