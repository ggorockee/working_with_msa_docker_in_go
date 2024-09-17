package services

import (
	"back-end/internals/core/domain"
	"back-end/internals/core/helpers"
	"back-end/internals/core/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

type MemoService struct {
	memoRepo ports.MemoRepository
}

func NewMemoService(repo ports.MemoRepository) *MemoService {
	return &MemoService{
		memoRepo: repo,
	}
}

func (s *MemoService) Create(createInput fiber.Map) error {

	model := s.memoRepo.GetModel()
	if err := mapstructure.Decode(createInput, &model); err != nil {
		return err
	}

	if err := s.memoRepo.Create(model.Title, model.Content); err != nil {
		return err
	}
	return nil
}

func (s *MemoService) GetAll() ([]*domain.Memo, error) {
	memos, err := s.memoRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return memos, nil
}

func (s *MemoService) Update(id int, updateUser helpers.UpdateMemoPayload) (*domain.Memo, error) {
	memo, err := s.memoRepo.Update(id, updateUser)
	if err != nil {
		return nil, err
	}

	return memo, nil
}

func (s *MemoService) Delete(id int) error {
	err := s.memoRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
