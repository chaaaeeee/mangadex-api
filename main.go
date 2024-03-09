package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const baseUrl = "https://api.mangadex.org"

func main() {
	app := fiber.New()

	addRoutes(app)

	fmt.Println("alr started")
	app.Listen(":9000")
}
