package handlers

import (
	"back-end/internals/core/helpers"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
