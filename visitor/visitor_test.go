package visitor

import (
	"testing"

	"github.com/ArturC03/r2d2/parser"
	// "github.com/ArturC03/r2d2/visitor"
	"github.com/antlr4-go/antlr/v4"
)

// Helper function to create a visitor and parse input
func setupVisitorWithInput(t *testing.T, input string) *R2D2Visitor {
	// Create the input stream
	inputStream := antlr.NewInputStream(input)

	// Create lexer
	lexer := parser.NewR2D2Lexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, 0)

	// Create parser and parse input
	p := parser.NewR2D2Parser(tokens)
	p.RemoveErrorListeners()
	tree := p.Program()

	// Create and use visitor
	v := NewR2D2Visitor()
	tree.Accept(v)

	return v
}

// TestNewR2D2Visitor tests the visitor creation
func TestNewR2D2Visitor(t *testing.T) {
	v := NewR2D2Visitor()

	if v == nil {
		t.Fatalf("Expected non-nil visitor")
	}

	if v.JsCode != "" {
		t.Errorf("Expected empty JsCode, got %q", v.JsCode)
	}

	if len(v.SymbolTable().Modules) != 0 {
		t.Errorf("Expected no modules in new visitor, got %d", len(v.SymbolTable().Modules))
	}
}

// TestEmptyProgram tests parsing an empty program
func TestEmptyProgram(t *testing.T) {
	v := setupVisitorWithInput(t, "")

	if v.JsCode != "" {
		t.Errorf("Expected empty JsCode for empty program, got %q", v.JsCode)
	}
}

// Expose the symbol table for testing
func (v *R2D2Visitor) SymbolTable() SymbolTable {
	return v.symbolTable
}
