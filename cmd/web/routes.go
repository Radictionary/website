package main

import (
	"net/http"

	"github.com/Radictionary/website/pkg/config"
	"github.com/Radictionary/website/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	if app.InProduction {
		mux.Use(middleware.Recoverer)
	}
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/delete", handlers.Repo.Delete)


	fileServer := http.FileServer(http.Dir("./frontend/"))
	mux.Handle("/frontend/*", http.StripPrefix("/frontend", fileServer))

	return mux
}
