package entity

type Task struct {
	id   *TaskId
	name *TaskName
	done bool
}

func NewTask(id *TaskId, name *TaskName) *Task {
	return &Task{id, name, false}
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
