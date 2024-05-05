package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/template/html/v2"
)

func main() {

	app := fiber.New()

	app.Static("/static", "./web/styles/")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./web/pages/index.html", true)
	})

	err := app.Listen(":42069")
	if err != nil {
		log.Print("App cannot be started")
		return
	}
}
