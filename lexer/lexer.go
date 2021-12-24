package lexer 
import "token"

type Lexer struct {
	input string  //要解析的源代码字符串
	position int  //当前读取的字符位置
    readPosition int //下一个要读取的字符位置，也就是position + 1
	ch  byte  //读取的字符
}

func New(input string) *Lexer { //生成一个词法解析器
    l := &Lexer{input: input}
	l.readChar() //先读取第一个字符
	return l
}

func (l *Lexer) readChar() { //读取当前字符
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition 
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token{//读取一个字符，判断是否属于特定分类
    var tok token.Token
	//忽略空格，回车，换行等特定字符
	l.skipSpecialChar()

	switch l.ch {
	case '=':
		tok = newToken(token.EQUAL, l.ch)
	case '(':
		tok = newToken(token.LPAR, l.ch)
	case ')':
		tok = newToken(token.RPAR, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch) 
	case 0:  //读取到末尾
		tok.Literal = ""
		tok.Type = token.EOF 
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			//看看变量名是否属于关键字
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.NUMBER
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok 
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

//如果读取到字符，数字或者下划线那么就继续读取下一个字符
func (l *Lexer) readIdentifier() string {
	position := l.position 
	for isLetter(l.ch) {
		l.readChar()
	}

	//获取到变量名字符串
	return l.input[position : l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipSpecialChar() {
	//不读取回车换行，空格等这些特定字符
	for l.ch == ' ' || l.ch == '\t' || l.ch =='\n' || l.ch == '\r' {
		l.readChar()
	}
}


func (l *Lexer) readNumber() string {
	position := l.position 
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position : l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && '9' >= ch
}
