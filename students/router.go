package students

import (
	"github.com/go-chi/chi"
)

func Router(router chi.Router) {
	router.Get("/", ListHandler)
	router.Post("/", CreateHandler)
	router.Get("/{id}", RetrieveHandler)
	router.Delete("/{id}", DeleteHandler)
}
