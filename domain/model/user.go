package model

import (
    "time"
)

type Users []User

type User struct {
	ID        int       `json:"id"`
	UserKey   string    `json:"user_key"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func EmptyUser() *User {
	return &User{}
}

func (user *User) IsEmpty() bool {
	return (
		user.ID == 0 &&
		user.UserKey == "" &&
		user.Username == "" &&
		user.Email == "" &&
		user.Password == "" &&
		user.Token == "" &&
		user.Status == "")
}
