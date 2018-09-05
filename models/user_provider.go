package models

import (
	"github.com/simplewayUA/weathereader/db"
)

// GetUserWithEmail ...
func GetUserWithEmail(e string) (*User, error) {
	var getDB = db.GetDB()
	var user User
	var err error

	if err = getDB.Preload("UserСities.WeatherCity").Where(&User{Email: e}).First(&user).Error; err != nil {
		return nil, err
	}
	if user.ID > 0 {
		return &user, err
	}
	return &user, nil
}

// GetUserWithID ...
func GetUserWithID(id uint) (*User, error) {
	var getDB = db.GetDB()
	var user User
	var err error

	if err = getDB.Preload("UserСities.WeatherCity").Where(&User{ID: id}).First(&user).Error; err != nil {
		return nil, err
	}
	if user.ID > 0 {
		return &user, err
	}
	return &user, nil
}
