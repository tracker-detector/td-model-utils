package main

import (
	"github.com/Tracking-Detector/td-model-utils/env"
	"github.com/Tracking-Detector/td-model-utils/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	handlers.NewModelHandler(app).RegisterRoutes()
	handlers.NewHomeHandler(app).RegisterRoutes()
	app.Listen(env.GetPort())
}
