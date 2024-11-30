package main

import (
	"fmt"
	"github.com/ArturC03/r2d2/lexer"
	"io"
	"log"
	"os"
)

func main() {
	// Specify the filename
	filename := "input.txt" // Change this to the path of your file

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

	// Create a new lexer instance with the file content
	l := lexer.New(string(content))

	// Tokenize the input until EOF
	for {
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)
		if tok.Type == lexer.EOF {
			break
		}
	}
}
