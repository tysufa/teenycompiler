package main

import (
	"fmt"
	"github.com/tysufa/teenycompiler/lexer"
	"github.com/tysufa/teenycompiler/token"
)

func main() {
	input := "+- */ >>= = != #un commentaire\n"
	l := lexer.New(input)

	tok := l.GetToken()

	for tok.Kind != token.EOF {
		fmt.Printf("%v\n", tok.Kind)
		tok = l.GetToken()
	}
}
