package ports

import (
	"back-end/internals/core/domain"
	"back-end/internals/core/helpers"
	"github.com/gofiber/fiber/v2"
)

type MemoRepository interface {
	GetModel() domain.Memo
	Create(title, content string) error
	Update(id int, updateUser helpers.UpdateMemoPayload) (*domain.Memo, error)
	GetById(id int) (*domain.Memo, error)
	GetAll() ([]*domain.Memo, error)
	Delete(id int) error
}

type MemoService interface {
	Create(createInput fiber.Map) error
	Update(id int, updateUser helpers.UpdateMemoPayload) (*domain.Memo, error)
	GetAll() ([]*domain.Memo, error)
	Delete(id int) error
}

type MemoHandler interface {
	Create(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}
