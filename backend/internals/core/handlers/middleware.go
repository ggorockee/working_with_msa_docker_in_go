package handlers

import (
	"back-end/internals/core/ports"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

type JWTHandler struct {
	jwtService ports.JWTService
}

func NewJWTHandler(jwtService ports.JWTService) *JWTHandler {
	return &JWTHandler{
		jwtService: jwtService,
	}
}

	

func (h *JWTHandler) AuthProtected() fiber.Handler{
	jwtConfig := h.jwtService.AuthProtected()
	return jwtware.New(*jwtConfig)
}


