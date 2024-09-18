package ports

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)



type JWTService interface {
	AuthProtected() *jwtware.Config
}


type JWTHandler interface {
	AuthProtected() fiber.Handler 
}