package main

import (
	"github.com/Tracking-Detector/td-model-utils/env"
	"github.com/Tracking-Detector/td-model-utils/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 8 * 1024 * 1024 * 1024, // 8 GB
	})
	app.Static("/static", "./static")
	handlers.NewModelHandler(app).RegisterRoutes()
	handlers.NewHomeHandler(app).RegisterRoutes()
	app.Listen(env.GetPort())
}
