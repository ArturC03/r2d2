package main

import (
	"fmt"
	"github.com/ArturC03/r2d2/lexer"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Specify the filename
	filename := "input.r2d2" // Change this to the path of your file

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close()

	// Read the contents of the file
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}

	// Debug: print file content
	fmt.Printf("File Content:\n%s\n", string(content))

	// Normalize line endings
	normalizedContent := strings.ReplaceAll(string(content), "\r\n", "\n")

	// Check if file is empty
	if len(normalizedContent) == 0 {
		log.Fatalf("Error: The file is empty.")
	}

	// Create a new lexer instance with the file content
	l := lexer.New(normalizedContent)

	// Tokenize the input until EOF
	for {
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)
		if tok.Type == lexer.EOF {
			break
		}
	}
}
