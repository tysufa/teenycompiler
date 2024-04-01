package parser

import (
	"os"

	"github.com/tysufa/teenycompiler/emitter"
	"github.com/tysufa/teenycompiler/lexer"
	"github.com/tysufa/teenycompiler/token"
)

type Parser struct {
	lexer          lexer.Lexer
	emitter        *emitter.Emitter
	curToken       token.Token
	peekToken      token.Token
	symbols        []string
	labelsDeclared []string
	labelsGotoed   []string
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

func New(l lexer.Lexer, e *emitter.Emitter) Parser {
	p := Parser{lexer: l, emitter: e}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Program() {
	p.emitter.HeaderLine("#include <stdio.h>")
	p.emitter.HeaderLine("int main (void){")

	for p.checkToken(token.NEWLINE) {
		p.nextToken()
	}

	for !p.checkToken(token.EOF) {
		p.statement()
	}

	p.emitter.EmitLine("return 0;")
	p.emitter.EmitLine("}")

	labelInDeclared := false
	for _, label := range p.labelsGotoed {
		labelInDeclared = false
		for _, labelDeclared := range p.labelsDeclared {
			if labelDeclared == label {
				labelInDeclared = true
			}
		}
		if !labelInDeclared {
			panic("Attempting to GOTO an undeclared label : " + label)
		}
	}
}

func (p *Parser) nl() {
	p.match(token.NEWLINE)
	for p.checkToken(token.NEWLINE) {
		p.nextToken()
	}
}

func (p *Parser) statement() {
	if p.checkToken(token.PRINT) {
		p.nextToken()

		if p.checkToken(token.STRING) {
			p.emitter.EmitLine("printf(\"" + p.curToken.Text + "\\n\");")
			p.nextToken()
		} else {
			p.emitter.Emit("printf(\"%" + ".2f\\n\", (float)(")
			p.expression()
			p.emitter.EmitLine("));")
		}
	} else if p.checkToken(token.IF) {
		p.nextToken()
		p.emitter.Emit("if(")
		p.comparison()

		p.match(token.THEN)
		p.nl()
		p.emitter.EmitLine("){")

		for !p.checkToken(token.ENDIF) {
			p.statement()
		}

		p.match(token.ENDIF)
		p.emitter.EmitLine("}")
	} else if p.checkToken(token.WHILE) {
		p.emitter.Emit("while(")
		p.nextToken()
		p.comparison()
		p.emitter.EmitLine("){")

		p.match(token.REPEAT)
		p.nl()

		for !p.checkToken(token.ENDWHILE) {
			p.statement()
		}

		p.match(token.ENDWHILE)
		p.emitter.EmitLine("}")
	} else if p.checkToken(token.LABEL) {
		p.nextToken()
		for _, el := range p.labelsDeclared {
			if el == p.curToken.Text {
				abort("Label already exists : " + p.curToken.Text)
			}
		}
		p.labelsDeclared = append(p.labelsDeclared, p.curToken.Text)

		p.emitter.EmitLine(p.curToken.Text + ":")
		p.match(token.IDENT)
	} else if p.checkToken(token.GOTO) {
		p.nextToken()
		p.labelsGotoed = append(p.labelsGotoed, p.curToken.Text)
		p.emitter.EmitLine("goto " + p.curToken.Text + ";")
		p.match(token.IDENT)
	} else if p.checkToken(token.LET) {
		p.nextToken()

		symbolsExist := false
		for _, symbol := range p.symbols {
			if p.curToken.Text == symbol {
				symbolsExist = true
			}
		}
		if !symbolsExist {
			p.symbols = append(p.symbols, p.curToken.Text)
			p.emitter.HeaderLine("float " + p.curToken.Text + ";")
		}

		p.emitter.Emit(p.curToken.Text + " = ")
		p.match(token.IDENT)
		p.match(token.EQ)
		p.expression()
		p.emitter.EmitLine(";")
	} else if p.checkToken(token.INPUT) {
		p.nextToken()

		symbolsExist := false
		for _, symbol := range p.symbols {
			if p.curToken.Text == symbol {
				symbolsExist = true
			}
		}
		if !symbolsExist {
			p.symbols = append(p.symbols, p.curToken.Text)
			p.emitter.HeaderLine("float " + p.curToken.Text + ";")
		}

		p.emitter.EmitLine("if(0==scanf(\"%" + "f\", &" + p.curToken.Text + ")) {")
		p.emitter.EmitLine(p.curToken.Text + " = 0;")
		p.emitter.Emit("scanf(\"%")
		p.emitter.EmitLine("*s\");")
		p.emitter.EmitLine("}")
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

	p.expression()
	if p.isComparisonOperator() {
		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
		p.expression()
	} else {
		abort("Expected comparison operator at : " + p.curToken.Text + "\n")
	}

	for p.isComparisonOperator() {
		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
		p.expression()
	}
}

func (p *Parser) term() {

	p.unary()

	for p.checkToken(token.SLASH) || p.checkToken(token.ASTERISK) {
		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
		p.unary()
	}
}

func (p *Parser) unary() {

	if p.checkToken(token.PLUS) || p.checkToken(token.MINUS) {
		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
	}
	p.primary()
}

func (p *Parser) primary() {

	if p.checkToken(token.NUMBER) {
		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
	} else if p.checkToken(token.IDENT) {
		symbolsExist := false
		for _, symbol := range p.symbols {
			if p.curToken.Text == symbol {
				symbolsExist = true
			}
		}
		if !symbolsExist {
			abort("Referencing variable before assignement : " + p.curToken.Text)
		}

		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
	} else {
		abort("Unexpected token at " + p.curToken.Text)
	}
}

func (p *Parser) expression() {

	p.term()

	for p.checkToken(token.MINUS) || p.checkToken(token.PLUS) {
		p.emitter.Emit(p.curToken.Text)
		p.nextToken()
		p.term()
	}
}

func abort(message string) {
	print("error. " + message)
	os.Exit(0)
}
