package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
		module example {
			let x = [1, 2, 3];
			export fn pseudo() {}
		}
	`

	// Tokens esperados
	expectedTokens := []Token{
		{Type: KeywordModule, Lexeme: "module", Line: 2, Column: 2},
		{Type: Identifier, Lexeme: "example", Line: 2, Column: 9},
		{Type: LeftBrace, Lexeme: "{", Line: 2, Column: 17},
		{Type: KeywordLet, Lexeme: "let", Line: 3, Column: 3},
		{Type: Identifier, Lexeme: "x", Line: 3, Column: 7},
		{Type: Assign, Lexeme: "=", Line: 3, Column: 9},
		{Type: LeftBracket, Lexeme: "[", Line: 3, Column: 11},
		{Type: Integer, Lexeme: "1", Line: 3, Column: 12},
		{Type: Comma, Lexeme: ",", Line: 3, Column: 13},
		{Type: Integer, Lexeme: "2", Line: 3, Column: 15},
		{Type: Comma, Lexeme: ",", Line: 3, Column: 16},
		{Type: Integer, Lexeme: "3", Line: 3, Column: 18},
		{Type: RightBracket, Lexeme: "]", Line: 3, Column: 19},
		{Type: Semicolon, Lexeme: ";", Line: 3, Column: 20},
		{Type: KeywordExport, Lexeme: "export", Line: 4, Column: 3},
		{Type: KeywordFn, Lexeme: "fn", Line: 4, Column: 10},
		{Type: Identifier, Lexeme: "pseudo", Line: 4, Column: 13},
		{Type: LeftParen, Lexeme: "(", Line: 4, Column: 19},
		{Type: RightParen, Lexeme: ")", Line: 4, Column: 20},
		{Type: LeftBrace, Lexeme: "{", Line: 4, Column: 22},
		{Type: RightBrace, Lexeme: "}", Line: 4, Column: 23},
		{Type: RightBrace, Lexeme: "}", Line: 5, Column: 1},
		{Type: EOF, Lexeme: "", Line: 5, Column: 2},
	}

	lexer := New(input)

	for i, expected := range expectedTokens {
		tok := lexer.NextToken()

		// Comparar cada campo
		if tok.Type != expected.Type {
			t.Fatalf("Erro no token %d - Tipo errado. Esperado %q, obteve %q", i, expected.Type, tok.Type)
		}
		if tok.Lexeme != expected.Lexeme {
			t.Fatalf("Erro no token %d - Lexeme errado. Esperado %q, obteve %q", i, expected.Lexeme, tok.Lexeme)
		}
		if tok.Line != expected.Line {
			t.Fatalf("Erro no token %d - Linha errada. Esperado %d, obteve %d", i, expected.Line, tok.Line)
		}
		if tok.Column != expected.Column {
			t.Fatalf("Erro no token %d - Coluna errada. Esperado %d, obteve %d", i, expected.Column, tok.Column)
		}
	}
}
