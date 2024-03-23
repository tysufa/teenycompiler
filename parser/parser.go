package parser

import (
	"github.com/tysufa/teenycompiler/lexer"
	"github.com/tysufa/teenycompiler/token"
	"os"
)

type Parser struct {
	lexer     lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) checkToken(k string) bool {
	return k == p.curToken.Kind
}

func (p *Parser) checkPeek(k string) bool {
	return k == p.peekToken.Kind
}

func (p *Parser) match(k string) {
	if !p.checkToken(k) {
		abort("Expected " + k + ", got " + p.curToken.Kind)
	}
	p.nextToken()
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.GetToken()
}

func New(l lexer.Lexer) Parser {
	p := Parser{lexer: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Program() {
	print("PROGRAM")

	for !p.checkToken(token.EOF) {
		p.statement()
	}
}

func (p *Parser) nl() {
	print("NEWLINE\n")
	p.match(token.NEWLINE)
	for p.checkToken(token.NEWLINE) {
		p.nextToken()
	}
}

func (p *Parser) statement() {
	if p.checkToken(token.PRINT) {
		print("STATEMENT-PRINT\n")
		p.nextToken()

		if p.checkToken(token.STRING) {
			p.nextToken()
		} else {
			p.expression()
		}
	}

	p.nl()
}

func (p *Parser) expression() {

}

func abort(message string) {
	print("error. " + message)
	os.Exit(0)
}
