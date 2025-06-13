package visitor

import (
	"strings"
	"testing"
)

// TestLiteralExpressions tests literal expressions translation
func TestLiteralExpressions(t *testing.T) {
	input := `
	module Test {
		fn literals() {
			var num = 42;
			var pi = 3.14;
			var message = "hello";
			var flag = true;
			var empty = null;
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check JS code for literal translations
	jsCode := v.JsCode

	if !strings.Contains(jsCode, "var num = 42") {
		t.Errorf("Expected integer literal in JS code")
	}

	if !strings.Contains(jsCode, "var pi = 3.14") {
		t.Errorf("Expected float literal in JS code")
	}

	if !strings.Contains(jsCode, "var message = \"hello\"") {
		t.Errorf("Expected string literal in JS code")
	}

	if !strings.Contains(jsCode, "var flag = true") {
		t.Errorf("Expected boolean literal in JS code")
	}

	if !strings.Contains(jsCode, "var empty = null") {
		t.Errorf("Expected null literal in JS code")
	}
}

// TestArithmeticExpressions tests arithmetic expressions
func TestArithmeticExpressions(t *testing.T) {
	input := `
	module Test {
		fn arithmetic() {
			var sum = 5 + 3;
			var diff = 10 - 2;
			var product = 4 * 2;
			var quotient = 20 / 5;
			var remainder = 10 % 3;
			var complex = (2 + 3) * 4;
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check JS code for arithmetic expressions
	jsCode := v.JsCode

	if !strings.Contains(jsCode, "var sum = 5+3") {
		t.Errorf("Expected addition in JS code")
	}

	if !strings.Contains(jsCode, "var diff = 10-2") {
		t.Errorf("Expected subtraction in JS code")
	}

	if !strings.Contains(jsCode, "var product = 4*2") {
		t.Errorf("Expected multiplication in JS code")
	}

	if !strings.Contains(jsCode, "var quotient = 20/5") {
		t.Errorf("Expected division in JS code")
	}

	if !strings.Contains(jsCode, "var remainder = 10%3") {
		t.Errorf("Expected modulo in JS code")
	}

	if !strings.Contains(jsCode, "var complex = (2+3)*4") {
		t.Errorf("Expected complex expression in JS code")
	}
}

// TestLogicalExpressions tests logical expressions
func TestLogicalExpressions(t *testing.T) {
	input := `
	module Test {
		fn logical() {
			var andResult = true && false;
			var orResult = true || false;
			var notResult = !true;
			var complex = (a > b) && (c < d); 
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check JS code for logical expressions
	jsCode := v.JsCode

	if !strings.Contains(jsCode, "var andResult = true&&false") {
		t.Errorf("Expected AND expression in JS code")
	}

	if !strings.Contains(jsCode, "var orResult = true||false") {
		t.Errorf("Expected OR expression in JS code")
	}

	if !strings.Contains(jsCode, "var notResult = !true") {
		t.Errorf("Expected NOT expression in JS code")
	}

	// Complex expression depends on variable declarations, might be flagged as error
}

// TestComparisonExpressions tests comparison expressions
func TestComparisonExpressions(t *testing.T) {
	input := `
	module Test {
		fn comparisons() {
			var equal = 2 == 2;
			var notEqual = 2 != 3;
			var less = 5 < 10;
			var greater = 10 > 5;
			var lessEqual = 5 <= 5;
			var greaterEqual = 5 >= 5;
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check JS code for comparison expressions
	jsCode := v.JsCode

	if !strings.Contains(jsCode, "var equal = 2==2") {
		t.Errorf("Expected equality comparison in JS code")
	}

	if !strings.Contains(jsCode, "var notEqual = 2!=3") {
		t.Errorf("Expected inequality comparison in JS code")
	}

	if !strings.Contains(jsCode, "var less = 5<10") {
		t.Errorf("Expected less than comparison in JS code")
	}

	if !strings.Contains(jsCode, "var greater = 10>5") {
		t.Errorf("Expected greater than comparison in JS code")
	}

	if !strings.Contains(jsCode, "var lessEqual = 5<=5") {
		t.Errorf("Expected less than or equal comparison in JS code")
	}

	if !strings.Contains(jsCode, "var greaterEqual = 5>=5") {
		t.Errorf("Expected greater than or equal comparison in JS code")
	}
}
