package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"toporet/hop/goclean/controller"
	"toporet/hop/goclean/gateway"
	"toporet/hop/goclean/usecase"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:Password1@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	booksRepository := gateway.NewBookRepository(db)
	useCase := usecase.NewAllBooksUseCase(booksRepository)

	mux := http.NewServeMux()

	mux.HandleFunc("/books", controller.Books(useCase))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
