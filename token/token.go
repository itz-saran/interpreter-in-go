package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string // Actual text that represents code
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "{"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookUpIdent(ident string) TokenType {
	fmt.Println("received ident", ident)
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
