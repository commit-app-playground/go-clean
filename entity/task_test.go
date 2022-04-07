package entity

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestDefaultTask(t *testing.T) {
	task := new(Task)

	assert.Assert(t, is.Nil(task.Id()))
	assert.Assert(t, is.Nil(task.Name()))
	assert.Equal(t, task.Done(), false)
}

func TestNewTask(t *testing.T) {
	id, _ := NewTaskId("id")
	name, _ := NewTaskName("Buy milk")

	task := NewTask(id, name)

	assert.Equal(t, task.Done(), false)
}

func TestMarkComplete(t *testing.T) {
	id, _ := NewTaskId("id")
	name, _ := NewTaskName("Buy milk")
	task := NewTask(id, name)

	task.MarkComplete()

	assert.Equal(t, task.Done(), true)
}

func TestMarkIncomplete(t *testing.T) {
	id, _ := NewTaskId("id")
	name, _ := NewTaskName("Buy milk")
	task := NewTask(id, name)
	task.MarkComplete()

	task.MarkIncomplete()

	assert.Equal(t, task.Done(), false)
}
