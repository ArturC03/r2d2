package lexer

import (
	"fmt"
)

// Lexer representa o estado do lexer.
type Lexer struct {
	input        string
	position     int  // Posição atual na entrada
	readPosition int  // Próxima posição a ser lida
	ch           byte // caracter atual sendo examinado
	line         int  // Número da linha atual
	column       int  // Número da coluna atual
}

// New cria um novo lexer para a entrada fornecida.
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 0}
	l.readChar() // Inicializa o primeiro caracter
	return l
}

// readChar avança o lexer por um caracter.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // Indica o fim do arquivo
	} else {
		l.ch = l.input[l.readPosition]
	}

	// Atualiza a posição e a leitura do próximo caracter
	l.position = l.readPosition
	l.readPosition++

	// Atualiza a linha e coluna corretamente
	if l.ch == '\n' {
		l.line++     // Incrementa a linha
		l.column = 0 // Zera a coluna
	} else {
		l.column++ // Incrementa a coluna para qualquer outro caracter
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
		if l.peekChar() == '/' || l.peekChar() == '*' {
			l.skipComment()
			return l.NextToken()
		}
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
			l.reportError("caracter inválido: '!'")
		}
	case '{':
		tok = l.newToken(LeftBrace, "{")
	case '}':
		tok = l.newToken(RightBrace, "}")
	case '(':
		tok = l.newToken(LeftParen, "(")
	case ')':
		tok = l.newToken(RightParen, ")")
	case '[':
		tok = l.newToken(LeftBracket, "[")
	case ']':
		tok = l.newToken(RightBracket, "]")
	case ';':
		tok = l.newToken(Semicolon, ";")
	case ',':
		tok = l.newToken(Comma, ",")
	case ':':
		tok = l.newToken(Colon, ":")
	case '"':
		stringLiteral := l.readString()
		if stringLiteral == "" {
			l.reportError("string não finalizada")
		}
		tok = l.newToken(String, stringLiteral)
	case 0:
		tok.Type = EOF
		tok.Lexeme = ""
	default:
		if isLetter(l.ch) {
			lexeme := l.readIdentifier()
			tok = l.newToken(lookupKeyword(lexeme), lexeme)
			return tok
		} else if isDigit(l.ch) {
			lexeme := l.readNumber()
			tok = l.newToken(Integer, lexeme)
			return tok
		} else {
			tok = l.newToken(Illegal, string(l.ch))
			l.reportError(fmt.Sprintf("caracter inválido: '%c'", l.ch))
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
	if isLetter(l.ch) {
		l.reportError(fmt.Sprintf("número inválido: '%s%c'", l.input[start:l.position], l.ch))
		for isLetter(l.ch) || isDigit(l.ch) {
			l.readChar()
		}
	}
	return l.input[start:l.position]
}

func (l *Lexer) readString() string {
	start := l.position + 1
	l.readChar() // Move past the opening quote
	for l.ch != '"' && l.ch != 0 {
		if l.ch == '\n' {
			return "" // String não finalizada
		}
		l.readChar()
	}
	str := l.input[start:l.position]
	l.readChar() // Move past the closing quote
	return str
}

func (l *Lexer) skipComment() {
	if l.ch == '/' && l.peekChar() == '/' {
		for l.ch != '\n' && l.ch != 0 {
			l.readChar()
		}
	} else if l.ch == '/' && l.peekChar() == '*' {
		l.readChar()
		l.readChar()
		for !(l.ch == '*' && l.peekChar() == '/') {
			if l.ch == 0 {
				l.reportError("comentário não finalizado")
				return
			}
			l.readChar()
		}
		l.readChar() // Avança sobre '*'
		l.readChar() // Avança sobre '/'
	}
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
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

// reportError exibe um erro léxico com a linha e coluna.
func (l *Lexer) reportError(message string) {
	fmt.Printf("Erro léxico na linha %d, coluna %d: %s\n", l.line, l.column, message)
}
