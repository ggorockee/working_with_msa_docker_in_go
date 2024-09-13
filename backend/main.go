package main

import (
	"back-end/internals/core/handlers"
	"back-end/internals/core/repositories"
	"back-end/internals/core/server"
	"back-end/internals/core/services"
)

func main() {
	// repositories
	userRepository := repositories.NewUserRepository()

	// services
	userService := services.NewUserService(userRepository)

	// handlers
	userHandlers := handlers.NewUserHandlers(userService)

	// server
	httpServer := server.NewServer(
		userHandlers,
	)

	httpServer.Initialize()
}
