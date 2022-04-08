package task

import (
	"fmt"
	"toporet/hop/goclean/entity"
	"toporet/hop/goclean/usecase"
)

type CreateTaskIn struct {
	taskName string
}

func NewCreateTaskIn(taskName string) (*CreateTaskIn, error) {
	if len(taskName) == 0 {
		return nil, fmt.Errorf("task name is required but got empty")
	}
	return &CreateTaskIn{taskName}, nil
}

type CreateTaskOut struct {
	taskId string
	err    error
}

func (o CreateTaskOut) TaskId() string {
	return o.taskId
}

func (o CreateTaskOut) Error() error {
	return o.err
}

type Presenter = usecase.Presenter[CreateTaskOut]

type CreateTaskUseCase struct {
	taskSaver NewTaskSaver
	presenter Presenter
}

func NewCreateTaskUseCase(
	s NewTaskSaver,
	p Presenter,
) CreateTaskUseCase {
	return CreateTaskUseCase{s, p}
}

func (u CreateTaskUseCase) Handle(in CreateTaskIn) {
	out := func() CreateTaskOut {
		tn, err := entity.NewTaskName(in.taskName)
		if err != nil {
			return CreateTaskOut{"", err}
		}

		task := entity.NewTask(tn)
		id, err := u.taskSaver.SaveNewTask(task)
		if err != nil {
			return CreateTaskOut{"", err}
		}

		return CreateTaskOut{id.String(), nil}
	}()

	u.presenter.Present(out)
}
