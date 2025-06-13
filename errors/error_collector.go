package errors

type CompileError struct {
	Message string
	Line    int
}

type ErrorCollector struct {
	Errors []CompileError
}

func (ec *ErrorCollector) Add(message string, line int) {
	ec.Errors = append(ec.Errors, CompileError{
		Message: message,
		Line:    line,
	})
}

func (ec *ErrorCollector) HasErrors() bool {
	return len(ec.Errors) > 0
}
