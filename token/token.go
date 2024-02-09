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
	EQ = "=="
	NOT_EQ = "=="

	PLUS = "+"
	MINUS = "_"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"

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
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)
var Keywords = map[string] TokenType {
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIndent(indent string) TokenType {
	if tok, ok := Keywords[indent]; ok {
		return tok
	}
	return IDENT
}