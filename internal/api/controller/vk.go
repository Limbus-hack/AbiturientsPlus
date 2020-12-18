package controller

import (
	"github.com/code7unner/vk-scrapper/internal/app"
	"net/http"
)

type VkController interface {
	GetByID(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)

	Ping(w http.ResponseWriter, r *http.Request)
}

type vkCtrl struct {
	app *app.App
}

func NewVkController(app *app.App) VkController {
	return &vkCtrl{app}
}

func (p vkCtrl) GetByID(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (p vkCtrl) GetAll(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (p vkCtrl) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (p vkCtrl) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ping"))
}
