package r2d2

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestBuildJSCode tests the JS code generation function with a simple input
func TestBuildJSCode(t *testing.T) {
	// Skip if we're only running short tests
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	// Test with a very simple input
	input := `
    module Test {
        export fn main() {
        @js <<
            console.log("Hello, world!");
            >>;
        }
    }
    `

	// Capture stdout to avoid output during tests
	oldStdout := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w

	jsCode := buildJSCode(input)

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Check if the JS code contains expected patterns
	expectedPatterns := []string{
		"const Test =",
		"function main()",
		"console.log",
		"return {main}",
		"Test.main()"}

	for _, pattern := range expectedPatterns {
		if !strings.Contains(jsCode, pattern) {
			t.Errorf("Expected JS code to contain %q, but it doesn't", pattern)
		}
	}
}

// TestRunCode tests the RunCode function with a minimum valid program
func TestRunCode(t *testing.T) {
	// Skip if deno is not installed
	if _, err := exec.LookPath("deno"); err != nil {
		t.Skip("Deno not installed, skipping test")
	}

	// Skip if we're only running short tests
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	// Create a simple valid program
	minimalProgram := `
    module Test {
        export fn main() {
            console.log("Test executed successfully");
        }
    }
    `

	// Capture stdout to check output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the code
	RunCode(minimalProgram)

	// Restore stdout and read captured output
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// Just check that the function executed without crashing
	// The exact output is complex to predict due to all the formatting and processing
	if !strings.Contains(output, "Execution completed") {
		t.Errorf("RunCode did not complete execution properly")
	}
}

// TestBuildJsFile tests the BuildJsFile function
func TestBuildJsFile(t *testing.T) {
	// Create a temporary directory for test
	tempDir, err := ioutil.TempDir("", "r2d2-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a simple test program
	testProgram := `
    module Test {
        export fn main() {
        @js <<
            console.log("JS File Test");
            >>;
        }
    }
    `

	// Test file paths
	testFilePath := filepath.Join(tempDir, "test.r2d2")
	expectedJsPath := filepath.Join(tempDir, "test.js")

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test the function
	BuildJsFile(testProgram, testFilePath)

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// Check output message
	if !strings.Contains(output, "Created JavaScript file") {
		t.Errorf("Expected output to mention JS file creation, got: %s", output)
	}

	// Check if file was created
	if _, err := os.Stat(expectedJsPath); os.IsNotExist(err) {
		t.Errorf("Expected JS file %s to be created, but it doesn't exist", expectedJsPath)
	}

	// Check file content
	content, err := ioutil.ReadFile(expectedJsPath)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	// Just verify that the file contains something that looks like JavaScript
	if !strings.Contains(string(content), "console.log") {
		t.Errorf("File content does not contain expected JavaScript")
	}
}

// TestBuildCode tests the BuildCode function
func TestBuildCode(t *testing.T) {
	// Skip if deno is not installed
	if _, err := exec.LookPath("deno"); err != nil {
		t.Skip("Deno not installed, skipping test")
	}

	// Skip in short mode
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	// Create a simple test program
	testProgram := `
    module Test {
        export fn main() {
            console.log("Build test");
        }
    }
    `

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test the function in the current directory
	// Make sure to clean up afterward
	outputName := "program"
	if runtime.GOOS == "windows" {
		outputName += ".exe"
	}
	defer os.Remove(outputName)

	// Run the function
	BuildCode(testProgram)

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// Check output messages
	if !strings.Contains(output, "Created executable") {
		t.Errorf("Expected output to contain 'Created executable', got: %s", output)
	}

	// Check if the executable was created
	if _, err := os.Stat(outputName); os.IsNotExist(err) {
		t.Errorf("Expected executable %s to be created, but it doesn't exist", outputName)
	}
}
