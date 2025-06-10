// Code generated from R2D2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // R2D2

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type R2D2Parser struct {
	*antlr.BaseParser
}

var R2D2ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func r2d2ParserInit() {
	staticData := &R2D2ParserStaticData
	staticData.LiteralNames = []string{
		"", "'type'", "'use'", "'import'", "'from'", "'interface'", "'module'",
		"'::'", "'export'", "'fn'", "'pseudo'", "'var'", "'let'", "'const'",
		"'if'", "'else'", "'loop'", "'for'", "'while'", "'break'", "'send'",
		"'continue'", "'return'", "'switch'", "'case'", "'default'", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'++'", "'--'", "'='", "'+='", "'-='", "'*='",
		"'/='", "'%='", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'",
		"'||'", "'!'", "'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'.'",
		"':'", "';'", "'@'", "'js'", "'=>'", "", "", "", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "USE", "IMPORT", "FROM", "INTERFACE", "MODULE", "IMPLEMENTS",
		"EXPORT", "FN", "PSEUDO", "VAR", "LET", "CONST", "IF", "ELSE", "LOOP",
		"FOR", "WHILE", "BREAK", "SEND", "CONTINUE", "RETURN", "SWITCH", "CASE",
		"DEFAULT", "PLUS", "MINUS", "MULT", "DIV", "MOD", "INCREMENT", "DECREMENT",
		"ASSIGN", "PLUS_ASSIGN", "MINUS_ASSIGN", "MULT_ASSIGN", "DIV_ASSIGN",
		"MOD_ASSIGN", "EQ", "NEQ", "LT", "GT", "LEQ", "GEQ", "AND", "OR", "NOT",
		"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "COMMA",
		"DOT", "COLON", "SEMI", "AT", "JS", "ARROW", "TYPE", "STRING_LITERAL",
		"BOOL_LITERAL", "NULL_LITERAL", "INT_LITERAL", "IDENTIFIER", "FLOAT_LITERAL",
		"COMMENT", "BLOCK_COMMENT", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"program", "declaration", "globalDeclaration", "importDeclaration",
		"interfaceDeclaration", "moduleDeclaration", "functionDeclaration",
		"functionCallStatement", "functionCall", "parameterList", "parameter",
		"typeExpression", "arrayDimensions", "baseType", "genericType", "typeDeclaration",
		"variableDeclaration", "statement", "expressionStatement", "ifStatement",
		"forStatement", "assignmentDeclaration", "assignment", "assignmentOperator",
		"simpleFor", "whileStatement", "loopStatement", "cicleControl", "breakStatement",
		"continueStatement", "returnStatement", "expression", "memberExpression",
		"memberPart", "argumentList", "primaryExpression", "arrayLiteral", "literal",
		"block", "switchStatement", "switchCase", "defaultCase", "jsStatement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 70, 552, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 5, 0, 88, 8, 0, 10, 0, 12, 0, 91, 9, 0, 1, 0, 5, 0, 94,
		8, 0, 10, 0, 12, 0, 97, 9, 0, 1, 0, 5, 0, 100, 8, 0, 10, 0, 12, 0, 103,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 110, 8, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 127, 8, 4, 1, 4, 1, 4, 1, 4, 5, 4, 132, 8, 4, 10, 4, 12, 4, 135,
		9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 143, 8, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 5, 5, 149, 8, 5, 10, 5, 12, 5, 152, 9, 5, 1, 5, 1, 5, 1, 6,
		3, 6, 157, 8, 6, 1, 6, 3, 6, 160, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 166,
		8, 6, 1, 6, 1, 6, 3, 6, 170, 8, 6, 1, 6, 1, 6, 3, 6, 174, 8, 6, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 5, 8, 181, 8, 8, 10, 8, 12, 8, 184, 9, 8, 1, 8, 1,
		8, 1, 8, 3, 8, 189, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 196, 8, 9,
		10, 9, 12, 9, 199, 9, 9, 1, 10, 1, 10, 3, 10, 203, 8, 10, 1, 11, 1, 11,
		3, 11, 207, 8, 11, 1, 11, 3, 11, 210, 8, 11, 1, 11, 3, 11, 213, 8, 11,
		1, 12, 1, 12, 3, 12, 217, 8, 12, 1, 12, 4, 12, 220, 8, 12, 11, 12, 12,
		12, 221, 1, 13, 1, 13, 1, 13, 3, 13, 227, 8, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 5, 14, 234, 8, 14, 10, 14, 12, 14, 237, 9, 14, 1, 14, 1, 14,
		1, 15, 3, 15, 242, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 248, 8, 15,
		10, 15, 12, 15, 251, 9, 15, 1, 15, 1, 15, 1, 16, 3, 16, 256, 8, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 261, 8, 16, 1, 16, 1, 16, 3, 16, 265, 8, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 3, 17, 281, 8, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 3, 19, 288, 8, 19, 1, 19, 1, 19, 3, 19, 292, 8, 19, 1, 19, 1, 19, 1,
		19, 3, 19, 297, 8, 19, 1, 19, 1, 19, 1, 19, 3, 19, 302, 8, 19, 1, 19, 1,
		19, 3, 19, 306, 8, 19, 1, 19, 1, 19, 1, 19, 3, 19, 311, 8, 19, 5, 19, 313,
		8, 19, 10, 19, 12, 19, 316, 9, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 322,
		8, 19, 3, 19, 324, 8, 19, 1, 20, 1, 20, 3, 20, 328, 8, 20, 1, 20, 1, 20,
		3, 20, 332, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 358, 8, 22, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 366, 8, 24, 1, 24, 1, 24, 1, 24,
		3, 24, 371, 8, 24, 1, 24, 3, 24, 374, 8, 24, 1, 25, 1, 25, 3, 25, 378,
		8, 25, 1, 25, 1, 25, 3, 25, 382, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 3, 27, 391, 8, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 29, 1, 30, 1, 30, 3, 30, 401, 8, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 415, 8, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 434, 8, 31, 10, 31,
		12, 31, 437, 9, 31, 1, 32, 1, 32, 5, 32, 441, 8, 32, 10, 32, 12, 32, 444,
		9, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 3, 33, 456, 8, 33, 1, 33, 3, 33, 459, 8, 33, 1, 34, 1, 34, 1, 34, 5,
		34, 464, 8, 34, 10, 34, 12, 34, 467, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 477, 8, 35, 1, 36, 1, 36, 1, 36, 1,
		36, 5, 36, 483, 8, 36, 10, 36, 12, 36, 486, 9, 36, 3, 36, 488, 8, 36, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 498, 8, 37,
		1, 38, 1, 38, 5, 38, 502, 8, 38, 10, 38, 12, 38, 505, 9, 38, 1, 38, 1,
		38, 1, 39, 1, 39, 3, 39, 511, 8, 39, 1, 39, 1, 39, 3, 39, 515, 8, 39, 1,
		39, 1, 39, 5, 39, 519, 8, 39, 10, 39, 12, 39, 522, 9, 39, 1, 39, 3, 39,
		525, 8, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 3, 40, 538, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		3, 41, 545, 8, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 0, 1, 62,
		43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 80, 82, 84, 0, 8, 1, 0, 11, 13, 1, 0, 31, 32, 1, 0, 33,
		38, 3, 0, 27, 27, 31, 32, 47, 47, 1, 0, 28, 30, 1, 0, 26, 27, 1, 0, 39,
		44, 1, 0, 45, 46, 609, 0, 89, 1, 0, 0, 0, 2, 109, 1, 0, 0, 0, 4, 111, 1,
		0, 0, 0, 6, 118, 1, 0, 0, 0, 8, 122, 1, 0, 0, 0, 10, 138, 1, 0, 0, 0, 12,
		156, 1, 0, 0, 0, 14, 175, 1, 0, 0, 0, 16, 182, 1, 0, 0, 0, 18, 192, 1,
		0, 0, 0, 20, 200, 1, 0, 0, 0, 22, 212, 1, 0, 0, 0, 24, 219, 1, 0, 0, 0,
		26, 226, 1, 0, 0, 0, 28, 228, 1, 0, 0, 0, 30, 241, 1, 0, 0, 0, 32, 255,
		1, 0, 0, 0, 34, 280, 1, 0, 0, 0, 36, 282, 1, 0, 0, 0, 38, 285, 1, 0, 0,
		0, 40, 325, 1, 0, 0, 0, 42, 335, 1, 0, 0, 0, 44, 357, 1, 0, 0, 0, 46, 359,
		1, 0, 0, 0, 48, 365, 1, 0, 0, 0, 50, 375, 1, 0, 0, 0, 52, 385, 1, 0, 0,
		0, 54, 390, 1, 0, 0, 0, 56, 392, 1, 0, 0, 0, 58, 395, 1, 0, 0, 0, 60, 398,
		1, 0, 0, 0, 62, 414, 1, 0, 0, 0, 64, 438, 1, 0, 0, 0, 66, 458, 1, 0, 0,
		0, 68, 460, 1, 0, 0, 0, 70, 476, 1, 0, 0, 0, 72, 478, 1, 0, 0, 0, 74, 497,
		1, 0, 0, 0, 76, 499, 1, 0, 0, 0, 78, 508, 1, 0, 0, 0, 80, 537, 1, 0, 0,
		0, 82, 544, 1, 0, 0, 0, 84, 546, 1, 0, 0, 0, 86, 88, 3, 6, 3, 0, 87, 86,
		1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0,
		90, 95, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 94, 3, 8, 4, 0, 93, 92, 1,
		0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96,
		101, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 100, 3, 2, 1, 0, 99, 98, 1, 0,
		0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0,
		102, 104, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 105, 5, 0, 0, 1, 105,
		1, 1, 0, 0, 0, 106, 110, 3, 10, 5, 0, 107, 110, 3, 4, 2, 0, 108, 110, 3,
		30, 15, 0, 109, 106, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 108, 1, 0,
		0, 0, 110, 3, 1, 0, 0, 0, 111, 112, 5, 13, 0, 0, 112, 113, 5, 66, 0, 0,
		113, 114, 3, 22, 11, 0, 114, 115, 5, 33, 0, 0, 115, 116, 3, 62, 31, 0,
		116, 117, 5, 57, 0, 0, 117, 5, 1, 0, 0, 0, 118, 119, 5, 2, 0, 0, 119, 120,
		5, 62, 0, 0, 120, 121, 5, 57, 0, 0, 121, 7, 1, 0, 0, 0, 122, 123, 5, 5,
		0, 0, 123, 126, 5, 66, 0, 0, 124, 125, 5, 7, 0, 0, 125, 127, 5, 66, 0,
		0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		133, 5, 50, 0, 0, 129, 132, 3, 12, 6, 0, 130, 132, 3, 32, 16, 0, 131, 129,
		1, 0, 0, 0, 131, 130, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0,
		0, 0, 133, 134, 1, 0, 0, 0, 134, 136, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0,
		136, 137, 5, 51, 0, 0, 137, 9, 1, 0, 0, 0, 138, 139, 5, 6, 0, 0, 139, 142,
		5, 66, 0, 0, 140, 141, 5, 7, 0, 0, 141, 143, 5, 66, 0, 0, 142, 140, 1,
		0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 150, 5, 50, 0,
		0, 145, 149, 3, 12, 6, 0, 146, 149, 3, 30, 15, 0, 147, 149, 3, 32, 16,
		0, 148, 145, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 147, 1, 0, 0, 0, 149,
		152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 153,
		1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 154, 5, 51, 0, 0, 154, 11, 1, 0,
		0, 0, 155, 157, 5, 8, 0, 0, 156, 155, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0,
		157, 159, 1, 0, 0, 0, 158, 160, 5, 10, 0, 0, 159, 158, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 5, 9, 0, 0, 162, 163,
		5, 66, 0, 0, 163, 165, 5, 48, 0, 0, 164, 166, 3, 18, 9, 0, 165, 164, 1,
		0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 169, 5, 49, 0,
		0, 168, 170, 3, 22, 11, 0, 169, 168, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0,
		170, 173, 1, 0, 0, 0, 171, 174, 3, 76, 38, 0, 172, 174, 5, 57, 0, 0, 173,
		171, 1, 0, 0, 0, 173, 172, 1, 0, 0, 0, 174, 13, 1, 0, 0, 0, 175, 176, 3,
		16, 8, 0, 176, 177, 5, 57, 0, 0, 177, 15, 1, 0, 0, 0, 178, 179, 5, 66,
		0, 0, 179, 181, 5, 55, 0, 0, 180, 178, 1, 0, 0, 0, 181, 184, 1, 0, 0, 0,
		182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 185, 1, 0, 0, 0, 184,
		182, 1, 0, 0, 0, 185, 186, 5, 66, 0, 0, 186, 188, 5, 48, 0, 0, 187, 189,
		3, 68, 34, 0, 188, 187, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 190, 1,
		0, 0, 0, 190, 191, 5, 49, 0, 0, 191, 17, 1, 0, 0, 0, 192, 197, 3, 20, 10,
		0, 193, 194, 5, 54, 0, 0, 194, 196, 3, 20, 10, 0, 195, 193, 1, 0, 0, 0,
		196, 199, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198,
		19, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 202, 5, 66, 0, 0, 201, 203,
		3, 22, 11, 0, 202, 201, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 21, 1, 0,
		0, 0, 204, 206, 3, 26, 13, 0, 205, 207, 3, 24, 12, 0, 206, 205, 1, 0, 0,
		0, 206, 207, 1, 0, 0, 0, 207, 213, 1, 0, 0, 0, 208, 210, 3, 24, 12, 0,
		209, 208, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211,
		213, 3, 26, 13, 0, 212, 204, 1, 0, 0, 0, 212, 209, 1, 0, 0, 0, 213, 23,
		1, 0, 0, 0, 214, 216, 5, 52, 0, 0, 215, 217, 5, 65, 0, 0, 216, 215, 1,
		0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 220, 5, 53, 0,
		0, 219, 214, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221,
		222, 1, 0, 0, 0, 222, 25, 1, 0, 0, 0, 223, 227, 5, 66, 0, 0, 224, 227,
		5, 61, 0, 0, 225, 227, 3, 28, 14, 0, 226, 223, 1, 0, 0, 0, 226, 224, 1,
		0, 0, 0, 226, 225, 1, 0, 0, 0, 227, 27, 1, 0, 0, 0, 228, 229, 5, 66, 0,
		0, 229, 230, 5, 41, 0, 0, 230, 235, 3, 22, 11, 0, 231, 232, 5, 54, 0, 0,
		232, 234, 3, 22, 11, 0, 233, 231, 1, 0, 0, 0, 234, 237, 1, 0, 0, 0, 235,
		233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 238, 1, 0, 0, 0, 237, 235,
		1, 0, 0, 0, 238, 239, 5, 42, 0, 0, 239, 29, 1, 0, 0, 0, 240, 242, 5, 8,
		0, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0,
		243, 244, 5, 1, 0, 0, 244, 245, 5, 66, 0, 0, 245, 249, 5, 50, 0, 0, 246,
		248, 3, 32, 16, 0, 247, 246, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249, 247,
		1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 252, 1, 0, 0, 0, 251, 249, 1, 0,
		0, 0, 252, 253, 5, 51, 0, 0, 253, 31, 1, 0, 0, 0, 254, 256, 5, 8, 0, 0,
		255, 254, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257,
		258, 7, 0, 0, 0, 258, 260, 5, 66, 0, 0, 259, 261, 3, 22, 11, 0, 260, 259,
		1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 264, 1, 0, 0, 0, 262, 263, 5, 33,
		0, 0, 263, 265, 3, 62, 31, 0, 264, 262, 1, 0, 0, 0, 264, 265, 1, 0, 0,
		0, 265, 266, 1, 0, 0, 0, 266, 267, 5, 57, 0, 0, 267, 33, 1, 0, 0, 0, 268,
		281, 3, 32, 16, 0, 269, 281, 3, 14, 7, 0, 270, 281, 3, 36, 18, 0, 271,
		281, 3, 38, 19, 0, 272, 281, 3, 40, 20, 0, 273, 281, 3, 50, 25, 0, 274,
		281, 3, 52, 26, 0, 275, 281, 3, 54, 27, 0, 276, 281, 3, 60, 30, 0, 277,
		281, 3, 78, 39, 0, 278, 281, 3, 42, 21, 0, 279, 281, 3, 84, 42, 0, 280,
		268, 1, 0, 0, 0, 280, 269, 1, 0, 0, 0, 280, 270, 1, 0, 0, 0, 280, 271,
		1, 0, 0, 0, 280, 272, 1, 0, 0, 0, 280, 273, 1, 0, 0, 0, 280, 274, 1, 0,
		0, 0, 280, 275, 1, 0, 0, 0, 280, 276, 1, 0, 0, 0, 280, 277, 1, 0, 0, 0,
		280, 278, 1, 0, 0, 0, 280, 279, 1, 0, 0, 0, 281, 35, 1, 0, 0, 0, 282, 283,
		3, 62, 31, 0, 283, 284, 5, 57, 0, 0, 284, 37, 1, 0, 0, 0, 285, 287, 5,
		14, 0, 0, 286, 288, 5, 48, 0, 0, 287, 286, 1, 0, 0, 0, 287, 288, 1, 0,
		0, 0, 288, 289, 1, 0, 0, 0, 289, 291, 3, 62, 31, 0, 290, 292, 5, 49, 0,
		0, 291, 290, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 296, 1, 0, 0, 0, 293,
		297, 3, 76, 38, 0, 294, 295, 5, 60, 0, 0, 295, 297, 3, 34, 17, 0, 296,
		293, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 297, 314, 1, 0, 0, 0, 298, 299,
		5, 15, 0, 0, 299, 301, 5, 14, 0, 0, 300, 302, 5, 48, 0, 0, 301, 300, 1,
		0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 305, 3, 62, 31,
		0, 304, 306, 5, 49, 0, 0, 305, 304, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306,
		310, 1, 0, 0, 0, 307, 311, 3, 76, 38, 0, 308, 309, 5, 60, 0, 0, 309, 311,
		3, 34, 17, 0, 310, 307, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 311, 313, 1,
		0, 0, 0, 312, 298, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0,
		0, 314, 315, 1, 0, 0, 0, 315, 323, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317,
		321, 5, 15, 0, 0, 318, 322, 3, 76, 38, 0, 319, 320, 5, 60, 0, 0, 320, 322,
		3, 34, 17, 0, 321, 318, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 322, 324, 1,
		0, 0, 0, 323, 317, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 39, 1, 0, 0,
		0, 325, 327, 5, 17, 0, 0, 326, 328, 5, 48, 0, 0, 327, 326, 1, 0, 0, 0,
		327, 328, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 331, 3, 48, 24, 0, 330,
		332, 5, 49, 0, 0, 331, 330, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 333,
		1, 0, 0, 0, 333, 334, 3, 76, 38, 0, 334, 41, 1, 0, 0, 0, 335, 336, 3, 44,
		22, 0, 336, 337, 5, 57, 0, 0, 337, 43, 1, 0, 0, 0, 338, 339, 5, 66, 0,
		0, 339, 340, 3, 46, 23, 0, 340, 341, 3, 62, 31, 0, 341, 358, 1, 0, 0, 0,
		342, 343, 5, 66, 0, 0, 343, 358, 7, 1, 0, 0, 344, 345, 5, 66, 0, 0, 345,
		346, 5, 52, 0, 0, 346, 347, 3, 62, 31, 0, 347, 348, 5, 53, 0, 0, 348, 349,
		3, 46, 23, 0, 349, 350, 3, 62, 31, 0, 350, 358, 1, 0, 0, 0, 351, 352, 5,
		66, 0, 0, 352, 353, 5, 52, 0, 0, 353, 354, 3, 62, 31, 0, 354, 355, 5, 53,
		0, 0, 355, 356, 7, 1, 0, 0, 356, 358, 1, 0, 0, 0, 357, 338, 1, 0, 0, 0,
		357, 342, 1, 0, 0, 0, 357, 344, 1, 0, 0, 0, 357, 351, 1, 0, 0, 0, 358,
		45, 1, 0, 0, 0, 359, 360, 7, 2, 0, 0, 360, 47, 1, 0, 0, 0, 361, 366, 3,
		32, 16, 0, 362, 363, 3, 44, 22, 0, 363, 364, 5, 57, 0, 0, 364, 366, 1,
		0, 0, 0, 365, 361, 1, 0, 0, 0, 365, 362, 1, 0, 0, 0, 365, 366, 1, 0, 0,
		0, 366, 370, 1, 0, 0, 0, 367, 368, 3, 62, 31, 0, 368, 369, 5, 57, 0, 0,
		369, 371, 1, 0, 0, 0, 370, 367, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371,
		373, 1, 0, 0, 0, 372, 374, 3, 44, 22, 0, 373, 372, 1, 0, 0, 0, 373, 374,
		1, 0, 0, 0, 374, 49, 1, 0, 0, 0, 375, 377, 5, 18, 0, 0, 376, 378, 5, 48,
		0, 0, 377, 376, 1, 0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0,
		379, 381, 3, 62, 31, 0, 380, 382, 5, 49, 0, 0, 381, 380, 1, 0, 0, 0, 381,
		382, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 384, 3, 76, 38, 0, 384, 51,
		1, 0, 0, 0, 385, 386, 5, 16, 0, 0, 386, 387, 3, 76, 38, 0, 387, 53, 1,
		0, 0, 0, 388, 391, 3, 56, 28, 0, 389, 391, 3, 58, 29, 0, 390, 388, 1, 0,
		0, 0, 390, 389, 1, 0, 0, 0, 391, 55, 1, 0, 0, 0, 392, 393, 5, 19, 0, 0,
		393, 394, 5, 57, 0, 0, 394, 57, 1, 0, 0, 0, 395, 396, 5, 21, 0, 0, 396,
		397, 5, 57, 0, 0, 397, 59, 1, 0, 0, 0, 398, 400, 5, 22, 0, 0, 399, 401,
		3, 62, 31, 0, 400, 399, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 402, 1,
		0, 0, 0, 402, 403, 5, 57, 0, 0, 403, 61, 1, 0, 0, 0, 404, 405, 6, 31, -1,
		0, 405, 415, 3, 74, 37, 0, 406, 415, 5, 66, 0, 0, 407, 415, 3, 16, 8, 0,
		408, 409, 5, 48, 0, 0, 409, 410, 3, 62, 31, 0, 410, 411, 5, 49, 0, 0, 411,
		415, 1, 0, 0, 0, 412, 413, 7, 3, 0, 0, 413, 415, 3, 62, 31, 5, 414, 404,
		1, 0, 0, 0, 414, 406, 1, 0, 0, 0, 414, 407, 1, 0, 0, 0, 414, 408, 1, 0,
		0, 0, 414, 412, 1, 0, 0, 0, 415, 435, 1, 0, 0, 0, 416, 417, 10, 4, 0, 0,
		417, 418, 7, 4, 0, 0, 418, 434, 3, 62, 31, 5, 419, 420, 10, 3, 0, 0, 420,
		421, 7, 5, 0, 0, 421, 434, 3, 62, 31, 4, 422, 423, 10, 2, 0, 0, 423, 424,
		7, 6, 0, 0, 424, 434, 3, 62, 31, 3, 425, 426, 10, 1, 0, 0, 426, 427, 7,
		7, 0, 0, 427, 434, 3, 62, 31, 2, 428, 429, 10, 7, 0, 0, 429, 430, 5, 52,
		0, 0, 430, 431, 3, 62, 31, 0, 431, 432, 5, 53, 0, 0, 432, 434, 1, 0, 0,
		0, 433, 416, 1, 0, 0, 0, 433, 419, 1, 0, 0, 0, 433, 422, 1, 0, 0, 0, 433,
		425, 1, 0, 0, 0, 433, 428, 1, 0, 0, 0, 434, 437, 1, 0, 0, 0, 435, 433,
		1, 0, 0, 0, 435, 436, 1, 0, 0, 0, 436, 63, 1, 0, 0, 0, 437, 435, 1, 0,
		0, 0, 438, 442, 3, 70, 35, 0, 439, 441, 3, 66, 33, 0, 440, 439, 1, 0, 0,
		0, 441, 444, 1, 0, 0, 0, 442, 440, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443,
		65, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 445, 446, 5, 52, 0, 0, 446, 447,
		3, 62, 31, 0, 447, 448, 5, 53, 0, 0, 448, 459, 1, 0, 0, 0, 449, 450, 5,
		55, 0, 0, 450, 459, 5, 66, 0, 0, 451, 459, 5, 31, 0, 0, 452, 459, 5, 32,
		0, 0, 453, 455, 5, 48, 0, 0, 454, 456, 3, 68, 34, 0, 455, 454, 1, 0, 0,
		0, 455, 456, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 459, 5, 49, 0, 0, 458,
		445, 1, 0, 0, 0, 458, 449, 1, 0, 0, 0, 458, 451, 1, 0, 0, 0, 458, 452,
		1, 0, 0, 0, 458, 453, 1, 0, 0, 0, 459, 67, 1, 0, 0, 0, 460, 465, 3, 62,
		31, 0, 461, 462, 5, 54, 0, 0, 462, 464, 3, 62, 31, 0, 463, 461, 1, 0, 0,
		0, 464, 467, 1, 0, 0, 0, 465, 463, 1, 0, 0, 0, 465, 466, 1, 0, 0, 0, 466,
		69, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0, 468, 477, 5, 66, 0, 0, 469, 477,
		3, 74, 37, 0, 470, 471, 5, 48, 0, 0, 471, 472, 3, 62, 31, 0, 472, 473,
		5, 49, 0, 0, 473, 477, 1, 0, 0, 0, 474, 477, 3, 72, 36, 0, 475, 477, 3,
		16, 8, 0, 476, 468, 1, 0, 0, 0, 476, 469, 1, 0, 0, 0, 476, 470, 1, 0, 0,
		0, 476, 474, 1, 0, 0, 0, 476, 475, 1, 0, 0, 0, 477, 71, 1, 0, 0, 0, 478,
		487, 5, 52, 0, 0, 479, 484, 3, 62, 31, 0, 480, 481, 5, 54, 0, 0, 481, 483,
		3, 62, 31, 0, 482, 480, 1, 0, 0, 0, 483, 486, 1, 0, 0, 0, 484, 482, 1,
		0, 0, 0, 484, 485, 1, 0, 0, 0, 485, 488, 1, 0, 0, 0, 486, 484, 1, 0, 0,
		0, 487, 479, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489,
		490, 5, 53, 0, 0, 490, 73, 1, 0, 0, 0, 491, 498, 5, 65, 0, 0, 492, 498,
		5, 67, 0, 0, 493, 498, 5, 62, 0, 0, 494, 498, 5, 63, 0, 0, 495, 498, 5,
		64, 0, 0, 496, 498, 3, 72, 36, 0, 497, 491, 1, 0, 0, 0, 497, 492, 1, 0,
		0, 0, 497, 493, 1, 0, 0, 0, 497, 494, 1, 0, 0, 0, 497, 495, 1, 0, 0, 0,
		497, 496, 1, 0, 0, 0, 498, 75, 1, 0, 0, 0, 499, 503, 5, 50, 0, 0, 500,
		502, 3, 34, 17, 0, 501, 500, 1, 0, 0, 0, 502, 505, 1, 0, 0, 0, 503, 501,
		1, 0, 0, 0, 503, 504, 1, 0, 0, 0, 504, 506, 1, 0, 0, 0, 505, 503, 1, 0,
		0, 0, 506, 507, 5, 51, 0, 0, 507, 77, 1, 0, 0, 0, 508, 510, 5, 23, 0, 0,
		509, 511, 5, 48, 0, 0, 510, 509, 1, 0, 0, 0, 510, 511, 1, 0, 0, 0, 511,
		512, 1, 0, 0, 0, 512, 514, 3, 62, 31, 0, 513, 515, 5, 49, 0, 0, 514, 513,
		1, 0, 0, 0, 514, 515, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 516, 520, 5, 50,
		0, 0, 517, 519, 3, 80, 40, 0, 518, 517, 1, 0, 0, 0, 519, 522, 1, 0, 0,
		0, 520, 518, 1, 0, 0, 0, 520, 521, 1, 0, 0, 0, 521, 524, 1, 0, 0, 0, 522,
		520, 1, 0, 0, 0, 523, 525, 3, 82, 41, 0, 524, 523, 1, 0, 0, 0, 524, 525,
		1, 0, 0, 0, 525, 526, 1, 0, 0, 0, 526, 527, 5, 51, 0, 0, 527, 79, 1, 0,
		0, 0, 528, 529, 5, 24, 0, 0, 529, 530, 3, 62, 31, 0, 530, 531, 3, 76, 38,
		0, 531, 538, 1, 0, 0, 0, 532, 533, 5, 24, 0, 0, 533, 534, 3, 62, 31, 0,
		534, 535, 5, 60, 0, 0, 535, 536, 3, 34, 17, 0, 536, 538, 1, 0, 0, 0, 537,
		528, 1, 0, 0, 0, 537, 532, 1, 0, 0, 0, 538, 81, 1, 0, 0, 0, 539, 540, 5,
		25, 0, 0, 540, 545, 3, 76, 38, 0, 541, 542, 5, 25, 0, 0, 542, 543, 5, 60,
		0, 0, 543, 545, 3, 34, 17, 0, 544, 539, 1, 0, 0, 0, 544, 541, 1, 0, 0,
		0, 545, 83, 1, 0, 0, 0, 546, 547, 5, 58, 0, 0, 547, 548, 5, 59, 0, 0, 548,
		549, 5, 62, 0, 0, 549, 550, 5, 57, 0, 0, 550, 85, 1, 0, 0, 0, 69, 89, 95,
		101, 109, 126, 131, 133, 142, 148, 150, 156, 159, 165, 169, 173, 182, 188,
		197, 202, 206, 209, 212, 216, 221, 226, 235, 241, 249, 255, 260, 264, 280,
		287, 291, 296, 301, 305, 310, 314, 321, 323, 327, 331, 357, 365, 370, 373,
		377, 381, 390, 400, 414, 433, 435, 442, 455, 458, 465, 476, 484, 487, 497,
		503, 510, 514, 520, 524, 537, 544,
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

// R2D2ParserInit initializes any static state used to implement R2D2Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewR2D2Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func R2D2ParserInit() {
	staticData := &R2D2ParserStaticData
	staticData.once.Do(r2d2ParserInit)
}

// NewR2D2Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewR2D2Parser(input antlr.TokenStream) *R2D2Parser {
	R2D2ParserInit()
	this := new(R2D2Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &R2D2ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "R2D2.g4"

	return this
}

// R2D2Parser tokens.
const (
	R2D2ParserEOF            = antlr.TokenEOF
	R2D2ParserT__0           = 1
	R2D2ParserUSE            = 2
	R2D2ParserIMPORT         = 3
	R2D2ParserFROM           = 4
	R2D2ParserINTERFACE      = 5
	R2D2ParserMODULE         = 6
	R2D2ParserIMPLEMENTS     = 7
	R2D2ParserEXPORT         = 8
	R2D2ParserFN             = 9
	R2D2ParserPSEUDO         = 10
	R2D2ParserVAR            = 11
	R2D2ParserLET            = 12
	R2D2ParserCONST          = 13
	R2D2ParserIF             = 14
	R2D2ParserELSE           = 15
	R2D2ParserLOOP           = 16
	R2D2ParserFOR            = 17
	R2D2ParserWHILE          = 18
	R2D2ParserBREAK          = 19
	R2D2ParserSEND           = 20
	R2D2ParserCONTINUE       = 21
	R2D2ParserRETURN         = 22
	R2D2ParserSWITCH         = 23
	R2D2ParserCASE           = 24
	R2D2ParserDEFAULT        = 25
	R2D2ParserPLUS           = 26
	R2D2ParserMINUS          = 27
	R2D2ParserMULT           = 28
	R2D2ParserDIV            = 29
	R2D2ParserMOD            = 30
	R2D2ParserINCREMENT      = 31
	R2D2ParserDECREMENT      = 32
	R2D2ParserASSIGN         = 33
	R2D2ParserPLUS_ASSIGN    = 34
	R2D2ParserMINUS_ASSIGN   = 35
	R2D2ParserMULT_ASSIGN    = 36
	R2D2ParserDIV_ASSIGN     = 37
	R2D2ParserMOD_ASSIGN     = 38
	R2D2ParserEQ             = 39
	R2D2ParserNEQ            = 40
	R2D2ParserLT             = 41
	R2D2ParserGT             = 42
	R2D2ParserLEQ            = 43
	R2D2ParserGEQ            = 44
	R2D2ParserAND            = 45
	R2D2ParserOR             = 46
	R2D2ParserNOT            = 47
	R2D2ParserLPAREN         = 48
	R2D2ParserRPAREN         = 49
	R2D2ParserLBRACE         = 50
	R2D2ParserRBRACE         = 51
	R2D2ParserLBRACK         = 52
	R2D2ParserRBRACK         = 53
	R2D2ParserCOMMA          = 54
	R2D2ParserDOT            = 55
	R2D2ParserCOLON          = 56
	R2D2ParserSEMI           = 57
	R2D2ParserAT             = 58
	R2D2ParserJS             = 59
	R2D2ParserARROW          = 60
	R2D2ParserTYPE           = 61
	R2D2ParserSTRING_LITERAL = 62
	R2D2ParserBOOL_LITERAL   = 63
	R2D2ParserNULL_LITERAL   = 64
	R2D2ParserINT_LITERAL    = 65
	R2D2ParserIDENTIFIER     = 66
	R2D2ParserFLOAT_LITERAL  = 67
	R2D2ParserCOMMENT        = 68
	R2D2ParserBLOCK_COMMENT  = 69
	R2D2ParserWHITESPACE     = 70
)

// R2D2Parser rules.
const (
	R2D2ParserRULE_program               = 0
	R2D2ParserRULE_declaration           = 1
	R2D2ParserRULE_globalDeclaration     = 2
	R2D2ParserRULE_importDeclaration     = 3
	R2D2ParserRULE_interfaceDeclaration  = 4
	R2D2ParserRULE_moduleDeclaration     = 5
	R2D2ParserRULE_functionDeclaration   = 6
	R2D2ParserRULE_functionCallStatement = 7
	R2D2ParserRULE_functionCall          = 8
	R2D2ParserRULE_parameterList         = 9
	R2D2ParserRULE_parameter             = 10
	R2D2ParserRULE_typeExpression        = 11
	R2D2ParserRULE_arrayDimensions       = 12
	R2D2ParserRULE_baseType              = 13
	R2D2ParserRULE_genericType           = 14
	R2D2ParserRULE_typeDeclaration       = 15
	R2D2ParserRULE_variableDeclaration   = 16
	R2D2ParserRULE_statement             = 17
	R2D2ParserRULE_expressionStatement   = 18
	R2D2ParserRULE_ifStatement           = 19
	R2D2ParserRULE_forStatement          = 20
	R2D2ParserRULE_assignmentDeclaration = 21
	R2D2ParserRULE_assignment            = 22
	R2D2ParserRULE_assignmentOperator    = 23
	R2D2ParserRULE_simpleFor             = 24
	R2D2ParserRULE_whileStatement        = 25
	R2D2ParserRULE_loopStatement         = 26
	R2D2ParserRULE_cicleControl          = 27
	R2D2ParserRULE_breakStatement        = 28
	R2D2ParserRULE_continueStatement     = 29
	R2D2ParserRULE_returnStatement       = 30
	R2D2ParserRULE_expression            = 31
	R2D2ParserRULE_memberExpression      = 32
	R2D2ParserRULE_memberPart            = 33
	R2D2ParserRULE_argumentList          = 34
	R2D2ParserRULE_primaryExpression     = 35
	R2D2ParserRULE_arrayLiteral          = 36
	R2D2ParserRULE_literal               = 37
	R2D2ParserRULE_block                 = 38
	R2D2ParserRULE_switchStatement       = 39
	R2D2ParserRULE_switchCase            = 40
	R2D2ParserRULE_defaultCase           = 41
	R2D2ParserRULE_jsStatement           = 42
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllImportDeclaration() []IImportDeclarationContext
	ImportDeclaration(i int) IImportDeclarationContext
	AllInterfaceDeclaration() []IInterfaceDeclarationContext
	InterfaceDeclaration(i int) IInterfaceDeclarationContext
	AllDeclaration() []IDeclarationContext
	Declaration(i int) IDeclarationContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(R2D2ParserEOF, 0)
}

func (s *ProgramContext) AllImportDeclaration() []IImportDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IImportDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportDeclarationContext); ok {
			tst[i] = t.(IImportDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) ImportDeclaration(i int) IImportDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportDeclarationContext)
}

