package handlers

import (
	"back-end/internals/core/helpers"
	"back-end/internals/core/ports"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

func (h *UserHandler) Login(c *fiber.Ctx) error {

	var loginInput helpers.LoginInput
	if err := c.BodyParser(&loginInput); err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	token, err := h.service.Login(loginInput.Email, loginInput.Password)

	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	jsonResp := helpers.JsonResponse{
		Error:   false,
		Message: "Welcome!!!",
		Data:    token,
	}
	return c.Status(http.StatusOK).JSON(jsonResp)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("userId")
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	if c.Locals("user") == nil {
		// local 값이 없으면
		c.Locals("user", "anonymousUser")

	} else {
		token := c.Locals("user").(*jwt.Token)
		if !h.service.ValidToken(token, id) {
			jsonResp := helpers.JsonResponse{
				Error:   true,
				Message: "Invalid token id",
				Data:    nil,
			}
			return c.Status(http.StatusInternalServerError).JSON(jsonResp)
		}
	}

	var updatePayload helpers.UpdateUserPayload
	if err := c.BodyParser(&updatePayload); err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	err = h.service.Update(id, updatePayload)
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	jsonResp := helpers.JsonResponse{
		Error:   false,
		Message: "successfully change!",
		Data:    updatePayload,
	}
	return c.Status(http.StatusOK).JSON(jsonResp)

}
