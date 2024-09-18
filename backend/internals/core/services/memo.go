package services

import (
	"back-end/internals/core/domain"
	"back-end/internals/core/helpers"
	"back-end/internals/core/ports"
)

type MemoService struct {
	memoRepo ports.MemoRepository
}

func NewMemoService(repo ports.MemoRepository) *MemoService {
	return &MemoService{
		memoRepo: repo,
	}
}

func (s *MemoService) Create(createInput *domain.Memo) error {
	if err := s.memoRepo.Create(createInput); err != nil {
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