func (s *ProgramContext) AllInterfaceDeclaration() []IInterfaceDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInterfaceDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IInterfaceDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInterfaceDeclarationContext); ok {
			tst[i] = t.(IInterfaceDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) InterfaceDeclaration(i int) IInterfaceDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceDeclarationContext)
}

func (s *ProgramContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, R2D2ParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserUSE {
		{
			p.SetState(86)
			p.ImportDeclaration()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserINTERFACE {
		{
			p.SetState(92)
			p.InterfaceDeclaration()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8514) != 0 {
		{
			p.SetState(98)
			p.Declaration()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(104)
		p.Match(R2D2ParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ModuleDeclaration() IModuleDeclarationContext
	GlobalDeclaration() IGlobalDeclarationContext
	TypeDeclaration() ITypeDeclarationContext

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_declaration
	return p
}

func InitEmptyDeclarationContext(p *DeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_declaration
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) ModuleDeclaration() IModuleDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModuleDeclarationContext)
}

func (s *DeclarationContext) GlobalDeclaration() IGlobalDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalDeclarationContext)
}

func (s *DeclarationContext) TypeDeclaration() ITypeDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, R2D2ParserRULE_declaration)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case R2D2ParserMODULE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.ModuleDeclaration()
		}

	case R2D2ParserCONST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.GlobalDeclaration()
		}

	case R2D2ParserT__0, R2D2ParserEXPORT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.TypeDeclaration()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGlobalDeclarationContext is an interface to support dynamic dispatch.
type IGlobalDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	TypeExpression() ITypeExpressionContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	SEMI() antlr.TerminalNode

	// IsGlobalDeclarationContext differentiates from other interfaces.
	IsGlobalDeclarationContext()
}

type GlobalDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalDeclarationContext() *GlobalDeclarationContext {
	var p = new(GlobalDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_globalDeclaration
	return p
}

func InitEmptyGlobalDeclarationContext(p *GlobalDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_globalDeclaration
}

func (*GlobalDeclarationContext) IsGlobalDeclarationContext() {}

func NewGlobalDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalDeclarationContext {
	var p = new(GlobalDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_globalDeclaration

	return p
}

func (s *GlobalDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalDeclarationContext) CONST() antlr.TerminalNode {
	return s.GetToken(R2D2ParserCONST, 0)
}

func (s *GlobalDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *GlobalDeclarationContext) TypeExpression() ITypeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExpressionContext)
}

func (s *GlobalDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserASSIGN, 0)
}

func (s *GlobalDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GlobalDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, 0)
}

func (s *GlobalDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterGlobalDeclaration(s)
	}
}

func (s *GlobalDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitGlobalDeclaration(s)
	}
}

func (s *GlobalDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitGlobalDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) GlobalDeclaration() (localctx IGlobalDeclarationContext) {
	localctx = NewGlobalDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, R2D2ParserRULE_globalDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(R2D2ParserCONST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.TypeExpression()
	}
	{
		p.SetState(114)
		p.Match(R2D2ParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.expression(0)
	}
	{
		p.SetState(116)
		p.Match(R2D2ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportDeclarationContext is an interface to support dynamic dispatch.
type IImportDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	USE() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsImportDeclarationContext differentiates from other interfaces.
	IsImportDeclarationContext()
}

type ImportDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclarationContext() *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_importDeclaration
	return p
}

func InitEmptyImportDeclarationContext(p *ImportDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_importDeclaration
}

func (*ImportDeclarationContext) IsImportDeclarationContext() {}

func NewImportDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_importDeclaration

	return p
}

func (s *ImportDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclarationContext) USE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserUSE, 0)
}

