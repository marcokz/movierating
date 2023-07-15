package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"movieraiting/database"
	"movieraiting/rest"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	pool, err := pgxpool.New(context.Background(), "postgres://postgres:pass@127.0.0.1:5432/todo")
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	db := database.NewDatabase(pool)
	handler := rest.NewHandler(db)

	router := chi.NewRouter()
	router.Post("/movie/users", handler.CreateUser)
	router.Get("/movie/getusers", handler.GetUsers)
}
