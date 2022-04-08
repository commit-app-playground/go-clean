package task

import (
	"net/http"
	"toporet/hop/goclean/controller/parser"
	"toporet/hop/goclean/usecase/task"
)

type payload struct {
	Name string
}

func create(w http.ResponseWriter, r *http.Request, f CreateTaskFactory) {

	in, err := parser.ParseAndTranslate(r, toInput)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	uc := f(w, r)
	uc.Handle(*in)
}

func toInput(p payload) (*task.CreateTaskIn, error) {
	return task.NewCreateTaskIn(p.Name)
}
