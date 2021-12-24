package lexer 

import(
	"testing"
	"token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},`

	tests := []struct {
		expectedType  token.TokenType 
		expectedLiteral string 
	} {
		{token.EQUAL, "="},
		{token.PLUS, "+"},
		{token.LPAR, "("},
		{token.RPAR, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokenType wrong. expected=%q, got=%q",
		    i, tt.expectedType, tok.Type)

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			    i, tt.expectedLiteral, tok.Literal)
			}
		}
	}
}

func TestNextToken2(t *testing.T) {
	input := `
	def add(x, y):
	    z = x + y
	    return z
	`

	tests := []struct {
		expectedType token.TokenType 
		expectedLiteral string 
	} {
		{token.DEF, "def"},
		{token.IDENTIFIER, "add"},
		{token.LPAR, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAR, ")"},
		{token.COLON, ":"},
		{token.IDENTIFIER, "z"},
		{token.EQUAL, "="},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "z"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "z"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokenType wrong. expected=%q, got=%q",
		    i, tt.expectedType, tok.Type)

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			    i, tt.expectedLiteral, tok.Literal)
			}
		}
	}
}

func TestNextToken3(t *testing.T) {
	input := `123abc`
	tests := []struct {
		expectedType token.TokenType 
		expectedLiteral string 
	} {
		{token.ILLEGAL, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokenType wrong. expected=%q, got=%q",
		    i, tt.expectedType, tok.Type)

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			    i, tt.expectedLiteral, tok.Literal)
			}
		}
	}
}