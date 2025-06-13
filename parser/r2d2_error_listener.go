package parser

import (
	"fmt"
	"os"

	"github.com/ArturC03/r2d2/errors"
	"github.com/ArturC03/r2d2Styles"
	"github.com/antlr4-go/antlr/v4"
)

// R2D2ErrorListener is a listener that listens to errors emitted by the parser.
type R2D2ErrorListener struct {
	*antlr.DefaultErrorListener
	ErrorCollector *errors.ErrorCollector
}

func NewR2D2ErrorListener(errorCollector *errors.ErrorCollector) *R2D2ErrorListener {
	return &R2D2ErrorListener{antlr.NewDefaultErrorListener(), errorCollector}
}

func (l *R2D2ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Println(r2d2Styles.ErrorMessage(fmt.Sprintf("Syntax error at %s: %s", r2d2Styles.Bold(fmt.Sprintf("line %d:%d", line, column)), msg)))

	l.ErrorCollector.Add(msg, line)
	// panic(nil)
	// Finaliza silenciosamente o processo
	os.Stdout.Close() // Evita saída residual
	os.Stderr.Close() // Evita saída de erro
	os.Exit(0)        // Sai silenciosamente
	return
}
