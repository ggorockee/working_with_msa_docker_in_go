package database

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey;unique;not null"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseUser struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseTinyUser struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

func (u *User) GetByEmail(email string) error {
	DBConn.Find(u, "email = ?", email)

	if u.Id == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (u *User) Create() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return errors.New("password must required")
	}

	u.Password = string(hashedPassword)

	if err := u.GetByEmail(u.Email); err == nil {
		// email을 찾으면 error
		return errors.New("email already exists")

	}

	if err := DBConn.Create(u).Error; err != nil {
		fmt.Println("err:", err.Error())
		return err
	}

	return nil
}

func (u *User) Serialize() ResponseUser {
	return ResponseUser{
		Id:        u.Id,
		Email:     u.Email,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
