package visitor

import (
	"strings"
	"testing"
)

// TestIfStatement tests if-else statements
func TestIfStatement(t *testing.T) {
	input := `
	module Test {
		fn conditional(x number) {
			if (x > 10) {
				return "greater";
			} else if (x == 10) {
				return "equal";
			} else {
				return "less";
			}
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check JS code for if statement
	jsCode := v.JsCode

	if !strings.Contains(jsCode, "if (x>10)") {
		t.Errorf("Expected if condition in JS code")
	}

	if !strings.Contains(jsCode, "return \"greater\"") {
		t.Errorf("Expected return statement in if block")
	}

	if !strings.Contains(jsCode, "} else {") {
		t.Errorf("Expected else statement in JS code")
	}
}

// TestLoopStatement tests loop statements
func TestLoopStatement(t *testing.T) {
	input := `
	module Test {
		fn infiniteLoop() {
			loop {
				if (true) {
					break;
				}
			}
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check JS code for loop statement
	jsCode := v.JsCode

	if !strings.Contains(jsCode, "while (true)") {
		t.Errorf("Expected infinite loop in JS code")
	}

	if !strings.Contains(jsCode, "break") {
		t.Errorf("Expected break statement in loop")
	}
}

// TestForStatement tests for statements
func TestForStatement(t *testing.T) {
	input := `
	module Test {
		fn forLoop() {
			for (var i = 0; i < 10; i++) {
				// Loop body
			}
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check JS code for for statement
	jsCode := v.JsCode

	if !strings.Contains(jsCode, "for (var i = 0; i<10; i++)") {
		t.Errorf("Expected for loop in JS code")
	}
}

// TestWhileStatement tests while statements
func TestWhileStatement(t *testing.T) {
	input := `
	module Test {
		fn whileLoop() {
			while (true) {
				// Loop body
			}
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check JS code for while statement
	jsCode := v.JsCode

	if !strings.Contains(jsCode, "while (true)") {
		t.Errorf("Expected while loop in JS code")
	}
}

// TestSwitchStatement tests switch statements
func TestSwitchStatement(t *testing.T) {
	input := `
	module Test {
		fn switchCase(value number) {
			switch (value) {
				case 1: {
					return "one";
				}
				case 2: {
					return "two";
				}
				default: {
					return "other";
				}
			}
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check JS code for switch statement
	jsCode := v.JsCode

	if !strings.Contains(jsCode, "switch (value)") {
		t.Errorf("Expected switch statement in JS code")
	}

	if !strings.Contains(jsCode, "case 1:") {
		t.Errorf("Expected case 1 in switch statement")
	}

	if !strings.Contains(jsCode, "case 2:") {
		t.Errorf("Expected case 2 in switch statement")
	}

	if !strings.Contains(jsCode, "default:") {
		t.Errorf("Expected default case in switch statement")
	}
}

// TestJsBlock tests JS block insertion
func TestJsBlock(t *testing.T) {
	input := `
	module Test {
		fn withJsBlock() {
		@js <<
			console.log("This is raw JavaScript");
			const x = document.getElementById("test");
			>>;
		}
	}
	`

	v := setupVisitorWithInput(t, input)

	// Check JS code for raw JS block
	jsCode := v.JsCode

	if !strings.Contains(jsCode, "console.log(\"This is raw JavaScript\")") {
		t.Errorf("Expected raw JS code in output")
	}

	if !strings.Contains(jsCode, "const x = document.getElementById(\"test\")") {
		t.Errorf("Expected raw JS code in output")
	}
}
