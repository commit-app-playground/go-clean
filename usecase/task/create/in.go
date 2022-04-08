package create

import "fmt"

type CreateTaskIn struct {
	taskName string
}

func NewCreateTaskIn(taskName string) (*CreateTaskIn, error) {
	if taskName == "" {
		return nil, fmt.Errorf("task name is required but got empty")
	}
	return &CreateTaskIn{taskName}, nil
}
