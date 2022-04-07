package task

import (
	"net/http"
	"toporet/hop/goclean/controller"
	"toporet/hop/goclean/usecase/task"
)

func create(w http.ResponseWriter, r *http.Request, f CreateTaskFactory) {

	in, err := controller.ParseRequestAs[task.CreateTaskIn](r)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	uc := f(w, r)
	uc.Handle(in)
}
