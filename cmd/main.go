package main

import (
	"github.com/Tracking-Detector/td-model-utils/env"
	"github.com/Tracking-Detector/td-model-utils/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/static", "./static")
	handlers.NewModelHandler(app).RegisterRoutes()
	handlers.NewHomeHandler(app).RegisterRoutes()
	app.Listen(env.GetPort())
}
