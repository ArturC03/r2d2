package parser

import (
	"fmt"
	"os"

	r2d2Styles "github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
)

// R2D2ErrorListener is a listener that listens to errors emitted by the parser.
type R2D2ErrorListener struct {
	*antlr.DefaultErrorListener
}

func NewR2D2ErrorListener() *R2D2ErrorListener {
	return &R2D2ErrorListener{antlr.NewDefaultErrorListener()}
}

func (l *R2D2ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Syntax error at line %d:%d: %s", line, column, msg)))
	// panic(nil)
	// Finaliza silenciosamente o processo
	os.Stdout.Close() // Evita saída residual
	os.Stderr.Close() // Evita saída de erro
	os.Exit(0)        // Sai silenciosamente
	return
}