func (s *ImportDeclarationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSTRING_LITERAL, 0)
}

func (s *ImportDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, 0)
}

func (s *ImportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitImportDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) ImportDeclaration() (localctx IImportDeclarationContext) {
	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, R2D2ParserRULE_importDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(R2D2ParserUSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(R2D2ParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(R2D2ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInterfaceDeclarationContext is an interface to support dynamic dispatch.
type IInterfaceDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTERFACE() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	IMPLEMENTS() antlr.TerminalNode
	AllFunctionDeclaration() []IFunctionDeclarationContext
	FunctionDeclaration(i int) IFunctionDeclarationContext
	AllVariableDeclaration() []IVariableDeclarationContext
	VariableDeclaration(i int) IVariableDeclarationContext

	// IsInterfaceDeclarationContext differentiates from other interfaces.
	IsInterfaceDeclarationContext()
}

type InterfaceDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceDeclarationContext() *InterfaceDeclarationContext {
	var p = new(InterfaceDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_interfaceDeclaration
	return p
}

func InitEmptyInterfaceDeclarationContext(p *InterfaceDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_interfaceDeclaration
}

func (*InterfaceDeclarationContext) IsInterfaceDeclarationContext() {}

func NewInterfaceDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceDeclarationContext {
	var p = new(InterfaceDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_interfaceDeclaration

	return p
}

func (s *InterfaceDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceDeclarationContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserINTERFACE, 0)
}

func (s *InterfaceDeclarationContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserIDENTIFIER)
}

func (s *InterfaceDeclarationContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, i)
}

func (s *InterfaceDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACE, 0)
}

func (s *InterfaceDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACE, 0)
}

func (s *InterfaceDeclarationContext) IMPLEMENTS() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIMPLEMENTS, 0)
}

func (s *InterfaceDeclarationContext) AllFunctionDeclaration() []IFunctionDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDeclarationContext); ok {
			tst[i] = t.(IFunctionDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceDeclarationContext) FunctionDeclaration(i int) IFunctionDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *InterfaceDeclarationContext) AllVariableDeclaration() []IVariableDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclarationContext); ok {
			tst[i] = t.(IVariableDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *InterfaceDeclarationContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *InterfaceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterInterfaceDeclaration(s)
	}
}

func (s *InterfaceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitInterfaceDeclaration(s)
	}
}

