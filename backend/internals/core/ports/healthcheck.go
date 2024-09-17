package ports

import "github.com/gofiber/fiber/v2"

type HealthCheckHandler interface {
	HealthCheck(c *fiber.Ctx) error
}
