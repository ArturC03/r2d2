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
	And           TokenType = "And"
	Or            TokenType = "Or"
	Not           TokenType = "Not"
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
	Comment       TokenType = "Comment" // // ou /* */
	KeywordModule TokenType = "KeywordModule"
	KeywordExport TokenType = "KeywordExport"
	KeywordFn     TokenType = "KeywordFn"
	KeywordLet    TokenType = "KeywordLet"
	KeywordVar    TokenType = "KeywordVar"
	KeywordIf     TokenType = "KeywordIf"
	KeywordElse   TokenType = "KeywordElse"
	KeywordLoop   TokenType = "KeywordLoop"
	KeywordBreak  TokenType = "KeywordBreak"
	KeywordSend   TokenType = "KeywordSend"
	KeywordFor    TokenType = "KeywordFor"    // for
	KeywordPseudo TokenType = "KeywordPseudo" // pseudo
	KeywordTrue   TokenType = "KeywordTrue"   // true
	KeywordFalse  TokenType = "KeywordFalse"  // false
	KeywordNil    TokenType = "KeywordNil"    // nil
)

// Token representa um token na entrada com tipo, lexema, linha e coluna.
type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
	Column int
}
