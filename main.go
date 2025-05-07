package r2d2

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/ArturC03/r2d2/parser"
	"github.com/ArturC03/r2d2/visitor"
	r2d2Styles "github.com/ArturC03/r2d2Styles"
	"github.com/ArturC03/r2d2Styles/spinner"
	"github.com/ArturC03/r2d2Styles/stopwatch"
	"github.com/antlr4-go/antlr/v4"
	tea "github.com/charmbracelet/bubbletea"
)

func buildJSCode(userInput string) string {
	// print(userInput)
	input := antlr.NewInputStream(userInput)

	lexer := parser.NewR2D2Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewR2D2Parser(stream)
	el := parser.NewR2D2ErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	p.BuildParseTrees = true
	tree := p.Program()
	v := visitor.NewR2D2Visitor()
	// fmt.Println(r2d2Styles.InfoMessage(v.JsCode))
	tree.Accept(v)
	// fmt.Println(r2d2Styles.InfoMessage(v.JsCode))

	// fmt.Println(r2d2Styles.InfoMessage("Running the code generated"))
	return v.JsCode
}

func RunCode(input string) {
	code := buildJSCode(input)
	// Create a temporary file to store the generated code
	tmpFile, err := os.CreateTemp(os.TempDir(), "deno_run_*.js")
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Error creating temporary file: %v", err)))
		return
	}
	defer os.Remove(tmpFile.Name())

	// Write the code to the temporary file
	tmpFile.WriteString(code)
	tmpFile.Close()

	// Prepare the command to run the code using Deno
	cmd := exec.Command("deno", "run", "--allow-all", tmpFile.Name())

	// Capture both stdout and stderr
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	// Initialize spinner using r2d2Styles/spinner
	spinnerModel := spinner.New()
	m := spinner.Model{Spinner: spinnerModel.Spinner}

	// Start BubbleTea program for spinner
	p := tea.NewProgram(&m)
	done := make(chan struct{})
	go func() {
		// p.Start()
		p.Run()
		close(done)
	}()

	// Measure execution time
	startTime := time.Now()

	// Run the command
	err = cmd.Run()
	p.Send(tea.Quit())
	<-done // Wait for spinner to close properly

	// Calculate execution time
	formattedTime := stopwatch.MeasureExecutionTime(startTime)

	// Handle command execution results
	if err != nil {
		// If there's an error, print error message
		log.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Deno run error: %v", err)))
		fmt.Println(r2d2Styles.ErrorMessage("Standard Error Output:"))
		fmt.Println(errBuf.String())
		return
	}

	// If successful, print output and execution time
	m.SetDone(true)
	if outBuf.Len() > 0 {
		fmt.Println(r2d2Styles.InfoMessage("Program Output:"))
		fmt.Println(outBuf.String())
	}
	fmt.Println(r2d2Styles.SuccessMessage(fmt.Sprintf("Execution completed in %s", formattedTime)))
}

// BuildCode executa o comando Deno compile com spinner
func BuildCode(input string, filename string) {
	fmt.Println(filename)
	code := buildJSCode(input)
	tmpFile, err := os.CreateTemp(os.TempDir(), "deno_code_*.js")
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Error creating temporary file: %v", err)))
		return
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString(code)
	tmpFile.Close()

	outputName := "program"
	if runtime.GOOS == "windows" {
		outputName += ".exe"
	}

	cmd := exec.Command("deno", "compile", "--allow-all", "--output", outputName, tmpFile.Name())

	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf

	// Inicializa o spinner usando o pacote r2d2Styles/spinner
	spinnerModel := spinner.New()

	m := spinner.Model{Spinner: spinnerModel.Spinner}

	// Inicia o programa BubbleTea
	p := tea.NewProgram(&m)
	done := make(chan struct{})
	go func() {
		// p.Start()
		p.Run()
		close(done)
	}()

	// Começa a medir o tempo agora
	startTime := time.Now() // Medir o tempo de início da compilação

	err = cmd.Run()
	p.Send(tea.Quit())
	<-done // Aguarda o encerramento correto do spinner

	if err != nil {
		log.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Deno compile: %v", err)))
		return
	}

	// Exibe o nome do executável gerado e o tempo com precisão de milissegundos
	m.SetDone(true)
	fmt.Println(r2d2Styles.SuccessMessage(fmt.Sprintf("Created executable: %s\n", outputName)))

	// Calcula e exibe o tempo de execução
	formattedTime := stopwatch.MeasureExecutionTime(startTime)
	fmt.Println(fmt.Sprintf("Compilation completed in %s", formattedTime)) // Exibe a mensagem atualizada
}

func BuildJsFile(input string, filename string) {
	code := buildJSCode(input)

	// Extract base filename without extension
	baseName := filename
	if len(filename) > 5 && filename[len(filename)-5:] == ".r2d2" {
		baseName = filename[:len(filename)-5]
	}

	// Create output JavaScript file with same base name
	outputName := baseName + ".js"
	outFile, err := os.Create(outputName)
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Error creating JavaScript file: %v", err)))
		return
	}
	defer outFile.Close()

	// Write JavaScript code to the file
	_, err = outFile.WriteString(code)
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Error writing to JavaScript file: %v", err)))
		return
	}

	// Display success message
	fmt.Println(r2d2Styles.SuccessMessage(fmt.Sprintf("Created JavaScript file: %s", outputName)))
}
