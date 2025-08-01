// Code generated from R2D2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type R2D2Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var R2D2LexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func r2d2lexerLexerInit() {
	staticData := &R2D2LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'type'", "'use'", "'import'", "'from'", "'interface'", "'module'",
		"'::'", "'export'", "'fn'", "'pseudo'", "'var'", "'let'", "'const'",
		"'if'", "'else'", "'loop'", "'for'", "'while'", "'break'", "'send'",
		"'continue'", "'return'", "'switch'", "'case'", "'default'", "'=>'",
		"'++'", "'--'", "'+='", "'-='", "'*='", "'/='", "'%='", "'=='", "'!='",
		"'<='", "'>='", "'&&'", "'||'", "'+'", "'-'", "'*'", "'/'", "'%'", "'='",
		"'<'", "'>'", "'!'", "'('", "')'", "'{'", "'}'", "'['", "']'", "','",
		"'.'", "':'", "';'", "'@'", "'js'", "", "", "", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "USE", "IMPORT", "FROM", "INTERFACE", "MODULE", "IMPLEMENTS",
		"EXPORT", "FN", "PSEUDO", "VAR", "LET", "CONST", "IF", "ELSE", "LOOP",
		"FOR", "WHILE", "BREAK", "SEND", "CONTINUE", "RETURN", "SWITCH", "CASE",
		"DEFAULT", "ARROW", "INCREMENT", "DECREMENT", "PLUS_ASSIGN", "MINUS_ASSIGN",
		"MULT_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN", "EQ", "NEQ", "LEQ", "GEQ",
		"AND", "OR", "PLUS", "MINUS", "MULT", "DIV", "MOD", "ASSIGN", "LT",
		"GT", "NOT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
		"COMMA", "DOT", "COLON", "SEMI", "AT", "JS", "TYPE", "STRING_LITERAL",
		"BOOL_LITERAL", "NULL_LITERAL", "INT_LITERAL", "FLOAT_LITERAL", "IDENTIFIER",
		"COMMENT", "BLOCK_COMMENT", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "USE", "IMPORT", "FROM", "INTERFACE", "MODULE", "IMPLEMENTS",
		"EXPORT", "FN", "PSEUDO", "VAR", "LET", "CONST", "IF", "ELSE", "LOOP",
		"FOR", "WHILE", "BREAK", "SEND", "CONTINUE", "RETURN", "SWITCH", "CASE",
		"DEFAULT", "ARROW", "INCREMENT", "DECREMENT", "PLUS_ASSIGN", "MINUS_ASSIGN",
		"MULT_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN", "EQ", "NEQ", "LEQ", "GEQ",
		"AND", "OR", "PLUS", "MINUS", "MULT", "DIV", "MOD", "ASSIGN", "LT",
		"GT", "NOT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK",
		"COMMA", "DOT", "COLON", "SEMI", "AT", "JS", "TYPE", "STRING_LITERAL",
		"ESCAPE_SEQUENCE", "BOOL_LITERAL", "NULL_LITERAL", "INT_LITERAL", "FLOAT_LITERAL",
		"IDENTIFIER", "DecimalIntegerLiteral", "HexIntegerLiteral", "OctalIntegerLiteral",
		"BinaryIntegerLiteral", "DecimalNumeral", "DecimalDigits", "DecimalDigit",
		"NonZeroDigit", "HexDigits", "HexDigit", "OctalDigits", "OctalDigit",
		"BinaryDigits", "BinaryDigit", "ExponentPart", "COMMENT", "BLOCK_COMMENT",
		"WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 70, 599, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 2, 85, 7, 85, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37,
		1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1,
		42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47,
		1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1,
		53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58,
		1, 58, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1,
		60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60,
		1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1,
		60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 3, 60, 435, 8, 60, 1, 61,
		1, 61, 1, 61, 1, 61, 1, 61, 5, 61, 442, 8, 61, 10, 61, 12, 61, 445, 9,
		61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 5, 61, 454, 8, 61,
		10, 61, 12, 61, 457, 9, 61, 1, 61, 3, 61, 460, 8, 61, 1, 62, 1, 62, 1,
		62, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 3, 63,
		474, 8, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1,
		65, 3, 65, 485, 8, 65, 1, 66, 1, 66, 1, 66, 3, 66, 490, 8, 66, 1, 66, 3,
		66, 493, 8, 66, 1, 66, 1, 66, 1, 66, 3, 66, 498, 8, 66, 1, 66, 1, 66, 1,
		66, 3, 66, 503, 8, 66, 1, 67, 1, 67, 5, 67, 507, 8, 67, 10, 67, 12, 67,
		510, 9, 67, 1, 68, 1, 68, 1, 69, 1, 69, 1, 69, 1, 69, 1, 70, 1, 70, 1,
		70, 1, 71, 1, 71, 1, 71, 1, 71, 1, 72, 1, 72, 1, 72, 3, 72, 528, 8, 72,
		3, 72, 530, 8, 72, 1, 73, 4, 73, 533, 8, 73, 11, 73, 12, 73, 534, 1, 74,
		1, 74, 1, 75, 1, 75, 1, 76, 4, 76, 542, 8, 76, 11, 76, 12, 76, 543, 1,
		77, 1, 77, 1, 78, 4, 78, 549, 8, 78, 11, 78, 12, 78, 550, 1, 79, 1, 79,
		1, 80, 4, 80, 556, 8, 80, 11, 80, 12, 80, 557, 1, 81, 1, 81, 1, 82, 1,
		82, 3, 82, 564, 8, 82, 1, 82, 1, 82, 1, 83, 1, 83, 1, 83, 1, 83, 5, 83,
		572, 8, 83, 10, 83, 12, 83, 575, 9, 83, 1, 83, 1, 83, 1, 84, 1, 84, 1,
		84, 1, 84, 5, 84, 583, 8, 84, 10, 84, 12, 84, 586, 9, 84, 1, 84, 1, 84,
		1, 84, 1, 84, 1, 84, 1, 85, 4, 85, 594, 8, 85, 11, 85, 12, 85, 595, 1,
		85, 1, 85, 2, 443, 584, 0, 86, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34,
		69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43,
		87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52,
		105, 53, 107, 54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60,
		121, 61, 123, 62, 125, 0, 127, 63, 129, 64, 131, 65, 133, 66, 135, 67,
		137, 0, 139, 0, 141, 0, 143, 0, 145, 0, 147, 0, 149, 0, 151, 0, 153, 0,
		155, 0, 157, 0, 159, 0, 161, 0, 163, 0, 165, 0, 167, 68, 169, 69, 171,
		70, 1, 0, 15, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 6, 0, 34, 34, 92, 92,
		98, 98, 110, 110, 114, 114, 116, 116, 3, 0, 65, 90, 95, 95, 97, 122, 4,
		0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 88, 88, 120, 120, 2, 0, 66, 66,
		98, 98, 1, 0, 48, 57, 1, 0, 49, 57, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0,
		48, 55, 1, 0, 48, 49, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 2,
		0, 10, 10, 13, 13, 5, 0, 9, 10, 12, 13, 32, 32, 160, 160, 8232, 8233, 611,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1,
		0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85,
		1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0,
		93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0,
		0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1,
		0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0,
		115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0,
		0, 0, 0, 123, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0, 131,
		1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0, 167, 1, 0, 0, 0,
		0, 169, 1, 0, 0, 0, 0, 171, 1, 0, 0, 0, 1, 173, 1, 0, 0, 0, 3, 178, 1,
		0, 0, 0, 5, 182, 1, 0, 0, 0, 7, 189, 1, 0, 0, 0, 9, 194, 1, 0, 0, 0, 11,
		204, 1, 0, 0, 0, 13, 211, 1, 0, 0, 0, 15, 214, 1, 0, 0, 0, 17, 221, 1,
		0, 0, 0, 19, 224, 1, 0, 0, 0, 21, 231, 1, 0, 0, 0, 23, 235, 1, 0, 0, 0,
		25, 239, 1, 0, 0, 0, 27, 245, 1, 0, 0, 0, 29, 248, 1, 0, 0, 0, 31, 253,
		1, 0, 0, 0, 33, 258, 1, 0, 0, 0, 35, 262, 1, 0, 0, 0, 37, 268, 1, 0, 0,
		0, 39, 274, 1, 0, 0, 0, 41, 279, 1, 0, 0, 0, 43, 288, 1, 0, 0, 0, 45, 295,
		1, 0, 0, 0, 47, 302, 1, 0, 0, 0, 49, 307, 1, 0, 0, 0, 51, 315, 1, 0, 0,
		0, 53, 318, 1, 0, 0, 0, 55, 321, 1, 0, 0, 0, 57, 324, 1, 0, 0, 0, 59, 327,
		1, 0, 0, 0, 61, 330, 1, 0, 0, 0, 63, 333, 1, 0, 0, 0, 65, 336, 1, 0, 0,
		0, 67, 339, 1, 0, 0, 0, 69, 342, 1, 0, 0, 0, 71, 345, 1, 0, 0, 0, 73, 348,
		1, 0, 0, 0, 75, 351, 1, 0, 0, 0, 77, 354, 1, 0, 0, 0, 79, 357, 1, 0, 0,
		0, 81, 359, 1, 0, 0, 0, 83, 361, 1, 0, 0, 0, 85, 363, 1, 0, 0, 0, 87, 365,
		1, 0, 0, 0, 89, 367, 1, 0, 0, 0, 91, 369, 1, 0, 0, 0, 93, 371, 1, 0, 0,
		0, 95, 373, 1, 0, 0, 0, 97, 375, 1, 0, 0, 0, 99, 377, 1, 0, 0, 0, 101,
		379, 1, 0, 0, 0, 103, 381, 1, 0, 0, 0, 105, 383, 1, 0, 0, 0, 107, 385,
		1, 0, 0, 0, 109, 387, 1, 0, 0, 0, 111, 389, 1, 0, 0, 0, 113, 391, 1, 0,
		0, 0, 115, 393, 1, 0, 0, 0, 117, 395, 1, 0, 0, 0, 119, 397, 1, 0, 0, 0,
		121, 434, 1, 0, 0, 0, 123, 459, 1, 0, 0, 0, 125, 461, 1, 0, 0, 0, 127,
		473, 1, 0, 0, 0, 129, 475, 1, 0, 0, 0, 131, 484, 1, 0, 0, 0, 133, 502,
		1, 0, 0, 0, 135, 504, 1, 0, 0, 0, 137, 511, 1, 0, 0, 0, 139, 513, 1, 0,
		0, 0, 141, 517, 1, 0, 0, 0, 143, 520, 1, 0, 0, 0, 145, 529, 1, 0, 0, 0,
		147, 532, 1, 0, 0, 0, 149, 536, 1, 0, 0, 0, 151, 538, 1, 0, 0, 0, 153,
		541, 1, 0, 0, 0, 155, 545, 1, 0, 0, 0, 157, 548, 1, 0, 0, 0, 159, 552,
		1, 0, 0, 0, 161, 555, 1, 0, 0, 0, 163, 559, 1, 0, 0, 0, 165, 561, 1, 0,
		0, 0, 167, 567, 1, 0, 0, 0, 169, 578, 1, 0, 0, 0, 171, 593, 1, 0, 0, 0,
		173, 174, 5, 116, 0, 0, 174, 175, 5, 121, 0, 0, 175, 176, 5, 112, 0, 0,
		176, 177, 5, 101, 0, 0, 177, 2, 1, 0, 0, 0, 178, 179, 5, 117, 0, 0, 179,
		180, 5, 115, 0, 0, 180, 181, 5, 101, 0, 0, 181, 4, 1, 0, 0, 0, 182, 183,
		5, 105, 0, 0, 183, 184, 5, 109, 0, 0, 184, 185, 5, 112, 0, 0, 185, 186,
		5, 111, 0, 0, 186, 187, 5, 114, 0, 0, 187, 188, 5, 116, 0, 0, 188, 6, 1,
		0, 0, 0, 189, 190, 5, 102, 0, 0, 190, 191, 5, 114, 0, 0, 191, 192, 5, 111,
		0, 0, 192, 193, 5, 109, 0, 0, 193, 8, 1, 0, 0, 0, 194, 195, 5, 105, 0,
		0, 195, 196, 5, 110, 0, 0, 196, 197, 5, 116, 0, 0, 197, 198, 5, 101, 0,
		0, 198, 199, 5, 114, 0, 0, 199, 200, 5, 102, 0, 0, 200, 201, 5, 97, 0,
		0, 201, 202, 5, 99, 0, 0, 202, 203, 5, 101, 0, 0, 203, 10, 1, 0, 0, 0,
		204, 205, 5, 109, 0, 0, 205, 206, 5, 111, 0, 0, 206, 207, 5, 100, 0, 0,
		207, 208, 5, 117, 0, 0, 208, 209, 5, 108, 0, 0, 209, 210, 5, 101, 0, 0,
		210, 12, 1, 0, 0, 0, 211, 212, 5, 58, 0, 0, 212, 213, 5, 58, 0, 0, 213,
		14, 1, 0, 0, 0, 214, 215, 5, 101, 0, 0, 215, 216, 5, 120, 0, 0, 216, 217,
		5, 112, 0, 0, 217, 218, 5, 111, 0, 0, 218, 219, 5, 114, 0, 0, 219, 220,
		5, 116, 0, 0, 220, 16, 1, 0, 0, 0, 221, 222, 5, 102, 0, 0, 222, 223, 5,
		110, 0, 0, 223, 18, 1, 0, 0, 0, 224, 225, 5, 112, 0, 0, 225, 226, 5, 115,
		0, 0, 226, 227, 5, 101, 0, 0, 227, 228, 5, 117, 0, 0, 228, 229, 5, 100,
		0, 0, 229, 230, 5, 111, 0, 0, 230, 20, 1, 0, 0, 0, 231, 232, 5, 118, 0,
		0, 232, 233, 5, 97, 0, 0, 233, 234, 5, 114, 0, 0, 234, 22, 1, 0, 0, 0,
		235, 236, 5, 108, 0, 0, 236, 237, 5, 101, 0, 0, 237, 238, 5, 116, 0, 0,
		238, 24, 1, 0, 0, 0, 239, 240, 5, 99, 0, 0, 240, 241, 5, 111, 0, 0, 241,
		242, 5, 110, 0, 0, 242, 243, 5, 115, 0, 0, 243, 244, 5, 116, 0, 0, 244,
		26, 1, 0, 0, 0, 245, 246, 5, 105, 0, 0, 246, 247, 5, 102, 0, 0, 247, 28,
		1, 0, 0, 0, 248, 249, 5, 101, 0, 0, 249, 250, 5, 108, 0, 0, 250, 251, 5,
		115, 0, 0, 251, 252, 5, 101, 0, 0, 252, 30, 1, 0, 0, 0, 253, 254, 5, 108,
		0, 0, 254, 255, 5, 111, 0, 0, 255, 256, 5, 111, 0, 0, 256, 257, 5, 112,
		0, 0, 257, 32, 1, 0, 0, 0, 258, 259, 5, 102, 0, 0, 259, 260, 5, 111, 0,
		0, 260, 261, 5, 114, 0, 0, 261, 34, 1, 0, 0, 0, 262, 263, 5, 119, 0, 0,
		263, 264, 5, 104, 0, 0, 264, 265, 5, 105, 0, 0, 265, 266, 5, 108, 0, 0,
		266, 267, 5, 101, 0, 0, 267, 36, 1, 0, 0, 0, 268, 269, 5, 98, 0, 0, 269,
		270, 5, 114, 0, 0, 270, 271, 5, 101, 0, 0, 271, 272, 5, 97, 0, 0, 272,
		273, 5, 107, 0, 0, 273, 38, 1, 0, 0, 0, 274, 275, 5, 115, 0, 0, 275, 276,
		5, 101, 0, 0, 276, 277, 5, 110, 0, 0, 277, 278, 5, 100, 0, 0, 278, 40,
		1, 0, 0, 0, 279, 280, 5, 99, 0, 0, 280, 281, 5, 111, 0, 0, 281, 282, 5,
		110, 0, 0, 282, 283, 5, 116, 0, 0, 283, 284, 5, 105, 0, 0, 284, 285, 5,
		110, 0, 0, 285, 286, 5, 117, 0, 0, 286, 287, 5, 101, 0, 0, 287, 42, 1,
		0, 0, 0, 288, 289, 5, 114, 0, 0, 289, 290, 5, 101, 0, 0, 290, 291, 5, 116,
		0, 0, 291, 292, 5, 117, 0, 0, 292, 293, 5, 114, 0, 0, 293, 294, 5, 110,
		0, 0, 294, 44, 1, 0, 0, 0, 295, 296, 5, 115, 0, 0, 296, 297, 5, 119, 0,
		0, 297, 298, 5, 105, 0, 0, 298, 299, 5, 116, 0, 0, 299, 300, 5, 99, 0,
		0, 300, 301, 5, 104, 0, 0, 301, 46, 1, 0, 0, 0, 302, 303, 5, 99, 0, 0,
		303, 304, 5, 97, 0, 0, 304, 305, 5, 115, 0, 0, 305, 306, 5, 101, 0, 0,
		306, 48, 1, 0, 0, 0, 307, 308, 5, 100, 0, 0, 308, 309, 5, 101, 0, 0, 309,
		310, 5, 102, 0, 0, 310, 311, 5, 97, 0, 0, 311, 312, 5, 117, 0, 0, 312,
		313, 5, 108, 0, 0, 313, 314, 5, 116, 0, 0, 314, 50, 1, 0, 0, 0, 315, 316,
		5, 61, 0, 0, 316, 317, 5, 62, 0, 0, 317, 52, 1, 0, 0, 0, 318, 319, 5, 43,
		0, 0, 319, 320, 5, 43, 0, 0, 320, 54, 1, 0, 0, 0, 321, 322, 5, 45, 0, 0,
		322, 323, 5, 45, 0, 0, 323, 56, 1, 0, 0, 0, 324, 325, 5, 43, 0, 0, 325,
		326, 5, 61, 0, 0, 326, 58, 1, 0, 0, 0, 327, 328, 5, 45, 0, 0, 328, 329,
		5, 61, 0, 0, 329, 60, 1, 0, 0, 0, 330, 331, 5, 42, 0, 0, 331, 332, 5, 61,
		0, 0, 332, 62, 1, 0, 0, 0, 333, 334, 5, 47, 0, 0, 334, 335, 5, 61, 0, 0,
		335, 64, 1, 0, 0, 0, 336, 337, 5, 37, 0, 0, 337, 338, 5, 61, 0, 0, 338,
		66, 1, 0, 0, 0, 339, 340, 5, 61, 0, 0, 340, 341, 5, 61, 0, 0, 341, 68,
		1, 0, 0, 0, 342, 343, 5, 33, 0, 0, 343, 344, 5, 61, 0, 0, 344, 70, 1, 0,
		0, 0, 345, 346, 5, 60, 0, 0, 346, 347, 5, 61, 0, 0, 347, 72, 1, 0, 0, 0,
		348, 349, 5, 62, 0, 0, 349, 350, 5, 61, 0, 0, 350, 74, 1, 0, 0, 0, 351,
		352, 5, 38, 0, 0, 352, 353, 5, 38, 0, 0, 353, 76, 1, 0, 0, 0, 354, 355,
		5, 124, 0, 0, 355, 356, 5, 124, 0, 0, 356, 78, 1, 0, 0, 0, 357, 358, 5,
		43, 0, 0, 358, 80, 1, 0, 0, 0, 359, 360, 5, 45, 0, 0, 360, 82, 1, 0, 0,
		0, 361, 362, 5, 42, 0, 0, 362, 84, 1, 0, 0, 0, 363, 364, 5, 47, 0, 0, 364,
		86, 1, 0, 0, 0, 365, 366, 5, 37, 0, 0, 366, 88, 1, 0, 0, 0, 367, 368, 5,
		61, 0, 0, 368, 90, 1, 0, 0, 0, 369, 370, 5, 60, 0, 0, 370, 92, 1, 0, 0,
		0, 371, 372, 5, 62, 0, 0, 372, 94, 1, 0, 0, 0, 373, 374, 5, 33, 0, 0, 374,
		96, 1, 0, 0, 0, 375, 376, 5, 40, 0, 0, 376, 98, 1, 0, 0, 0, 377, 378, 5,
		41, 0, 0, 378, 100, 1, 0, 0, 0, 379, 380, 5, 123, 0, 0, 380, 102, 1, 0,
		0, 0, 381, 382, 5, 125, 0, 0, 382, 104, 1, 0, 0, 0, 383, 384, 5, 91, 0,
		0, 384, 106, 1, 0, 0, 0, 385, 386, 5, 93, 0, 0, 386, 108, 1, 0, 0, 0, 387,
		388, 5, 44, 0, 0, 388, 110, 1, 0, 0, 0, 389, 390, 5, 46, 0, 0, 390, 112,
		1, 0, 0, 0, 391, 392, 5, 58, 0, 0, 392, 114, 1, 0, 0, 0, 393, 394, 5, 59,
		0, 0, 394, 116, 1, 0, 0, 0, 395, 396, 5, 64, 0, 0, 396, 118, 1, 0, 0, 0,
		397, 398, 5, 106, 0, 0, 398, 399, 5, 115, 0, 0, 399, 120, 1, 0, 0, 0, 400,
		401, 5, 110, 0, 0, 401, 402, 5, 117, 0, 0, 402, 403, 5, 109, 0, 0, 403,
		404, 5, 98, 0, 0, 404, 405, 5, 101, 0, 0, 405, 435, 5, 114, 0, 0, 406,
		407, 5, 98, 0, 0, 407, 408, 5, 111, 0, 0, 408, 409, 5, 111, 0, 0, 409,
		410, 5, 108, 0, 0, 410, 411, 5, 101, 0, 0, 411, 412, 5, 97, 0, 0, 412,
		435, 5, 110, 0, 0, 413, 414, 5, 115, 0, 0, 414, 415, 5, 116, 0, 0, 415,
		416, 5, 114, 0, 0, 416, 417, 5, 105, 0, 0, 417, 418, 5, 110, 0, 0, 418,
		435, 5, 103, 0, 0, 419, 420, 5, 97, 0, 0, 420, 421, 5, 114, 0, 0, 421,
		422, 5, 114, 0, 0, 422, 423, 5, 97, 0, 0, 423, 435, 5, 121, 0, 0, 424,
		425, 5, 111, 0, 0, 425, 426, 5, 98, 0, 0, 426, 427, 5, 106, 0, 0, 427,
		428, 5, 101, 0, 0, 428, 429, 5, 99, 0, 0, 429, 435, 5, 116, 0, 0, 430,
		431, 5, 118, 0, 0, 431, 432, 5, 111, 0, 0, 432, 433, 5, 105, 0, 0, 433,
		435, 5, 100, 0, 0, 434, 400, 1, 0, 0, 0, 434, 406, 1, 0, 0, 0, 434, 413,
		1, 0, 0, 0, 434, 419, 1, 0, 0, 0, 434, 424, 1, 0, 0, 0, 434, 430, 1, 0,
		0, 0, 435, 122, 1, 0, 0, 0, 436, 437, 5, 34, 0, 0, 437, 438, 5, 34, 0,
		0, 438, 439, 5, 34, 0, 0, 439, 443, 1, 0, 0, 0, 440, 442, 9, 0, 0, 0, 441,
		440, 1, 0, 0, 0, 442, 445, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 443, 441,
		1, 0, 0, 0, 444, 446, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446, 447, 5, 34,
		0, 0, 447, 448, 5, 34, 0, 0, 448, 460, 5, 34, 0, 0, 449, 455, 5, 34, 0,
		0, 450, 451, 5, 92, 0, 0, 451, 454, 9, 0, 0, 0, 452, 454, 8, 0, 0, 0, 453,
		450, 1, 0, 0, 0, 453, 452, 1, 0, 0, 0, 454, 457, 1, 0, 0, 0, 455, 453,
		1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 458, 1, 0, 0, 0, 457, 455, 1, 0,
		0, 0, 458, 460, 5, 34, 0, 0, 459, 436, 1, 0, 0, 0, 459, 449, 1, 0, 0, 0,
		460, 124, 1, 0, 0, 0, 461, 462, 5, 92, 0, 0, 462, 463, 7, 1, 0, 0, 463,
		126, 1, 0, 0, 0, 464, 465, 5, 116, 0, 0, 465, 466, 5, 114, 0, 0, 466, 467,
		5, 117, 0, 0, 467, 474, 5, 101, 0, 0, 468, 469, 5, 102, 0, 0, 469, 470,
		5, 97, 0, 0, 470, 471, 5, 108, 0, 0, 471, 472, 5, 115, 0, 0, 472, 474,
		5, 101, 0, 0, 473, 464, 1, 0, 0, 0, 473, 468, 1, 0, 0, 0, 474, 128, 1,
		0, 0, 0, 475, 476, 5, 110, 0, 0, 476, 477, 5, 117, 0, 0, 477, 478, 5, 108,
		0, 0, 478, 479, 5, 108, 0, 0, 479, 130, 1, 0, 0, 0, 480, 485, 3, 137, 68,
		0, 481, 485, 3, 139, 69, 0, 482, 485, 3, 141, 70, 0, 483, 485, 3, 143,
		71, 0, 484, 480, 1, 0, 0, 0, 484, 481, 1, 0, 0, 0, 484, 482, 1, 0, 0, 0,
		484, 483, 1, 0, 0, 0, 485, 132, 1, 0, 0, 0, 486, 487, 3, 145, 72, 0, 487,
		489, 5, 46, 0, 0, 488, 490, 3, 147, 73, 0, 489, 488, 1, 0, 0, 0, 489, 490,
		1, 0, 0, 0, 490, 492, 1, 0, 0, 0, 491, 493, 3, 165, 82, 0, 492, 491, 1,
		0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 503, 1, 0, 0, 0, 494, 495, 5, 46, 0,
		0, 495, 497, 3, 147, 73, 0, 496, 498, 3, 165, 82, 0, 497, 496, 1, 0, 0,
		0, 497, 498, 1, 0, 0, 0, 498, 503, 1, 0, 0, 0, 499, 500, 3, 145, 72, 0,
		500, 501, 3, 165, 82, 0, 501, 503, 1, 0, 0, 0, 502, 486, 1, 0, 0, 0, 502,
		494, 1, 0, 0, 0, 502, 499, 1, 0, 0, 0, 503, 134, 1, 0, 0, 0, 504, 508,
		7, 2, 0, 0, 505, 507, 7, 3, 0, 0, 506, 505, 1, 0, 0, 0, 507, 510, 1, 0,
		0, 0, 508, 506, 1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 136, 1, 0, 0, 0,
		510, 508, 1, 0, 0, 0, 511, 512, 3, 145, 72, 0, 512, 138, 1, 0, 0, 0, 513,
		514, 5, 48, 0, 0, 514, 515, 7, 4, 0, 0, 515, 516, 3, 153, 76, 0, 516, 140,
		1, 0, 0, 0, 517, 518, 5, 48, 0, 0, 518, 519, 3, 157, 78, 0, 519, 142, 1,
		0, 0, 0, 520, 521, 5, 48, 0, 0, 521, 522, 7, 5, 0, 0, 522, 523, 3, 161,
		80, 0, 523, 144, 1, 0, 0, 0, 524, 530, 5, 48, 0, 0, 525, 527, 3, 151, 75,
		0, 526, 528, 3, 147, 73, 0, 527, 526, 1, 0, 0, 0, 527, 528, 1, 0, 0, 0,
		528, 530, 1, 0, 0, 0, 529, 524, 1, 0, 0, 0, 529, 525, 1, 0, 0, 0, 530,
		146, 1, 0, 0, 0, 531, 533, 3, 149, 74, 0, 532, 531, 1, 0, 0, 0, 533, 534,
		1, 0, 0, 0, 534, 532, 1, 0, 0, 0, 534, 535, 1, 0, 0, 0, 535, 148, 1, 0,
		0, 0, 536, 537, 7, 6, 0, 0, 537, 150, 1, 0, 0, 0, 538, 539, 7, 7, 0, 0,
		539, 152, 1, 0, 0, 0, 540, 542, 3, 155, 77, 0, 541, 540, 1, 0, 0, 0, 542,
		543, 1, 0, 0, 0, 543, 541, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544, 154,
		1, 0, 0, 0, 545, 546, 7, 8, 0, 0, 546, 156, 1, 0, 0, 0, 547, 549, 3, 159,
		79, 0, 548, 547, 1, 0, 0, 0, 549, 550, 1, 0, 0, 0, 550, 548, 1, 0, 0, 0,
		550, 551, 1, 0, 0, 0, 551, 158, 1, 0, 0, 0, 552, 553, 7, 9, 0, 0, 553,
		160, 1, 0, 0, 0, 554, 556, 3, 163, 81, 0, 555, 554, 1, 0, 0, 0, 556, 557,
		1, 0, 0, 0, 557, 555, 1, 0, 0, 0, 557, 558, 1, 0, 0, 0, 558, 162, 1, 0,
		0, 0, 559, 560, 7, 10, 0, 0, 560, 164, 1, 0, 0, 0, 561, 563, 7, 11, 0,
		0, 562, 564, 7, 12, 0, 0, 563, 562, 1, 0, 0, 0, 563, 564, 1, 0, 0, 0, 564,
		565, 1, 0, 0, 0, 565, 566, 3, 147, 73, 0, 566, 166, 1, 0, 0, 0, 567, 568,
		5, 47, 0, 0, 568, 569, 5, 47, 0, 0, 569, 573, 1, 0, 0, 0, 570, 572, 8,
		13, 0, 0, 571, 570, 1, 0, 0, 0, 572, 575, 1, 0, 0, 0, 573, 571, 1, 0, 0,
		0, 573, 574, 1, 0, 0, 0, 574, 576, 1, 0, 0, 0, 575, 573, 1, 0, 0, 0, 576,
		577, 6, 83, 0, 0, 577, 168, 1, 0, 0, 0, 578, 579, 5, 47, 0, 0, 579, 580,
		5, 42, 0, 0, 580, 584, 1, 0, 0, 0, 581, 583, 9, 0, 0, 0, 582, 581, 1, 0,
		0, 0, 583, 586, 1, 0, 0, 0, 584, 585, 1, 0, 0, 0, 584, 582, 1, 0, 0, 0,
		585, 587, 1, 0, 0, 0, 586, 584, 1, 0, 0, 0, 587, 588, 5, 42, 0, 0, 588,
		589, 5, 47, 0, 0, 589, 590, 1, 0, 0, 0, 590, 591, 6, 84, 0, 0, 591, 170,
		1, 0, 0, 0, 592, 594, 7, 14, 0, 0, 593, 592, 1, 0, 0, 0, 594, 595, 1, 0,
		0, 0, 595, 593, 1, 0, 0, 0, 595, 596, 1, 0, 0, 0, 596, 597, 1, 0, 0, 0,
		597, 598, 6, 85, 0, 0, 598, 172, 1, 0, 0, 0, 23, 0, 434, 443, 453, 455,
		459, 473, 484, 489, 492, 497, 502, 508, 527, 529, 534, 543, 550, 557, 563,
		573, 584, 595, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// R2D2LexerInit initializes any static state used to implement R2D2Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewR2D2Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func R2D2LexerInit() {
	staticData := &R2D2LexerLexerStaticData
	staticData.once.Do(r2d2lexerLexerInit)
}

