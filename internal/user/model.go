package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"index"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func NewUser(email, password, name string) *User {
	user := &User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	return user
}
