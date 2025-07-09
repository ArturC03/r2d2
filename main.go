package r2d2

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/ArturC03/r2d2/errors"
	"github.com/ArturC03/r2d2/parser"
	"github.com/ArturC03/r2d2/visitor"
	"github.com/ArturC03/r2d2Styles"
	"github.com/ArturC03/r2d2Styles/spinner"
	"github.com/ArturC03/r2d2Styles/stopwatch"
	"github.com/antlr4-go/antlr/v4"
	tea "github.com/charmbracelet/bubbletea"
)

// Builds the JavaScript code from the user input
func BuildJSCode(userInput string) (string, errors.ErrorCollector) {
	errorCollector := &errors.ErrorCollector{HasSyntaxError: false}
	input := antlr.NewInputStream(userInput)

	lexer := parser.NewR2D2Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewR2D2Parser(stream)

	el := parser.NewR2D2ErrorListener(errorCollector)
	p.RemoveErrorListeners()
	p.AddErrorListener(el)

	p.BuildParseTrees = true
	tree := p.Program()

	v := visitor.NewR2D2Visitor(errorCollector)
	v.LoadStdModules()

	tree.Accept(v)

	return v.JsCode, *errorCollector
}

func RunCode(input string) error {
	code, errorCollector := BuildJSCode(input)
	if errorCollector.HasErrors() {
		return fmt.Errorf("R2D2 build errors occurred")
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "deno_run_*.js")
	if err != nil {
		message := fmt.Sprintf("System Error: Could not create temporary file for execution: %v", r2d2Styles.Bold(err.Error()))
		fmt.Println(r2d2Styles.ErrorMessage(message))
		return err
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(code)
	if err != nil {
		message := fmt.Sprintf("System Error: Could not write code to temporary file: %v", r2d2Styles.Bold(err.Error()))
		fmt.Println(r2d2Styles.ErrorMessage(message))
		return err
	}
	tmpFile.Close()

	cmd := exec.Command("deno", "run", "--allow-all", tmpFile.Name())

	var stderrBuffer bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = &stderrBuffer
	cmd.Stdout = os.Stdout

	startTime := time.Now()
	cmdErr := cmd.Run()

	formattedTime := stopwatch.MeasureExecutionTime(startTime)

	if cmdErr != nil {
		rawDenoError := stderrBuffer.String()

		// Remove ANSI color codes that interfere with regex matching
		ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
		cleanError := ansiRegex.ReplaceAllString(rawDenoError, "")

		// Extract the core error message
		reCoreError := regexp.MustCompile(`error: (?:Uncaught \(in promise\) )?(.*?)(?:\n\s*at |$)`)
		matches := reCoreError.FindStringSubmatch(cleanError)

		if len(matches) > 1 {
			denoMainError := strings.TrimSpace(matches[1])

			// Pattern matching for specific error types
			var errorMessage string
			switch {
			// Property access errors on undefined/null
			case strings.Contains(denoMainError, "Cannot read properties of undefined"):
				property := extractProperty(denoMainError, "undefined")
				errorMessage = fmt.Sprintf("Cannot read property %s of %s - variable is undefined. Check if it was properly initialized.", r2d2Styles.Bold("'"+property+"'"), r2d2Styles.Bold("undefined"))

			case strings.Contains(denoMainError, "Cannot read properties of null"):
				property := extractProperty(denoMainError, "null")
				errorMessage = fmt.Sprintf("Cannot read property %s of %s - variable is null. Check if it was properly assigned.", r2d2Styles.Bold("'"+property+"'"), r2d2Styles.Bold("null"))

			// Property access errors on primitives
			case strings.Contains(denoMainError, "Cannot read properties of"):
				errorMessage = fmt.Sprintf("%s - trying to access a property on a value that doesn't have properties.", denoMainError)

			// Function call errors
			case strings.Contains(denoMainError, "is not a function"):
				funcName := extractFunctionName(denoMainError)
				errorMessage = fmt.Sprintf("%s is not a function - you're trying to call something that isn't a function.", r2d2Styles.Bold("'"+funcName+"'"))

			// Variable not defined
			case strings.Contains(denoMainError, "is not defined"):
				varName := extractVariableName(denoMainError)
				errorMessage = fmt.Sprintf("%s is not defined - this variable or function hasn't been declared.", r2d2Styles.Bold("'"+varName+"'"))

			// Assignment errors
			case strings.Contains(denoMainError, "Assignment to constant variable"):
				errorMessage = fmt.Sprintf("Cannot reassign a %s variable - you're trying to change a variable declared with 'const'.", r2d2Styles.Bold("constant"))

			// Syntax errors
			case strings.Contains(denoMainError, "Unexpected token"):
				errorMessage = fmt.Sprintf("%s in code - there's a syntax error in your code structure.", r2d2Styles.Bold("Unexpected token"))

			case strings.Contains(denoMainError, "Missing initializer"):
				errorMessage = fmt.Sprintf("Missing initializer for %s declaration - constants must be initialized when declared.", r2d2Styles.Bold("const"))

			// Range errors
			case strings.Contains(denoMainError, "Maximum call stack size exceeded"):
				errorMessage = fmt.Sprintf("%s - infinite recursion detected. Check your recursive function calls.", r2d2Styles.Bold("Maximum call stack size exceeded"))

			case strings.Contains(denoMainError, "Invalid array length"):
				errorMessage = fmt.Sprintf("%s - array length must be a positive integer.", r2d2Styles.Bold("Invalid array length"))

			// Network/Permission errors (Deno specific)
			case strings.Contains(denoMainError, "PermissionDenied"):
				errorMessage = fmt.Sprintf("%s - the operation requires additional permissions.", r2d2Styles.Bold("Permission denied"))

			case strings.Contains(denoMainError, "NetworkError"):
				errorMessage = fmt.Sprintf("%s - failed to connect to the network resource.", r2d2Styles.Bold("Network error"))

			// Generic TypeError
			case strings.Contains(denoMainError, "TypeError:"):
				cleanMessage := strings.TrimPrefix(denoMainError, "TypeError: ")
				errorMessage = fmt.Sprintf("%s - type-related error in your code.", cleanMessage)

			// Generic ReferenceError
			case strings.Contains(denoMainError, "ReferenceError:"):
				cleanMessage := strings.TrimPrefix(denoMainError, "ReferenceError: ")
				errorMessage = fmt.Sprintf("%s - reference error, check variable names and scopes.", cleanMessage)

			// Generic SyntaxError
			case strings.Contains(denoMainError, "SyntaxError:"):
				cleanMessage := strings.TrimPrefix(denoMainError, "SyntaxError: ")
				errorMessage = fmt.Sprintf("%s - syntax error in your code structure.", cleanMessage)

			// Default fallback
			default:
				errorMessage = fmt.Sprintf("%s - an unexpected issue occurred during execution of your R2D2 program.", denoMainError)
			}

			// Print the single consolidated error message
			fmt.Println(r2d2Styles.ErrorMessage(errorMessage))
		} else {
			// No regex match - generic error
			fmt.Println(r2d2Styles.ErrorMessage("An unexpected issue occurred during execution of your R2D2 program."))
		}

		return cmdErr
	}

	fmt.Println(r2d2Styles.SuccessMessage(fmt.Sprintf("Execution completed in %s", r2d2Styles.Bold(formattedTime))))
	return nil
}

// Helper functions to extract specific parts of error messages
func extractProperty(errorMsg, valueType string) string {
	// Extract property name from "Cannot read properties of undefined (reading 'property')"
	re := regexp.MustCompile(`\(reading '([^']+)'\)`)
	matches := re.FindStringSubmatch(errorMsg)
	if len(matches) > 1 {
		return matches[1]
	}
	return "unknown"
}

func extractFunctionName(errorMsg string) string {
	// Extract function name from "someVar is not a function"
	parts := strings.Split(errorMsg, " is not a function")
	if len(parts) > 0 {
		return strings.TrimSpace(parts[0])
	}
	return "unknown"
}

func extractVariableName(errorMsg string) string {
	// Extract variable name from "someVar is not defined"
	parts := strings.Split(errorMsg, " is not defined")
	if len(parts) > 0 {
		return strings.TrimSpace(parts[0])
	}
	return "unknown"
}

// BuildCode executa o comando Deno compile com spinner
func BuildCode(input string, filename string) error {
	code, errorCollector := BuildJSCode(input)

	if errorCollector.HasErrors() {
		return fmt.Errorf("build errors occurred")
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "deno_code_*.js")
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Error creating temporary file at %s: %v", r2d2Styles.Bold(os.TempDir()), err)))
		return err
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString(code)
	tmpFile.Close()

	baseName := filename
	if len(filename) > 5 && filename[len(filename)-5:] == ".r2d2" {
		baseName = filename[:len(filename)-5]
	}

	outputName := baseName
	if runtime.GOOS == "windows" {
		outputName += ".exe"
	}

	cmd := exec.Command("deno", "compile", "--allow-all", "--output", outputName, tmpFile.Name())

	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf

	spinnerModel := spinner.New()
	m := spinner.Model{Spinner: spinnerModel.Spinner}
	p := tea.NewProgram(&m)
	done := make(chan struct{})
	go func() {
		p.Run()
		close(done)
	}()

	startTime := time.Now()
	cmdErr := cmd.Run()
	p.Send(tea.Quit())
	<-done
	if cmdErr != nil {
		rawDenoCompileError := stderrBuf.String()
		log.Printf("Internal: Raw Deno Compile Error Output:\n%s", rawDenoCompileError)

		displayError := "R2D2 compilation failed. An internal error occurred during executable generation."
		if len(rawDenoCompileError) > 0 {
			// Attempt to get the most relevant line of the compile error
			lines := strings.Split(rawDenoCompileError, "\n")
			for _, line := range lines {
				trimmedLine := strings.TrimSpace(line)
				if len(trimmedLine) > 0 &&
					!strings.HasPrefix(trimmedLine, "Check file:") &&
					!strings.HasPrefix(trimmedLine, "error: Download") &&
					!strings.HasPrefix(trimmedLine, "Check http") {
					displayError = fmt.Sprintf("Compilation Error: %s", trimmedLine)
					// Remove temporary file paths from compile errors too
					displayError = strings.ReplaceAll(displayError, fmt.Sprintf("file://%s", os.TempDir()), "your R2D2 code base")
					break
				}
			}
		}
		fmt.Println(r2d2Styles.ErrorMessage(r2d2Styles.Bold(displayError)))
		return cmdErr
	}

	m.SetDone(true)
	fmt.Println(r2d2Styles.SuccessMessage(fmt.Sprintf("Created executable: %s\n", r2d2Styles.Bold(outputName))))

	formattedTime := stopwatch.MeasureExecutionTime(startTime)
	fmt.Println(fmt.Sprintf("Compilation completed in %s", r2d2Styles.Bold(formattedTime)))
	return nil
}

func BuildJsFile(input string, filename string) error {
	code, errorCollector := BuildJSCode(input)

	if errorCollector.HasErrors() {
		return fmt.Errorf("build errors occurred")
	}

	baseName := filename
	if len(filename) > 5 && filename[len(filename)-5:] == ".r2d2" {
		baseName = filename[:len(filename)-5]
	}

	outputName := baseName + ".js"
	outFile, err := os.Create(outputName)
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Error creating JavaScript file %s: %v", r2d2Styles.Bold(outputName), err)))
		return err
	}
	defer outFile.Close()

	_, err = outFile.WriteString(code)
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Error writing to JavaScript file: %v", err)))
		return err
	}

	fmt.Println(r2d2Styles.SuccessMessage(fmt.Sprintf("Created JavaScript file: %s", r2d2Styles.Bold(outputName))))
	return nil
}
