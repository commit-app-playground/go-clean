package task

import (
	"encoding/json"
	"log"
	"net/http"

	"toporet/hop/goclean/usecase/task/create"
)

type CreateTaskPresenter struct {
	w http.ResponseWriter
}

func NewCreateTaskPresenter(w http.ResponseWriter) CreateTaskPresenter {
	return CreateTaskPresenter{w: w}
}

func (p CreateTaskPresenter) Present(o create.CreateTaskOut) {
	taskId := o.TaskId()
	err := o.Error()
	w := p.w

	if err != nil {

		// TODO: interpret the error, e.g. duplicate task name => 400, not 500
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	} else {

		w.Header().Set("Content-Type", "application/json")

		// TODO: set location header using the taskId

		w.WriteHeader(http.StatusOK)

		resp := make(map[string]any)
		resp["status"] = http.StatusOK
		resp["statusText"] = "Status OK"
		resp["data"] = taskId

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
	}
}
