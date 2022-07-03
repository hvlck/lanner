package main

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	r := strings.NewReader("20+20+(3*10)\n")

	l := Lex(r)
	l.Start()
	l.removeWhitespace()

	v, err := l.eval()

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fl, is := IsFloat(v)
	if !is {
		t.Log(err)
		t.Fail()
	}

	i, _ := fl.Int64()
	if i != 40 {
		t.Log("addition failed with value ", v)
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

	fl, is := IsFloat(v)
	if !is {
		t.Log(err)
		t.Fail()
	}

	i, _ := fl.Int64()

	if i != 0 {
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

	fl, is := IsFloat(v)
	if !is {
		t.Log(err)
		t.Fail()
	}

	i, _ := fl.Int64()

	if i != 400 {
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

	fl, is := IsFloat(v)
	if !is {
		t.Log(err)
		t.Fail()
	}

	i, _ := fl.Int64()

	if i != 1 {
		t.Log("division failed")
		t.Fail()
	}
}
