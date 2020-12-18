package controller

import (
	"github.com/code7unner/vk-scrapper/internal/app"
	"net/http"
)

type PostsController interface {
	GetByID(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)

	Ping(w http.ResponseWriter, r *http.Request)
}

type postsCtrl struct {
	app *app.App
}

func NewPostsController(app *app.App) PostsController {
	return &postsCtrl{app}
}

func (p postsCtrl) GetByID(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (p postsCtrl) GetAll(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (p postsCtrl) Create(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (p postsCtrl) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ping"))
}
