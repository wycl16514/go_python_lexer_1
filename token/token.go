package token 

type TokenType string 

type Token struct {
	Type TokenType  //类型
	Literal  string  //对应字符串
}

//例如数值”1“对应的实例为Token {"INT", "1"}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	IDENTIFIER = "IDENTIFIER"  //变量类型对应的归类
    NUMBER = "NUMBER"  //数值类型对应的归类
	EQUAL = "=" //赋值操作符
	PLUS = "+" //加号操作符
	LPAR = "("
	RPAR = ")"
    LBRACE = "{"
	RBRACE = "}"
    COMMA = ","
    COLON = ":"
	DEF = "def"  //关键字
	INT = "int"
	RETURN = "return"
	ASSERT = "assert"
	AND = "and"
)

var keywords = map[string]TokenType {
	"def" : DEF,
	"int" : INT,
    "return" : RETURN,
	"assert" : ASSERT,
	"and" : AND,
}

//将关键字从变量名中区别开来
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok 
	}

	return IDENTIFIER
}
