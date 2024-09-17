package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

const PORT string = "8000"

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Static("/assets",
		"./templates/assets")

	//app.Get("/signup", func(c *fiber.Ctx) error {
	//	return c.Render("signup", nil)
	//})

	app.Get("/signin", func(c *fiber.Ctx) error {
		return c.Render("signin", nil)
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup", nil)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{"Title": "Hello world"})
	})

	app.Listen(fmt.Sprintf(":%s", PORT))
}
