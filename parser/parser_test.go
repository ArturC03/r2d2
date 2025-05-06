package parser

import (
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

// TestValidParsing tests that valid R2D2 code can be parsed without errors
func TestValidParsing(t *testing.T) {
	testCases := []struct {
		name string
		code string
	}{
		{
			name: "Empty Program",
			code: "",
		},
		{
			name: "Import Statement",
			code: "use \"console.r2d2\";\n",
		},
		{
			name: "Module Declaration",
			code: `module TestModule {
				export fn main() {
					return;
				}
			}`,
		},
		{
			name: "Function Declaration in Module",
			code: `module TestModule {
				fn test(a number, b string): number {
					return a;
				}
			}`,
		},
		{
			name: "Variable Declaration in Function",
			code: `module TestModule {
				fn main() {
					var x number = 10;
					let y string = "hello";
					const z boolean = true;
				}
			}`,
		},
		{
			name: "If Statement in Function",
			code: `module TestModule {
				fn main() {
					if (x > 10) {
						return true;
					} else if (x < 5) {
						return false;
					} else {
						return null;
					}
				}
			}`,
		},
		{
			name: "For Loop in Function",
			code: `module TestModule {
				fn main() {
					for (let i number = 0; i < 10; i++) {
						console.log(i);
					}
				}
			}`,
		},
		{
			name: "While Loop in Function",
			code: `module TestModule {
				fn main() {
					while (x > 0) {
						x--;
					}
				}
			}`,
		},
		{
			name: "Function Call in Function",
			code: `module TestModule {
				fn main() {
					console.log("Hello, World!");
				}
			}`,
		},
		{
			name: "Type Declaration in Module",
			code: `module TestModule {
				type Person {
					var name string;
					var age number;
				}

				fn main() {
					return;
				}
			}`,
		},
		{
			name: "Array Declaration in Function",
			code: `module TestModule {
				fn main() {
					let numbers number[] = [1, 2, 3, 4, 5];
				}
			}`,
		},
		{
			name: "Interface Declaration",
			code: `interface Printable {
				fn print(message string): void;
			}

			module TestModule {
				fn main() {
					return;
				}
			}`,
		},
		{
			name: "Switch Statement in Function",
			code: `module TestModule {
				fn main() {
					switch (day) {
						case 1: {
							return "Monday";
						}
						case 2: {
							return "Tuesday";
						}
						default: {
							return "Other day";
						}
					}
				}
			}`,
		},
		{
			name: "JavaScript Block in Function",
			code: `module TestModule {
				fn main() {
					@js << console.log("This is JavaScript code"); >>;
				}
			}`,
		},
		{
			name: "Compound expressions in Function",
			code: `module TestModule {
				fn main() {
					let result number = (a + b) * (c - d) / e;
				}
			}`,
		},
		{
			name: "Loop Statement in Function",
			code: `module TestModule {
				fn main() {
					loop {
						if (x > 10) {
							break;
						}
						x++;
					}
				}
			}`,
		},
		// {
		// 	name: "Generic Types in Module",
		// 	code: `module TestModule {
		// 		type List<T> {
		// 			var items T[];
		// 		}

		// 		fn main() {
		// 			return;
		// 		}
		// 	}`,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := antlr.NewInputStream(tc.code)
			lexer := NewR2D2Lexer(input)
			tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

			p := NewR2D2Parser(tokenStream)

			// Add an error listener that will fail the test on any syntax error
			errorListener := &testErrorListener{t: t}
			p.RemoveErrorListeners()
			p.AddErrorListener(errorListener)

			// Try to parse the input
			p.Program()

			if errorListener.hasError {
				t.Errorf("Parser reported error on valid code: %s", tc.code)
			}
		})
	}
}

