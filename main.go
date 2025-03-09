package main

import (
	"bytes"
	"fmt"
	"github.com/ArturC03/r2d2/parser"
	"github.com/ArturC03/r2d2/visitor"
	r2d2Styles "github.com/ArturC03/r2d2Styles"
	"github.com/ArturC03/r2d2Styles/spinner"
	"github.com/ArturC03/r2d2Styles/stopwatch"
	"github.com/antlr4-go/antlr/v4"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	input := antlr.NewInputStream(`
import a from "/home/rutra/Documentos/CODE/PAP/r2d2/main.go";

const as: i32 = 1;

module cookie {
    type a {
        var a: i32 = 1;
    }

    export var a: i32 = 1;

    export fn main() {
        loop {
            var i: i32 = 3;
            cookie();
						for (var i: i32 = 0; i < 10; i++) {
		return 1;
							cookie();
						}
            
            if (2 == 3) {
                return 1;
            } else {
                return 2;
            }
		break;
        }
    }
}
`)

	lexer := parser.NewR2D2Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewR2D2Parser(stream)
	el := parser.NewR2D2ErrorListener()
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	p.BuildParseTrees = true
	tree := p.Program()
	v := visitor.NewR2D2Visitor()
	tree.Accept(v)
	fmt.Println(r2d2Styles.InfoMessage(v.JsCode))

	fmt.Println(r2d2Styles.InfoMessage("Running the code generated"))
	BuildCode(v.JsCode)
}

// BuildCode executa o comando Deno compile com spinner
func BuildCode(code string) {
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
		p.Start()
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
