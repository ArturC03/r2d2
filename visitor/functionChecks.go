package visitor

import (
	parser "github.com/ArturC03/r2d2/parser"
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
