package task

import (
	"fmt"
	"net/http"
	"net/url"

	"toporet/hop/goclean/usecase/task/create"
)

type CreateTaskPresenter struct {
	w http.ResponseWriter
}

func NewCreateTaskPresenter(w http.ResponseWriter) CreateTaskPresenter {
	return CreateTaskPresenter{w: w}
}

func (p CreateTaskPresenter) Present(o create.CreateTaskOut) {
	w := p.w

	statusCode, err := func() (int, error) {

		if yes, err := o.IsInputError(); yes {
			return http.StatusBadRequest, err
		} else if yes, err := o.IsDatabaseError(); yes {
			// OR could return http.StatusBadGateway
			return http.StatusInternalServerError, err
		} else {
			return http.StatusCreated, nil
		}
	}()

	if err != nil {

		w.WriteHeader(statusCode)
		w.Write([]byte(err.Error()))

	} else {

		id := url.PathEscape(o.TaskId())

		// TODO: the route /tasks can be a constant
		// declared somewhere in the common web api related code
		w.Header().Set("Location", fmt.Sprintf("/tasks/%s", id))

		w.WriteHeader(statusCode)
	}
}
