package database

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string
	Admin    bool   `json:"admin"`
	Email    string `json:"email"`
	Data     string `json:"data"`
	Banned   bool   `json:"banned"`
}

func GetUser(id int) *User {
	user := &User{Id: id}
	result := db.First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return user
}

func SetUser(user *User) {
	db.Save(user)
}
