package handlers

import (
	"back-end/internals/core/helpers"
	"back-end/internals/core/ports"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct {
	service ports.UserService
}

// UserHandler 초기화
func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// UserHandler 인터페이스사용
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var registerUserPayload helpers.RegisterUserPayload

	if err := c.BodyParser(&registerUserPayload); err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}
	// Extract the body and get the email and password
	err := h.service.Register(
		registerUserPayload.Email,
		registerUserPayload.Password,
		registerUserPayload.PasswordConfirm,
	)

	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	return nil
}
