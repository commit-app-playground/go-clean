package create

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
