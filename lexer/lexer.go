package lexer

import (
	"fmt"
	"unicode"
)

// Tipos de token
type TokenType string

const (
	TokenEOF      TokenType = "EOF"      // Fim do ficheiro
	TokenIdent    TokenType = "IDENT"    // Identificador
	TokenNumber   TokenType = "NUMBER"   // Literal numérico
	TokenString   TokenType = "STRING"   // Literal string
	TokenKeyword  TokenType = "KEYWORD"  // Palavra reservada
	TokenSymbol   TokenType = "SYMBOL"   // Símbolo (ex.: {}, (), =)
	TokenColon    TokenType = "COLON"    // ":"
	TokenOperator TokenType = "OPERATOR" // Operadores
	TokenError    TokenType = "ERROR"    // Erro
	TokenComment  TokenType = "COMMENT"  // Comentário
)

// Token representa um elemento lexical
type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
	Column int
}

// Lexer estrutura principal
type Lexer struct {
	input  string
	pos    int
	line   int
	column int
}

// Nova instância do lexer (Construtor)
func NewLexer(input string) *Lexer {
	return &Lexer{input: input, line: 1, column: 0}
}

// Avança para o próximo caracter
func (l *Lexer) next() rune {
	if l.pos >= len(l.input) {
		return 0
	}

	ch := rune(l.input[l.pos])
	l.pos++
	if ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
	return ch
}

// Verifica o próximo caracter sem passar por ele
func (l *Lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return rune(l.input[l.pos])
}

// Retorna o próximo token
func (l *Lexer) NextToken() Token {
	ch := l.next()

	// Ignorar espaços em branco e comentários
	for unicode.IsSpace(ch) {
		ch = l.next()
	}

	if ch == '/' && l.peek() == '/' { // Comentário de linha
		for ch != '\n' && ch != 0 {
			ch = l.next()
		}

		return Token{Type: TokenComment, Lexeme: "//", Line: l.line, Column: l.column}
	}

	// Fim do ficheiro
	if ch == 0 {
		return Token{Type: TokenEOF, Lexeme: "", Line: l.line, Column: l.column}
	}

	// Identificadores e palavras reservadas
	if unicode.IsLetter(ch) || ch == '_' {
		start := l.pos - 1
		for unicode.IsLetter(l.peek()) || unicode.IsDigit(l.peek()) || l.peek() == '_' {
			l.next()
		}
		lexeme := l.input[start:l.pos]
		if isKeyword(lexeme) {
			return Token{Type: TokenKeyword, Lexeme: lexeme, Line: l.line, Column: l.column}
		}
		return Token{Type: TokenIdent, Lexeme: lexeme, Line: l.line, Column: l.column}
	}

	// Literais numéricos
	if unicode.IsDigit(ch) {
		start := l.pos - 1
		for unicode.IsDigit(l.peek()) || l.peek() == '.' {
			l.next()
		}
		lexeme := l.input[start:l.pos]
		return Token{Type: TokenNumber, Lexeme: lexeme, Line: l.line, Column: l.column}
	}

	// Literais string
	if ch == '"' {
		start := l.pos
		for l.peek() != '"' && l.peek() != 0 {
			l.next()
		}
		if l.peek() == '"' {
			l.next()
			return Token{Type: TokenString, Lexeme: l.input[start : l.pos-1], Line: l.line, Column: l.column}
		}
		return Token{Type: TokenError, Lexeme: "String não terminada", Line: l.line, Column: l.column}
	}

	// Dois pontos ":"
	if ch == ':' {
		return Token{Type: TokenColon, Lexeme: ":", Line: l.line, Column: l.column}
	}

	// Operadores e símbolos
	if isOperator(ch) {
		return Token{Type: TokenOperator, Lexeme: string(ch), Line: l.line, Column: l.column}
	}
	if isSymbol(ch) {
		return Token{Type: TokenSymbol, Lexeme: string(ch), Line: l.line, Column: l.column}
	}

	// Erro de caractere inválido
	return Token{Type: TokenError, Lexeme: fmt.Sprintf("Caractere inesperado: %c", ch), Line: l.line, Column: l.column}
}

// Verifica palavras reservadas
func isKeyword(lexeme string) bool {
	keywords := []string{
		"import", "from", "module", "export", "fn", "pseudo",
		"let", "var", "const", "if", "else", "loop", "for", "while", "do", "break",
	}
	for _, kw := range keywords {
		if lexeme == kw {
			return true
		}
	}
	return false
}

// Verifica operadores (+, -, *, /, etc.)
func isOperator(ch rune) bool {
	operators := "+-*/=<>"
	return runeContains(operators, ch)
}

// Verifica símbolos (ex.: {}, (), ;)
func isSymbol(ch rune) bool {
	symbols := "{}[]();,."
	return runeContains(symbols, ch)
}

// Helper: verifica se o caractere está num conjunto
func runeContains(set string, r rune) bool {
	for _, s := range set {
		if r == s {
			return true
		}
	}
	return false
}
