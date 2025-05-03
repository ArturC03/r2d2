package visitor

import (
	"strings"
	"testing"
)

// TestVariableDeclaration tests variable declarations
func TestVariableDeclaration(t *testing.T) {
	input := `
	module Test {
		var simpleVar = 42;
		let stringVar string = "hello";
		const boolVar boolean = true;

		fn withLocalVars() {
			var localNum = 10;
			let localStr string = "local";
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check module variables
	module, exists := v.SymbolTable().Modules["Test"]
	if !exists {
		t.Fatalf("Module 'Test' not found in symbol table")
	}

	// Check simpleVar
	if _, exists := module.Variables["simpleVar"]; !exists {
		t.Errorf("Variable 'simpleVar' not found in module")
	}

	// Check stringVar
	stringVar, exists := module.Variables["stringVar"]
	if !exists {
		t.Errorf("Variable 'stringVar' not found in module")
	} else if stringVar.Type != "string" {
		t.Errorf("Expected stringVar to have type string, got %s", stringVar.Type)
	}

	// Check boolVar
	boolVar, exists := module.Variables["boolVar"]
	if !exists {
		t.Errorf("Variable 'boolVar' not found in module")
	} else if boolVar.Type != "boolean" {
		t.Errorf("Expected boolVar to have type boolean, got %s", boolVar.Type)
	}

	// Check function local variables
	withLocalVars, exists := module.Functions["withLocalVars"]
	if !exists {
		t.Errorf("Function 'withLocalVars' not found in module")
	} else {
		if _, exists := withLocalVars.Variables["localNum"]; !exists {
			t.Errorf("Local variable 'localNum' not found in function")
		}

		localStr, exists := withLocalVars.Variables["localStr"]
		if !exists {
			t.Errorf("Local variable 'localStr' not found in function")
		} else if localStr.Type != "string" {
			t.Errorf("Expected localStr to have type string, got %s", localStr.Type)
		}
	}

	// Check JS code generation
	if !strings.Contains(v.JsCode, "var simpleVar = 42") {
		t.Errorf("Expected simpleVar declaration in JS code")
	}

	if !strings.Contains(v.JsCode, "let stringVar = \"hello\"") {
		t.Errorf("Expected stringVar declaration in JS code")
	}

	if !strings.Contains(v.JsCode, "const boolVar = true") {
		t.Errorf("Expected boolVar declaration in JS code")
	}

	if !strings.Contains(v.JsCode, "var localNum = 10") {
		t.Errorf("Expected localNum declaration in JS code")
	}

	if !strings.Contains(v.JsCode, "let localStr = \"local\"") {
		t.Errorf("Expected localStr declaration in JS code")
	}
}

// TestGlobalVariables tests global variable declarations
func TestGlobalVariables(t *testing.T) {
	input := `
	const globalVar number = 100;

	module Test {
		fn useGlobal() {
			// Use global variable
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check global variable
	if _, exists := v.SymbolTable().Globals["globalVar"]; !exists {
		t.Errorf("Global variable 'globalVar' not found")
	}

	// Check JS code generation
	if !strings.Contains(v.JsCode, "const globalVar = 100") {
		t.Errorf("Expected globalVar declaration in JS code")
	}
}

// TestVariableAssignment tests variable assignments
func TestVariableAssignment(t *testing.T) {
	input := `
	module Test {
		var counter = 0;

		fn updateCounter() {
			counter = 1;
			counter += 5;
			counter++;
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check module variable
	module, exists := v.SymbolTable().Modules["Test"]
	if !exists {
		t.Fatalf("Module 'Test' not found in symbol table")
	}

	if _, exists := module.Variables["counter"]; !exists {
		t.Errorf("Variable 'counter' not found in module")
	}

	// Check JS code generation for assignments
	if !strings.Contains(v.JsCode, "counter = 1") {
		t.Errorf("Expected assignment in JS code")
	}

	if !strings.Contains(v.JsCode, "counter += 5") {
		t.Errorf("Expected compound assignment in JS code")
	}

	if !strings.Contains(v.JsCode, "counter++") {
		t.Errorf("Expected increment in JS code")
	}
}
