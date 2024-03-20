package lexer

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

func (l *Lexer) GetToken() {
	if l.CurChar == '+' {

	} else if l.CurChar == '-' {

	} else if l.CurChar == '*' {

	} else if l.CurChar == '/' {

	} else if l.CurChar == '\n' {

	} else if l.CurChar == 0 {

	}

	l.NextChar()
}
