package create

import "toporet/hop/goclean/entity"

type CreateTaskOut struct {
	taskId   *string
	inputErr error
	dbErr    error
}

func (o CreateTaskOut) InputError(err error) CreateTaskOut {
	return CreateTaskOut{inputErr: err}
}

func (o CreateTaskOut) DatabaseError(err error) CreateTaskOut {
	return CreateTaskOut{dbErr: err}
}

func (o CreateTaskOut) Success(id *entity.TaskId) CreateTaskOut {
	idStr := id.String()
	return CreateTaskOut{taskId: &idStr}
}

func (o CreateTaskOut) TaskId() string {
	//
	// will panic if accessed when either
	// of the inputErr or dbErr aren't nil
	//
	return *o.taskId
}

func (o CreateTaskOut) IsInputError() (bool, error) {
	return o.inputErr != nil, o.inputErr
}

func (o CreateTaskOut) IsDatabaseError() (bool, error) {
	return o.dbErr != nil, o.dbErr
}
