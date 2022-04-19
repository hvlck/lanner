package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	r, _ := Divide(10, 20, 2)

	if r != 0.5 {
		t.Fail()
	}
}
