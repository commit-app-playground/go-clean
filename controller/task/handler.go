package task

import (
	"net/http"
	"strings"
	"toporet/hop/goclean/controller"
	"toporet/hop/goclean/usecase/task"
)

type CreateTaskFactory = controller.UseCaseFactory[task.CreateTaskUseCase]

func Handle(f CreateTaskFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch strings.ToUpper(r.Method) {
		case http.MethodPost:

			create(w, r, f)

		default:
			http.NotFound(w, r)
		}
	}
}
