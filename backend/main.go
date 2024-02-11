package main

import (
	"api/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controller.Index)
	r.Get("/search", controller.Search)

	http.ListenAndServe(":8080", r)
}
