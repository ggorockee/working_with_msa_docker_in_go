package domain

import (
	"time"
)

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey;unique;not null" mapstructure:"id"`
	Email     string    `json:"email" mapstructure:"email"`
	Password  string    `json:"-" gorm:"type:varchar(255)"`
	Name      string    `json:"name" mapstructure:"name"`
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"`
}

var UserModel User

//func (u *User) GetByEmail(email string) (*User, error) {
//	var user User
//	err := database.DB.Conn.Where("email = ?", email).First(&user).Error
//	if err != nil {
//		return nil, err
//	}
//
//	return &user, nil
//}
//
//func (u *User) GetOne(id int) (*User, error) {
//	var user User
//	err := database.DB.Conn.Where("id = ?", id).First(&user).Error
//	if err != nil {
//		return nil, err
//	}
//
//	return &user, nil
//}
//
//type userMap map[string]any
//
//func (u *User) Update(updateUserMap userMap) (*User, error) {
//
//	var user User
//	if err := mapstructure.Decode(updateUserMap, &user); err != nil {
//		log.Println("cannot parse map to structre")
//	}
//
//	user.UpdatedAt = time.Now()
//
//	log.Println("Update ==============================")
//	log.Println("id", user.Id)
//	log.Println("id", user.UpdatedAt)
//	log.Println("id", user.Email)
//	log.Println("id", user.Name)
//	log.Println("Update end ==============================")
//
//	if err := database.DB.Conn.Model(&user).Updates(user).Error; err != nil {
//		return nil, err
//	}
//
//	return &user, nil
//}
//
//// Delete deletes one user from the database, by User.ID
//func (u *User) Delete() error {
//
//	err := database.DB.Conn.Delete(u, u.Id).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//// DeleteByID deletes one user from the database, by ID
//func (u *User) DeleteByID(id int) error {
//
//	err := database.DB.Conn.Delete(u, id).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//// Insert inserts a new user into the database, and returns the ID of the newly inserted row
//func (u *User) Insert(user User) error {
//
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
//	if err != nil {
//		return err
//	}
//
//	u.Password = string(hashedPassword)
//
//	if err != nil {
//		return err
//	}
//
//	if err := database.DB.Conn.Create(&u).Error; err != nil {
//		return err
//	}
//
//	return nil
//
//}
//
//// ResetPassword is the method we will use to change a user's password.
//func (u *User) ResetPassword(password string) error {
//
//	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
//	//if err != nil {
//	//	return err
//	//}
//	//
//	return nil
//}
//
//// PasswordMatches uses Go's bcrypt package to compare a user supplied password
//// with the hash we have stored for a given user in the database. If the password
//// and hash match, we return true; otherwise, we return false.
//func (u *User) PasswordMatches(plainText string) (bool, error) {
//	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
//	if err != nil {
//		switch {
//		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
//			// invalid password
//			return false, nil
//		default:
//			return false, err
//		}
//	}
//
//	return true, nil
//}
