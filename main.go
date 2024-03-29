package main

import (
	"fmt"
	"os"

	"github.com/tysufa/teenycompiler/emitter"
	"github.com/tysufa/teenycompiler/lexer"
	"github.com/tysufa/teenycompiler/parser"
)

func main() {
	fmt.Println("teeny tiny compiler")
	args := os.Args[1:]
	if len(args) == 0 {
		panic("veuillez fournir un fichier à compiler")
	}

	dat, _ := os.ReadFile(string(args[0]))

	l := lexer.New(string(dat))
	e := emitter.New("out.c")
	p := parser.New(l, &e)

	p.Program()
	e.WriteFile()
	fmt.Println("Compiling completed")
}
