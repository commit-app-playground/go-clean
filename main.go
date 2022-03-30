package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	"toporet/hop/goclean/entity"
	"toporet/hop/goclean/infrastructure"
	"toporet/hop/goclean/usecase"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:Password1@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	userRepository := infrastructure.NewUserRepository(db)
	useCase := usecase.NewUserRegisterUseCase(userRepository)

	mux := http.NewServeMux()

	mux.HandleFunc("/", helloWorld(useCase))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func helloWorld(u *usecase.UserRegisterUseCase) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		u.RegisterUser(entity.User{})
		io.WriteString(rw, `{"message": "hello world..."}`)
	}
}
