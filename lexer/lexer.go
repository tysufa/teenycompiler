package lexer

import (
	"os"

	"github.com/tysufa/teenycompiler/token"
)

type Lexer struct {
	curPos  int
	nextPos int
	input   string
	CurChar byte
}

func New(input string) Lexer {
	l := Lexer{input: input}
	l.NextChar()
	return l
}

func (l *Lexer) NextChar() {
	if l.nextPos >= len(l.input) {
		l.CurChar = 0
	} else {
		l.curPos = l.nextPos
		l.CurChar = l.input[l.curPos]
		l.nextPos++
	}
}

func (l *Lexer) PeekChar() byte {
	if l.nextPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPos]
	}
}

func createToken(text byte, kind string) token.Token {
	return token.Token{Text: string(text), Kind: kind}
}

func abort(message string) {
	print(message)
	os.Exit(0)
}

func (l *Lexer) skipWhitespace() {
	for l.CurChar == ' ' || l.CurChar == '\t' || l.CurChar == '\r' {
		l.NextChar()
	}
}

func (l *Lexer) skipComment() {
	if l.CurChar == '#' {
		for l.CurChar != '\n' {
			l.NextChar()
		}
	}
}

func (l *Lexer) GetToken() token.Token {
	tok := token.Token{}
	l.skipWhitespace()
	l.skipComment()

	if l.CurChar == '+' {
		tok = createToken(l.CurChar, token.PLUS)
	} else if l.CurChar == '-' {
		tok = createToken(l.CurChar, token.MINUS)
	} else if l.CurChar == '*' {
		tok = createToken(l.CurChar, token.ASTERISK)
	} else if l.CurChar == '/' {
		tok = createToken(l.CurChar, token.SLASH)
	} else if l.CurChar == '\n' {
		tok = createToken(l.CurChar, token.NEWLINE)
	} else if l.CurChar == 0 {
		tok = createToken(l.CurChar, token.EOF)
	} else if l.CurChar == '"' {
		l.NextChar()
	} else if l.CurChar == '>' {
		if l.PeekChar() == '=' {
			lastChar := l.CurChar
			l.NextChar()
			tok = token.Token{Text: string(lastChar + l.CurChar), Kind: token.GTEQ}
		} else {
			tok = createToken(l.CurChar, token.GT)
		}
	} else if l.CurChar == '<' {
		if l.PeekChar() == '=' {
			lastChar := l.CurChar
			l.NextChar()
			tok = token.Token{Text: string(lastChar + l.CurChar), Kind: token.LTEQ}
		} else {
			tok = createToken(l.CurChar, token.LT)
		}
	} else if l.CurChar == '!' {
		if l.PeekChar() == '=' {
			lastChar := l.CurChar
			l.NextChar()
			tok = token.Token{Text: string(lastChar + l.CurChar), Kind: token.NOTEQ}
		} else {
			abort("Expeted !=, got '!" + string(l.CurChar) + "'")
		}
	} else if l.CurChar == '=' {
		if l.PeekChar() == '=' {
			lastChar := l.CurChar
			l.NextChar()
			tok = token.Token{Text: string(lastChar + l.CurChar), Kind: token.EQEQ}
		} else {
			tok = createToken(l.CurChar, token.EQ)
		}
	} else {
		abort("Unknown token : '" + string(l.CurChar) + "'")
	}

	l.NextChar()
	return tok
}
