package bootstrap

import (
	"database/sql"
	"net/http"

	c "toporet/hop/goclean/controller/task"
	g "toporet/hop/goclean/gateway"
	p "toporet/hop/goclean/presenter/task"
	u "toporet/hop/goclean/usecase/task"
)

func Task(db *sql.DB) c.CreateTaskFactory {
	return func(w http.ResponseWriter, r *http.Request) u.CreateTaskUseCase {
		store := g.NewTaskStore(db)

		ucCreate := u.NewCreateTaskUseCase(store, p.NewCreateTaskPresenter(w))

		return ucCreate
	}
}
