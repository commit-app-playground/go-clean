package entity

import "errors"

type Task struct {
	id   *TaskId
	name *TaskName
	done bool
}

func NewTask(name *TaskName) *Task {
	return &Task{nil, name, false}
}

func NewTaskFromExisting(id *TaskId, name *TaskName, done bool) (*Task, error) {
	if id == nil {
		return nil, errors.New("nil task id")
	}

	if name == nil {
		return nil, errors.New("nil task name")
	}

	return &Task{id, name, done}, nil
}

func (t *Task) Id() *TaskId {
	return t.id
}

func (t *Task) Name() *TaskName {
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
