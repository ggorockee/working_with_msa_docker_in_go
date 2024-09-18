package ports

import (
	"back-end/internals/core/domain"
	"back-end/internals/core/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserRepository interface {
	GetModel() *domain.User
	Register(email, password string) error
	// Login(email, password string) error
	ValidToken(t *jwt.Token, id int) bool
	ValidUser(id string, p string) bool
	GetFindById(id int) (*domain.User, error)
	GetFindByEmail(email string) (*domain.User, error)
	PasswordMatches(plainText string, targetText string) (bool, error)
	Update(id int, updateUserPayload helpers.UpdateUserPayload) error 
}

type UserService interface {
	Register(email, password, conformPassword string) error
	Login(email, password string) (jwtToken string, err error)
	Update(id int, updateUser helpers.UpdateUserPayload) error 
	ValidToken(token *jwt.Token, id int) bool 
}

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

