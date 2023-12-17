package main

import (
	"fmt"
	"net/http"

	"github.com/MatheusMoronari/Desafio/configs"
	"github.com/MatheusMoronari/Desafio/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{"https://*", "http://*"},
			AllowedHeaders: []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		},
	))
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.List)
	r.Get("/{id}", handlers.Get)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
