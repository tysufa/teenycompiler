package token

const (
	EOF     = "EOF"
	NEWLINE = "NEWLINE"
	NUMBER  = "NUMBER"
	IDENT   = "IDENT"
	STRING  = "STRING"
	//keywords
	LABEL    = "LABEL"
	GOTO     = "GOTO"
	PRINT    = "PRINT"
	INPUT    = "INPUT"
	LET      = "LET"
	IF       = "IF"
	THEN     = "THEN"
	ENDIF    = "ENDIF"
	WHILE    = "WHILE"
	REPEAT   = "REPEAT"
	ENDWHILE = "ENDWHILE"
	//operators
	EQ       = "EQ"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	ASTERISK = "ASTERISK"
	SLASH    = "SLASH"
	EQEQ     = "EQEQ"
	NOTEQ    = "NOTEQ"
	LT       = "LT"
	LTEQ     = "LTEQ"
	GT       = "GT"
	GTEQ     = "GTEQ"
)

type Token struct {
	Text string
	Kind string
}