// TestInvalidParsing tests that invalid R2D2 code results in parsing errors
func TestInvalidParsing(t *testing.T) {
	testCases := []struct {
		name string
		code string
	}{
		{
			name: "Missing Semicolon in Function",
			code: `module TestModule {
				fn main() {
					let x number = 10
				}
			}`,
		},
		{
			name: "Mismatched Brackets in Function",
			code: `module TestModule {
				fn main() {
					if (x > 10) {
						return true;
				}
			}`,
		},
		{
			name: "Invalid Variable Declaration in Function",
			code: `module TestModule {
				fn main() {
					let = 10;
				}
			}`,
		},
		{
			name: "Incomplete Function Declaration",
			code: `module TestModule {
				fn test(a number,
			}`,
		},
		{
			name: "Invalid Expression in Function",
			code: `module TestModule {
				fn main() {
					let result number = 10 + * 5;
				}
			}`,
		},
		{
			name: "Missing Closing Tag in JS Block",
			code: `module TestModule {
				fn main() {
					@js << console.log("test");
				}
			}`,
		},
		{
			name: "Invalid If Statement in Function",
			code: "if x > 10 { return true; }",
		},
		{
			name: "Invalid For Loop in Function",
			code: "for let i = 0; i < 10; i++ { }"},
		{
			name: "Incomplete Type Declaration in Module",
			code: `module TestModule {
				type Person {
					var name string
				fn main() {
					return;
				}
			}`,
		},
		{
			name: "Missing Type In Parameter",
			code: `module TestModule {
				fn test(a, b string) { }
			}`,
		},
		// {
		// 	name: "Standalone Statement Without Module",
		// 	code: `let x number = 10;`,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := antlr.NewInputStream(tc.code)
			lexer := NewR2D2Lexer(input)
			tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

			p := NewR2D2Parser(tokenStream)

			// Add a custom error listener that will track if an error was reported
			errorListener := &testErrorListener{t: t, expectError: true}
			p.RemoveErrorListeners()
			p.AddErrorListener(errorListener)

			// Try to parse the input
			p.Program()

			if !errorListener.hasError {
				t.Errorf("Parser did not report error on invalid code: %s", tc.code)
			}
		})
	}
}

// Test complex program with multiple features
func TestComplexProgram(t *testing.T) {
	code := `
	use "console.r2d2";
use "math.r2d2";

interface Drawable {
    fn draw(): void;
}

module Geometry implements Drawable {
    export const PI number = 3.14159;

    export fn calculateDistance(p1_x number, p1_y number, p2_x number, p2_y number): number {
        let dx number = p2_x - p1_x;
        let dy number = p2_y - p1_y;
        return Math.sqrt(dx * dx + dy * dy);
    }

    export fn draw(): void {
        Console.log("Drawing geometry...");
    }

    fn private_helper(): void {
        // This is a private function
    }
}

module Main {
    export fn main(): void {
        // Using simple variables for points
        let origin_x number = 0;
        let origin_y number = 0;
        let p_x number = 3;
        let p_y number = 4;

        let distance number = Geometry.calculateDistance(origin_x, origin_y, p_x, p_y);

        if (distance > 5) {
            Console.log("Point is far away!");
        } else {
            Console.log("Point is nearby!");
        }

        let i number = 1;
        while (i < 10) {
            if (i % 2 != 0) {
                Console.log(i);
            }
            i++;
        }

        let i number = 1;
        while (i <= 5) {
            Console.log(i);
            i++;
        }

        // JavaScript integration
        @js <<
            console.log("This is native JavaScript code");
        >>;
    }
}
	`

	input := antlr.NewInputStream(code)
	lexer := NewR2D2Lexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewR2D2Parser(tokenStream)

	// Add an error listener that will fail the test on any syntax error
	errorListener := &testErrorListener{t: t}
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Try to parse the input
	p.Program()

	if errorListener.hasError {
		t.Errorf("Parser reported error on complex program")
	}
}

// Test that the custom error listener is working correctly
// func TestErrorListener(t *testing.T) {
// 	input := antlr.NewInputStream(`module Test {
// 		fn main() {
// 			let x number =
// 		}
// 	}`)
// 	lexer := NewR2D2Lexer(input)
// 	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

// 	p := NewR2D2Parser(tokenStream)

// 	// Create a new error listener and add it to the parser
// 	customErrorListener := NewR2D2ErrorListener()
// 	p.RemoveErrorListeners()
// 	p.AddErrorListener(customErrorListener)

