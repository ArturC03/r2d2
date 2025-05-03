package visitor

import (
	"strings"
	"testing"
)

// TestBasicModuleDeclaration tests a simple module declaration
func TestBasicModuleDeclaration(t *testing.T) {
	input := `
	module Test {
		fn hello() {
			// Empty function
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check if module was created
	if _, exists := v.SymbolTable().Modules["Test"]; !exists {
		t.Fatalf("Module 'Test' not found in symbol table")
	}

	// Check if function was registered
	module := v.SymbolTable().Modules["Test"]
	if _, exists := module.Functions["hello"]; !exists {
		t.Errorf("Function 'hello' not found in module 'Test'")
	}

	// Check if JS code was generated
	if !strings.Contains(v.JsCode, "const Test = (function ()") {
		t.Errorf("Expected module initialization in JS code, got: %s", v.JsCode)
	}

	if !strings.Contains(v.JsCode, "function hello(") {
		t.Errorf("Expected function declaration in JS code, got: %s", v.JsCode)
	}
}

// TestModuleWithExports tests module with exported items
func TestModuleWithExports(t *testing.T) {
	input := `
	module ExportTest {
		export fn publicFn() {
			// Public function
		}

		fn privateFn() {
			// Private function
		}

		export var publicVar = 42;
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check if module was created
	module, exists := v.SymbolTable().Modules["ExportTest"]
	if !exists {
		t.Fatalf("Module 'ExportTest' not found in symbol table")
	}

	// Check exports
	exports := module.Exports()

	// Should contain publicFn and publicVar
	found := make(map[string]bool)
	for _, name := range exports {
		found[name] = true
	}

	if !found["publicFn"] {
		t.Errorf("Expected 'publicFn' to be exported")
	}

	if !found["publicVar"] {
		t.Errorf("Expected 'publicVar' to be exported")
	}

	if found["privateFn"] {
		t.Errorf("Expected 'privateFn' to NOT be exported")
	}

	// Check module return statement
	if !strings.Contains(v.JsCode, "return {publicFn, publicVar}") {
		t.Errorf("Expected export statement in JS code, got: %s", v.JsCode)
	}
}

// TestMultipleModules tests defining multiple modules
func TestMultipleModules(t *testing.T) {
	input := `
	module First {
		export fn hello() {
			// Hello function
		}
	}

	module Second {
		export fn world() {
			// World function
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check if both modules were created
	if _, exists := v.SymbolTable().Modules["First"]; !exists {
		t.Fatalf("Module 'First' not found in symbol table")
	}

	if _, exists := v.SymbolTable().Modules["Second"]; !exists {
		t.Fatalf("Module 'Second' not found in symbol table")
	}

	// Check functions in each module
	if _, exists := v.SymbolTable().Modules["First"].Functions["hello"]; !exists {
		t.Errorf("Function 'hello' not found in module 'First'")
	}

	if _, exists := v.SymbolTable().Modules["Second"].Functions["world"]; !exists {
		t.Errorf("Function 'world' not found in module 'Second'")
	}

	// Check if JS code contains both module declarations
	if !strings.Contains(v.JsCode, "const First = (function ()") {
		t.Errorf("Expected First module initialization in JS code")
	}

	if !strings.Contains(v.JsCode, "const Second = (function ()") {
		t.Errorf("Expected Second module initialization in JS code")
	}
}
