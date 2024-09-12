package main

import (
	"back-end/database"
	"back-end/routes"
	"github.com/gofiber/fiber/v2"
)

const (
	PORT string = "3000"
)

func main() {
	database.ConnectDB()
	app := routes.Config{
		Server: fiber.New(),
	}

	app.SetupMiddleWare()
	app.SetupRoutes()
	app.Listen(PORT)

}