// NewR2D2Lexer produces a new lexer instance for the optional input antlr.CharStream.
func NewR2D2Lexer(input antlr.CharStream) *R2D2Lexer {
	R2D2LexerInit()
	l := new(R2D2Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &R2D2LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "R2D2.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// R2D2Lexer tokens.
const (
	R2D2LexerT__0           = 1
	R2D2LexerUSE            = 2
	R2D2LexerIMPORT         = 3
	R2D2LexerFROM           = 4
	R2D2LexerINTERFACE      = 5
	R2D2LexerMODULE         = 6
	R2D2LexerIMPLEMENTS     = 7
	R2D2LexerEXPORT         = 8
	R2D2LexerFN             = 9
	R2D2LexerPSEUDO         = 10
	R2D2LexerVAR            = 11
	R2D2LexerLET            = 12
	R2D2LexerCONST          = 13
	R2D2LexerIF             = 14
	R2D2LexerELSE           = 15
	R2D2LexerLOOP           = 16
	R2D2LexerFOR            = 17
	R2D2LexerWHILE          = 18
	R2D2LexerBREAK          = 19
	R2D2LexerSEND           = 20
	R2D2LexerCONTINUE       = 21
	R2D2LexerRETURN         = 22
	R2D2LexerSWITCH         = 23
	R2D2LexerCASE           = 24
	R2D2LexerDEFAULT        = 25
	R2D2LexerARROW          = 26
	R2D2LexerINCREMENT      = 27
	R2D2LexerDECREMENT      = 28
	R2D2LexerPLUS_ASSIGN    = 29
	R2D2LexerMINUS_ASSIGN   = 30
	R2D2LexerMULT_ASSIGN    = 31
	R2D2LexerDIV_ASSIGN     = 32
	R2D2LexerMOD_ASSIGN     = 33
	R2D2LexerEQ             = 34
	R2D2LexerNEQ            = 35
	R2D2LexerLEQ            = 36
	R2D2LexerGEQ            = 37
	R2D2LexerAND            = 38
	R2D2LexerOR             = 39
	R2D2LexerPLUS           = 40
	R2D2LexerMINUS          = 41
	R2D2LexerMULT           = 42
	R2D2LexerDIV            = 43
	R2D2LexerMOD            = 44
	R2D2LexerASSIGN         = 45
	R2D2LexerLT             = 46
	R2D2LexerGT             = 47
	R2D2LexerNOT            = 48
	R2D2LexerLPAREN         = 49
	R2D2LexerRPAREN         = 50
	R2D2LexerLBRACE         = 51
	R2D2LexerRBRACE         = 52
	R2D2LexerLBRACK         = 53
	R2D2LexerRBRACK         = 54
	R2D2LexerCOMMA          = 55
	R2D2LexerDOT            = 56
	R2D2LexerCOLON          = 57
	R2D2LexerSEMI           = 58
	R2D2LexerAT             = 59
	R2D2LexerJS             = 60
	R2D2LexerTYPE           = 61
	R2D2LexerSTRING_LITERAL = 62
	R2D2LexerBOOL_LITERAL   = 63
	R2D2LexerNULL_LITERAL   = 64
	R2D2LexerINT_LITERAL    = 65
	R2D2LexerFLOAT_LITERAL  = 66
	R2D2LexerIDENTIFIER     = 67
	R2D2LexerCOMMENT        = 68
	R2D2LexerBLOCK_COMMENT  = 69
	R2D2LexerWHITESPACE     = 70
)
