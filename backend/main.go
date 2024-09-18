package main

import (
	"back-end/configs"
	"back-end/database"
	"back-end/internals/core/handlers"
	"back-end/internals/core/repositories"
	"back-end/internals/core/server"
	"back-end/internals/core/services"
)

func main() {
	database.Connect()
	userRepository := repositories.NewUserRepository()

	userService := services.NewUserService(userRepository)
	userHandlers := handlers.NewUserHandler(userService)

	healthCheckHandlers := handlers.NewHealthCheckHandler()

	memoRepository := repositories.NewMemoRepository()
	memoService := services.NewMemoService(memoRepository)
	memoHandlers := handlers.NewMemoHandler(memoService)

	
	jwtService := services.NewJWTService(configs.New().JWTSecret)
	midlewWareHandler := handlers.NewJWTHandler(jwtService)

	httpServer := server.NewServer(
		userHandlers,
		healthCheckHandlers,
		memoHandlers,
		midlewWareHandler,
	)
	httpServer.SetupRoute()
	httpServer.Listen()
}
