package lexer

// TokenType representa os diferentes tipos de tokens.
type TokenType int

// Token representa um token com seu tipo, lexema, linha e coluna.
type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
	Column int
}

const (
	Illegal TokenType = iota
	EOF
	Identifier
	KeywordModule
	KeywordExport
	KeywordFn
	KeywordLet
	KeywordVar
	KeywordIf
	KeywordElse
	KeywordLoop
	KeywordBreak
	KeywordFor
	KeywordPseudo
	Assign
	Plus
	Minus
	Multiply
	Divide
	LessThan
	GreaterThan
	Equals
	NotEquals
	Semicolon
	Comma
	LeftBrace
	RightBrace
	LeftParen
	RightParen
	Integer
	String
	Colon
)
