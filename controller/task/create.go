package task

import (
	"encoding/json"
	"errors"
	"net/http"
	"toporet/hop/goclean/usecase/task"
)

func create(w http.ResponseWriter, r *http.Request, f CreateTaskFactory) {

	parseRequest := func() (*task.CreateTaskIn, error) {
		var in task.CreateTaskIn
		if r.Body == nil {
			return nil, errors.New("Missing request body")
		}
		err := json.NewDecoder(r.Body).Decode(&in)
		return &in, err
	}

	in, err := parseRequest()

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	uc := f(w, r)
	uc.Handle(*in)
}