// 	// Try to parse the input (should result in an error)
// 	// This is a bit tricky to test since the error listener
// 	// exits the program. We'll modify the test to not expect
// 	// a failure and only verify it calls the error method.

// 	// Since our error listener exits the program, we're not directly
// 	// testing its behavior here, but this ensures the listener's
// 	// integration with the parser
// 	p.Program()
// }

// TestLexerTokens tests that the lexer produces the expected tokens
func TestLexerTokens(t *testing.T) {
	code := `module TestModule {
		let x number = 42;
		const PI number = 3.14159;
	}`

	input := antlr.NewInputStream(code)
	lexer := NewR2D2Lexer(input)

	// Expected tokens in order
	expectedTokens := []struct {
		tokenType int
		text      string
	}{
		{R2D2LexerMODULE, "module"},
		{R2D2LexerIDENTIFIER, "TestModule"},
		{R2D2LexerLBRACE, "{"},
		{R2D2LexerLET, "let"},
		{R2D2LexerIDENTIFIER, "x"},
		{R2D2LexerTYPE, "number"},
		{R2D2LexerASSIGN, "="},
		{R2D2LexerINT_LITERAL, "42"},
		{R2D2LexerSEMI, ";"},
		{R2D2LexerCONST, "const"},
		{R2D2LexerIDENTIFIER, "PI"},
		{R2D2LexerTYPE, "number"},
		{R2D2LexerASSIGN, "="},
		{R2D2LexerFLOAT_LITERAL, "3.14159"},
		{R2D2LexerSEMI, ";"},
		{R2D2LexerRBRACE, "}"},
	}

	for i, expected := range expectedTokens {
		token := lexer.NextToken()
		if token.GetTokenType() == antlr.TokenEOF {
			t.Errorf("Unexpected EOF at position %d", i)
			break
		}

		if token.GetTokenType() != expected.tokenType {
			t.Errorf("Token type mismatch at position %d: expected %d, got %d",
				i, expected.tokenType, token.GetTokenType())
		}

		if token.GetText() != expected.text {
			t.Errorf("Token text mismatch at position %d: expected '%s', got '%s'",
				i, expected.text, token.GetText())
		}
	}

	// Make sure we've reached the end of the input
	token := lexer.NextToken()
	if token.GetTokenType() != antlr.TokenEOF {
		t.Errorf("Expected EOF, got token type %d, text '%s'", token.GetTokenType(), token.GetText())
	}
}

// testErrorListener is a custom error listener that allows tests to check for parsing errors
type testErrorListener struct {
	antlr.DefaultErrorListener
	t           *testing.T
	hasError    bool
	expectError bool
}

func (l *testErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any,
	line, column int, msg string, e antlr.RecognitionException) {

	l.hasError = true

	if !l.expectError {
		l.t.Logf("Syntax error at line %d:%d: %s", line, column, msg)
	}
}

// Test parser error recovery
// func TestParserErrorRecovery(t *testing.T) {
// 	// Code with multiple errors that the parser should be able to recover from
// 	code := `
// 	import Console from "console";

// 	module TestModule {
// 		// Error 1: Missing semicolon
// 		let x number = 10

// 		fn main() {
// 			// Error 2: Invalid expression
// 			let y number = 10 + * 5;

// 			// Should still parse this statement
// 			return 42;
// 		}
// 	}
// 	`

// 	input := antlr.NewInputStream(code)
// 	lexer := NewR2D2Lexer(input)
// 	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

// 	p := NewR2D2Parser(tokenStream)

// 	// Create error listener that counts errors but allows parsing to continue
// 	countingErrorListener := &countingErrorListener{}
// 	p.RemoveErrorListeners()
// 	p.AddErrorListener(countingErrorListener)

// 	// Set recovery mode
// 	p.SetErrorHandler(antlr.NewBailErrorStrategy())

// 	// Try to parse
// 	p.Program()

// 	// Verify we found the expected number of errors
// 	if countingErrorListener.errorCount == 0 {
// 		t.Errorf("Expected parser to detect errors, but none were reported")
// 	}
// }

// countingErrorListener counts syntax errors but allows parsing to continue
type countingErrorListener struct {
	antlr.DefaultErrorListener
	errorCount int
}

