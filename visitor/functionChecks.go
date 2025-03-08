package visitor

import (
	parser "github.com/ArturC03/r2d2/parser"
	"github.com/antlr4-go/antlr/v4"
	"reflect"
)

func isExported(ctx any) bool {
	if funcDecl, ok := ctx.(*parser.FunctionDeclarationContext); ok {
		if funcDecl.EXPORT() != nil {
			return true
		}
	}

	if varDecl, ok := ctx.(*parser.VariableDeclarationContext); ok {
		if varDecl.EXPORT() != nil {
			return true
		}
	}
	return false
}

func isPseudo(ctx any) bool {
	if funcDecl, ok := ctx.(*parser.FunctionDeclarationContext); ok {
		if funcDecl.PSEUDO() != nil {
			return true
		}
	}
	return false
}

func (module Module) Exports() []string {
	var exports []string

	for _, function := range module.Functions {
		if function.isExported {
			exports = append(exports, function.Name)
		}
	}

	for _, variable := range module.Variables {
		if variable.isExported {
			exports = append(exports, variable.Name)
		}
	}
	return exports
}

func findParent(ctx antlr.Tree, targetTypes ...any) bool {
	parent := ctx.GetParent()
	for parent != nil {
		parentType := reflect.TypeOf(parent)
		for _, targetType := range targetTypes {
			if reflect.TypeOf(targetType) == parentType {
				return true
			}
		}
		parent = parent.GetParent()
	}
	return false
}

func findChild(ctx antlr.Tree, targets ...any) bool {
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)

		// Verifica se o filho é do tipo desejado
		for _, target := range targets {
			if reflect.TypeOf(child) == reflect.TypeOf(target) {
				return true
			}
		}

		// Chama recursivamente para procurar nos filhos
		if findChild(child, targets...) {
			return true
		}
	}
	return false
}