func (s *InterfaceDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitInterfaceDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) InterfaceDeclaration() (localctx IInterfaceDeclarationContext) {
	localctx = NewInterfaceDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, R2D2ParserRULE_interfaceDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(R2D2ParserINTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserIMPLEMENTS {
		{
			p.SetState(124)
			p.Match(R2D2ParserIMPLEMENTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(128)
		p.Match(R2D2ParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16128) != 0 {
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(129)
				p.FunctionDeclaration()
			}

		case 2:
			{
				p.SetState(130)
				p.VariableDeclaration()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(R2D2ParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IModuleDeclarationContext is an interface to support dynamic dispatch.
type IModuleDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MODULE() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	IMPLEMENTS() antlr.TerminalNode
	AllFunctionDeclaration() []IFunctionDeclarationContext
	FunctionDeclaration(i int) IFunctionDeclarationContext
	AllTypeDeclaration() []ITypeDeclarationContext
	TypeDeclaration(i int) ITypeDeclarationContext
	AllVariableDeclaration() []IVariableDeclarationContext
	VariableDeclaration(i int) IVariableDeclarationContext

	// IsModuleDeclarationContext differentiates from other interfaces.
	IsModuleDeclarationContext()
}

type ModuleDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleDeclarationContext() *ModuleDeclarationContext {
	var p = new(ModuleDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_moduleDeclaration
	return p
}

func InitEmptyModuleDeclarationContext(p *ModuleDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_moduleDeclaration
}

func (*ModuleDeclarationContext) IsModuleDeclarationContext() {}

func NewModuleDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDeclarationContext {
	var p = new(ModuleDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_moduleDeclaration

	return p
}

func (s *ModuleDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDeclarationContext) MODULE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserMODULE, 0)
}

func (s *ModuleDeclarationContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserIDENTIFIER)
}

func (s *ModuleDeclarationContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, i)
}

func (s *ModuleDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACE, 0)
}

func (s *ModuleDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACE, 0)
}

func (s *ModuleDeclarationContext) IMPLEMENTS() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIMPLEMENTS, 0)
}

func (s *ModuleDeclarationContext) AllFunctionDeclaration() []IFunctionDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDeclarationContext); ok {
			tst[i] = t.(IFunctionDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ModuleDeclarationContext) FunctionDeclaration(i int) IFunctionDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *ModuleDeclarationContext) AllTypeDeclaration() []ITypeDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclarationContext); ok {
			tst[i] = t.(ITypeDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ModuleDeclarationContext) TypeDeclaration(i int) ITypeDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *ModuleDeclarationContext) AllVariableDeclaration() []IVariableDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclarationContext); ok {
			tst[i] = t.(IVariableDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ModuleDeclarationContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ModuleDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterModuleDeclaration(s)
	}
}

func (s *ModuleDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitModuleDeclaration(s)
	}
}

func (s *ModuleDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitModuleDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) ModuleDeclaration() (localctx IModuleDeclarationContext) {
	localctx = NewModuleDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, R2D2ParserRULE_moduleDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(R2D2ParserMODULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserIMPLEMENTS {
		{
			p.SetState(140)
			p.Match(R2D2ParserIMPLEMENTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(144)
		p.Match(R2D2ParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16130) != 0 {
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(145)
				p.FunctionDeclaration()
			}

		case 2:
			{
				p.SetState(146)
				p.TypeDeclaration()
			}

		case 3:
			{
				p.SetState(147)
				p.VariableDeclaration()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
		p.Match(R2D2ParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	SEMI() antlr.TerminalNode
	EXPORT() antlr.TerminalNode
	PSEUDO() antlr.TerminalNode
	ParameterList() IParameterListContext
	TypeExpression() ITypeExpressionContext

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_functionDeclaration
	return p
}

func InitEmptyFunctionDeclarationContext(p *FunctionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_functionDeclaration
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) FN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserFN, 0)
}

func (s *FunctionDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *FunctionDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLPAREN, 0)
}

func (s *FunctionDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRPAREN, 0)
}

func (s *FunctionDeclarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, 0)
}

func (s *FunctionDeclarationContext) EXPORT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserEXPORT, 0)
}

func (s *FunctionDeclarationContext) PSEUDO() antlr.TerminalNode {
	return s.GetToken(R2D2ParserPSEUDO, 0)
}

func (s *FunctionDeclarationContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDeclarationContext) TypeExpression() ITypeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExpressionContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, R2D2ParserRULE_functionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserEXPORT {
		{
			p.SetState(155)
			p.Match(R2D2ParserEXPORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserPSEUDO {
		{
			p.SetState(158)
			p.Match(R2D2ParserPSEUDO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(161)
		p.Match(R2D2ParserFN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(R2D2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserIDENTIFIER {
		{
			p.SetState(164)
			p.ParameterList()
		}

	}
	{
		p.SetState(167)
		p.Match(R2D2ParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-52)) & ^0x3f) == 0 && ((int64(1)<<(_la-52))&16897) != 0 {
		{
			p.SetState(168)
			p.TypeExpression()
		}

	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case R2D2ParserLBRACE:
		{
			p.SetState(171)
			p.Block()
		}

	case R2D2ParserSEMI:
		{
			p.SetState(172)
			p.Match(R2D2ParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallStatementContext is an interface to support dynamic dispatch.
type IFunctionCallStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionCall() IFunctionCallContext
	SEMI() antlr.TerminalNode

	// IsFunctionCallStatementContext differentiates from other interfaces.
	IsFunctionCallStatementContext()
}

type FunctionCallStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallStatementContext() *FunctionCallStatementContext {
	var p = new(FunctionCallStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_functionCallStatement
	return p
}

func InitEmptyFunctionCallStatementContext(p *FunctionCallStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_functionCallStatement
}

func (*FunctionCallStatementContext) IsFunctionCallStatementContext() {}

func NewFunctionCallStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallStatementContext {
	var p = new(FunctionCallStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_functionCallStatement

	return p
}

func (s *FunctionCallStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallStatementContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, 0)
}

func (s *FunctionCallStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterFunctionCallStatement(s)
	}
}

func (s *FunctionCallStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitFunctionCallStatement(s)
	}
}

func (s *FunctionCallStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitFunctionCallStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) FunctionCallStatement() (localctx IFunctionCallStatementContext) {
	localctx = NewFunctionCallStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, R2D2ParserRULE_functionCallStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.FunctionCall()
	}
	{
		p.SetState(176)
		p.Match(R2D2ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserIDENTIFIER)
}

func (s *FunctionCallContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, i)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRPAREN, 0)
}

func (s *FunctionCallContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserDOT)
}

func (s *FunctionCallContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserDOT, i)
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, R2D2ParserRULE_functionCall)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(178)
				p.Match(R2D2ParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(179)
				p.Match(R2D2ParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(R2D2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-27)) & ^0x3f) == 0 && ((int64(1)<<(_la-27))&2164700217393) != 0 {
		{
			p.SetState(187)
			p.ArgumentList()
		}

	}
	{
		p.SetState(190)
		p.Match(R2D2ParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserCOMMA)
}

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserCOMMA, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, R2D2ParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Parameter()
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserCOMMA {
		{
			p.SetState(193)
			p.Match(R2D2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Parameter()
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	TypeExpression() ITypeExpressionContext

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *ParameterContext) TypeExpression() ITypeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExpressionContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, R2D2ParserRULE_parameter)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-52)) & ^0x3f) == 0 && ((int64(1)<<(_la-52))&16897) != 0 {
		{
			p.SetState(201)
			p.TypeExpression()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeExpressionContext is an interface to support dynamic dispatch.
type ITypeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BaseType() IBaseTypeContext
	ArrayDimensions() IArrayDimensionsContext

	// IsTypeExpressionContext differentiates from other interfaces.
	IsTypeExpressionContext()
}

type TypeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExpressionContext() *TypeExpressionContext {
	var p = new(TypeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_typeExpression
	return p
}

func InitEmptyTypeExpressionContext(p *TypeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_typeExpression
}

func (*TypeExpressionContext) IsTypeExpressionContext() {}

func NewTypeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExpressionContext {
	var p = new(TypeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_typeExpression

	return p
}

func (s *TypeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExpressionContext) BaseType() IBaseTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *TypeExpressionContext) ArrayDimensions() IArrayDimensionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDimensionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayDimensionsContext)
}

func (s *TypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterTypeExpression(s)
	}
}

func (s *TypeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitTypeExpression(s)
	}
}

func (s *TypeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitTypeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) TypeExpression() (localctx ITypeExpressionContext) {
	localctx = NewTypeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, R2D2ParserRULE_typeExpression)
	var _la int

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.BaseType()
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == R2D2ParserLBRACK {
			{
				p.SetState(205)
				p.ArrayDimensions()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == R2D2ParserLBRACK {
			{
				p.SetState(208)
				p.ArrayDimensions()
			}

		}
		{
			p.SetState(211)
			p.BaseType()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayDimensionsContext is an interface to support dynamic dispatch.
type IArrayDimensionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLBRACK() []antlr.TerminalNode
	LBRACK(i int) antlr.TerminalNode
	AllRBRACK() []antlr.TerminalNode
	RBRACK(i int) antlr.TerminalNode
	AllINT_LITERAL() []antlr.TerminalNode
	INT_LITERAL(i int) antlr.TerminalNode

	// IsArrayDimensionsContext differentiates from other interfaces.
	IsArrayDimensionsContext()
}

type ArrayDimensionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDimensionsContext() *ArrayDimensionsContext {
	var p = new(ArrayDimensionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_arrayDimensions
	return p
}

func InitEmptyArrayDimensionsContext(p *ArrayDimensionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_arrayDimensions
}

func (*ArrayDimensionsContext) IsArrayDimensionsContext() {}

func NewArrayDimensionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDimensionsContext {
	var p = new(ArrayDimensionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_arrayDimensions

	return p
}

func (s *ArrayDimensionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDimensionsContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserLBRACK)
}

func (s *ArrayDimensionsContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACK, i)
}

func (s *ArrayDimensionsContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserRBRACK)
}

func (s *ArrayDimensionsContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACK, i)
}

func (s *ArrayDimensionsContext) AllINT_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserINT_LITERAL)
}

func (s *ArrayDimensionsContext) INT_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserINT_LITERAL, i)
}

func (s *ArrayDimensionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDimensionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDimensionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterArrayDimensions(s)
	}
}

func (s *ArrayDimensionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitArrayDimensions(s)
	}
}

func (s *ArrayDimensionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitArrayDimensions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) ArrayDimensions() (localctx IArrayDimensionsContext) {
	localctx = NewArrayDimensionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, R2D2ParserRULE_arrayDimensions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == R2D2ParserLBRACK {
		{
			p.SetState(214)
			p.Match(R2D2ParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == R2D2ParserINT_LITERAL {
			{
				p.SetState(215)
				p.Match(R2D2ParserINT_LITERAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(218)
			p.Match(R2D2ParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	GenericType() IGenericTypeContext

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_baseType
	return p
}

func InitEmptyBaseTypeContext(p *BaseTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_baseType
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *BaseTypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserTYPE, 0)
}

func (s *BaseTypeContext) GenericType() IGenericTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericTypeContext)
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (s *BaseTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitBaseType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) BaseType() (localctx IBaseTypeContext) {
	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, R2D2ParserRULE_baseType)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Match(R2D2ParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(225)
			p.GenericType()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGenericTypeContext is an interface to support dynamic dispatch.
type IGenericTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LT() antlr.TerminalNode
	AllTypeExpression() []ITypeExpressionContext
	TypeExpression(i int) ITypeExpressionContext
	GT() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsGenericTypeContext differentiates from other interfaces.
	IsGenericTypeContext()
}

type GenericTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericTypeContext() *GenericTypeContext {
	var p = new(GenericTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_genericType
	return p
}

func InitEmptyGenericTypeContext(p *GenericTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_genericType
}

func (*GenericTypeContext) IsGenericTypeContext() {}

func NewGenericTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericTypeContext {
	var p = new(GenericTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_genericType

	return p
}

func (s *GenericTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *GenericTypeContext) LT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLT, 0)
}

func (s *GenericTypeContext) AllTypeExpression() []ITypeExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeExpressionContext); ok {
			len++
		}
	}

	tst := make([]ITypeExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeExpressionContext); ok {
			tst[i] = t.(ITypeExpressionContext)
			i++
		}
	}

	return tst
}

func (s *GenericTypeContext) TypeExpression(i int) ITypeExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExpressionContext)
}

