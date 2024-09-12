package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Config struct {
	Server *fiber.App
}

func (app *Config) SetupMiddleWare() {

}

func (app *Config) SetupRoutes() {
	apiGroup := app.Server.Group("/api/v1")

	// HealthCheck
	apiGroup.Get("/healthcheck", app.HealthCheck)

	// UserRoute
	userGroup := apiGroup.Group("/users")
	userGroup.Post("/", app.CreateUser)

}

func (app *Config) Listen(port string) {

	log.Printf("Listening Server: %s\n", port)

	if err := app.Server.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Println("Cannot Listening Server")
		panic(err)
	}
}
