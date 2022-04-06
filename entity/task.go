package entity

import (
	"fmt"
	"strings"
)

type Task struct {
	id   int
	name string
	done bool
}

func NewTask(id int, name string) (*Task, error) {
	trim := strings.TrimSpace(name)

	if trim == "" {
		return nil, fmt.Errorf("task name is empty or whitespace (%q)", name)
	}
	return &Task{id, name, false}, nil
}

func (t *Task) Id() int {
	return t.id
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) Done() bool {
	return t.done
}

func (t *Task) MarkComplete() {
	t.done = true
}

func (t *Task) MarkIncomplete() {
	t.done = false
}
