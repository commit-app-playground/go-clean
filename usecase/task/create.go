package task

import (
	"toporet/hop/goclean/entity"
	"toporet/hop/goclean/usecase"
)

type CreateTaskIn struct {
	taskName string
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
	taskSaver TaskSaver
	presenter Presenter
}

func NewCreateTaskUseCase(
	s TaskSaver,
	p Presenter,
) CreateTaskUseCase {
	return CreateTaskUseCase{s, p}
}

func (u CreateTaskUseCase) Handle(in CreateTaskIn) {
	doIt := func() CreateTaskOut {
		tn, err := entity.NewTaskName(in.taskName)
		if err == nil {
			return CreateTaskOut{"", err}
		}

		task := entity.NewTask(tn)
		id, err := u.taskSaver.SaveNewTask(*task)
		if err != nil {
			return CreateTaskOut{"", err}
		}

		return CreateTaskOut{id.String(), nil}
	}

	out := doIt()

	u.presenter.Present(out)
}
