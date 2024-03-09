package main

import (
	"github.com/gofiber/fiber/v2"
)

func addRoutes(app *fiber.App) {
	app.Post("/getMangaTitlesByTitle", getMangaTitlesByTitle)
}
