package main

import (
	"bytes"
	"fmt"
	"github.com/ArturC03/r2d2/parser"
	"github.com/ArturC03/r2d2/visitor"
	r2d2Styles "github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	strings "strings"
)

func main() {
	// Criar um stream de entrada
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
            
            if (2 == 3) {
                return 1;
            } else {
                return 2;
            }

            continue;
        }
    }
}
`) // Ajuste conforme a gramática

	// fmt.Println(r2d2Styles.InfoMessage("Input stream criado."))

	// Criar Lexer
	lexer := parser.NewR2D2Lexer(input)
	if lexer == nil {
		fmt.Println(r2d2Styles.ErrorMessage("Lexer não foi inicializado corretamente!"))
	}
	// fmt.Println(r2d2Styles.InfoMessage("Lexer criado."))

	// Criar Token Stream
	stream := antlr.NewCommonTokenStream(lexer, 0)
	// fmt.Println(r2d2Styles.InfoMessage("Token stream criado."))

	// Criar Parser
	p := parser.NewR2D2Parser(stream)

	// Criar Error Listener
	el := parser.NewR2D2ErrorListener()

	// Adicionar error listener ao parser
	p.RemoveErrorListeners()
	p.AddErrorListener(el)

	if p == nil {
		fmt.Println(r2d2Styles.ErrorMessage("Parser não foi inicializado corretamente!"))
	}
	// fmt.Println(r2d2Styles.InfoMessage("Parser criado."))

	// Construir a árvore de parse
	p.BuildParseTrees = true
	tree := p.Program() // Ajuste conforme sua regra principal
	// Imprimir a árvore sintática
	// fmt.Println(tree.ToStringTree(p.GetRuleNames(), nil))
	// return
	if tree == nil {
		fmt.Println(r2d2Styles.ErrorMessage("Árvore sintática não foi criada corretamente!"))
	}
	// fmt.Println(r2d2Styles.InfoMessage("Árvore sintática criada."))

	// Criar e aplicar o Visitor
	v := visitor.NewR2D2Visitor()
	if v == nil {
		fmt.Println(r2d2Styles.ErrorMessage("Visitor não foi inicializado corretamente!"))
	}
	// fmt.Println(r2d2Styles.InfoMessage("Visitor criado."))

	// Aplicar o visitor na árvore - esta é a forma correta para ANTLR em Go
	result := tree.Accept(v)

	// fmt.Println(r2d2Styles.InfoMessage("Visitor aplicado com sucesso."))

	// Exibir o resultado
	fmt.Println(r2d2Styles.InfoMessage(fmt.Sprintf("Resultado do Visitor: %v", result)))
	fmt.Println(r2d2Styles.InfoMessage(v.JsCode))

	fmt.Println(r2d2Styles.InfoMessage("Running the code generated"))
	BuildCode(v.JsCode)
}

func RunCode(code string) {
	cmd := exec.Command("deno", "eval", "--quiet", code)

	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf

	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o Deno eval: %v\n", err)
		return
	}

	// Expressão regular para capturar a mensagem de erro relevante
	re := regexp.MustCompile(`error: (Uncaught [^\n]+)`)
	matches := re.FindStringSubmatch(stderrBuf.String())
	if len(matches) > 1 {
		fmt.Println(matches[1]) // Exibe a mensagem de erro simplificada
	} else {
		fmt.Println("Nenhuma mensagem de erro encontrada.")
	}
}

// BuildCode executa o comando Deno bundle com o código fornecido.
func BuildCode(code string) {
	// Create a temporary file to store the code
	tmpFile, err := ioutil.TempFile(os.TempDir(), "deno_code_*.js")
	if err != nil {
		log.Fatalf("Error creating temporary file: %v", err)
		return
	}
	defer os.Remove(tmpFile.Name()) // Remove the temporary file after use

	// Write the code (JsCode string) to the temporary file
	_, err = tmpFile.WriteString(code)
	if err != nil {
		log.Fatalf("Error writing to temporary file: %v", err)
		return
	}
	tmpFile.Close()

	// Output executable name
	outputName := "program"
	if runtime.GOOS == "windows" {
		outputName += ".exe"
	}

	// Create the Deno compile command
	cmd := exec.Command("deno", "compile", "--allow-all", "--output", outputName, tmpFile.Name())

	// Print the command before executing it
	fmt.Println("Deno compile command:")
	fmt.Println(strings.Join(cmd.Args, " "))

	// Create buffers to capture standard output and errors
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	// Execute the command and check if any errors occurred
	if err := cmd.Run(); err != nil {
		log.Printf("Error executing Deno compile: %v", err)
		fmt.Println("Error output:", stderrBuf.String())
		return
	}

	// Display the command's standard output
	fmt.Println("Deno compile command output:")
	fmt.Println(stdoutBuf.String())

	fmt.Printf("Created executable: %s\n", outputName)
}