func (s *GenericTypeContext) GT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserGT, 0)
}

func (s *GenericTypeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserCOMMA)
}

func (s *GenericTypeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserCOMMA, i)
}

func (s *GenericTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterGenericType(s)
	}
}

func (s *GenericTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitGenericType(s)
	}
}

func (s *GenericTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitGenericType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) GenericType() (localctx IGenericTypeContext) {
	localctx = NewGenericTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, R2D2ParserRULE_genericType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(R2D2ParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.TypeExpression()
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserCOMMA {
		{
			p.SetState(231)
			p.Match(R2D2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.TypeExpression()
		}

		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(238)
		p.Match(R2D2ParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	EXPORT() antlr.TerminalNode
	AllVariableDeclaration() []IVariableDeclarationContext
	VariableDeclaration(i int) IVariableDeclarationContext

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_typeDeclaration
	return p
}

func InitEmptyTypeDeclarationContext(p *TypeDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_typeDeclaration
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *TypeDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACE, 0)
}

func (s *TypeDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACE, 0)
}

func (s *TypeDeclarationContext) EXPORT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserEXPORT, 0)
}

func (s *TypeDeclarationContext) AllVariableDeclaration() []IVariableDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclarationContext); ok {
			tst[i] = t.(IVariableDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *TypeDeclarationContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, R2D2ParserRULE_typeDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserEXPORT {
		{
			p.SetState(240)
			p.Match(R2D2ParserEXPORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(243)
		p.Match(R2D2ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Match(R2D2ParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14592) != 0 {
		{
			p.SetState(246)
			p.VariableDeclaration()
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(252)
		p.Match(R2D2ParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	VAR() antlr.TerminalNode
	LET() antlr.TerminalNode
	CONST() antlr.TerminalNode
	EXPORT() antlr.TerminalNode
	TypeExpression() ITypeExpressionContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *VariableDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, 0)
}

func (s *VariableDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(R2D2ParserVAR, 0)
}

func (s *VariableDeclarationContext) LET() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLET, 0)
}

func (s *VariableDeclarationContext) CONST() antlr.TerminalNode {
	return s.GetToken(R2D2ParserCONST, 0)
}

func (s *VariableDeclarationContext) EXPORT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserEXPORT, 0)
}

func (s *VariableDeclarationContext) TypeExpression() ITypeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExpressionContext)
}

func (s *VariableDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserASSIGN, 0)
}

func (s *VariableDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, R2D2ParserRULE_variableDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserEXPORT {
		{
			p.SetState(254)
			p.Match(R2D2ParserEXPORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(257)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14336) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(258)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-52)) & ^0x3f) == 0 && ((int64(1)<<(_la-52))&16897) != 0 {
		{
			p.SetState(259)
			p.TypeExpression()
		}

	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserASSIGN {
		{
			p.SetState(262)
			p.Match(R2D2ParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.expression(0)
		}

	}
	{
		p.SetState(266)
		p.Match(R2D2ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDeclaration() IVariableDeclarationContext
	FunctionCallStatement() IFunctionCallStatementContext
	ExpressionStatement() IExpressionStatementContext
	IfStatement() IIfStatementContext
	ForStatement() IForStatementContext
	WhileStatement() IWhileStatementContext
	LoopStatement() ILoopStatementContext
	CicleControl() ICicleControlContext
	ReturnStatement() IReturnStatementContext
	SwitchStatement() ISwitchStatementContext
	AssignmentDeclaration() IAssignmentDeclarationContext
	JsStatement() IJsStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) FunctionCallStatement() IFunctionCallStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallStatementContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) LoopStatement() ILoopStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopStatementContext)
}

func (s *StatementContext) CicleControl() ICicleControlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICicleControlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICicleControlContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) AssignmentDeclaration() IAssignmentDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentDeclarationContext)
}

func (s *StatementContext) JsStatement() IJsStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJsStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJsStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, R2D2ParserRULE_statement)
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.VariableDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(269)
			p.FunctionCallStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(270)
			p.ExpressionStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(271)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(272)
			p.ForStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(273)
			p.WhileStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(274)
			p.LoopStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(275)
			p.CicleControl()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(276)
			p.ReturnStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(277)
			p.SwitchStatement()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(278)
			p.AssignmentDeclaration()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(279)
			p.JsStatement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	SEMI() antlr.TerminalNode

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_expressionStatement
	return p
}

func InitEmptyExpressionStatementContext(p *ExpressionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_expressionStatement
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, 0)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, R2D2ParserRULE_expressionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.expression(0)
	}
	{
		p.SetState(283)
		p.Match(R2D2ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIF() []antlr.TerminalNode
	IF(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllARROW() []antlr.TerminalNode
	ARROW(i int) antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	AllELSE() []antlr.TerminalNode
	ELSE(i int) antlr.TerminalNode

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserIF)
}

func (s *IfStatementContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserIF, i)
}

func (s *IfStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementContext) AllARROW() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserARROW)
}

func (s *IfStatementContext) ARROW(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserARROW, i)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserLPAREN)
}

func (s *IfStatementContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserLPAREN, i)
}

func (s *IfStatementContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserRPAREN)
}

func (s *IfStatementContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserRPAREN, i)
}

func (s *IfStatementContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserELSE)
}

