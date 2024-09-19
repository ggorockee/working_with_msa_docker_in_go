package repositories

import (
	"back-end/database"
	"back-end/internals/core/domain"
	"back-end/internals/core/helpers"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	conn  *gorm.DB
	model *domain.User
}

// UserRepository를 초기화
func NewUserRepository() *UserRepository {
	return &UserRepository{
		conn:  database.Connect().Conn,
		model: &domain.UserModel,
	}
}

// UserRepository의 인터페이스를 구현
func (r *UserRepository) Register(email, password string) error {
	// UserRepository에는 *gorm.DB가 conn의 이름으로 있고
	// 이걸 이용해서 찾을 수 있음

	_, err := r.GetFindByEmail(email)
	if err == nil {
		return errors.New("email already use")
	}

	user := r.GetModel()
	user.Email = email

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = r.conn.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// func (r *UserRepository) Login(email, password string) error       {
// 	return nil
// }

func (r *UserRepository) GetModel() *domain.User {
	return r.model
}

func (r *UserRepository) ValidToken(t *jwt.Token, id int) bool {

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == id
}
func (r *UserRepository) ValidUser(id string, p string) bool {
	user := r.GetModel()
	if err := r.conn.First(user, id).Error; err != nil {
		return false
	}

	return true

}

func (r *UserRepository) GetFindById(id int) (*domain.User, error) {
	user := r.GetModel()

	if err := r.conn.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetFindByEmail(email string) (*domain.User, error) {
	user := r.GetModel()

	if err := r.conn.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (r *UserRepository) PasswordMatches(hashedPassword string, plainText string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, errors.New("password not matched")
		default:
			return false, err
		}
	}

	return true, nil
}

func (r *UserRepository) Update(id int, updateUserPayload helpers.UpdateUserPayload) error {
	user, err := r.GetFindById(id)

	if err != nil {
		return err
	}

	if updateUserPayload.Name != "" {
		user.Name = updateUserPayload.Name
	}

	err = r.conn.Save(user).Error
	if err != nil {
		return err
	}

	return nil
}
