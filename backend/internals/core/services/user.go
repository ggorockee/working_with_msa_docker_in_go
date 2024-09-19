package services

import (
	"back-end/configs"
	"back-end/internals/core/helpers"
	"back-end/internals/core/ports"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	// userRepo는 UserRepository 인터페이스
	// 따라서 Register(email, password string) error 메소드를 가지고 있는 structure는
	// userRepo에 쓸수 있음
	repo ports.UserRepository
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
	err := s.repo.Register(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(email, password string) (jwtToken string, err error) {
	// password가 같은지 비교
	user, err := s.repo.GetFindByEmail(email)
	if err != nil {
		return "", err
	}

	ok, err := s.repo.PasswordMatches(user.Password, password)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	jwtToken, err = token.SignedString([]byte(configs.New().JWTSecret))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (s *UserService) Update(id int, updateUser helpers.UpdateUserPayload) error {
	err := s.repo.Update(id, updateUser)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) ValidToken(token *jwt.Token, id int) bool {
	return s.repo.ValidToken(token, id)
}
