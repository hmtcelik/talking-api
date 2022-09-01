package router

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

type HandledUser struct{
	Id int `json:"id"`
	ReversedName string `json:"name"`
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
			return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func SetupRoutes(app *fiber.App) {

	// base
	app.Get("/", func (c *fiber.Ctx) error{
		return c.SendString("HOMEPAGE")
	})
	
	// get json res
	app.Get("/getjson", func(c *fiber.Ctx) error {
		user := User{
			Id: 10,
			Name: "John",
		}
		c.JSON(user)
		return c.SendStatus(200)
	})

	app.Get("/reversed-body", func(c *fiber.Ctx) error{
		body := new(User)
		getJson("http://localhost:3000/getjson", body)
		handledUser := HandledUser{
			Id: body.Id * body.Id,
			ReversedName: Reverse(body.Name),
		}
		c.JSON(handledUser)
		return c.SendStatus((200))
	})
}
