package repositories

import (
	"back-end/database"
	"back-end/internals/core/domain"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

// UserRepository를 초기화
func NewUserRepository() *UserRepository {
	return &UserRepository{
		conn: database.Connect().Conn,
	}
}

// UserRepository의 인터페이스를 구현
func (r *UserRepository) Register(email, password string) error {
	// UserRepository에는 *gorm.DB가 conn의 이름으로 있고
	// 이걸 이용해서 찾을 수 있음

	user := domain.User{
		Email:    email,
		Password: password,
	}

	err := r.conn.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Login(email, password string) error       {}
func (r *UserRepository) Logout(id int) error                      {}
func (r *UserRepository) ValidToken(t *jwt.Token, id string) bool  {}
func (r *UserRepository) validUser(id string, p string) bool       {}
func (r *UserRepository) GetFindById(id int) (*domain.User, error) {}
