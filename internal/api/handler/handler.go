package handler

import (
	"github.com/code7unner/vk-scrapper/internal/api/controller"
	"github.com/code7unner/vk-scrapper/internal/app"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"net/http"
	"os"
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
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Accept-Encoding", "Authorization"},
	}).Handler)

	ctrl := controller.New(app)
	fs := http.FileServer(http.Dir("./public"))
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat("./public" + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
	r.MethodFunc("GET", "/ping", ctrl.Vk.Ping)
	r.MethodFunc("POST", "/prediction", ctrl.Prediction.Get)
	r.MethodFunc("PATCH", "/status", ctrl.Prediction.UpdateStatus)
	r.MethodFunc("GET", "/prediction", ctrl.Prediction.GetCached)

	return &Handler{
		Mux: r,
		app: app,
	}
}
