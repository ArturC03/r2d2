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
		"", "'type'", "'import'", "'from'", "'interface'", "'module'", "'implements'",
		"'export'", "'fn'", "'pseudo'", "'var'", "'let'", "'const'", "'if'",
		"'else'", "'loop'", "'for'", "'while'", "'break'", "'send'", "'continue'",
		"'return'", "'switch'", "'case'", "'default'", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'++'", "'--'", "'='", "'+='", "'-='", "'*='", "'/='",
		"'%='", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'", "'||'",
		"'!'", "'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'.'", "':'",
		"';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "IMPORT", "FROM", "INTERFACE", "MODULE", "IMPLEMENTS", "EXPORT",
		"FN", "PSEUDO", "VAR", "LET", "CONST", "IF", "ELSE", "LOOP", "FOR",
		"WHILE", "BREAK", "SEND", "CONTINUE", "RETURN", "SWITCH", "CASE", "DEFAULT",
		"PLUS", "MINUS", "MULT", "DIV", "MOD", "INCREMENT", "DECREMENT", "ASSIGN",
		"PLUS_ASSIGN", "MINUS_ASSIGN", "MULT_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN",
		"EQ", "NEQ", "LT", "GT", "LEQ", "GEQ", "AND", "OR", "NOT", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "COMMA", "DOT", "COLON",
		"SEMI", "IDENTIFIER", "TYPE", "STRING_LITERAL", "BOOL_LITERAL", "INT_LITERAL",
		"FLOAT_LITERAL", "COMMENT", "BLOCK_COMMENT", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"program", "declaration", "globalDeclaration", "importDeclaration",
		"interfaceDeclaration", "moduleDeclaration", "functionDeclaration",
		"functionCallStatement", "functionCall", "parameterList", "parameter",
		"typeExpression", "arrayDimensions", "baseType", "genericType", "typeDeclaration",
		"variableDeclaration", "statement", "expressionStatement", "ifStatement",
		"forStatement", "assignmentDeclaration", "assignment", "assignmentOperator",
		"simpleFor", "whileStatement", "loopStatement", "cicleControl", "breakStatement",
		"continueStatement", "returnStatement", "expression", "logicalExpression",
		"comparisonExpression", "additiveExpression", "multiplicativeExpression",
		"unaryExpression", "memberExpression", "memberPart", "argumentList",
		"primaryExpression", "arrayLiteral", "literal", "block", "switchStatement",
		"switchCase", "defaultCase",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 65, 497, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 1, 0,
		5, 0, 96, 8, 0, 10, 0, 12, 0, 99, 9, 0, 1, 0, 5, 0, 102, 8, 0, 10, 0, 12,
		0, 105, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 112, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 5, 4, 131, 8, 4, 10, 4, 12, 4, 134, 9, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 142, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5,
		148, 8, 5, 10, 5, 12, 5, 151, 9, 5, 1, 5, 1, 5, 1, 6, 3, 6, 156, 8, 6,
		1, 6, 3, 6, 159, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 165, 8, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 170, 8, 6, 1, 6, 1, 6, 3, 6, 174, 8, 6, 1, 7, 1, 7, 1, 7,
		1, 8, 1, 8, 5, 8, 181, 8, 8, 10, 8, 12, 8, 184, 9, 8, 1, 8, 1, 8, 1, 8,
		3, 8, 189, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 196, 8, 9, 10, 9,
		12, 9, 199, 9, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 3, 11, 206, 8, 11,
		1, 12, 1, 12, 3, 12, 210, 8, 12, 1, 12, 4, 12, 213, 8, 12, 11, 12, 12,
		12, 214, 1, 13, 1, 13, 1, 13, 3, 13, 220, 8, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 5, 14, 227, 8, 14, 10, 14, 12, 14, 230, 9, 14, 1, 14, 1, 14,
		1, 15, 3, 15, 235, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 241, 8, 15,
		10, 15, 12, 15, 244, 9, 15, 1, 15, 1, 15, 1, 16, 3, 16, 249, 8, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 254, 8, 16, 1, 16, 1, 16, 3, 16, 258, 8, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 3, 17, 273, 8, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 3,
		19, 280, 8, 19, 1, 19, 1, 19, 3, 19, 284, 8, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 3, 19, 290, 8, 19, 1, 19, 1, 19, 3, 19, 294, 8, 19, 1, 19, 1, 19, 5,
		19, 298, 8, 19, 10, 19, 12, 19, 301, 9, 19, 1, 19, 1, 19, 3, 19, 305, 8,
		19, 1, 20, 1, 20, 3, 20, 309, 8, 20, 1, 20, 1, 20, 3, 20, 313, 8, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		3, 22, 326, 8, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 334,
		8, 24, 1, 24, 1, 24, 1, 24, 3, 24, 339, 8, 24, 1, 24, 3, 24, 342, 8, 24,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1,
		27, 3, 27, 355, 8, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30,
		1, 30, 3, 30, 365, 8, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 5, 32, 374, 8, 32, 10, 32, 12, 32, 377, 9, 32, 1, 33, 1, 33, 1, 33,
		5, 33, 382, 8, 33, 10, 33, 12, 33, 385, 9, 33, 1, 34, 1, 34, 1, 34, 5,
		34, 390, 8, 34, 10, 34, 12, 34, 393, 9, 34, 1, 35, 1, 35, 1, 35, 5, 35,
		398, 8, 35, 10, 35, 12, 35, 401, 9, 35, 1, 36, 1, 36, 1, 36, 3, 36, 406,
		8, 36, 1, 37, 1, 37, 5, 37, 410, 8, 37, 10, 37, 12, 37, 413, 9, 37, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38,
		425, 8, 38, 1, 38, 3, 38, 428, 8, 38, 1, 39, 1, 39, 1, 39, 5, 39, 433,
		8, 39, 10, 39, 12, 39, 436, 9, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 3, 40, 446, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41,
		452, 8, 41, 10, 41, 12, 41, 455, 9, 41, 3, 41, 457, 8, 41, 1, 41, 1, 41,
		1, 42, 1, 42, 1, 43, 1, 43, 5, 43, 465, 8, 43, 10, 43, 12, 43, 468, 9,
		43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 5, 44, 478,
		8, 44, 10, 44, 12, 44, 481, 9, 44, 1, 44, 3, 44, 484, 8, 44, 1, 44, 1,
		44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46,
		0, 0, 47, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 0, 9, 1, 0, 10, 12, 1,
		0, 30, 31, 1, 0, 32, 37, 1, 0, 44, 45, 1, 0, 38, 43, 1, 0, 25, 26, 1, 0,
		27, 29, 3, 0, 26, 26, 30, 31, 46, 46, 1, 0, 59, 62, 523, 0, 97, 1, 0, 0,
		0, 2, 111, 1, 0, 0, 0, 4, 113, 1, 0, 0, 0, 6, 120, 1, 0, 0, 0, 8, 126,
		1, 0, 0, 0, 10, 137, 1, 0, 0, 0, 12, 155, 1, 0, 0, 0, 14, 175, 1, 0, 0,
		0, 16, 182, 1, 0, 0, 0, 18, 192, 1, 0, 0, 0, 20, 200, 1, 0, 0, 0, 22, 203,
		1, 0, 0, 0, 24, 212, 1, 0, 0, 0, 26, 219, 1, 0, 0, 0, 28, 221, 1, 0, 0,
		0, 30, 234, 1, 0, 0, 0, 32, 248, 1, 0, 0, 0, 34, 272, 1, 0, 0, 0, 36, 274,
		1, 0, 0, 0, 38, 277, 1, 0, 0, 0, 40, 306, 1, 0, 0, 0, 42, 316, 1, 0, 0,
		0, 44, 325, 1, 0, 0, 0, 46, 327, 1, 0, 0, 0, 48, 333, 1, 0, 0, 0, 50, 343,
		1, 0, 0, 0, 52, 349, 1, 0, 0, 0, 54, 354, 1, 0, 0, 0, 56, 356, 1, 0, 0,
		0, 58, 359, 1, 0, 0, 0, 60, 362, 1, 0, 0, 0, 62, 368, 1, 0, 0, 0, 64, 370,
		1, 0, 0, 0, 66, 378, 1, 0, 0, 0, 68, 386, 1, 0, 0, 0, 70, 394, 1, 0, 0,
		0, 72, 405, 1, 0, 0, 0, 74, 407, 1, 0, 0, 0, 76, 427, 1, 0, 0, 0, 78, 429,
		1, 0, 0, 0, 80, 445, 1, 0, 0, 0, 82, 447, 1, 0, 0, 0, 84, 460, 1, 0, 0,
		0, 86, 462, 1, 0, 0, 0, 88, 471, 1, 0, 0, 0, 90, 487, 1, 0, 0, 0, 92, 492,
		1, 0, 0, 0, 94, 96, 3, 6, 3, 0, 95, 94, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0,
		97, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 103, 1, 0, 0, 0, 99, 97, 1,
		0, 0, 0, 100, 102, 3, 2, 1, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0,
		0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 106, 1, 0, 0, 0, 105,
		103, 1, 0, 0, 0, 106, 107, 5, 0, 0, 1, 107, 1, 1, 0, 0, 0, 108, 112, 3,
		10, 5, 0, 109, 112, 3, 8, 4, 0, 110, 112, 3, 4, 2, 0, 111, 108, 1, 0, 0,
		0, 111, 109, 1, 0, 0, 0, 111, 110, 1, 0, 0, 0, 112, 3, 1, 0, 0, 0, 113,
		114, 5, 12, 0, 0, 114, 115, 5, 57, 0, 0, 115, 116, 3, 22, 11, 0, 116, 117,
		5, 32, 0, 0, 117, 118, 3, 62, 31, 0, 118, 119, 5, 56, 0, 0, 119, 5, 1,
		0, 0, 0, 120, 121, 5, 2, 0, 0, 121, 122, 5, 57, 0, 0, 122, 123, 5, 3, 0,
		0, 123, 124, 5, 59, 0, 0, 124, 125, 5, 56, 0, 0, 125, 7, 1, 0, 0, 0, 126,
		127, 5, 4, 0, 0, 127, 128, 5, 57, 0, 0, 128, 132, 5, 49, 0, 0, 129, 131,
		3, 12, 6, 0, 130, 129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0,
		0, 0, 132, 133, 1, 0, 0, 0, 133, 135, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0,
		135, 136, 5, 50, 0, 0, 136, 9, 1, 0, 0, 0, 137, 138, 5, 5, 0, 0, 138, 141,
		5, 57, 0, 0, 139, 140, 5, 6, 0, 0, 140, 142, 5, 57, 0, 0, 141, 139, 1,
		0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 149, 5, 49, 0,
		0, 144, 148, 3, 12, 6, 0, 145, 148, 3, 30, 15, 0, 146, 148, 3, 32, 16,
		0, 147, 144, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 146, 1, 0, 0, 0, 148,
		151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 152,
		1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 153, 5, 50, 0, 0, 153, 11, 1, 0,
		0, 0, 154, 156, 5, 7, 0, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0,
		156, 158, 1, 0, 0, 0, 157, 159, 5, 9, 0, 0, 158, 157, 1, 0, 0, 0, 158,
		159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 5, 8, 0, 0, 161, 162,
		5, 57, 0, 0, 162, 164, 5, 47, 0, 0, 163, 165, 3, 18, 9, 0, 164, 163, 1,
		0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 169, 5, 48, 0,
		0, 167, 168, 5, 55, 0, 0, 168, 170, 3, 22, 11, 0, 169, 167, 1, 0, 0, 0,
		169, 170, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 174, 3, 86, 43, 0, 172,
		174, 5, 56, 0, 0, 173, 171, 1, 0, 0, 0, 173, 172, 1, 0, 0, 0, 174, 13,
		1, 0, 0, 0, 175, 176, 3, 16, 8, 0, 176, 177, 5, 56, 0, 0, 177, 15, 1, 0,
		0, 0, 178, 179, 5, 57, 0, 0, 179, 181, 5, 54, 0, 0, 180, 178, 1, 0, 0,
		0, 181, 184, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183,
		185, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 185, 186, 5, 57, 0, 0, 186, 188,
		5, 47, 0, 0, 187, 189, 3, 78, 39, 0, 188, 187, 1, 0, 0, 0, 188, 189, 1,
		0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 5, 48, 0, 0, 191, 17, 1, 0, 0,
		0, 192, 197, 3, 20, 10, 0, 193, 194, 5, 53, 0, 0, 194, 196, 3, 20, 10,
		0, 195, 193, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197,
		198, 1, 0, 0, 0, 198, 19, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 201, 5,
		57, 0, 0, 201, 202, 3, 22, 11, 0, 202, 21, 1, 0, 0, 0, 203, 205, 3, 26,
		13, 0, 204, 206, 3, 24, 12, 0, 205, 204, 1, 0, 0, 0, 205, 206, 1, 0, 0,
		0, 206, 23, 1, 0, 0, 0, 207, 209, 5, 51, 0, 0, 208, 210, 5, 61, 0, 0, 209,
		208, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 213,
		5, 52, 0, 0, 212, 207, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 212, 1, 0,
		0, 0, 214, 215, 1, 0, 0, 0, 215, 25, 1, 0, 0, 0, 216, 220, 5, 57, 0, 0,
		217, 220, 5, 58, 0, 0, 218, 220, 3, 28, 14, 0, 219, 216, 1, 0, 0, 0, 219,
		217, 1, 0, 0, 0, 219, 218, 1, 0, 0, 0, 220, 27, 1, 0, 0, 0, 221, 222, 5,
		57, 0, 0, 222, 223, 5, 40, 0, 0, 223, 228, 3, 22, 11, 0, 224, 225, 5, 53,
		0, 0, 225, 227, 3, 22, 11, 0, 226, 224, 1, 0, 0, 0, 227, 230, 1, 0, 0,
		0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 231, 1, 0, 0, 0, 230,
		228, 1, 0, 0, 0, 231, 232, 5, 41, 0, 0, 232, 29, 1, 0, 0, 0, 233, 235,
		5, 7, 0, 0, 234, 233, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 1, 0,
		0, 0, 236, 237, 5, 1, 0, 0, 237, 238, 5, 57, 0, 0, 238, 242, 5, 49, 0,
		0, 239, 241, 3, 32, 16, 0, 240, 239, 1, 0, 0, 0, 241, 244, 1, 0, 0, 0,
		242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 245, 1, 0, 0, 0, 244,
		242, 1, 0, 0, 0, 245, 246, 5, 50, 0, 0, 246, 31, 1, 0, 0, 0, 247, 249,
		5, 7, 0, 0, 248, 247, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 1, 0,
		0, 0, 250, 251, 7, 0, 0, 0, 251, 253, 5, 57, 0, 0, 252, 254, 3, 22, 11,
		0, 253, 252, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255,
		256, 5, 32, 0, 0, 256, 258, 3, 62, 31, 0, 257, 255, 1, 0, 0, 0, 257, 258,
		1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 260, 5, 56, 0, 0, 260, 33, 1, 0,
		0, 0, 261, 273, 3, 32, 16, 0, 262, 273, 3, 14, 7, 0, 263, 273, 3, 36, 18,
		0, 264, 273, 3, 38, 19, 0, 265, 273, 3, 40, 20, 0, 266, 273, 3, 50, 25,
		0, 267, 273, 3, 52, 26, 0, 268, 273, 3, 54, 27, 0, 269, 273, 3, 60, 30,
		0, 270, 273, 3, 88, 44, 0, 271, 273, 3, 42, 21, 0, 272, 261, 1, 0, 0, 0,
		272, 262, 1, 0, 0, 0, 272, 263, 1, 0, 0, 0, 272, 264, 1, 0, 0, 0, 272,
		265, 1, 0, 0, 0, 272, 266, 1, 0, 0, 0, 272, 267, 1, 0, 0, 0, 272, 268,
		1, 0, 0, 0, 272, 269, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 271, 1, 0,
		0, 0, 273, 35, 1, 0, 0, 0, 274, 275, 3, 62, 31, 0, 275, 276, 5, 56, 0,
		0, 276, 37, 1, 0, 0, 0, 277, 279, 5, 13, 0, 0, 278, 280, 5, 47, 0, 0, 279,
		278, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 283,
		3, 62, 31, 0, 282, 284, 5, 48, 0, 0, 283, 282, 1, 0, 0, 0, 283, 284, 1,
		0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 299, 3, 86, 43, 0, 286, 287, 5, 14,
		0, 0, 287, 289, 5, 13, 0, 0, 288, 290, 5, 47, 0, 0, 289, 288, 1, 0, 0,
		0, 289, 290, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 293, 3, 62, 31, 0,
		292, 294, 5, 48, 0, 0, 293, 292, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294,
		295, 1, 0, 0, 0, 295, 296, 3, 86, 43, 0, 296, 298, 1, 0, 0, 0, 297, 286,
		1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 299, 300, 1, 0,
		0, 0, 300, 304, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 302, 303, 5, 14, 0, 0,
		303, 305, 3, 86, 43, 0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305,
		39, 1, 0, 0, 0, 306, 308, 5, 16, 0, 0, 307, 309, 5, 47, 0, 0, 308, 307,
		1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 312, 3, 48,
		24, 0, 311, 313, 5, 48, 0, 0, 312, 311, 1, 0, 0, 0, 312, 313, 1, 0, 0,
		0, 313, 314, 1, 0, 0, 0, 314, 315, 3, 86, 43, 0, 315, 41, 1, 0, 0, 0, 316,
		317, 3, 44, 22, 0, 317, 318, 5, 56, 0, 0, 318, 43, 1, 0, 0, 0, 319, 320,
		5, 57, 0, 0, 320, 321, 3, 46, 23, 0, 321, 322, 3, 62, 31, 0, 322, 326,
		1, 0, 0, 0, 323, 324, 5, 57, 0, 0, 324, 326, 7, 1, 0, 0, 325, 319, 1, 0,
		0, 0, 325, 323, 1, 0, 0, 0, 326, 45, 1, 0, 0, 0, 327, 328, 7, 2, 0, 0,
		328, 47, 1, 0, 0, 0, 329, 334, 3, 32, 16, 0, 330, 331, 3, 44, 22, 0, 331,
		332, 5, 56, 0, 0, 332, 334, 1, 0, 0, 0, 333, 329, 1, 0, 0, 0, 333, 330,
		1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 338, 1, 0, 0, 0, 335, 336, 3, 62,
		31, 0, 336, 337, 5, 56, 0, 0, 337, 339, 1, 0, 0, 0, 338, 335, 1, 0, 0,
		0, 338, 339, 1, 0, 0, 0, 339, 341, 1, 0, 0, 0, 340, 342, 3, 44, 22, 0,
		341, 340, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 49, 1, 0, 0, 0, 343, 344,
		5, 17, 0, 0, 344, 345, 5, 47, 0, 0, 345, 346, 3, 62, 31, 0, 346, 347, 5,
		48, 0, 0, 347, 348, 3, 86, 43, 0, 348, 51, 1, 0, 0, 0, 349, 350, 5, 15,
		0, 0, 350, 351, 3, 86, 43, 0, 351, 53, 1, 0, 0, 0, 352, 355, 3, 56, 28,
		0, 353, 355, 3, 58, 29, 0, 354, 352, 1, 0, 0, 0, 354, 353, 1, 0, 0, 0,
		355, 55, 1, 0, 0, 0, 356, 357, 5, 18, 0, 0, 357, 358, 5, 56, 0, 0, 358,
		57, 1, 0, 0, 0, 359, 360, 5, 20, 0, 0, 360, 361, 5, 56, 0, 0, 361, 59,
		1, 0, 0, 0, 362, 364, 5, 21, 0, 0, 363, 365, 3, 62, 31, 0, 364, 363, 1,
		0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 5, 56, 0,
		0, 367, 61, 1, 0, 0, 0, 368, 369, 3, 64, 32, 0, 369, 63, 1, 0, 0, 0, 370,
		375, 3, 66, 33, 0, 371, 372, 7, 3, 0, 0, 372, 374, 3, 66, 33, 0, 373, 371,
		1, 0, 0, 0, 374, 377, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375, 376, 1, 0,
		0, 0, 376, 65, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 383, 3, 68, 34, 0,
		379, 380, 7, 4, 0, 0, 380, 382, 3, 68, 34, 0, 381, 379, 1, 0, 0, 0, 382,
		385, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 67, 1,
		0, 0, 0, 385, 383, 1, 0, 0, 0, 386, 391, 3, 70, 35, 0, 387, 388, 7, 5,
		0, 0, 388, 390, 3, 70, 35, 0, 389, 387, 1, 0, 0, 0, 390, 393, 1, 0, 0,
		0, 391, 389, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 69, 1, 0, 0, 0, 393,
		391, 1, 0, 0, 0, 394, 399, 3, 72, 36, 0, 395, 396, 7, 6, 0, 0, 396, 398,
		3, 72, 36, 0, 397, 395, 1, 0, 0, 0, 398, 401, 1, 0, 0, 0, 399, 397, 1,
		0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 71, 1, 0, 0, 0, 401, 399, 1, 0, 0,
		0, 402, 403, 7, 7, 0, 0, 403, 406, 3, 72, 36, 0, 404, 406, 3, 74, 37, 0,
		405, 402, 1, 0, 0, 0, 405, 404, 1, 0, 0, 0, 406, 73, 1, 0, 0, 0, 407, 411,
		3, 80, 40, 0, 408, 410, 3, 76, 38, 0, 409, 408, 1, 0, 0, 0, 410, 413, 1,
		0, 0, 0, 411, 409, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 75, 1, 0, 0,
		0, 413, 411, 1, 0, 0, 0, 414, 415, 5, 51, 0, 0, 415, 416, 3, 62, 31, 0,
		416, 417, 5, 52, 0, 0, 417, 428, 1, 0, 0, 0, 418, 419, 5, 54, 0, 0, 419,
		428, 5, 57, 0, 0, 420, 428, 5, 30, 0, 0, 421, 428, 5, 31, 0, 0, 422, 424,
		5, 47, 0, 0, 423, 425, 3, 78, 39, 0, 424, 423, 1, 0, 0, 0, 424, 425, 1,
		0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 428, 5, 48, 0, 0, 427, 414, 1, 0, 0,
		0, 427, 418, 1, 0, 0, 0, 427, 420, 1, 0, 0, 0, 427, 421, 1, 0, 0, 0, 427,
		422, 1, 0, 0, 0, 428, 77, 1, 0, 0, 0, 429, 434, 3, 62, 31, 0, 430, 431,
		5, 53, 0, 0, 431, 433, 3, 62, 31, 0, 432, 430, 1, 0, 0, 0, 433, 436, 1,
		0, 0, 0, 434, 432, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 79, 1, 0, 0,
		0, 436, 434, 1, 0, 0, 0, 437, 446, 5, 57, 0, 0, 438, 446, 3, 84, 42, 0,
		439, 440, 5, 47, 0, 0, 440, 441, 3, 62, 31, 0, 441, 442, 5, 48, 0, 0, 442,
		446, 1, 0, 0, 0, 443, 446, 3, 82, 41, 0, 444, 446, 3, 16, 8, 0, 445, 437,
		1, 0, 0, 0, 445, 438, 1, 0, 0, 0, 445, 439, 1, 0, 0, 0, 445, 443, 1, 0,
		0, 0, 445, 444, 1, 0, 0, 0, 446, 81, 1, 0, 0, 0, 447, 456, 5, 51, 0, 0,
		448, 453, 3, 62, 31, 0, 449, 450, 5, 53, 0, 0, 450, 452, 3, 62, 31, 0,
		451, 449, 1, 0, 0, 0, 452, 455, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 453,
		454, 1, 0, 0, 0, 454, 457, 1, 0, 0, 0, 455, 453, 1, 0, 0, 0, 456, 448,
		1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 458, 1, 0, 0, 0, 458, 459, 5, 52,
		0, 0, 459, 83, 1, 0, 0, 0, 460, 461, 7, 8, 0, 0, 461, 85, 1, 0, 0, 0, 462,
		466, 5, 49, 0, 0, 463, 465, 3, 34, 17, 0, 464, 463, 1, 0, 0, 0, 465, 468,
		1, 0, 0, 0, 466, 464, 1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 469, 1, 0,
		0, 0, 468, 466, 1, 0, 0, 0, 469, 470, 5, 50, 0, 0, 470, 87, 1, 0, 0, 0,
		471, 472, 5, 22, 0, 0, 472, 473, 5, 47, 0, 0, 473, 474, 3, 62, 31, 0, 474,
		475, 5, 48, 0, 0, 475, 479, 5, 49, 0, 0, 476, 478, 3, 90, 45, 0, 477, 476,
		1, 0, 0, 0, 478, 481, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 479, 480, 1, 0,
		0, 0, 480, 483, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 482, 484, 3, 92, 46,
		0, 483, 482, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484, 485, 1, 0, 0, 0, 485,
		486, 5, 50, 0, 0, 486, 89, 1, 0, 0, 0, 487, 488, 5, 23, 0, 0, 488, 489,
		3, 62, 31, 0, 489, 490, 5, 55, 0, 0, 490, 491, 3, 86, 43, 0, 491, 91, 1,
		0, 0, 0, 492, 493, 5, 24, 0, 0, 493, 494, 5, 55, 0, 0, 494, 495, 3, 86,
		43, 0, 495, 93, 1, 0, 0, 0, 55, 97, 103, 111, 132, 141, 147, 149, 155,
		158, 164, 169, 173, 182, 188, 197, 205, 209, 214, 219, 228, 234, 242, 248,
		253, 257, 272, 279, 283, 289, 293, 299, 304, 308, 312, 325, 333, 338, 341,
		354, 364, 375, 383, 391, 399, 405, 411, 424, 427, 434, 445, 453, 456, 466,
		479, 483,
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
	R2D2ParserIMPORT         = 2
	R2D2ParserFROM           = 3
	R2D2ParserINTERFACE      = 4
	R2D2ParserMODULE         = 5
	R2D2ParserIMPLEMENTS     = 6
	R2D2ParserEXPORT         = 7
	R2D2ParserFN             = 8
	R2D2ParserPSEUDO         = 9
	R2D2ParserVAR            = 10
	R2D2ParserLET            = 11
	R2D2ParserCONST          = 12
	R2D2ParserIF             = 13
	R2D2ParserELSE           = 14
	R2D2ParserLOOP           = 15
	R2D2ParserFOR            = 16
	R2D2ParserWHILE          = 17
	R2D2ParserBREAK          = 18
	R2D2ParserSEND           = 19
	R2D2ParserCONTINUE       = 20
	R2D2ParserRETURN         = 21
	R2D2ParserSWITCH         = 22
	R2D2ParserCASE           = 23
	R2D2ParserDEFAULT        = 24
	R2D2ParserPLUS           = 25
	R2D2ParserMINUS          = 26
	R2D2ParserMULT           = 27
	R2D2ParserDIV            = 28
	R2D2ParserMOD            = 29
	R2D2ParserINCREMENT      = 30
	R2D2ParserDECREMENT      = 31
	R2D2ParserASSIGN         = 32
	R2D2ParserPLUS_ASSIGN    = 33
	R2D2ParserMINUS_ASSIGN   = 34
	R2D2ParserMULT_ASSIGN    = 35
	R2D2ParserDIV_ASSIGN     = 36
	R2D2ParserMOD_ASSIGN     = 37
	R2D2ParserEQ             = 38
	R2D2ParserNEQ            = 39
	R2D2ParserLT             = 40
	R2D2ParserGT             = 41
	R2D2ParserLEQ            = 42
	R2D2ParserGEQ            = 43
	R2D2ParserAND            = 44
	R2D2ParserOR             = 45
	R2D2ParserNOT            = 46
	R2D2ParserLPAREN         = 47
	R2D2ParserRPAREN         = 48
	R2D2ParserLBRACE         = 49
	R2D2ParserRBRACE         = 50
	R2D2ParserLBRACK         = 51
	R2D2ParserRBRACK         = 52
	R2D2ParserCOMMA          = 53
	R2D2ParserDOT            = 54
	R2D2ParserCOLON          = 55
	R2D2ParserSEMI           = 56
	R2D2ParserIDENTIFIER     = 57
	R2D2ParserTYPE           = 58
	R2D2ParserSTRING_LITERAL = 59
	R2D2ParserBOOL_LITERAL   = 60
	R2D2ParserINT_LITERAL    = 61
	R2D2ParserFLOAT_LITERAL  = 62
	R2D2ParserCOMMENT        = 63
	R2D2ParserBLOCK_COMMENT  = 64
	R2D2ParserWHITESPACE     = 65
)

