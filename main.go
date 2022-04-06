package main

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"toporet/hop/goclean/bootstrap"
	"toporet/hop/goclean/controller"
	// "toporet/hop/goclean/controller/task"
)

func main() {
	db, err := sql.Open("postgres",
		"postgres://postgres:Password1@localhost/bookstore?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	driver.ErrBadConn = errors.New("sdfgh")

	mux := http.NewServeMux()

	getAll := bootstrap.Book(db)

	mux.HandleFunc("/books", controller.Books(getAll))
	// mux.HandleFunc("/tasks", task.Handler(getAll))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
