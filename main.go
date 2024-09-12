package main

import (
	"back-end/routes"
	"github.com/gofiber/fiber/v2"
)

const (
	PORT string = "3000"
)

func main() {
	app := routes.Config{
		Server: fiber.New(),
	}

	app.SetupMiddleWare()
	app.SetupRoutes()
	app.Listen(PORT, true)
}
