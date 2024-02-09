package lexer

import (
	"testing"

	"github.com/RusimbiPatrick/Eleganta/token"
)


func TestNextToken(t *testing.T){
	input := `let six = 6;
		let seven = 7;

		let add = fn(x,y){
			x + y;
		};
		let result = add(six, seven);
		!-/*6;
		6 < 20 >6;
		
		if(6 < 20){
			return true;
		} else {
			return false;
		}

		12 == 12;
		14 != 7;
		`

	tests :=[]struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "six"},
		{token.ASSIGN, "="},
		{token.INT, "6"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "seven"},
		{token.ASSIGN, "="},
		{token.INT, "7"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "six"},
		{token.COMMA, ","},
		{token.IDENT, "seven"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "6"},
		{token.SEMICOLON, ";"},
		{token.INT, "6"},
		{token.LT, "<"},
		{token.INT, "20"},
		{token.GT, ">"},
		{token.INT, "6"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "6"},
		{token.LT, "<"},
		{token.INT, "20"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "12"},
		{token.EQ, "=="},
		{token.INT, "12"},
		{token.SEMICOLON, ";"},
		{token.INT, "14"},
		{token.NOT_EQ, "!="},
		{token.INT, "7"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l:= New(input)
	//compare raw input with tokenised input
	for i, tt := range tests {
		tok := l.NextToken()
		//Check for token type
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i,tt.expectedType, tok.Type)
		}
		//Check for the token Literal
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}