package handlers

import (
	"back-end/internals/core/ports"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService ports.UserService
}

// var _ ports.UserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService ports.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var email string
	var password string
	var passwordConfirm string

	err := h.userService.Register(email, password, passwordConfirm)
	if err != nil {
		return err
	}
	return nil
}
