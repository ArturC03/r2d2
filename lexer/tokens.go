package lexer

// TokenType representa os tipos de tokens possíveis.
type TokenType string

const (
	// Tipos de tokens
	Illegal       TokenType = "Illegal"
	EOF           TokenType = "EOF"
	Identifier    TokenType = "Identifier"
	Integer       TokenType = "Integer"
	String        TokenType = "String"
	Assign        TokenType = "Assign"
	Equals        TokenType = "Equals"
	NotEquals     TokenType = "NotEquals"
	Plus          TokenType = "Plus"
	Minus         TokenType = "Minus"
	Multiply      TokenType = "Multiply"
	Divide        TokenType = "Divide"
	LessThan      TokenType = "LessThan"
	GreaterThan   TokenType = "GreaterThan"
	LeftBrace     TokenType = "LeftBrace"
	RightBrace    TokenType = "RightBrace"
	LeftParen     TokenType = "LeftParen"
	RightParen    TokenType = "RightParen"
	LeftBracket   TokenType = "LeftBracket"  // [
	RightBracket  TokenType = "RightBracket" // ]
	Semicolon     TokenType = "Semicolon"
	Comma         TokenType = "Comma"
	Colon         TokenType = "Colon"
	KeywordModule TokenType = "KeywordModule"
	KeywordExport TokenType = "KeywordExport"
	KeywordFn     TokenType = "KeywordFn"
	KeywordLet    TokenType = "KeywordLet"
	KeywordVar    TokenType = "KeywordVar"
	KeywordIf     TokenType = "KeywordIf"
	KeywordElse   TokenType = "KeywordElse"
	KeywordLoop   TokenType = "KeywordLoop"
	KeywordBreak  TokenType = "KeywordBreak"
	KeywordFor    TokenType = "KeywordFor"
	KeywordPseudo TokenType = "KeywordPseudo"
)

// Token representa um token na entrada com tipo, lexema, linha e coluna.
type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
	Column int
}
