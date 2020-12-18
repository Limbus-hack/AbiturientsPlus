package controller

import "github.com/code7unner/rest-api-template/internal/app"

type Controller struct {
	Posts PostsController
}

func New(app *app.App) *Controller {
	return &Controller{
		Posts: NewPostsController(app),
	}
}