func (s *IfStatementContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserELSE, i)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, R2D2ParserRULE_ifStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Match(R2D2ParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(286)
			p.Match(R2D2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(289)
		p.expression(0)
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserRPAREN {
		{
			p.SetState(290)
			p.Match(R2D2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case R2D2ParserLBRACE:
		{
			p.SetState(293)
			p.Block()
		}

	case R2D2ParserARROW:
		{
			p.SetState(294)
			p.Match(R2D2ParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.Statement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(298)
				p.Match(R2D2ParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(299)
				p.Match(R2D2ParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(301)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(300)
					p.Match(R2D2ParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			} else if p.HasError() { // JIM
				goto errorExit
			}
			{
				p.SetState(303)
				p.expression(0)
			}
			p.SetState(305)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == R2D2ParserRPAREN {
				{
					p.SetState(304)
					p.Match(R2D2ParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(310)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case R2D2ParserLBRACE:
				{
					p.SetState(307)
					p.Block()
				}

			case R2D2ParserARROW:
				{
					p.SetState(308)
					p.Match(R2D2ParserARROW)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(309)
					p.Statement()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(317)
			p.Match(R2D2ParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case R2D2ParserLBRACE:
			{
				p.SetState(318)
				p.Block()
			}

		case R2D2ParserARROW:
			{
				p.SetState(319)
				p.Match(R2D2ParserARROW)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(320)
				p.Statement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	SimpleFor() ISimpleForContext
	Block() IBlockContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(R2D2ParserFOR, 0)
}

func (s *ForStatementContext) SimpleFor() ISimpleForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleForContext)
}

func (s *ForStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLPAREN, 0)
}

func (s *ForStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRPAREN, 0)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, R2D2ParserRULE_forStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Match(R2D2ParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(326)
			p.Match(R2D2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(329)
		p.SimpleFor()
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserRPAREN {
		{
			p.SetState(330)
			p.Match(R2D2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(333)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentDeclarationContext is an interface to support dynamic dispatch.
type IAssignmentDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	SEMI() antlr.TerminalNode

	// IsAssignmentDeclarationContext differentiates from other interfaces.
	IsAssignmentDeclarationContext()
}

type AssignmentDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentDeclarationContext() *AssignmentDeclarationContext {
	var p = new(AssignmentDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_assignmentDeclaration
	return p
}

func InitEmptyAssignmentDeclarationContext(p *AssignmentDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_assignmentDeclaration
}

func (*AssignmentDeclarationContext) IsAssignmentDeclarationContext() {}

func NewAssignmentDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentDeclarationContext {
	var p = new(AssignmentDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_assignmentDeclaration

	return p
}

func (s *AssignmentDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentDeclarationContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *AssignmentDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, 0)
}

func (s *AssignmentDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterAssignmentDeclaration(s)
	}
}

func (s *AssignmentDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitAssignmentDeclaration(s)
	}
}

func (s *AssignmentDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitAssignmentDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) AssignmentDeclaration() (localctx IAssignmentDeclarationContext) {
	localctx = NewAssignmentDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, R2D2ParserRULE_assignmentDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Assignment()
	}
	{
		p.SetState(336)
		p.Match(R2D2ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AssignmentOperator() IAssignmentOperatorContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	INCREMENT() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *AssignmentContext) AssignmentOperator() IAssignmentOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentOperatorContext)
}

func (s *AssignmentContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserINCREMENT, 0)
}

func (s *AssignmentContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserDECREMENT, 0)
}

func (s *AssignmentContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACK, 0)
}

func (s *AssignmentContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACK, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, R2D2ParserRULE_assignment)
	var _la int

	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.AssignmentOperator()
		}
		{
			p.SetState(340)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			_la = p.GetTokenStream().LA(1)

			if !(_la == R2D2ParserINCREMENT || _la == R2D2ParserDECREMENT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(344)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Match(R2D2ParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.expression(0)
		}
		{
			p.SetState(347)
			p.Match(R2D2ParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
			p.AssignmentOperator()
		}
		{
			p.SetState(349)
			p.expression(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(351)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Match(R2D2ParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.expression(0)
		}
		{
			p.SetState(354)
			p.Match(R2D2ParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			_la = p.GetTokenStream().LA(1)

			if !(_la == R2D2ParserINCREMENT || _la == R2D2ParserDECREMENT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentOperatorContext is an interface to support dynamic dispatch.
type IAssignmentOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASSIGN() antlr.TerminalNode
	PLUS_ASSIGN() antlr.TerminalNode
	MINUS_ASSIGN() antlr.TerminalNode
	MULT_ASSIGN() antlr.TerminalNode
	DIV_ASSIGN() antlr.TerminalNode
	MOD_ASSIGN() antlr.TerminalNode

	// IsAssignmentOperatorContext differentiates from other interfaces.
	IsAssignmentOperatorContext()
}

type AssignmentOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOperatorContext() *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_assignmentOperator
	return p
}

func InitEmptyAssignmentOperatorContext(p *AssignmentOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_assignmentOperator
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentOperatorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserASSIGN, 0)
}

func (s *AssignmentOperatorContext) PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserPLUS_ASSIGN, 0)
}

func (s *AssignmentOperatorContext) MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserMINUS_ASSIGN, 0)
}

func (s *AssignmentOperatorContext) MULT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserMULT_ASSIGN, 0)
}

func (s *AssignmentOperatorContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserDIV_ASSIGN, 0)
}

func (s *AssignmentOperatorContext) MOD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserMOD_ASSIGN, 0)
}

func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitAssignmentOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, R2D2ParserRULE_assignmentOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&541165879296) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleForContext is an interface to support dynamic dispatch.
type ISimpleForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableDeclaration() IVariableDeclarationContext
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode
	Expression() IExpressionContext

	// IsSimpleForContext differentiates from other interfaces.
	IsSimpleForContext()
}

type SimpleForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleForContext() *SimpleForContext {
	var p = new(SimpleForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_simpleFor
	return p
}

func InitEmptySimpleForContext(p *SimpleForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_simpleFor
}

func (*SimpleForContext) IsSimpleForContext() {}

func NewSimpleForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleForContext {
	var p = new(SimpleForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_simpleFor

	return p
}

func (s *SimpleForContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleForContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *SimpleForContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *SimpleForContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *SimpleForContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserSEMI)
}

func (s *SimpleForContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, i)
}

func (s *SimpleForContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SimpleForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterSimpleFor(s)
	}
}

func (s *SimpleForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitSimpleFor(s)
	}
}

func (s *SimpleForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitSimpleFor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) SimpleFor() (localctx ISimpleForContext) {
	localctx = NewSimpleForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, R2D2ParserRULE_simpleFor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(365)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(361)
			p.VariableDeclaration()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(362)
			p.Assignment()
		}
		{
			p.SetState(363)
			p.Match(R2D2ParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(367)
			p.expression(0)
		}
		{
			p.SetState(368)
			p.Match(R2D2ParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserIDENTIFIER {
		{
			p.SetState(372)
			p.Assignment()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserWHILE, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLPAREN, 0)
}

func (s *WhileStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRPAREN, 0)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, R2D2ParserRULE_whileStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		p.Match(R2D2ParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(377)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(376)
			p.Match(R2D2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(379)
		p.expression(0)
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserRPAREN {
		{
			p.SetState(380)
			p.Match(R2D2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(383)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopStatementContext is an interface to support dynamic dispatch.
type ILoopStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOOP() antlr.TerminalNode
	Block() IBlockContext

	// IsLoopStatementContext differentiates from other interfaces.
	IsLoopStatementContext()
}

type LoopStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopStatementContext() *LoopStatementContext {
	var p = new(LoopStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_loopStatement
	return p
}

func InitEmptyLoopStatementContext(p *LoopStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_loopStatement
}

func (*LoopStatementContext) IsLoopStatementContext() {}

func NewLoopStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopStatementContext {
	var p = new(LoopStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_loopStatement

	return p
}

func (s *LoopStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopStatementContext) LOOP() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLOOP, 0)
}

func (s *LoopStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterLoopStatement(s)
	}
}

func (s *LoopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitLoopStatement(s)
	}
}

func (s *LoopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitLoopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) LoopStatement() (localctx ILoopStatementContext) {
	localctx = NewLoopStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, R2D2ParserRULE_loopStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(R2D2ParserLOOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(386)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICicleControlContext is an interface to support dynamic dispatch.
type ICicleControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BreakStatement() IBreakStatementContext
	ContinueStatement() IContinueStatementContext

	// IsCicleControlContext differentiates from other interfaces.
	IsCicleControlContext()
}

type CicleControlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCicleControlContext() *CicleControlContext {
	var p = new(CicleControlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_cicleControl
	return p
}

func InitEmptyCicleControlContext(p *CicleControlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_cicleControl
}

func (*CicleControlContext) IsCicleControlContext() {}

func NewCicleControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CicleControlContext {
	var p = new(CicleControlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_cicleControl

	return p
}

func (s *CicleControlContext) GetParser() antlr.Parser { return s.parser }

func (s *CicleControlContext) BreakStatement() IBreakStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *CicleControlContext) ContinueStatement() IContinueStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *CicleControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CicleControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CicleControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterCicleControl(s)
	}
}

func (s *CicleControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitCicleControl(s)
	}
}

func (s *CicleControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitCicleControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) CicleControl() (localctx ICicleControlContext) {
	localctx = NewCicleControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, R2D2ParserRULE_cicleControl)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case R2D2ParserBREAK:
		{
			p.SetState(388)
			p.BreakStatement()
		}

	case R2D2ParserCONTINUE:
		{
			p.SetState(389)
			p.ContinueStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_breakStatement
	return p
}

func InitEmptyBreakStatementContext(p *BreakStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_breakStatement
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(R2D2ParserBREAK, 0)
}

func (s *BreakStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, 0)
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, R2D2ParserRULE_breakStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Match(R2D2ParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(393)
		p.Match(R2D2ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_continueStatement
	return p
}

func InitEmptyContinueStatementContext(p *ContinueStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_continueStatement
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserCONTINUE, 0)
}

func (s *ContinueStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, 0)
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, R2D2ParserRULE_continueStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(R2D2ParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.Match(R2D2ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRETURN, 0)
}

func (s *ReturnStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, R2D2ParserRULE_returnStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Match(R2D2ParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-27)) & ^0x3f) == 0 && ((int64(1)<<(_la-27))&2164700217393) != 0 {
		{
			p.SetState(399)
			p.expression(0)
		}

	}
	{
		p.SetState(402)
		p.Match(R2D2ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayAccessExpressionContext struct {
	ExpressionContext
}

func NewArrayAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayAccessExpressionContext {
	var p = new(ArrayAccessExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAccessExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayAccessExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayAccessExpressionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACK, 0)
}

func (s *ArrayAccessExpressionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACK, 0)
}

func (s *ArrayAccessExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterArrayAccessExpression(s)
	}
}

func (s *ArrayAccessExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitArrayAccessExpression(s)
	}
}

func (s *ArrayAccessExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitArrayAccessExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditiveExpressionContext struct {
	ExpressionContext
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AdditiveExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(R2D2ParserPLUS, 0)
}

func (s *AdditiveExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(R2D2ParserMINUS, 0)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExpressionContext struct {
	ExpressionContext
}

func NewFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionContext {
	var p = new(FunctionCallExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitFunctionCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenthesisExpressionContext struct {
	ExpressionContext
}

func NewParenthesisExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisExpressionContext {
	var p = new(ParenthesisExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLPAREN, 0)
}

func (s *ParenthesisExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesisExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRPAREN, 0)
}

func (s *ParenthesisExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterParenthesisExpression(s)
	}
}

func (s *ParenthesisExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitParenthesisExpression(s)
	}
}

func (s *ParenthesisExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitParenthesisExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparisonExpressionContext struct {
	ExpressionContext
}

func NewComparisonExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ComparisonExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparisonExpressionContext) EQ() antlr.TerminalNode {
	return s.GetToken(R2D2ParserEQ, 0)
}

func (s *ComparisonExpressionContext) NEQ() antlr.TerminalNode {
	return s.GetToken(R2D2ParserNEQ, 0)
}

func (s *ComparisonExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLT, 0)
}

func (s *ComparisonExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserGT, 0)
}

func (s *ComparisonExpressionContext) LEQ() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLEQ, 0)
}

func (s *ComparisonExpressionContext) GEQ() antlr.TerminalNode {
	return s.GetToken(R2D2ParserGEQ, 0)
}

func (s *ComparisonExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterComparisonExpression(s)
	}
}

func (s *ComparisonExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitComparisonExpression(s)
	}
}

func (s *ComparisonExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitComparisonExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiplicativeExpressionContext struct {
	ExpressionContext
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultiplicativeExpressionContext) MULT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserMULT, 0)
}

func (s *MultiplicativeExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(R2D2ParserDIV, 0)
}

func (s *MultiplicativeExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(R2D2ParserMOD, 0)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExpressionContext struct {
	ExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExpressionContext struct {
	ExpressionContext
}

func NewUnaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserNOT, 0)
}

func (s *UnaryExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(R2D2ParserMINUS, 0)
}

func (s *UnaryExpressionContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserINCREMENT, 0)
}

func (s *UnaryExpressionContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserDECREMENT, 0)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalExpressionContext struct {
	ExpressionContext
}

func NewLogicalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(R2D2ParserAND, 0)
}

func (s *LogicalExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(R2D2ParserOR, 0)
}

func (s *LogicalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitLogicalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *R2D2Parser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, R2D2ParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(405)
			p.Literal()
		}

	case 2:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(406)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFunctionCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(407)
			p.FunctionCall()
		}

	case 4:
		localctx = NewParenthesisExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(408)
			p.Match(R2D2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.expression(0)
		}
		{
			p.SetState(410)
			p.Match(R2D2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(412)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140744065024000) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(413)
			p.expression(5)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(433)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, R2D2ParserRULE_expression)
				p.SetState(416)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(417)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1879048192) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(418)
					p.expression(5)
				}

			case 2:
				localctx = NewAdditiveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, R2D2ParserRULE_expression)
				p.SetState(419)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(420)
					_la = p.GetTokenStream().LA(1)

					if !(_la == R2D2ParserPLUS || _la == R2D2ParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(421)
					p.expression(4)
				}

			case 3:
				localctx = NewComparisonExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, R2D2ParserRULE_expression)
				p.SetState(422)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(423)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34634616274944) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(424)
					p.expression(3)
				}

			case 4:
				localctx = NewLogicalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, R2D2ParserRULE_expression)
				p.SetState(425)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(426)
					_la = p.GetTokenStream().LA(1)

					if !(_la == R2D2ParserAND || _la == R2D2ParserOR) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(427)
					p.expression(2)
				}

			case 5:
				localctx = NewArrayAccessExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, R2D2ParserRULE_expression)
				p.SetState(428)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(429)
					p.Match(R2D2ParserLBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(430)
					p.expression(0)
				}
				{
					p.SetState(431)
					p.Match(R2D2ParserRBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMemberExpressionContext is an interface to support dynamic dispatch.
type IMemberExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	AllMemberPart() []IMemberPartContext
	MemberPart(i int) IMemberPartContext

	// IsMemberExpressionContext differentiates from other interfaces.
	IsMemberExpressionContext()
}

type MemberExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberExpressionContext() *MemberExpressionContext {
	var p = new(MemberExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_memberExpression
	return p
}

func InitEmptyMemberExpressionContext(p *MemberExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_memberExpression
}

func (*MemberExpressionContext) IsMemberExpressionContext() {}

func NewMemberExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberExpressionContext {
	var p = new(MemberExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_memberExpression

	return p
}

func (s *MemberExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *MemberExpressionContext) AllMemberPart() []IMemberPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMemberPartContext); ok {
			len++
		}
	}

	tst := make([]IMemberPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMemberPartContext); ok {
			tst[i] = t.(IMemberPartContext)
			i++
		}
	}

	return tst
}

