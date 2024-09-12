package handlers

import (
	"back-end/database"
	"back-end/helpers"

	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CreateUser(c *fiber.Ctx) error {
	var user database.User

	if err := c.BodyParser(&user); err != nil {
		jsonResponse := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResponse)
	}

	var userRepository database.Repository
	userRepository = &user

	if err := userRepository.Create(); err != nil {
		jsonResponse := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResponse)
	}

	return nil
}
