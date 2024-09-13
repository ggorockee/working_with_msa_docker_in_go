package ports

import "github.com/gofiber/fiber/v2"

type UserService interface {
	Register(email string, password string, passwordConfirm string) error
}

type UserRepository interface {
	Register(email string, password string) error
}

type UserHandler interface {
	Register(c *fiber.Ctx) error
}
