package routes

import (
	"back-end/handlers"
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
	apiGroup.Post("/", handlers.CreateUser)
}

func (app *Config) Listen(port string, hostAll ...bool) {
	h := false // localhost
	if len(hostAll) > 0 {
		for _, host := range hostAll {
			h = host
		}
	}

	var listenPort string
	if !h {
		listenPort = fmt.Sprintf(":%s", port)
	} else {
		listenPort = fmt.Sprintf("0.0.0.0:%s", port)
	}

	log.Printf("Listening Server: %s\n", listenPort)

	if err := app.Server.Listen(listenPort); err != nil {
		log.Println("Cannot Listening Server")
		panic(err)
	}
}
