package main

import (
	"log"
	"os"
)

func main() {
	file, _ := os.Open("./test.lnn")

	l := Lex(file)
	l.Start()

	l.removeWhitespace()
	for i, t := range l.tokens {
		log.Printf("tok #%v :: '%s' of token type %v on ln %v:%v", i, t.text, t.token, t.span.line, t.span.column)

		f, err := l.eval()
		if err != nil {
			panic(err)
		}

		log.Printf("value: %v", f)
	}
}
