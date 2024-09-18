package server

import (
	"back-end/internals/core/ports"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	userHandler        ports.UserHandler
	healthCheckHandler ports.HealthCheckHandler
	memoHandler        ports.MemoHandler
	jwtHandler	ports.JWTHandler
}

const PORT = "3000"

var (
	server *Server
	app    *fiber.App
)

//var once sync.Once

func NewServer(
	userHandler ports.UserHandler,
	healthCheckHandler ports.HealthCheckHandler,
	memoHandler ports.MemoHandler,
	jwtHandler ports.JWTHandler,
) *Server {
	server = &Server{
		userHandler:        userHandler,
		healthCheckHandler: healthCheckHandler,
		memoHandler:        memoHandler,
		jwtHandler: jwtHandler,
	}
	return server
}

func (s *Server) SetupRoute() {
	app = fiber.New()

	v1 := app.Group("/api/v1")

	// healthCheck
	v1.Get("/healthcheck", s.healthCheckHandler.HealthCheck)

	// userroute
	userRoutes := v1.Group("/users")
	userRoutes.Post("/", s.userHandler.Register)
	userRoutes.Put("/:userId", s.userHandler.Update)
	userRoutes.Patch("/:userId", s.jwtHandler.AuthProtected(), s.userHandler.Update)
	userRoutes.Post("/login", s.userHandler.Login)

	// memoroute
	memoRoutes := v1.Group("/memos")
	memoRoutes.Post("/", s.jwtHandler.AuthProtected(), s.memoHandler.Create)
	memoRoutes.Get("/", s.memoHandler.GetAll)
	memoRoutes.Put("/:memoId", s.memoHandler.Update)
	memoRoutes.Patch("/:memoId", s.memoHandler.Update)
	memoRoutes.Delete("/:memoId", s.memoHandler.Delete)
}

func (s *Server) Listen(port ...string) {
	var p string

	switch len(port) {
	case 0:
		log.Println("default port: 3000")
		p = "3000" // default
	case 1:
		log.Printf("listening port %s\n", port[0])
		p = port[0]
	default:
		log.Println("argument number just one!")
		os.Exit(1)
	}

	app.Listen(fmt.Sprintf(":%s", p))
}
