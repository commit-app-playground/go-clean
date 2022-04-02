package presenter

import (
	"net/http"
	"toporet/hop/goclean/usecase"
)

type HttpPresenter[TOut any] interface {
	usecase.Presenter[TOut]
	WriteResponse(w http.ResponseWriter)
}
