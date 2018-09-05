package models

import (
	"github.com/simplewayUA/weathereader/db"
	"time"
)

// User model
type User struct {
	ID        uint   `gorm:"primary_key"`
	UserName  string `gorm:"NOT NULL"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"NOT NULL"`
	CreatedAt time.Time

	UserСities []UserCity `gorm:"foreignkey:UserID"`
}

// Save ...
func (u *User) Save() error {
	var getDB = db.GetDB()

	return getDB.Save(u).Error
}

// UserCreateParams ...
type UserCreateParams struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserLoginParams ...
type UserLoginParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ToUser ...
func (u *UserCreateParams) ToUser() *User {
	user := User{
		UserName: u.UserName,
		Email:    u.Email,
		Password: u.Password,
	}
	return &user
}

// TransformedUser ...
type TransformedUser struct {
	ID       uint   `json:"ID"`
	UserName string `json:"User_name"`
	Email    string `json:"Email"`
}
