package ports

import (
	"back-end/internals/core/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserRepository interface {
	Register(email, password string) error
	Login(email, password string) error
	Logout(id int) error
	ValidToken(t *jwt.Token, id string) bool
	validUser(id string, p string) bool
	GetFindById(id int) (*domain.User, error)
}

type UserService interface {
	Register(email, password, conformPassword string) error
}

type UserHandler interface {
	Register(c *fiber.Ctx) error
}
