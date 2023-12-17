package main

import (
	"fmt"
	"net/http"

	"github.com/MatheusMoronari/Desafio/configs"
	"github.com/MatheusMoronari/Desafio/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Post("/", CORS(handlers.Create))
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.List)
	r.Get("/{id}", handlers.Get)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
