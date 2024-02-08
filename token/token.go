package token

//define token data structure

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//Identiers + literals
	IDENT = "IDENT"
	INT = "INT"

	//operators
	ASSIGN = "="
	PLUS = "+"

	//Delimeters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)
var Keywords = map[string] TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func LookupIndent(indent string) TokenType {
	if tok, ok := Keywords[indent]; ok {
		return tok
	}
	return IDENT
}