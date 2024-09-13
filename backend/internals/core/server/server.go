package server

import (
	"back-end/internals/core/ports"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	userHandler ports.UserHandler
	// middlewares ports.Middlewares
}

func NewServer(uHandler ports.UserHandler) *Server {
	return &Server{
		userHandler: uHandler,
	}
}

func (s *Server) Initialize(port ...string) {
	var p string

	switch len(port) {
	case 0:
		p = "3000"
	case 1:
		p = port[0]
	default:
		log.Println("port argument is just one!")
		os.Exit(1)
	}

	app := fiber.New()
	v1 := app.Group("/api/v1")

	userRoutes := v1.Group("/user")
	userRoutes.Post("/", s.userHandler.Register)

	err := app.Listen(fmt.Sprintf(":%s", p))
	if err != nil {
		log.Fatal(err)
	}
}
