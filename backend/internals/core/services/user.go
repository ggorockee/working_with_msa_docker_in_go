package services

import (
	"back-end/internals/core/ports"
	"errors"
)

type UserService struct {
	userRepository ports.UserRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
// var _ ports.UserService = (*UserService)(nil)

func NewUserService(repository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Register(email string, password string, passwordConfirm string) error {
	if password != passwordConfirm {
		return errors.New("the passwords are not equal")
	}

	err := s.userRepository.Register(email, password)
	if err != nil {
		return err
	}
	return nil
}
