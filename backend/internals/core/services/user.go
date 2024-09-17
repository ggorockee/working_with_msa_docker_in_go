package services

import (
	"back-end/internals/core/ports"
	"errors"
)

type UserService struct {
	// userRepo는 UserRepository 인터페이스
	// 따라서 Register(email, password string) error 메소드를 가지고 있는 structure는
	// userRepo에 쓸수 있음
	userRepo ports.UserRepository
}

// Register(email, password string) error를 구현한 type을 인자로 받을 수 있음
// 그리고 UserService 구조체를 초기화
func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo}
}

// UserService Interface를 사용하기 위해 Register(email, password, conformPassword string) error 메소드를 구현
// UserService 구조체는 UserService의 interface로 사용할수 있음
func (s *UserService) Register(email, password, conformPassword string) error {
	if password != conformPassword {
		return errors.New("the passwords are not equal")
	}

	// UserService -> userRepo
	// userRepo는 ports.UserRepository의 interface이고
	// 이 interface는 Register(email, password string) error 메소드를 가지고 있음
	err := s.userRepo.Register(email, password)
	if err != nil {
		return err
	}
	return nil
}
