package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"toporet/hop/goclean/usecase/book"
)

func Books(u book.GetAllBooksUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.ToUpper(r.Method) != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		books, err := u.GetAllBooks()

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		} else {

			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]any)
			resp["status"] = http.StatusOK
			resp["statusText"] = "Status OK"
			resp["data"] = books

			err = json.NewEncoder(w).Encode(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.WriteHeader(http.StatusOK)
		}
	}
}
