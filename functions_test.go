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
	fTwo, isT := IsInt(two)
	iT := fTwo.Int64()

	fifteen := Factorial(15)

	fFift, is := IsInt(fifteen)
	iF := fFift.Int64()
	if !is || !isT {
		t.Log("return value not an int")
	}

	if iT != 2 || iF != 1307674368000 {
		t.Log("factorials incorrect")
		t.Fail()
	}
}
