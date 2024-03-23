package main

import (
	"fmt"
	"os"

	"github.com/tysufa/teenycompiler/lexer"
	"github.com/tysufa/teenycompiler/parser"
)

func main() {
	fmt.Println("teeny tiny compiler")
	args := os.Args[1:]

	dat, _ := os.ReadFile(string(args[0]))

	l := lexer.New(string(dat))
	p := parser.New(l)

	p.Program()
	fmt.Println("Parsing completed")
}
