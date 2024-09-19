package handlers

import (
	"back-end/internals/core/helpers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct {
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h HealthCheckHandler) HealthCheck(c *fiber.Ctx) error {
	jsonResp := helpers.JsonResponse{
		Error:   false,
		Message: "Health Check... OK!!",
		Data:    nil,
	}
	return c.Status(http.StatusOK).JSON(jsonResp)
}
