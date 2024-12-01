package tests

import (
	"github.com/ArturC03/r2d2/lexer" // Importa o pacote lexer da sua estrutura
	"testing"
)

func TestNextToken(t *testing.T) {
	// Exemplo de entrada que será analisada pelo lexer
	input := `
		module example {
			let x = [1, 2, 3];
			export fn pseudo() {}
		}
	`

	// Tokens esperados
	expectedTokens := []lexer.Token{
		{Type: lexer.KeywordModule, Lexeme: "module", Line: 2, Column: 2},
		{Type: lexer.Identifier, Lexeme: "example", Line: 2, Column: 9},
		{Type: lexer.LeftBrace, Lexeme: "{", Line: 2, Column: 17},
		{Type: lexer.KeywordLet, Lexeme: "let", Line: 3, Column: 3},
		{Type: lexer.Identifier, Lexeme: "x", Line: 3, Column: 7},
		{Type: lexer.Assign, Lexeme: "=", Line: 3, Column: 9},
		{Type: lexer.LeftBracket, Lexeme: "[", Line: 3, Column: 11},
		{Type: lexer.Integer, Lexeme: "1", Line: 3, Column: 12},
		{Type: lexer.Comma, Lexeme: ",", Line: 3, Column: 13},
		{Type: lexer.Integer, Lexeme: "2", Line: 3, Column: 15},
		{Type: lexer.Comma, Lexeme: ",", Line: 3, Column: 16},
		{Type: lexer.Integer, Lexeme: "3", Line: 3, Column: 18},
		{Type: lexer.RightBracket, Lexeme: "]", Line: 3, Column: 19},
		{Type: lexer.Semicolon, Lexeme: ";", Line: 3, Column: 20},
		{Type: lexer.KeywordExport, Lexeme: "export", Line: 4, Column: 3},
		{Type: lexer.KeywordFn, Lexeme: "fn", Line: 4, Column: 10},
		{Type: lexer.Identifier, Lexeme: "pseudo", Line: 4, Column: 13},
		{Type: lexer.LeftParen, Lexeme: "(", Line: 4, Column: 19},
		{Type: lexer.RightParen, Lexeme: ")", Line: 4, Column: 20},
		{Type: lexer.LeftBrace, Lexeme: "{", Line: 4, Column: 22},
		{Type: lexer.RightBrace, Lexeme: "}", Line: 4, Column: 23},
		{Type: lexer.RightBrace, Lexeme: "}", Line: 5, Column: 1},
		{Type: lexer.EOF, Lexeme: "", Line: 5, Column: 2},
	}

	// Inicializa o lexer com a entrada
	lex := lexer.New(input)

	// Itera sobre os tokens esperados e os compara com os tokens gerados
	for i, expected := range expectedTokens {
		tok := lex.NextToken()

		// Verifica se o tipo do token está correto
		if tok.Type != expected.Type {
			t.Fatalf("Erro no token %d - Tipo errado. Esperado %q, obteve %q", i, expected.Type, tok.Type)
		}
		// Verifica se o lexema está correto
		if tok.Lexeme != expected.Lexeme {
			t.Fatalf("Erro no token %d - Lexema errado. Esperado %q, obteve %q", i, expected.Lexeme, tok.Lexeme)
		}
		// Verifica se a linha está correta
		if tok.Line != expected.Line {
			t.Fatalf("Erro no token %d - Linha errada. Esperado %d, obteve %d", i, expected.Line, tok.Line)
		}
		// Verifica se a coluna está correta
		if tok.Column != expected.Column {
			t.Fatalf("Erro no token %d - Coluna errada. Esperado %d, obteve %d", i, expected.Column, tok.Column)
		}
	}
}