// R2D2Parser rules.
const (
	R2D2ParserRULE_program                  = 0
	R2D2ParserRULE_declaration              = 1
	R2D2ParserRULE_globalDeclaration        = 2
	R2D2ParserRULE_importDeclaration        = 3
	R2D2ParserRULE_interfaceDeclaration     = 4
	R2D2ParserRULE_moduleDeclaration        = 5
	R2D2ParserRULE_functionDeclaration      = 6
	R2D2ParserRULE_functionCallStatement    = 7
	R2D2ParserRULE_functionCall             = 8
	R2D2ParserRULE_parameterList            = 9
	R2D2ParserRULE_parameter                = 10
	R2D2ParserRULE_typeExpression           = 11
	R2D2ParserRULE_arrayDimensions          = 12
	R2D2ParserRULE_baseType                 = 13
	R2D2ParserRULE_genericType              = 14
	R2D2ParserRULE_typeDeclaration          = 15
	R2D2ParserRULE_variableDeclaration      = 16
	R2D2ParserRULE_statement                = 17
	R2D2ParserRULE_expressionStatement      = 18
	R2D2ParserRULE_ifStatement              = 19
	R2D2ParserRULE_forStatement             = 20
	R2D2ParserRULE_assignmentDeclaration    = 21
	R2D2ParserRULE_assignment               = 22
	R2D2ParserRULE_assignmentOperator       = 23
	R2D2ParserRULE_simpleFor                = 24
	R2D2ParserRULE_whileStatement           = 25
	R2D2ParserRULE_loopStatement            = 26
	R2D2ParserRULE_cicleControl             = 27
	R2D2ParserRULE_breakStatement           = 28
	R2D2ParserRULE_continueStatement        = 29
	R2D2ParserRULE_returnStatement          = 30
	R2D2ParserRULE_expression               = 31
	R2D2ParserRULE_logicalExpression        = 32
	R2D2ParserRULE_comparisonExpression     = 33
	R2D2ParserRULE_additiveExpression       = 34
	R2D2ParserRULE_multiplicativeExpression = 35
	R2D2ParserRULE_unaryExpression          = 36
	R2D2ParserRULE_memberExpression         = 37
	R2D2ParserRULE_memberPart               = 38
	R2D2ParserRULE_argumentList             = 39
	R2D2ParserRULE_primaryExpression        = 40
	R2D2ParserRULE_arrayLiteral             = 41
	R2D2ParserRULE_literal                  = 42
	R2D2ParserRULE_block                    = 43
	R2D2ParserRULE_switchStatement          = 44
	R2D2ParserRULE_switchCase               = 45
	R2D2ParserRULE_defaultCase              = 46
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
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserIMPORT {
		{
			p.SetState(94)
			p.ImportDeclaration()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4144) != 0 {
		{
			p.SetState(100)
			p.Declaration()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
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
	InterfaceDeclaration() IInterfaceDeclarationContext
	GlobalDeclaration() IGlobalDeclarationContext

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

func (s *DeclarationContext) InterfaceDeclaration() IInterfaceDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInterfaceDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInterfaceDeclarationContext)
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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case R2D2ParserMODULE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.ModuleDeclaration()
		}

	case R2D2ParserINTERFACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.InterfaceDeclaration()
		}

	case R2D2ParserCONST:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.GlobalDeclaration()
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
		p.SetState(113)
		p.Match(R2D2ParserCONST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.TypeExpression()
	}
	{
		p.SetState(116)
		p.Match(R2D2ParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Expression()
	}
	{
		p.SetState(118)
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
	IMPORT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	FROM() antlr.TerminalNode
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

func (s *ImportDeclarationContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIMPORT, 0)
}

func (s *ImportDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *ImportDeclarationContext) FROM() antlr.TerminalNode {
	return s.GetToken(R2D2ParserFROM, 0)
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
		p.SetState(120)
		p.Match(R2D2ParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(R2D2ParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(R2D2ParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
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
	IDENTIFIER() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllFunctionDeclaration() []IFunctionDeclarationContext
	FunctionDeclaration(i int) IFunctionDeclarationContext

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

func (s *InterfaceDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(R2D2ParserIDENTIFIER, 0)
}

func (s *InterfaceDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACE, 0)
}

func (s *InterfaceDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACE, 0)
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
		p.SetState(126)
		p.Match(R2D2ParserINTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
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
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0 {
		{
			p.SetState(129)
			p.FunctionDeclaration()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(135)
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
		p.SetState(137)
		p.Match(R2D2ParserMODULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserIMPLEMENTS {
		{
			p.SetState(139)
			p.Match(R2D2ParserIMPLEMENTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(143)
		p.Match(R2D2ParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8066) != 0 {
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(144)
				p.FunctionDeclaration()
			}

		case 2:
			{
				p.SetState(145)
				p.TypeDeclaration()
			}

		case 3:
			{
				p.SetState(146)
				p.VariableDeclaration()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
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
	COLON() antlr.TerminalNode
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

func (s *FunctionDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(R2D2ParserCOLON, 0)
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
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserEXPORT {
		{
			p.SetState(154)
			p.Match(R2D2ParserEXPORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserPSEUDO {
		{
			p.SetState(157)
			p.Match(R2D2ParserPSEUDO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(160)
		p.Match(R2D2ParserFN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(R2D2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserIDENTIFIER {
		{
			p.SetState(163)
			p.ParameterList()
		}

	}
	{
		p.SetState(166)
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

	if _la == R2D2ParserCOLON {
		{
			p.SetState(167)
			p.Match(R2D2ParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8793489381961760768) != 0 {
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.TypeExpression()
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.BaseType()
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserLBRACK {
		{
			p.SetState(204)
			p.ArrayDimensions()
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
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == R2D2ParserLBRACK {
		{
			p.SetState(207)
			p.Match(R2D2ParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == R2D2ParserINT_LITERAL {
			{
				p.SetState(208)
				p.Match(R2D2ParserINT_LITERAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(211)
			p.Match(R2D2ParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(214)
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
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(R2D2ParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(218)
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
		p.SetState(221)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(R2D2ParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.TypeExpression()
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserCOMMA {
		{
			p.SetState(224)
			p.Match(R2D2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.TypeExpression()
		}

		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(231)
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
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserEXPORT {
		{
			p.SetState(233)
			p.Match(R2D2ParserEXPORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(236)
		p.Match(R2D2ParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.Match(R2D2ParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7296) != 0 {
		{
			p.SetState(239)
			p.VariableDeclaration()
		}

		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(245)
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
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserEXPORT {
		{
			p.SetState(247)
			p.Match(R2D2ParserEXPORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(250)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7168) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(251)
		p.Match(R2D2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserIDENTIFIER || _la == R2D2ParserTYPE {
		{
			p.SetState(252)
			p.TypeExpression()
		}

	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserASSIGN {
		{
			p.SetState(255)
			p.Match(R2D2ParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.Expression()
		}

	}
	{
		p.SetState(259)
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
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(261)
			p.VariableDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(262)
			p.FunctionCallStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(263)
			p.ExpressionStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(264)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(265)
			p.ForStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(266)
			p.WhileStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(267)
			p.LoopStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(268)
			p.CicleControl()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(269)
			p.ReturnStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(270)
			p.SwitchStatement()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(271)
			p.AssignmentDeclaration()
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
		p.SetState(274)
		p.Expression()
	}
	{
		p.SetState(275)
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
		p.SetState(277)
		p.Match(R2D2ParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(278)
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
		p.SetState(281)
		p.Expression()
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserRPAREN {
		{
			p.SetState(282)
			p.Match(R2D2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(285)
		p.Block()
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(286)
				p.Match(R2D2ParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(287)
				p.Match(R2D2ParserIF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(289)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(288)
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
				p.SetState(291)
				p.Expression()
			}
			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == R2D2ParserRPAREN {
				{
					p.SetState(292)
					p.Match(R2D2ParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(295)
				p.Block()
			}

		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserELSE {
		{
			p.SetState(302)
			p.Match(R2D2ParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.Block()
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
		p.SetState(306)
		p.Match(R2D2ParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(307)
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
		p.SetState(310)
		p.SimpleFor()
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserRPAREN {
		{
			p.SetState(311)
			p.Match(R2D2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(314)
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
		p.SetState(316)
		p.Assignment()
	}
	{
		p.SetState(317)
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
	Expression() IExpressionContext
	INCREMENT() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode

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

func (s *AssignmentContext) Expression() IExpressionContext {
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

func (s *AssignmentContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserINCREMENT, 0)
}

func (s *AssignmentContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(R2D2ParserDECREMENT, 0)
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

	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.AssignmentOperator()
		}
		{
			p.SetState(321)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
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
		p.SetState(327)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&270582939648) != 0) {
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
	p.SetState(333)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(329)
			p.VariableDeclaration()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(330)
			p.Assignment()
		}
		{
			p.SetState(331)
			p.Match(R2D2ParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(335)
			p.Expression()
		}
		{
			p.SetState(336)
			p.Match(R2D2ParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserIDENTIFIER {
		{
			p.SetState(340)
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
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext

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

func (s *WhileStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLPAREN, 0)
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

func (s *WhileStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRPAREN, 0)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.Match(R2D2ParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.Match(R2D2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Expression()
	}
	{
		p.SetState(346)
		p.Match(R2D2ParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
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
		p.SetState(349)
		p.Match(R2D2ParserLOOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(350)
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
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case R2D2ParserBREAK:
		{
			p.SetState(352)
			p.BreakStatement()
		}

	case R2D2ParserCONTINUE:
		{
			p.SetState(353)
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
		p.SetState(356)
		p.Match(R2D2ParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
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
		p.SetState(359)
		p.Match(R2D2ParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
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
		p.SetState(362)
		p.Match(R2D2ParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8793489381961760768) != 0 {
		{
			p.SetState(363)
			p.Expression()
		}

	}
	{
		p.SetState(366)
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

	// Getter signatures
	LogicalExpression() ILogicalExpressionContext

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

func (s *ExpressionContext) LogicalExpression() ILogicalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(R2D2Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case R2D2Visitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *R2D2Parser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, R2D2ParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.LogicalExpression()
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

// ILogicalExpressionContext is an interface to support dynamic dispatch.
type ILogicalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComparisonExpression() []IComparisonExpressionContext
	ComparisonExpression(i int) IComparisonExpressionContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsLogicalExpressionContext differentiates from other interfaces.
	IsLogicalExpressionContext()
}

type LogicalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalExpressionContext() *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_logicalExpression
	return p
}

func InitEmptyLogicalExpressionContext(p *LogicalExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_logicalExpression
}

func (*LogicalExpressionContext) IsLogicalExpressionContext() {}

func NewLogicalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_logicalExpression

	return p
}

func (s *LogicalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalExpressionContext) AllComparisonExpression() []IComparisonExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparisonExpressionContext); ok {
			len++
		}
	}

	tst := make([]IComparisonExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparisonExpressionContext); ok {
			tst[i] = t.(IComparisonExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalExpressionContext) ComparisonExpression(i int) IComparisonExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonExpressionContext); ok {
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

	return t.(IComparisonExpressionContext)
}

func (s *LogicalExpressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserAND)
}

func (s *LogicalExpressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserAND, i)
}

func (s *LogicalExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserOR)
}

func (s *LogicalExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserOR, i)
}

func (s *LogicalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *R2D2Parser) LogicalExpression() (localctx ILogicalExpressionContext) {
	localctx = NewLogicalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, R2D2ParserRULE_logicalExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.ComparisonExpression()
	}
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserAND || _la == R2D2ParserOR {
		{
			p.SetState(371)
			_la = p.GetTokenStream().LA(1)

			if !(_la == R2D2ParserAND || _la == R2D2ParserOR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(372)
			p.ComparisonExpression()
		}

		p.SetState(377)
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

// IComparisonExpressionContext is an interface to support dynamic dispatch.
type IComparisonExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAdditiveExpression() []IAdditiveExpressionContext
	AdditiveExpression(i int) IAdditiveExpressionContext
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllNEQ() []antlr.TerminalNode
	NEQ(i int) antlr.TerminalNode
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllLEQ() []antlr.TerminalNode
	LEQ(i int) antlr.TerminalNode
	AllGEQ() []antlr.TerminalNode
	GEQ(i int) antlr.TerminalNode

	// IsComparisonExpressionContext differentiates from other interfaces.
	IsComparisonExpressionContext()
}

type ComparisonExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonExpressionContext() *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_comparisonExpression
	return p
}

func InitEmptyComparisonExpressionContext(p *ComparisonExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_comparisonExpression
}

func (*ComparisonExpressionContext) IsComparisonExpressionContext() {}

func NewComparisonExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_comparisonExpression

	return p
}

func (s *ComparisonExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonExpressionContext) AllAdditiveExpression() []IAdditiveExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExpressionContext); ok {
			tst[i] = t.(IAdditiveExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonExpressionContext) AdditiveExpression(i int) IAdditiveExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
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

	return t.(IAdditiveExpressionContext)
}

func (s *ComparisonExpressionContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserEQ)
}

func (s *ComparisonExpressionContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserEQ, i)
}

func (s *ComparisonExpressionContext) AllNEQ() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserNEQ)
}

func (s *ComparisonExpressionContext) NEQ(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserNEQ, i)
}

func (s *ComparisonExpressionContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserLT)
}

func (s *ComparisonExpressionContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserLT, i)
}

func (s *ComparisonExpressionContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserGT)
}

func (s *ComparisonExpressionContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserGT, i)
}

func (s *ComparisonExpressionContext) AllLEQ() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserLEQ)
}

func (s *ComparisonExpressionContext) LEQ(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserLEQ, i)
}

func (s *ComparisonExpressionContext) AllGEQ() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserGEQ)
}

func (s *ComparisonExpressionContext) GEQ(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserGEQ, i)
}

func (s *ComparisonExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *R2D2Parser) ComparisonExpression() (localctx IComparisonExpressionContext) {
	localctx = NewComparisonExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, R2D2ParserRULE_comparisonExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.AdditiveExpression()
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17317308137472) != 0 {
		{
			p.SetState(379)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17317308137472) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(380)
			p.AdditiveExpression()
		}

		p.SetState(385)
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

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicativeExpression() []IMultiplicativeExpressionContext
	MultiplicativeExpression(i int) IMultiplicativeExpressionContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_additiveExpression
	return p
}

func InitEmptyAdditiveExpressionContext(p *AdditiveExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_additiveExpression
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExpressionContext); ok {
			tst[i] = t.(IMultiplicativeExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
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

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserPLUS)
}

func (s *AdditiveExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserPLUS, i)
}

func (s *AdditiveExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserMINUS)
}

func (s *AdditiveExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserMINUS, i)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *R2D2Parser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, R2D2ParserRULE_additiveExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.MultiplicativeExpression()
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserPLUS || _la == R2D2ParserMINUS {
		{
			p.SetState(387)
			_la = p.GetTokenStream().LA(1)

			if !(_la == R2D2ParserPLUS || _la == R2D2ParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(388)
			p.MultiplicativeExpression()
		}

		p.SetState(393)
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

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnaryExpression() []IUnaryExpressionContext
	UnaryExpression(i int) IUnaryExpressionContext
	AllMULT() []antlr.TerminalNode
	MULT(i int) antlr.TerminalNode
	AllDIV() []antlr.TerminalNode
	DIV(i int) antlr.TerminalNode
	AllMOD() []antlr.TerminalNode
	MOD(i int) antlr.TerminalNode

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_multiplicativeExpression
	return p
}

func InitEmptyMultiplicativeExpressionContext(p *MultiplicativeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_multiplicativeExpression
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) AllUnaryExpression() []IUnaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IUnaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryExpressionContext); ok {
			tst[i] = t.(IUnaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) UnaryExpression(i int) IUnaryExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
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

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicativeExpressionContext) AllMULT() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserMULT)
}

func (s *MultiplicativeExpressionContext) MULT(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserMULT, i)
}

func (s *MultiplicativeExpressionContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserDIV)
}

func (s *MultiplicativeExpressionContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserDIV, i)
}

func (s *MultiplicativeExpressionContext) AllMOD() []antlr.TerminalNode {
	return s.GetTokens(R2D2ParserMOD)
}

func (s *MultiplicativeExpressionContext) MOD(i int) antlr.TerminalNode {
	return s.GetToken(R2D2ParserMOD, i)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *R2D2Parser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, R2D2ParserRULE_multiplicativeExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		p.UnaryExpression()
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&939524096) != 0 {
		{
			p.SetState(395)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&939524096) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(396)
			p.UnaryExpression()
		}

		p.SetState(401)
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

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryExpression() IUnaryExpressionContext
	NOT() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	INCREMENT() antlr.TerminalNode
	DECREMENT() antlr.TerminalNode
	MemberExpression() IMemberExpressionContext

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = R2D2ParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = R2D2ParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
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

func (s *UnaryExpressionContext) MemberExpression() IMemberExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberExpressionContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *R2D2Parser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, R2D2ParserRULE_unaryExpression)
	var _la int

	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case R2D2ParserMINUS, R2D2ParserINCREMENT, R2D2ParserDECREMENT, R2D2ParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(402)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70372032512000) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(403)
			p.UnaryExpression()
		}

	case R2D2ParserLPAREN, R2D2ParserLBRACK, R2D2ParserIDENTIFIER, R2D2ParserSTRING_LITERAL, R2D2ParserBOOL_LITERAL, R2D2ParserINT_LITERAL, R2D2ParserFLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(404)
			p.MemberExpression()
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
	p.EnterRule(localctx, 74, R2D2ParserRULE_memberExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
		p.PrimaryExpression()
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&20406939032748032) != 0 {
		{
			p.SetState(408)
			p.MemberPart()
		}

		p.SetState(413)
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
	p.EnterRule(localctx, 76, R2D2ParserRULE_memberPart)
	var _la int

	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case R2D2ParserLBRACK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(414)
			p.Match(R2D2ParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.Expression()
		}
		{
			p.SetState(416)
			p.Match(R2D2ParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(418)
			p.Match(R2D2ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserINCREMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(420)
			p.Match(R2D2ParserINCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserDECREMENT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(421)
			p.Match(R2D2ParserDECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case R2D2ParserLPAREN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(422)
			p.Match(R2D2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8793489381961760768) != 0 {
			{
				p.SetState(423)
				p.ArgumentList()
			}

		}
		{
			p.SetState(426)
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
	p.EnterRule(localctx, 78, R2D2ParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		p.Expression()
	}
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserCOMMA {
		{
			p.SetState(430)
			p.Match(R2D2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Expression()
		}

		p.SetState(436)
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
	p.EnterRule(localctx, 80, R2D2ParserRULE_primaryExpression)
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(437)
			p.Match(R2D2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
			p.Literal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(439)
			p.Match(R2D2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(440)
			p.Expression()
		}
		{
			p.SetState(441)
			p.Match(R2D2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(443)
			p.ArrayLiteral()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(444)
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
	p.EnterRule(localctx, 82, R2D2ParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
		p.Match(R2D2ParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8793489381961760768) != 0 {
		{
			p.SetState(448)
			p.Expression()
		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == R2D2ParserCOMMA {
			{
				p.SetState(449)
				p.Match(R2D2ParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(450)
				p.Expression()
			}

			p.SetState(455)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(458)
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
	p.EnterRule(localctx, 84, R2D2ParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8646911284551352320) != 0) {
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
	p.EnterRule(localctx, 86, R2D2ParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.Match(R2D2ParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8793489381969607808) != 0 {
		{
			p.SetState(463)
			p.Statement()
		}

		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(469)
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
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
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

func (s *SwitchStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLPAREN, 0)
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

func (s *SwitchStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRPAREN, 0)
}

func (s *SwitchStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserLBRACE, 0)
}

func (s *SwitchStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(R2D2ParserRBRACE, 0)
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
	p.EnterRule(localctx, 88, R2D2ParserRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(471)
		p.Match(R2D2ParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(472)
		p.Match(R2D2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(473)
		p.Expression()
	}
	{
		p.SetState(474)
		p.Match(R2D2ParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(475)
		p.Match(R2D2ParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == R2D2ParserCASE {
		{
			p.SetState(476)
			p.SwitchCase()
		}

		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == R2D2ParserDEFAULT {
		{
			p.SetState(482)
			p.DefaultCase()
		}

	}
	{
		p.SetState(485)
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
	COLON() antlr.TerminalNode
	Block() IBlockContext

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

func (s *SwitchCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(R2D2ParserCOLON, 0)
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
	p.EnterRule(localctx, 90, R2D2ParserRULE_switchCase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(487)
		p.Match(R2D2ParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(488)
		p.Expression()
	}
	{
		p.SetState(489)
		p.Match(R2D2ParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(490)
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

// IDefaultCaseContext is an interface to support dynamic dispatch.
type IDefaultCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Block() IBlockContext

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

func (s *DefaultCaseContext) COLON() antlr.TerminalNode {
	return s.GetToken(R2D2ParserCOLON, 0)
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
	p.EnterRule(localctx, 92, R2D2ParserRULE_defaultCase)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.Match(R2D2ParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(493)
		p.Match(R2D2ParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(494)
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
