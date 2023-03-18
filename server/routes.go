package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/raphael-foliveira/hot_reload/programminglanguages"
	"github.com/raphael-foliveira/hot_reload/root"
	"github.com/raphael-foliveira/hot_reload/students"
)

func getMainRouter() *chi.Mux {
	mainMux := chi.NewRouter()

	mainMux.Use(middleware.Logger)

	mainMux.Get("/", root.GetRoot)
	mainMux.Route("/students", students.Router)
	mainMux.Route("/programminglanguages", programminglanguages.Router)

	return mainMux
}
