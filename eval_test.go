package main

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	r := strings.NewReader("20+20\n")

	l := Lex(r)
	l.Start()
	l.removeWhitespace()

	v, err := l.eval()

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if v != 40 {
		t.Log("addition failed")
		t.Fail()
	}
}

func TestSubtract(t *testing.T) {
	r := strings.NewReader("20-20\n")

	l := Lex(r)
	l.Start()
	l.removeWhitespace()

	v, err := l.eval()

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if v != 0 {
		t.Log("subtraction failed")
		t.Fail()
	}
}

func TestMultiply(t *testing.T) {
	r := strings.NewReader("20*20\n")

	l := Lex(r)
	l.Start()
	l.removeWhitespace()

	v, err := l.eval()

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if v != 400 {
		t.Log("multiplication failed")
		t.Fail()
	}
}

func TestDivide(t *testing.T) {
	r := strings.NewReader("20/20\n")

	l := Lex(r)
	l.Start()
	l.removeWhitespace()

	v, err := l.eval()

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if v != 1 {
		t.Log("division failed")
		t.Fail()
	}
}
