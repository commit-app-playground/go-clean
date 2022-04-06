package entity

import (
	"testing"

	"gotest.tools/assert"
)

func check[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Error(got, "!=", want)
	}
}

func TestDefaultTask(t *testing.T) {
	var task Task

	check(t, task.Name(), "")
	check(t, task.Id(), 0)
	check(t, task.Done(), false)
}

func TestNewTask_ValidName(t *testing.T) {
	task, err := NewTask(123, "Buy milk")

	assert.NilError(t, err)

	check(t, task.Id(), 123)
	check(t, task.Name(), "Buy milk")
	check(t, task.Done(), false)
}

func TestNewTask_InvalidName(t *testing.T) {
	task, err := NewTask(123, " \t ")

	assert.Error(t, err, "task name is empty or whitespace (\" \\t \")")

	check(t, task, nil)
}

func TestMarkComplete(t *testing.T) {
	task, _ := NewTask(123, "Buy milk")

	task.MarkComplete()

	check(t, task.Done(), true)
}

func TestMarkIncomplete(t *testing.T) {
	task, _ := NewTask(123, "Buy milk")
	task.MarkComplete()

	task.MarkIncomplete()

	check(t, task.Done(), false)
}

// func TestNewTask(t *testing.T) {
// 	id := TaskId(42)
// 	name := TaskName("Foo Bar")

// 	todo := NewTask(&id, &name, true)

// 	check(t, todo.Id(), id)
// 	check(t, todo.Name(), "Foo Bar")
// 	check(t, todo.Done(), true)
// }

// func TestNewTask_InvalidName(t *testing.T) {
// 	id := TaskId(42)
// 	name := TaskName(" \t ")

// 	todo := NewTask(&id, &name, true)

// 	check(t, todo.Id(), 42)
// 	check(t, todo.Name(), " \t ")
// 	check(t, todo.Done(), true)
// }
