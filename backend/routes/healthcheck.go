package routes

import (
	"back-end/helpers"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (app *Config) HealthCheck(c *fiber.Ctx) error {

	response := helpers.JsonResponse{
		Error:   false,
		Message: "ok",
		Data:    nil,
	}
	return c.Status(http.StatusOK).JSON(response)
}
