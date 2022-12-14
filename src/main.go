package main

import (
	"talkyapi/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)
	

	app.Listen(":3000")
}
