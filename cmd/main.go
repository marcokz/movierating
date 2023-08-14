package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"movieraiting/database"
	"movieraiting/rest"
	"net/http"
)

func main() {
	pool, err := database.Connect(context.Background(), "postgres://postgres:pass@127.0.0.1:5432/todo")
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	db := database.NewDatabase(pool)
	handler := rest.NewHandler(db)

	router := chi.NewRouter()
	router.Post("/users/signup", handler.SignUp)
	router.Get("/users/signin", handler.SignIn)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
