package main

import (
	"bytes"
	"fmt"
	"github.com/ArturC03/r2d2/parser"
	"github.com/ArturC03/r2d2/visitor"
	r2d2Styles "github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
	"log"
	"os/exec"
	"regexp"
)

func main() {
	// Criar um stream de entrada
	input := antlr.NewInputStream(`import a from "/home/rutra/Documentos/CODE/PAP/r2d2/main.go";

		const as: i32 = 1;

		module cookie{

		type a{
			var a: i32 = 1;
		}

		export var a: i32 = 1;

		export fn main(){
		loop {
		var i: i32 = 3;
		cookie();
		return 1 ;
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
	RunCode(v.JsCode)
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

// BuildCode executa o comando Deno build com o código fornecido.
func BuildCode(code string) {
	cmd := exec.Command("deno", "bundle", code)

	// Cria buffers para capturar a saída padrão e os erros.
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	// Executa o comando e verifica se ocorreu algum erro.
	if err := cmd.Run(); err != nil {
		log.Printf("Erro ao executar o Deno build: %v", err)
		fmt.Println("Saída de erro:", stderrBuf.String())
		return
	}

	// Exibe a saída padrão do comando.
	fmt.Println("Saída do Deno build:", stdoutBuf.String())
}
