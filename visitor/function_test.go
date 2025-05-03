package visitor

import (
	"strings"
	"testing"
)

// TestFunctionDeclaration tests basic function declaration
func TestFunctionDeclaration(t *testing.T) {
	input := `
	module Test {
		fn simple() {
			// Empty function
		}

		fn withParams(x number, y string) {
			// Function with params
		}

		fn withReturn() {
			return 42;
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check if functions were registered
	module, exists := v.SymbolTable().Modules["Test"]
	if !exists {
		t.Fatalf("Module 'Test' not found in symbol table")
	}

	// Check simple function
	if _, exists := module.Functions["simple"]; !exists {
		t.Errorf("Function 'simple' not found in module")
	}

	// Check function with parameters
	withParams, exists := module.Functions["withParams"]
	if !exists {
		t.Errorf("Function 'withParams' not found in module")
	} else {
		if len(withParams.Arguments) != 2 {
			t.Errorf("Expected 2 parameters in 'withParams', got %d", len(withParams.Arguments))
		} else {
			xArg, xExists := withParams.Arguments["x"]
			if !xExists || xArg.Type != "number" {
				if !xExists {
					t.Errorf("Expected parameter 'x' not found")
				} else {
					t.Errorf("Expected parameter x:number, got x:%s", xArg.Type)
				}
			}

			yArg, yExists := withParams.Arguments["y"]
			if !yExists || yArg.Type != "string" {
				if !yExists {
					t.Errorf("Expected parameter 'y' not found")
				} else {
					t.Errorf("Expected parameter y:string, got y:%s", yArg.Type)
				}
			}
		}
	}

	// Check JS code generation
	if !strings.Contains(v.JsCode, "function simple()") {
		t.Errorf("Expected simple function declaration in JS code")
	}

	if !strings.Contains(v.JsCode, "function withParams(x, y)") {
		t.Errorf("Expected withParams function declaration in JS code")
	}

	if !strings.Contains(v.JsCode, "return 42") {
		t.Errorf("Expected return statement in withReturn function")
	}
}

// TestPseudoFunction tests pseudo function declaration and validation
func TestPseudoFunction(t *testing.T) {
	input := `
	module Test {
		fn someFunction() {
			// Function body
		}
		fn anotherFunction() {
			// Function body
		}
		pseudo fn onlyFunctionCalls() {
			someFunction();
			anotherFunction();
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check if function was registered
	module, exists := v.SymbolTable().Modules["Test"]
	if !exists {
		t.Fatalf("Module 'Test' not found in symbol table")
	}

	// Check if function exists and is marked as pseudo
	pseudoFn, exists := module.Functions["onlyFunctionCalls"]
	if !exists {
		t.Errorf("Function 'onlyFunctionCalls' not found in module")
	} else if !pseudoFn.IsPseudo() {
		t.Errorf("Expected 'onlyFunctionCalls' to be marked as pseudo")
	}

	// JS code should still be generated
	if !strings.Contains(v.JsCode, "function onlyFunctionCalls()") {
		t.Errorf("Expected pseudo function declaration in JS code")
	}
}

// TestFunctionCall tests function call translation
func TestFunctionCall(t *testing.T) {
	input := `
	module Test {
	fn print() {
		// Print function
	}
		fn caller() {
			print("Hello");
			helper();
		}

		fn helper() {
			// Helper function
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check if functions were registered
	module, exists := v.SymbolTable().Modules["Test"]
	if !exists {
		t.Fatalf("Module 'Test' not found in symbol table")
	}

	if _, exists := module.Functions["caller"]; !exists {
		t.Errorf("Function 'caller' not found in module")
	}

	if _, exists := module.Functions["helper"]; !exists {
		t.Errorf("Function 'helper' not found in module")
	}

	// JS code should contain function calls
	if !strings.Contains(v.JsCode, "print(\"Hello\")") {
		t.Errorf("Expected function call to print in JS code")
	}

	if !strings.Contains(v.JsCode, "helper()") {
		t.Errorf("Expected function call to Test.helper in JS code")
	}
}

// Expose IsPseudo for testing
func (f Function) IsPseudo() bool {
	return f.isPseudo
}
