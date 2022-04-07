package task

import (
	"testing"
	"toporet/hop/goclean/entity"

	"gotest.tools/assert"
)

type MockNewTaskSaver struct{}

type MockPresenter struct {
	out CreateTaskOut
}

func TestHandle(t *testing.T) {
	s := MockNewTaskSaver{}
	p := MockPresenter{}
	uc := NewCreateTaskUseCase(s, p)

	uc.Handle(CreateTaskIn{})

	got := p.Received()

	want := CreateTaskOut{}
	assert.Equal(t, got, want)
}

func (s MockNewTaskSaver) SaveNewTask(t *entity.Task) (*entity.TaskId, error) {
	return nil, nil
}

func (p MockPresenter) Present(out CreateTaskOut) {
	p.out = out
}

func (p *MockPresenter) Received() CreateTaskOut {
	return p.out
}
