package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	engine := handlebars.New("./templates", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index.html", fiber.Map{
			"Page": "~",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
