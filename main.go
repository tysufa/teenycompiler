package main

import (
	"fmt"
	"github.com/tysufa/teenycompiler/lexer"
)

func main() {
	input := "LET foobar = 123"
	l := lexer.New(input)

	for l.PeekChar() != 0 {
		fmt.Printf("%c\n", l.CurChar)
		l.NextChar()
	}
}
