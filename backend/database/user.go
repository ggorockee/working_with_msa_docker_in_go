package database

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey;unique"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

	if err := DBConn.Create(u).Error; err != nil {
		return err
	}

	return nil
}
