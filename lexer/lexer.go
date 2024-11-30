package lexer

import (
	"unicode"
)

// Lexer represents the state of the lexer.
type Lexer struct {
	input        string
	position     int  // Current position in the input
	readPosition int  // Next position to read
	ch           byte // Current character under examination
	line         int  // Current line number
	column       int  // Current column number
}

// New creates a new lexer for the given input.
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 0}
	l.readChar() // Initialize first character
	return l
}

// readChar advances the lexer by one character.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // Indicates end of file
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

// NextToken extracts the next token from the input.
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
	for isLetter(l.ch) {
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
	case "const":
		return KeywordConst
	case "let":
		return KeywordLet
	case "var":
		return KeywordVar
	case "fn":
		return KeywordFn
	case "loop":
		return KeywordLoop
	case "break":
		return KeywordBreak
	case "if":
		return KeywordIf
	case "else":
		return KeywordElse
	case "do":
		return KeywordDo
	case "while":
		return KeywordWhile
	case "for":
		return KeywordFor
	case "pseudo":
		return KeywordPseudo
	default:
		return Identifier
	}
}
