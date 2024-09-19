package handlers

import (
	"back-end/internals/core/helpers"
	"back-end/internals/core/ports"
	"back-end/internals/core/utils"
	"errors"
	"fmt"
	"net/http"

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

	refer := helpers.NewRefer()
	refer.SetUserId(userId)

	if err := h.service.Create(requestPayload, *refer); err != nil {
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
	// get current login id
	userId, err := h.GetCurrentLoginUserId(c)
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	// url parsing
	memoId, err := c.ParamsInt("memoId")
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	memo, err := h.service.Get(memoId)
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	// check the owner of the memo ID and the login user are the same
	if !memo.CheckOwner(userId) {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusForbidden).JSON(jsonResp)
	}

	//h.service.
	jsonResp := helpers.JsonResponse{
		Error:   false,
		Message: "success",
		Data:    memo,
	}
	return c.Status(http.StatusOK).JSON(jsonResp)
}

func (h *MemoHandler) GetAll(c *fiber.Ctx) error {
	userId, err := h.GetCurrentLoginUserId(c)
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	refer := helpers.NewRefer()
	refer.SetUserId(userId)

	memos, err := h.service.GetAll(*refer)
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
	// get current login id
	userId, err := h.GetCurrentLoginUserId(c)
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	// url parsing
	memoId, err := c.ParamsInt("memoId")
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	// find memo
	memo, err := h.service.Get(memoId)
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	if !memo.CheckOwner(userId) {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: "permission denied",
			Data:    nil,
		}
		return c.Status(http.StatusForbidden).JSON(jsonResp)
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

	memo, err = h.service.Update(memoId, updatePayload)

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
	// get current login id
	userId, err := h.GetCurrentLoginUserId(c)
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	// url parsing
	memoId, err := c.ParamsInt("memoId")
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	// find memo
	memo, err := h.service.Get(memoId)
	if err != nil {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		return c.Status(http.StatusBadRequest).JSON(jsonResp)
	}

	if !memo.CheckOwner(userId) {
		jsonResp := helpers.JsonResponse{
			Error:   true,
			Message: "permission denied",
			Data:    nil,
		}
		return c.Status(http.StatusForbidden).JSON(jsonResp)
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
