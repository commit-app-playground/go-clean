package task

import (
	"net/http"
	"toporet/hop/goclean/controller"
	"toporet/hop/goclean/usecase/task"
)

func create(w http.ResponseWriter, r *http.Request, f CreateTaskFactory) {

	in, err := func() (*task.CreateTaskIn, error) {
		p, err := controller.ParseRequestAs[payload](r)
		if err != nil {
			return nil, err
		}
		return task.NewCreateTaskIn(p.Name)
	}()

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	uc := f(w, r)
	uc.Handle(*in)
}

type payload struct {
	Name string
}
