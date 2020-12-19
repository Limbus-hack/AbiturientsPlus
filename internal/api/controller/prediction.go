package controller

import (
	"encoding/json"

	"github.com/code7unner/vk-scrapper/internal/app"
	"io/ioutil"
	"net/http"
)

type PredictionController interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type PredictionCtrl struct {
	app *app.App
}

func NewPredictionController(app *app.App) PredictionController {
	return &PredictionCtrl{app}
}

func (p PredictionCtrl) Get(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		p.error(w, r, http.StatusBadRequest, err)
	}
	r.Body.Close()

	predict, err := p.app.Vws.Predict(string(data))
	if err != nil {
		p.error(w, r, http.StatusInternalServerError, err)
	}

	p.respond(w, r, http.StatusOK, predict)
}

// respond with error
func (p PredictionCtrl) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	p.respond(w, r, code, map[string]string{"error": err.Error()})
}

// abstract respond
func (p PredictionCtrl) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
