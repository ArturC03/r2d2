package tests

import (
	"github.com/ArturC03/r2d2/lexer" // Caminho correto do módulo
	"testing"
)

// Função auxiliar para rodar o lexer e comparar com os tokens esperados
func runLexerWithExpectedTokens(input string, expectedTokens []lexer.TokenType, t *testing.T) {
	l := lexer.New(input) // Inicializa o lexer

	for i, expected := range expectedTokens {
		tok := l.NextToken()

		if tok.Type != expected {
			t.Errorf("Token esperado: %v, mas obtido: %v (na posição %d)", expected, tok.Type, i+1)
		}
	}
}

// Teste de lexer básico com atribuição
func TestLexerBasic(t *testing.T) {
	input := `let x = 5;`
	expectedTokens := []lexer.TokenType{
		lexer.KeywordLet, lexer.Identifier, lexer.Assign, lexer.Integer, lexer.Semicolon, lexer.EOF,
	}

	t.Run("Atribuição simples", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de operadores
func TestLexerOperators(t *testing.T) {
	input := `x + y - z * (a / b) == 10 != 20;`
	expectedTokens := []lexer.TokenType{
		lexer.Identifier, lexer.Plus, lexer.Identifier, lexer.Minus, lexer.Identifier,
		lexer.Multiply, lexer.LeftParen, lexer.Identifier, lexer.Divide, lexer.Identifier,
		lexer.RightParen, lexer.Equals, lexer.Integer, lexer.NotEquals, lexer.Integer, lexer.Semicolon, lexer.EOF,
	}

	t.Run("Operadores básicos", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de comentários de linha
func TestLexerCommentLine(t *testing.T) {
	input := `let x = 5; // Isto é um comentário`
	expectedTokens := []lexer.TokenType{
		lexer.KeywordLet, lexer.Identifier, lexer.Assign, lexer.Integer, lexer.Semicolon, lexer.EOF,
	}

	t.Run("Comentário de linha", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de comentários de bloco
func TestLexerCommentBlock(t *testing.T) {
	input := `let x = 5; /* Comentário de bloco */`
	expectedTokens := []lexer.TokenType{
		lexer.KeywordLet, lexer.Identifier, lexer.Assign, lexer.Integer, lexer.Semicolon, lexer.EOF,
	}

	t.Run("Comentário de bloco", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de strings
func TestLexerString(t *testing.T) {
	input := `"Olá, Mundo!"`
	expectedTokens := []lexer.TokenType{
		lexer.String, lexer.EOF,
	}

	t.Run("String simples", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de números
func TestLexerNumber(t *testing.T) {
	input := `12345`
	expectedTokens := []lexer.TokenType{
		lexer.Integer, lexer.EOF,
	}

	t.Run("Número inteiro", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de erro ao ter letras após números
func TestLexerNumberWithLetters(t *testing.T) {
	input := `123abc`
	expectedTokens := []lexer.TokenType{
		lexer.Integer, lexer.Illegal, lexer.EOF,
	}

	t.Run("Número com letras (erro esperado)", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de palavras-chave
func TestLexerKeywords(t *testing.T) {
	input := `let x = 5; if x > 10 { return true; } else { return false; }`
	expectedTokens := []lexer.TokenType{
		lexer.KeywordLet, lexer.Identifier, lexer.Assign, lexer.Integer, lexer.Semicolon,
		lexer.KeywordIf, lexer.Identifier, lexer.GreaterThan, lexer.Integer, lexer.LeftBrace,
		lexer.KeywordReturn, lexer.KeywordTrue, lexer.Semicolon, lexer.KeywordElse,
		lexer.LeftBrace, lexer.KeywordReturn, lexer.KeywordFalse, lexer.Semicolon, lexer.RightBrace, lexer.RightBrace, lexer.EOF,
	}

	t.Run("Palavras-chave", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de caractere ilegal
func TestLexerIllegalCharacter(t *testing.T) {
	input := `let x = 5; $var = 10;`
	expectedTokens := []lexer.TokenType{
		lexer.KeywordLet, lexer.Identifier, lexer.Assign, lexer.Integer, lexer.Semicolon,
		lexer.Illegal, lexer.Identifier, lexer.Assign, lexer.Integer, lexer.Semicolon, lexer.EOF,
	}

	t.Run("Caractere ilegal '$'", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de múltiplas palavras-chave
func TestLexerMultipleKeywords(t *testing.T) {
	input := `let var if else fn module`
	expectedTokens := []lexer.TokenType{
		lexer.KeywordLet, lexer.KeywordVar, lexer.KeywordIf, lexer.KeywordElse, lexer.KeywordFn, lexer.KeywordModule, lexer.EOF,
	}

	t.Run("Múltiplas palavras-chave", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de comentário aninhado (deve gerar erro)
func TestLexerNestedComments(t *testing.T) {
	input := `/* Início comentário /* comentário interno */ Fim comentário */`
	expectedTokens := []lexer.TokenType{
		lexer.Illegal, lexer.EOF,
	}

	t.Run("Comentário aninhado (erro esperado)", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de string não finalizada
func TestLexerUnfinishedString(t *testing.T) {
	input := `"string não finalizada`
	expectedTokens := []lexer.TokenType{
		lexer.Illegal, lexer.EOF,
	}

	t.Run("String não finalizada (erro esperado)", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de comentário de bloco não finalizado
func TestLexerUnfinishedBlockComment(t *testing.T) {
	input := `/* Comentário não finalizado`
	expectedTokens := []lexer.TokenType{
		lexer.Illegal, lexer.EOF,
	}

	t.Run("Comentário de bloco não finalizado (erro esperado)", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}

// Teste de expressão complexa
func TestLexerComplexExpression(t *testing.T) {
	input := `let x = (a + b) * (c / d) - (e && f);`
	expectedTokens := []lexer.TokenType{
		lexer.KeywordLet, lexer.Identifier, lexer.Assign, lexer.LeftParen, lexer.Identifier,
		lexer.Plus, lexer.Identifier, lexer.RightParen, lexer.Multiply, lexer.LeftParen, lexer.Identifier,
		lexer.Divide, lexer.Identifier, lexer.RightParen, lexer.Minus, lexer.LeftParen, lexer.Identifier,
		lexer.Identifier, lexer.RightParen, lexer.Semicolon, lexer.EOF,
	}

	t.Run("Expressão complexa", func(t *testing.T) {
		runLexerWithExpectedTokens(input, expectedTokens, t)
	})
}
