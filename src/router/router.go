package router

import (
	"github.com/gofiber/fiber/v2"
)

type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

func SetupRoutes(app *fiber.App) {

	// base
	app.Get("/", func (c *fiber.Ctx) error{
		return c.SendString("HOMEPAGE")
	})
	
	// get json res
	app.Get("/getjson", func(c *fiber.Ctx) error {
		user := User{
			Id: 1,
			Name: "John",
		}
		c.JSON(user)
		return c.SendStatus(200)
	})
}