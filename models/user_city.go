package models

import (
	"github.com/simplewayUA/weathereader/db"
)

// UserCity ...
type UserCity struct {
	UserID uint `gorm:"primary_key"`
	CityID uint `gorm:"primary_key"`

	WeatherCity Weather `gorm:"foreignkey:ID;association_foreignkey:cityID"`
}

// GetUserWithWeather ...
func GetUserWithWeather(u uint, w uint) []UserCity {
	var user User
	user.UserСities = append(user.UserСities, UserCity{
		UserID: u,
		CityID: w,
	})
	return user.UserСities
}

// DeleteUserWithCity ...
func DeleteUserWithCity(u uint, w uint) (*UserCity, error) {
	var getDB = db.GetDB()
	var cities UserCity
	var err error

	if err = getDB.Where(&UserCity{UserID: u, CityID: w}).First(&cities).Error; err != nil {
		return nil, err
	}

	userCity := &UserCity{
		UserID: u,
		CityID: w,
	}

	if err = getDB.Delete(&userCity).Error; err != nil {
		return nil, err
	}
	return userCity, nil
}

// GetUserCities ...
func GetUserCities(cityID uint) ([]UserCity, error) {
	var getDB = db.GetDB()
	var usercities []UserCity
	var err error

	if err = getDB.Where("city_id = ?", cityID).Find(&usercities).Error; err != nil {
		return nil, err
	}
	return usercities, nil
}

// GetAllUsersWithID ...
func GetAllUsersWithID(userID uint) ([]User, error) {
	var getDB = db.GetDB()
	var users []User
	var err error

	if err = getDB.Where("id = ?", userID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
