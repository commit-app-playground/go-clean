package create

import (
	"errors"
	"testing"

	"gotest.tools/assert"
)

func TestTaskId_Panic_When_InputError_IsNot_Nil(t *testing.T) {
	out := CreateTaskOut{}.InputError(
		errors.New("invalid input"))

	arg := shouldPanic(t, func() {
		out.TaskId()
	})

	s, ok := arg.(string)
	assert.Check(t, ok)
	assert.Equal(t, s,
		"invalid operation accessing task id "+
			"in the presense of input error: invalid input")
}

func TestTaskId_Panic_When_DatabaseError_IsNot_Nil(t *testing.T) {
	out := CreateTaskOut{}.DatabaseError(
		errors.New("save failed"))

	arg := shouldPanic(t, func() {
		out.TaskId()
	})

	s, ok := arg.(string)
	assert.Check(t, ok)
	assert.Equal(t, s,
		"invalid operation accessing task id "+
			"in the presense of input error: save failed")
}

func shouldPanic(t *testing.T, f func()) (panicArg any) {
	t.Helper()
	defer func() {
		panicArg = recover()
	}()
	f()
	t.Errorf("should have panicked")
	return
}
