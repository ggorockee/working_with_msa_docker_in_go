package services

import (
	"back-end/internals/core/domain"
	"back-end/internals/core/helpers"
	"back-end/internals/core/ports"
	"errors"
)

type MemoService struct {
	memoRepo ports.MemoRepository
}

func NewMemoService(repo ports.MemoRepository) *MemoService {
	return &MemoService{
		memoRepo: repo,
	}
}

func (s *MemoService) Create(createInput helpers.CreateMemoPayload, referOption ...helpers.Refer) error {
	var refer helpers.Refer
	if len(referOption) > 0 {
		refer = referOption[0]
	}
	if err := s.memoRepo.Create(createInput, refer); err != nil {
		return err
	}
	return nil
}

func (s *MemoService) GetAll(referOption ...helpers.Refer) ([]*domain.Memo, error) {
	refer := helpers.NewRefer()
	if len(referOption) > 0 {
		*refer = referOption[0]
	} else {
		return nil, errors.New("referOption must provided")
	}

	memos, err := s.memoRepo.GetAll(*refer)
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

func (s *MemoService) Get(id int) (*domain.Memo, error) {
	memo, err := s.memoRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return memo, nil
}
