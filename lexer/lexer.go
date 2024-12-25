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

// Função para retornar os erros de forma amigável ao WASM
func (l *Lexer) wasmReportError(message string) string {
	return fmt.Sprintf("Erro léxico na linha %d, coluna %d: %s", l.line, l.column, message)
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

// nextToken extrai o próximo token da entrada.
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
		} else if l.peekChar() == '!' {
			l.readChar()
			tok = l.newToken(Not, "!")
		} else {
			tok = l.newToken(Not, "!")
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
			l.wasmReportError("string não finalizada")
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
			l.wasmReportError(fmt.Sprintf("caracter inválido: '%c'", l.ch))
		}
	}

	l.readChar()
	return tok
}

// skipWhitespace pula os espaços e as novas linhas.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// peekChar retorna o próximo caracter na sequência.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// newToken cria um novo token com o tipo e o lexema fornecido.
func (l *Lexer) newToken(tokenType TokenType, lexeme string) Token {
	return Token{
		Type:   tokenType,
		Lexeme: lexeme,
		Line:   l.line,
		Column: l.column,
	}
}

// readIdentifier lê um identificador e retorna como string.
func (l *Lexer) readIdentifier() string {
	startPos := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

// readNumber lê um número inteiro e retorna como string.
func (l *Lexer) readNumber() string {
	startPos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

// readString lê uma string delimitada por aspas.
func (l *Lexer) readString() string {
	l.readChar()
	startPos := l.position
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

// skipComment ignora os comentários.
func (l *Lexer) skipComment() {
	if l.ch == '/' {
		for l.ch != '\n' && l.ch != 0 {
			l.readChar()
		}
	} else if l.ch == '*' {
		l.readChar()
		for l.ch != '*' && l.peekChar() != '/' {
			l.readChar()
		}
		l.readChar()
	}
}

// isLetter verifica se um caracter é uma letra.
func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

// isDigit verifica se um caracter é um dígito numérico.
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// lookupKeyword retorna o tipo do token baseado na palavra reservada.
func lookupKeyword(lexeme string) TokenType {
	switch lexeme {
	case "var":
		return KeywordVar
	case "let":
		return KeywordLet
	case "fn":
		return KeywordFn
	case "if":
		return KeywordIf
	case "else":
		return KeywordElse
	case "true":
		return KeywordTrue
	case "false":
		return KeywordFalse
	case "nil":
		return KeywordNil
	}
	return Identifier
}
