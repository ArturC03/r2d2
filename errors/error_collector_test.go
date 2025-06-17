package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorCollector_Add(t *testing.T) {
	t.Parallel()

	ec := ErrorCollector{}

	ec.Add("message", 1)

	assert.Equal(t, 1, len(ec.Errors))
	assert.Equal(t, "message", ec.Errors[0].Message)
	assert.Equal(t, 1, ec.Errors[0].Line)
}

func TestErrorCollector_GetErrors(t *testing.T) {
	t.Parallel()

	ec := ErrorCollector{}

	ec.Add("message", 1)

	errors := ec.Errors

	assert.Equal(t, 1, len(errors))
	assert.Equal(t, "message", errors[0].Message)
	assert.Equal(t, 1, errors[0].Line)
}

func TestErrorCollector_HasErrors(t *testing.T) {
	t.Parallel()

	ec := ErrorCollector{}
	assert.False(t, ec.HasErrors())

	ec.Add("message", 1)
	assert.True(t, ec.HasErrors())
}
