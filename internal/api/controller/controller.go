package controller

import "github.com/code7unner/vk-scrapper/internal/app"

type Controller struct {
	Vk VkController
}

func New(app *app.App) *Controller {
	return &Controller{
		Vk: NewVkController(app),
	}
}