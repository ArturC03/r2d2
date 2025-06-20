package r2d2

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
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

// Executes the code generated by the parser
func RunCode(input string) error {
	code, errorCollector := BuildJSCode(input)
	if errorCollector.HasErrors() {
		return fmt.Errorf("build errors occurred")
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "deno_run_*.js")
	if err != nil {
		message := fmt.Sprintf("Error creating temporary file: %v", r2d2Styles.Bold(err.Error()))
		fmt.Println(r2d2Styles.ErrorMessage(message))
		return err
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(code)
	if err != nil {
		return err
	}
	tmpFile.Close()

	// 	// Show spinner briefly, then run command with direct stdio
	// 	spinnerModel := spinner.New()
	// 	m := spinner.Model{Spinner: spinnerModel.Spinner}
	// 	p := tea.NewProgram(&m)
	//
	// // Start spinner
	// 	done := make(chan struct{})
	// 	go func() {
	// 		p.Run()
	// 		close(done)
	// 	}()
	//
	// 	// Brief delay to show spinner, then stop it
	// 	time.Sleep(100 * time.Millisecond)
	// 	p.Send(tea.Quit())
	// 	<-done
	// fmt.Println() // Clean line

	// Now run the command with direct stdio access
	cmd := exec.Command("deno", "run", "--allow-all", tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	startTime := time.Now()
	err = cmd.Run()

	formattedTime := stopwatch.MeasureExecutionTime(startTime)
	if err != nil {
		log.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Deno run error: %s", r2d2Styles.Bold(err.Error()))))
		return err
	}

	fmt.Println(r2d2Styles.SuccessMessage(fmt.Sprintf("Execution completed in %s", r2d2Styles.Bold(formattedTime))))
	return nil
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
		log.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Deno compile: %v", r2d2Styles.Bold(err.Error()))))
		return err
	}

	// Exibe o nome do executável gerado e o tempo com precisão de milissegundos
	m.SetDone(true)
	fmt.Println(r2d2Styles.SuccessMessage(fmt.Sprintf("Created executable: %s\n", r2d2Styles.Bold(outputName))))

	// Calcula e exibe o tempo de execução
	formattedTime := stopwatch.MeasureExecutionTime(startTime)
	fmt.Println(fmt.Sprintf("Compilation completed in %s", r2d2Styles.Bold(formattedTime))) // Exibe a mensagem atualizada
	return nil
}

func BuildJsFile(input string, filename string) error {
	code, errorCollector := BuildJSCode(input)

	if errorCollector.HasErrors() {
		return fmt.Errorf("build errors occurred")
	}

	// Extract base filename without extension
	baseName := filename
	if len(filename) > 5 && filename[len(filename)-5:] == ".r2d2" {
		baseName = filename[:len(filename)-5]
	}

	// Create output JavaScript file with same base name
	outputName := baseName + ".js"
	outFile, err := os.Create(outputName)
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Error creating JavaScript file %s: %v", r2d2Styles.Bold(outputName), err)))
		return err
	}
	defer outFile.Close()

	// Write JavaScript code to the file
	_, err = outFile.WriteString(code)
	if err != nil {
		fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Error writing to JavaScript file: %v", err)))
		return err
	}

	// Display success message
	fmt.Println(r2d2Styles.SuccessMessage(fmt.Sprintf("Created JavaScript file: %s", r2d2Styles.Bold(outputName))))
	return nil
}