func (l *countingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any,
	line, column int, msg string, e antlr.RecognitionException) {

	l.errorCount++
}

// TestTreeWalking tests using a parse tree listener to walk the tree
func TestTreeWalking(t *testing.T) {
	code := `
	module Test {
		fn sayHello(name string): void {
			console.log("Hello, " + name);
		}
	}
	`

	input := antlr.NewInputStream(code)
	lexer := NewR2D2Lexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewR2D2Parser(tokenStream)
	tree := p.Program()

	// Create a listener to verify tree structure
	listener := &testListener{t: t}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	// Check that we found what we were looking for
	if !listener.foundModule {
		t.Error("Parse tree walker did not find module declaration")
	}

	if !listener.foundFunction {
		t.Error("Parse tree walker did not find function declaration")
	}
}

// testListener implements R2D2Listener to walk the tree and verify its structure
type testListener struct {
	BaseR2D2Listener
	t             *testing.T
	foundModule   bool
	foundFunction bool
}

// Corrected method names according to the grammar rule names
func (l *testListener) EnterModuleDeclaration(ctx *ModuleDeclarationContext) {
	l.foundModule = true

	// Verify module name
	if ctx.IDENTIFIER(0).GetText() != "Test" {
		l.t.Errorf("Expected module name 'Test', got '%s'", ctx.IDENTIFIER(0).GetText())
	}
}

func (l *testListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {
	l.foundFunction = true

	// Verify function name
	if ctx.IDENTIFIER().GetText() != "sayHello" {
		l.t.Errorf("Expected function name 'sayHello', got '%s'", ctx.IDENTIFIER().GetText())
	}
}

// Test the visitor pattern
func TestVisitorPattern(t *testing.T) {
	code := `module TestModule {
		fn main() {
			let x number = 42;
		}
	}`

	input := antlr.NewInputStream(code)
	lexer := NewR2D2Lexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewR2D2Parser(tokenStream)
	tree := p.Program()

	// Create a visitor to check tree structure
	visitor := &testVisitor{t: t}
	visitor.Visit(tree)

	if !visitor.foundVarDecl {
		t.Error("Visitor did not find variable declaration")
	}
}

// testVisitor implements R2D2Visitor to verify tree structure
type testVisitor struct {
	BaseR2D2Visitor
	t            *testing.T
	foundVarDecl bool
}

func (v *testVisitor) Visit(tree antlr.ParseTree) any {
	switch t := tree.(type) {
	case *VariableDeclarationContext:
		v.foundVarDecl = true

		// Verify variable name
		if t.IDENTIFIER().GetText() != "x" {
			v.t.Errorf("Expected variable name 'x', got '%s'", t.IDENTIFIER().GetText())
		}

		// Verify type - check if it's using TYPE token rule instead
		if t.TypeExpression().BaseType().TYPE() != nil && t.TypeExpression().BaseType().TYPE().GetText() != "number" {
			v.t.Errorf("Expected variable type 'number', got '%s'", t.TypeExpression().BaseType().TYPE().GetText())
		}

		return nil
	default:
		// Continue visiting children
		for i := range tree.GetChildCount() {
			child := tree.GetChild(i)
			v.Visit(child.(antlr.ParseTree))
		}
		return nil
	}
}

// Helper function to verify that a syntax tree contains certain text
func TestParseTreeToString(t *testing.T) {
	code := `module TestModule {
		fn main() {
			let x number = 42;
		}
	}`

	input := antlr.NewInputStream(code)
	lexer := NewR2D2Lexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := NewR2D2Parser(tokenStream)
	tree := p.Program()

	// Convert tree to string form
	treeString := antlr.TreesStringTree(tree, nil, p)

	// Check for expected components in the tree string
	expectedParts := []string{
		"moduleDeclaration",
		"TestModule",
		"variableDeclaration",
		"let",
		"x",
		"number",
		"expression",
	}

	for _, part := range expectedParts {
		if !strings.Contains(treeString, part) {
			t.Errorf("Expected parse tree string to contain '%s', but it didn't.\nTree: %s",
				part, treeString)
		}
	}
}
