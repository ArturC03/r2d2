package lexer

import (
	"unicode"
)

// Lexer representa o estado do lexer.
type Lexer struct {
	input        string
	position     int  // Posição atual na entrada
	readPosition int  // Próxima posição a ser lida
	ch           byte // Caractere atual sendo examinado
	line         int  // Número da linha atual
	column       int  // Número da coluna atual
}

// New cria um novo lexer para a entrada fornecida.
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 0}
	l.readChar() // Inicializa o primeiro caractere
	return l
}

// readChar avança o lexer por um caractere.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // Indica o fim do arquivo
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
}

// NextToken extrai o próximo token da entrada.
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(Equals, "==")
		} else {
			tok = l.newToken(Assign, "=")
		}
	case '+':
		tok = l.newToken(Plus, "+")
	case '-':
		tok = l.newToken(Minus, "-")
	case '*':
		tok = l.newToken(Multiply, "*")
	case '/':
		tok = l.newToken(Divide, "/")
	case '<':
		tok = l.newToken(LessThan, "<")
	case '>':
		tok = l.newToken(GreaterThan, ">")
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = l.newToken(NotEquals, "!=")
		} else {
			tok = l.newToken(Illegal, string(l.ch))
		}
	case '{':
		tok = l.newToken(LeftBrace, "{")
	case '}':
		tok = l.newToken(RightBrace, "}")
	case '(':
		tok = l.newToken(LeftParen, "(")
	case ')':
		tok = l.newToken(RightParen, ")")
	case ';':
		tok = l.newToken(Semicolon, ";")
	case ',':
		tok = l.newToken(Comma, ",")
	case ':':
		tok = l.newToken(Colon, ":")
	case '"':
		tok = l.newToken(String, l.readString())
	case 0:
		tok.Type = EOF
		tok.Lexeme = ""
	default:
		if isLetter(l.ch) {
			lexeme := l.readIdentifier()
			tok.Type = lookupKeyword(lexeme)
			tok.Lexeme = lexeme
			return tok
		} else if isDigit(l.ch) {
			tok.Type = Integer
			tok.Lexeme = l.readNumber()
			return tok
		} else {
			tok = l.newToken(Illegal, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

// Helper methods

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) newToken(tokenType TokenType, lexeme string) Token {
	return Token{
		Type:   tokenType,
		Lexeme: lexeme,
		Line:   l.line,
		Column: l.column,
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readString() string {
	start := l.position + 1
	l.readChar() // Move past the opening quote
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	str := l.input[start:l.position]
	l.readChar() // Move past the closing quote
	return str
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

func isDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}

func lookupKeyword(lexeme string) TokenType {
	switch lexeme {
	case "module":
		return KeywordModule
	case "export":
		return KeywordExport
	case "fn":
		return KeywordFn
	case "let":
		return KeywordLet
	case "var":
		return KeywordVar
	case "if":
		return KeywordIf
	case "else":
		return KeywordElse
	case "loop":
		return KeywordLoop
	case "break":
		return KeywordBreak
	case "for":
		return KeywordFor
	case "pseudo":
		return KeywordPseudo
	default:
		return Identifier
	}
}
