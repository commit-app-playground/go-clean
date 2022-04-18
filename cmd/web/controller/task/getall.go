package task

import (
	"net/http"
	"toporet/hop/goclean/cmd/web/controller"
	"toporet/hop/goclean/pkg/usecase/task/getall"
)

type GetAllTasksUseCaseFactory controller.UseCaseFactory[getall.GetAllTasksUseCase]

func (f GetAllTasksUseCaseFactory) getAll(w http.ResponseWriter, r *http.Request) {

	// toUseCaseInput := func(p payload) (*getall.GetAllTasksIn, error) {
	// 	in, err := getall.NewGetAllTasksIn()

	// 	return &in, err
	// }
	// in, err := parser.ParseAndTranslate(r, toUseCaseInput)

	in, err := getall.NewGetAllTasksIn()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	uc := f(w, r)
	uc.Handle(in)
}
