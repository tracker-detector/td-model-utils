package handlers

import (
	"github.com/Tracking-Detector/td-model-utils/views"
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	app *fiber.App
}

func NewHomeHandler(app *fiber.App) *HomeHandler {
	return &HomeHandler{app: app}
}

func (hh *HomeHandler) GetHome(c *fiber.Ctx) error {
	return Render(c, views.Manual())
}

func (hh *HomeHandler) RegisterRoutes() {
	hh.app.Get("/", hh.GetHome)
}
