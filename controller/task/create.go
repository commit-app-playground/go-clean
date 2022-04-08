package task

import (
	"net/http"
	"toporet/hop/goclean/controller"
	"toporet/hop/goclean/controller/parser"
	"toporet/hop/goclean/usecase/task/create"
)

type payload struct {
	Name string
}

type CreateTaskUseCaseFactory controller.UseCaseFactory[create.CreateTaskUseCase]

func (f CreateTaskUseCaseFactory) create(w http.ResponseWriter, r *http.Request) {

	in, err := parser.ParseAndTranslate(r, toInput)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	uc := f(w, r)
	uc.Handle(*in)
}

func toInput(p payload) (*create.CreateTaskIn, error) {
	return create.NewCreateTaskIn(p.Name)
}
