package controller

import (
	"encoding/json"
	"github.com/code7unner/vk-scrapper/internal/api/repository"
	"github.com/code7unner/vk-scrapper/internal/app"
	"net/http"
	"strconv"
)

type PredictionController interface {
	GetWithFilter(w http.ResponseWriter, r *http.Request)
}

type PredictionCtrl struct {
	app *app.App
}

func NewPredictionController(app *app.App) PredictionController {
	return &PredictionCtrl{app}
}

func (p PredictionCtrl) GetWithFilter(w http.ResponseWriter, r *http.Request) {
	keys, _ := r.URL.Query()["city"]
	city, _ := strconv.Atoi(keys[0])
	//school := keys[1]
	vkUserImpl := repository.NewVkUserImpl(p.app.Conf)
	users, _ := vkUserImpl.GetVkUsers(city)
	p.respond(w, r, http.StatusOK, users)
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
