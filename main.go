package main

import (
	"fmt"
	"github.com/tysufa/teenycompiler/lexer"
	"github.com/tysufa/teenycompiler/token"
)

func main() {
	input := "IF+-123foo*THEN/"
	l := lexer.New(input)

	tok := l.GetToken()

	for tok.Kind != token.EOF {
		fmt.Printf("%v : %s\n", tok.Kind, tok.Text)
		tok = l.GetToken()
	}
}
