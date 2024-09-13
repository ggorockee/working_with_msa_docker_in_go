package domain

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey;unique;not null"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) GetByEmail(email string) *User {
	return u

	// some login
}
