package lexer

// TokenType represents the different types of tokens in the language.
type TokenType int

const (
	// Special tokens
	Illegal TokenType = iota
	EOF

	// Identifiers + literals
	Identifier
	Integer
	String

	// Keywords
	KeywordModule
	KeywordExport
	KeywordConst
	KeywordLet
	KeywordVar
	KeywordFn
	KeywordLoop
	KeywordBreak
	KeywordIf
	KeywordElse
	KeywordDo
	KeywordWhile
	KeywordFor
	KeywordPseudo

	// Operators
	Assign      // =
	Plus        // +
	Minus       // -
	Multiply    // *
	Divide      // /
	LessThan    // <
	GreaterThan // >
	Equals      // ==
	NotEquals   // !=

	// Symbols
	LeftBrace   // {
	RightBrace  // }
	LeftParen   // (
	RightParen  // )
	Semicolon   // ;
	Colon       // :
	Comma       // ,
	DoubleColon // ::
	Quote       // "

	// Other
	Comment
)

// Token represents a single token with its type, lexeme, and position.
type Token struct {
	Type   TokenType // Type of token
	Lexeme string    // Actual text of the token
	Line   int       // Line number where the token is found
	Column int       // Column number for debugging
}

// TokenType.String provides readable names for token types.
func (t TokenType) String() string {
	switch t {
	case Illegal:
		return "Illegal"
	case EOF:
		return "EOF"
	case Identifier:
		return "Identifier"
	case Integer:
		return "Integer"
	case String:
		return "String"
	case KeywordModule:
		return "Module"
	case KeywordExport:
		return "Export"
	case KeywordConst:
		return "Const"
	case KeywordLet:
		return "Let"
	case KeywordVar:
		return "Var"
	case KeywordFn:
		return "Fn"
	case KeywordLoop:
		return "Loop"
	case KeywordBreak:
		return "Break"
	case KeywordIf:
		return "If"
	case KeywordElse:
		return "Else"
	case KeywordDo:
		return "Do"
	case KeywordWhile:
		return "While"
	case KeywordFor:
		return "For"
	case KeywordPseudo:
		return "Pseudo"
	case Assign:
		return "Assign"
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Multiply:
		return "Multiply"
	case Divide:
		return "Divide"
	case LessThan:
		return "LessThan"
	case GreaterThan:
		return "GreaterThan"
	case Equals:
		return "Equals"
	case NotEquals:
		return "NotEquals"
	case LeftBrace:
		return "LeftBrace"
	case RightBrace:
		return "RightBrace"
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case Semicolon:
		return "Semicolon"
	case Colon:
		return "Colon"
	case Comma:
		return "Comma"
	case DoubleColon:
		return "DoubleColon"
	case Quote:
		return "Quote"
	case Comment:
		return "Comment"
	default:
		return "Unknown"
	}
}
