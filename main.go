package main

import (
	"fmt"
	"github.com/ArturC03/r2d2/lexer"
)

func main() {
	source := `
	import "function" from "path"
	module cookie {
		export fn main(argv: string[]): int {
			const constante = 34;
			var variavel_mut = constante;
			let variavel_imut = variavel_mut;

			if (i >= 2) {
			} else if (i == 2) {
			} else {
			}

			loop {
				break;
			}

			for (var i = 0; i < 3; i++) {
			}
		}
		export pseudo fn PseudoFuncao() {}
	}`
	l := lexer.NewLexer(source)
	for {
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)
		if tok.Type == lexer.TokenEOF {
			break
		}
	}
}
