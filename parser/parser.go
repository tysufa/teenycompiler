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
	print("PROGRAM\n")

	for p.checkToken(token.NEWLINE) {
		p.nextToken()
	}

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
	} else if p.checkToken(token.IF) {
		print("STATEMENT-IF\n")
		p.nextToken()
		p.comparison()

		p.match(token.THEN)
		p.nl()

		for !p.checkToken(token.ENDIF) {
			p.statement()
		}

		p.match(token.ENDIF)
	} else if p.checkToken(token.WHILE) {
		print("STATEMENT-WHILE\n")
		p.nextToken()
		p.comparison()

		p.match(token.REPEAT)
		p.nl()

		for !p.checkToken(token.ENDWHILE) {
			p.statement()
		}

		p.match(token.ENDWHILE)
	} else if p.checkToken(token.LABEL) {
		print("STATEMENT-LABEL\n")
		p.nextToken()
		p.match(token.IDENT)
	} else if p.checkToken(token.GOTO) {
		print("STATEMENT-GOTO\n")
		p.nextToken()
		p.match(token.IDENT)
	} else if p.checkToken(token.LET) {
		print("STATEMENT-LET\n")
		p.nextToken()
		p.match(token.IDENT)
		p.match(token.EQ)
		p.expression()
	} else if p.checkToken(token.INPUT) {
		print("STATEMENT-INPUT\n")
		p.nextToken()
		p.match(token.IDENT)
	} else {
		abort("Invalid statement at " + p.curToken.Text + "(" + p.curToken.Kind + ")")
	}

	p.nl()
}

func (p *Parser) isComparisonOperator() bool {
	return p.checkToken(token.GT) || p.checkToken(token.LT) || p.checkToken(token.GTEQ) || p.checkToken(token.LTEQ) || p.checkToken(token.EQEQ) || p.checkToken(token.NOTEQ)
}

func (p *Parser) comparison() {
	print("COMPARISON\n")

	p.expression()
	if p.isComparisonOperator() {
		p.nextToken()
		p.expression()
	} else {
		abort("Expected comparison operator at : " + p.curToken.Text + "\n")
	}

	for p.isComparisonOperator() {
		p.nextToken()
		p.expression()
	}
}

func (p *Parser) term() {
	print("TERM\n")

	p.unary()

	for p.checkToken(token.SLASH) || p.checkToken(token.ASTERISK) {
		p.nextToken()
		p.unary()
	}
}

func (p *Parser) unary() {
	print("UNARY\n")

	if p.checkToken(token.PLUS) || p.checkToken(token.MINUS) {
		p.nextToken()
	}
	p.primary()
}

func (p *Parser) primary() {
	print("PRIMARY : " + p.curToken.Text + "\n")

	if p.checkToken(token.NUMBER) {
		p.nextToken()
	} else if p.checkToken(token.IDENT) {
		p.nextToken()
	} else {
		abort("Unexpected token at " + p.curToken.Text)
	}
}

func (p *Parser) expression() {
	print("EXPRESSION\n")

	p.term()

	for p.checkToken(token.MINUS) || p.checkToken(token.PLUS) {
		p.nextToken()
		p.term()
	}
}

func abort(message string) {
	print("error. " + message)
	os.Exit(0)
}