func (s *MemberExpressionContext) MemberPart(i int) IMemberPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberPartContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberPartContext)
}

func (s *MemberExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterMemberExpression(s)
	}
}

func (s *MemberExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitMemberExpression(s)
	}
}

func (s *MemberExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitMemberExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) MemberExpression() (localctx IMemberExpressionContext) {
	localctx = NewMemberExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, R2D2ParserRULE_memberExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		p.PrimaryExpression()
	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&40813878065496064) != 0 {
		{
			p.SetState(439)
			p.MemberPart()
		}

		p.SetState(444)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMemberPartContext is an interface to support dynamic dispatch.
type IMemberPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	Expression() IExpressionContext
	RBRACK() antlr.TerminalNode
	DOT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	INCREMENT() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsMemberPartContext differentiates from other interfaces.
	IsMemberPartContext()
}

type MemberPartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberPartContext() *MemberPartContext {
	var p = new(MemberPartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_memberPart
	return p
}

func InitEmptyMemberPartContext(p *MemberPartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_memberPart
}

func (*MemberPartContext) IsMemberPartContext() {}

func NewMemberPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberPartContext {
	var p = new(MemberPartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_memberPart

	return p
}

func (s *MemberPartContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberPartContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACK, 0)
}

func (s *MemberPartContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemberPartContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACK, 0)
}

func (s *MemberPartContext) DOT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserDOT, 0)
}

func (s *MemberPartContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *MemberPartContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserINCREMENT, 0)
}

func (s *MemberPartContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserDECREMENT, 0)
}

func (s *MemberPartContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLPAREN, 0)
}

func (s *MemberPartContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRPAREN, 0)
}

func (s *MemberPartContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *MemberPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterMemberPart(s)
	}
}

func (s *MemberPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitMemberPart(s)
	}
}

func (s *MemberPartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitMemberPart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) MemberPart() (localctx IMemberPartContext) {
	localctx = NewMemberPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, R2D2ParserRULE_memberPart)
	var _la int

	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case R2D2ParserLBRACK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(445)
			p.Match(R2D2ParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.expression(0)
		}
		{
			p.SetState(447)
			p.Match(R2D2ParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(449)
			p.Match(R2D2ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(450)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserINCREMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(451)
			p.Match(R2D2ParserINCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserDECREMENT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(452)
			p.Match(R2D2ParserDECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserLPAREN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(453)
			p.Match(R2D2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(455)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-27)) & ^0x3f) == 0 && ((int64(1)<<(_la-27))&2164700217393) != 0 {
			{
				p.SetState(454)
				p.ArgumentList()
			}

		}
		{
			p.SetState(457)
			p.Match(R2D2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, R2D2ParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		p.expression(0)
	}
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserCOMMA {
		{
			p.SetState(461)
			p.Match(R2D2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.expression(0)
		}

		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Literal() ILiteralContext
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	ArrayLiteral() IArrayLiteralContext
	FunctionCall() IFunctionCallContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *PrimaryExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLPAREN, 0)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRPAREN, 0)
}

func (s *PrimaryExpressionContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *PrimaryExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, R2D2ParserRULE_primaryExpression)
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(468)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(469)
			p.Literal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(470)
			p.Match(R2D2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
			p.expression(0)
		}
		{
			p.SetState(472)
			p.Match(R2D2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(474)
			p.ArrayLiteral()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(475)
			p.FunctionCall()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACK, 0)
}

func (s *ArrayLiteralContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACK, 0)
}

func (s *ArrayLiteralContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserCOMMA)
}

func (s *ArrayLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserCOMMA, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, R2D2ParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(478)
		p.Match(R2D2ParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-27)) & ^0x3f) == 0 && ((int64(1)<<(_la-27))&2164700217393) != 0 {
		{
			p.SetState(479)
			p.expression(0)
		}
		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == R2D2ParserCOMMA {
			{
				p.SetState(480)
				p.Match(R2D2ParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(481)
				p.expression(0)
			}

			p.SetState(486)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(489)
		p.Match(R2D2ParserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_LITERAL() antlr.TerminalNode
	FLOAT_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	BOOL_LITERAL() antlr.TerminalNode
	NULL_LITERAL() antlr.TerminalNode
	ArrayLiteral() IArrayLiteralContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(R2D2ParserINT_LITERAL, 0)
}

func (s *LiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(R2D2ParserFLOAT_LITERAL, 0)
}

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(R2D2ParserBOOL_LITERAL, 0)
}

func (s *LiteralContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(R2D2ParserNULL_LITERAL, 0)
}

func (s *LiteralContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, R2D2ParserRULE_literal)
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case R2D2ParserINT_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(491)
			p.Match(R2D2ParserINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserFLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(492)
			p.Match(R2D2ParserFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(493)
			p.Match(R2D2ParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserBOOL_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(494)
			p.Match(R2D2ParserBOOL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserNULL_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(495)
			p.Match(R2D2ParserNULL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserLBRACK:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(496)
			p.ArrayLiteral()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACE, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, R2D2ParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)
		p.Match(R2D2ParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-8)) & ^0x3f) == 0 && ((int64(1)<<(_la-8))&1136052247483445113) != 0 {
		{
			p.SetState(500)
			p.Statement()
		}

		p.SetState(505)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(506)
		p.Match(R2D2ParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expression() IExpressionContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllSwitchCase() []ISwitchCaseContext
	SwitchCase(i int) ISwitchCaseContext
	DefaultCase() IDefaultCaseContext

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_switchStatement
	return p
}

func InitEmptySwitchStatementContext(p *SwitchStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_switchStatement
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSWITCH, 0)
}

func (s *SwitchStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACE, 0)
}

func (s *SwitchStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACE, 0)
}

func (s *SwitchStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLPAREN, 0)
}

func (s *SwitchStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRPAREN, 0)
}

func (s *SwitchStatementContext) AllSwitchCase() []ISwitchCaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			len++
		}
	}

	tst := make([]ISwitchCaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchCaseContext); ok {
			tst[i] = t.(ISwitchCaseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStatementContext) SwitchCase(i int) ISwitchCaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchCaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchCaseContext)
}

func (s *SwitchStatementContext) DefaultCase() IDefaultCaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultCaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultCaseContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, R2D2ParserRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(508)
		p.Match(R2D2ParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(510)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(509)
			p.Match(R2D2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(512)
		p.expression(0)
	}
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserRPAREN {
		{
			p.SetState(513)
			p.Match(R2D2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(516)
		p.Match(R2D2ParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(520)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserCASE {
		{
			p.SetState(517)
			p.SwitchCase()
		}

		p.SetState(522)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(524)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserDEFAULT {
		{
			p.SetState(523)
			p.DefaultCase()
		}

	}
	{
		p.SetState(526)
		p.Match(R2D2ParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchCaseContext is an interface to support dynamic dispatch.
type ISwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CASE() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext
	ARROW() antlr.TerminalNode
	Statement() IStatementContext

	// IsSwitchCaseContext differentiates from other interfaces.
	IsSwitchCaseContext()
}

type SwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchCaseContext() *SwitchCaseContext {
	var p = new(SwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_switchCase
	return p
}

func InitEmptySwitchCaseContext(p *SwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_switchCase
}

func (*SwitchCaseContext) IsSwitchCaseContext() {}

func NewSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_switchCase

	return p
}

func (s *SwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchCaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserCASE, 0)
}

func (s *SwitchCaseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchCaseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SwitchCaseContext) ARROW() antlr.TerminalNode {
	return s.GetToken(R2D2ParserARROW, 0)
}

func (s *SwitchCaseContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterSwitchCase(s)
	}
}

func (s *SwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitSwitchCase(s)
	}
}

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) SwitchCase() (localctx ISwitchCaseContext) {
	localctx = NewSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, R2D2ParserRULE_switchCase)
	p.SetState(537)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(528)
			p.Match(R2D2ParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(529)
			p.expression(0)
		}
		{
			p.SetState(530)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(532)
			p.Match(R2D2ParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(533)
			p.expression(0)
		}
		{
			p.SetState(534)
			p.Match(R2D2ParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(535)
			p.Statement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultCaseContext is an interface to support dynamic dispatch.
type IDefaultCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	Block() IBlockContext
	ARROW() antlr.TerminalNode
	Statement() IStatementContext

	// IsDefaultCaseContext differentiates from other interfaces.
	IsDefaultCaseContext()
}

type DefaultCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultCaseContext() *DefaultCaseContext {
	var p = new(DefaultCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_defaultCase
	return p
}

func InitEmptyDefaultCaseContext(p *DefaultCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_defaultCase
}

func (*DefaultCaseContext) IsDefaultCaseContext() {}

func NewDefaultCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultCaseContext {
	var p = new(DefaultCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_defaultCase

	return p
}

func (s *DefaultCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultCaseContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserDEFAULT, 0)
}

func (s *DefaultCaseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DefaultCaseContext) ARROW() antlr.TerminalNode {
	return s.GetToken(R2D2ParserARROW, 0)
}

func (s *DefaultCaseContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterDefaultCase(s)
	}
}

func (s *DefaultCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitDefaultCase(s)
	}
}

func (s *DefaultCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitDefaultCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) DefaultCase() (localctx IDefaultCaseContext) {
	localctx = NewDefaultCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, R2D2ParserRULE_defaultCase)
	p.SetState(544)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 68, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(539)
			p.Match(R2D2ParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(541)
			p.Match(R2D2ParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(542)
			p.Match(R2D2ParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(543)
			p.Statement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJsStatementContext is an interface to support dynamic dispatch.
type IJsStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	JS() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsJsStatementContext differentiates from other interfaces.
	IsJsStatementContext()
}

type JsStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsStatementContext() *JsStatementContext {
	var p = new(JsStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_jsStatement
	return p
}

func InitEmptyJsStatementContext(p *JsStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_jsStatement
}

func (*JsStatementContext) IsJsStatementContext() {}

func NewJsStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsStatementContext {
	var p = new(JsStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_jsStatement

	return p
}

func (s *JsStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *JsStatementContext) AT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserAT, 0)
}

func (s *JsStatementContext) JS() antlr.TerminalNode {
	return s.GetToken(R2D2ParserJS, 0)
}

func (s *JsStatementContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSTRING_LITERAL, 0)
}

func (s *JsStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(R2D2ParserSEMI, 0)
}

func (s *JsStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterJsStatement(s)
	}
}

func (s *JsStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitJsStatement(s)
	}
}

func (s *JsStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitJsStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) JsStatement() (localctx IJsStatementContext) {
	localctx = NewJsStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, R2D2ParserRULE_jsStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(546)
		p.Match(R2D2ParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(547)
		p.Match(R2D2ParserJS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(548)
		p.Match(R2D2ParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(549)
		p.Match(R2D2ParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *R2D2Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 31:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *R2D2Parser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
