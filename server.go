package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/template/handlebars/v2"
)

type Body struct {
	Text string `form:"text"`
}

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
	}).Name("index")

	app.Post("/command", func(c *fiber.Ctx) error {
		body := new(Body)

		if err := c.BodyParser(body); err != nil {
			return err
		}

		cmd := strings.Trim(body.Text, " ");

		if cmd == "" {
			return c.SendString("")
		}

		if strings.HasPrefix(cmd, "ls") {
			return c.Render("commands/ls", fiber.Map{})
		}

		return c.SendString(strings.Split(cmd, " ")[0] + ": command not found")
	}).Name("command")

	log.Fatal(app.Listen(":3000"))
}
