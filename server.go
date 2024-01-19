package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	engine := handlebars.New("./templates", ".hbs")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(compress.New())

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Page": "~",
		}, "layout")
	})

	log.Fatal(app.Listen(":3000"))
}
