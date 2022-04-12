package create

import (
	"fmt"
	"toporet/hop/goclean/entity"
)

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

//
// Will panic if accessed when either of inputErr or dbErr are not nil
//
func (o CreateTaskOut) TaskId() string {
	panicWithError := func(err error) {
		panic(
			fmt.Sprintf(
				"invalid operation accessing task id in the presense of input error: %v",
				err))
	}
	if o.inputErr != nil {
		panicWithError(o.inputErr)

	}
	if o.dbErr != nil {
		panicWithError(o.dbErr)
	}
	return *o.taskId
}

func (o CreateTaskOut) IsInputError() (bool, error) {
	return o.inputErr != nil, o.inputErr
}

func (o CreateTaskOut) IsDatabaseError() (bool, error) {
	return o.dbErr != nil, o.dbErr
}

func (o CreateTaskOut) IsSuccess() bool {
	return o.taskId != nil
}
