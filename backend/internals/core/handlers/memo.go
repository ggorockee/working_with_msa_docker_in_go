package handlers

import (
	"back-end/internals/core/domain"
	"back-end/internals/core/helpers"
	"back-end/internals/core/ports"
	"back-end/internals/core/utils"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type MemoHandler struct {
	service ports.MemoService
}

func NewMemoHandler(service ports.MemoService) *MemoHandler {
	return &MemoHandler{service}
}

func (h *MemoHandler) Create(c *fiber.Ctx) error {
	userId, err := h.GetCurrentLoginUserId(c)
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusUnauthorized).JSON(jsonResp)
	}

	
	var requestPayload helpers.CreateMemoPayload
	if err := c.BodyParser(&requestPayload); err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	var memo domain.Memo
	
	memo.UserRefer = userId
	
	if requestPayload.Content != "" {
		memo.Content = requestPayload.Content
		memo.UpdatedAt = time.Now()
	}

	if requestPayload.Title != "" {
		memo.Title = requestPayload.Title
		memo.UpdatedAt = time.Now()
	}


	if err := h.service.Create(&memo); err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	jsonResp := helpers.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("create note %s", requestPayload.Title),
		Data:    nil,
	}
	return c.Status(http.StatusOK).JSON(jsonResp)
}
func (h *MemoHandler) Get(c *fiber.Ctx) error {
	//memoId, _ := c.ParamsInt("memoId")
	//h.service.
	return nil
}

func (h *MemoHandler) GetAll(c *fiber.Ctx) error {

	memos, err := h.service.GetAll()
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
		Message: "success",
		Data:    memos,
	}
	return c.Status(http.StatusOK).JSON(jsonResp)
}

func (h *MemoHandler) Update(c *fiber.Ctx) error {
	memoId, err := c.ParamsInt("memoId")
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	var updatePayload helpers.UpdateMemoPayload

	if err := c.BodyParser(&updatePayload); err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	memo, err := h.service.Update(memoId, updatePayload)

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
		Message: "update success",
		Data:    memo,
	}
	return c.Status(http.StatusOK).JSON(jsonResp)
}

func (h *MemoHandler) Delete(c *fiber.Ctx) error {
	memoId, err := c.ParamsInt("memoId")
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   false,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	if err := h.service.Delete(memoId); err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   false,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	jsonResp := helpers.JsonResponse{
		Error:   false,
		Message: "memo delete",
		Data:    nil,
	}
	return c.Status(http.StatusOK).JSON(jsonResp)
}

func (h *MemoHandler) GetCurrentLoginUserId(c *fiber.Ctx) (userId int, err error) {
	token := c.Locals("user").(*jwt.Token)
	if token == nil {
		return 0, errors.New("cannot find user")
	}

	userId = utils.GetUserIdFromJwtToken(token)
	return userId, nil
}