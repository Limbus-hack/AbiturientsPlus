package handler

import (
	"github.com/code7unner/rest-api-template/internal/api/controller"
	"github.com/code7unner/rest-api-template/internal/app"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"net/http"
)

type Handler struct {
	*chi.Mux
	app *app.App
}

func New(app *app.App) *Handler {
	r := chi.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Accept-Encoding", "Authorization"},
	}).Handler)

	ctrl := controller.New(app)
	r.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("./public"))))
	r.MethodFunc("GET", "/ping", ctrl.Posts.Ping)

	return &Handler{
		Mux: r,
		app: app,
	}
}
