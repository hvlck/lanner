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
