package ports

import (
	"back-end/internals/core/domain"

	"github.com/gofiber/fiber/v2"
)

type AuthRepository interface {
	Login(email, password string) (*domain.User, error)
}

type AuthService interface {
	Login(email, password string) (*domain.User, error)
}

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}
