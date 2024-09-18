package services

import (
	"back-end/internals/core/helpers"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

type JWTService struct {
	jwtConfig jwtware.Config
}

func NewJWTService(signingKey string, errorHandler ...func(c *fiber.Ctx, err error) error) *JWTService{
	var jwtConfig jwtware.Config
	if len(errorHandler) > 0 {
		jwtConfig = jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte(signingKey)},
			ErrorHandler: errorHandler[0],
		}

		return &JWTService{
			jwtConfig: jwtConfig,
		}
	}

	jwtConfig = jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(signingKey)},
		ErrorHandler: jwtError,
	}

	return &JWTService{
		jwtConfig: jwtConfig,
	}
}


func (s *JWTService) AuthProtected() *jwtware.Config{
	return &s.jwtConfig
}


func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		jsonResponse := helpers.JsonResponse{
			Error:   true,
			Message: "Missing or malformed JWT",
			Data:    nil,
		}
		return c.Status(fiber.StatusBadRequest).
			JSON(jsonResponse)
	}

	jsonResponse := helpers.JsonResponse{
		Error:   true,
		Message: "Invalid or expired JWT",
		Data:    nil,
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(jsonResponse)
}