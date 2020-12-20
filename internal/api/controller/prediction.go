package controller

import (
	"encoding/json"
	"errors"
	"github.com/code7unner/vk-scrapper/internal/api/service"
	"github.com/code7unner/vk-scrapper/internal/app"
	"io/ioutil"
	"net/http"
	"strconv"
)

type PredictionController interface {
	GetInRealTime(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetCached(w http.ResponseWriter, r *http.Request)
	UpdateStatus(w http.ResponseWriter, r *http.Request)
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

func (p PredictionCtrl) GetInRealTime(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	keys, ok := params["city"]
	if !ok {
		p.error(w, r, http.StatusBadRequest, errors.New("city query is required"))
	}

	city, _ := strconv.Atoi(keys[0])
	users, err := service.GetVkUsers(city, &p.app.Conf)
	if err != nil {
		p.error(w, r, http.StatusInternalServerError, err)
	}
	subs, err := service.BulkGetVkUserSubs(users, &p.app.Conf)
	if err != nil {
		p.error(w, r, http.StatusInternalServerError, err)
	}

	p.respond(w, r, http.StatusOK, subs)
}

func (p PredictionCtrl) GetCached(w http.ResponseWriter, r *http.Request) {
	var city int
	params := r.URL.Query()
	cityKey, ok := params["city"]
	if ok {
		city, _ = strconv.Atoi(cityKey[0])
	} else {
		city = 0
	}
	schoolKey, ok := params["school"]
	if !ok {
		p.error(w, r, http.StatusBadRequest, errors.New("school query is required"))
	}

	school, _ := strconv.Atoi(schoolKey[0])

	users, err := p.app.Repo.Users.Retrieve(p.app.Ctx, city, school)
	if err != nil {
		p.error(w, r, http.StatusInternalServerError, err)
	}

	p.respond(w, r, http.StatusOK, users)
}

func (p PredictionCtrl) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id     int    `json:"id"`
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		p.error(w, r, http.StatusInternalServerError, err)
	}
	defer r.Body.Close()

	rowsUpdated, err := p.app.Repo.Users.Update(p.app.Ctx, req.Id, req.Status)
	if err != nil {
		p.error(w, r, http.StatusInternalServerError, err)
	}

	p.respond(w, r, http.StatusOK, rowsUpdated)
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
