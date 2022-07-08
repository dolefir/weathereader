package models

import (
	"strings"

	"github.com/dolefir/weathereader/db"
	"github.com/jinzhu/gorm"
)

// Weather model
type Weather struct {
	gorm.Model
	CityName string
	TempMax  float64
	Temp     float64
	TempMin  float64
	Desc     string
}

// Save weather
func (w *Weather) Save() error {
	var getDB = db.GetDB()
	return getDB.Save(w).Error
}

// FindCityByItsName ...
func FindCityByItsName(n string) (*Weather, error) {
	var getDB = db.GetDB()
	var weather Weather
	var err error

	if err = getDB.Where(&Weather{CityName: strings.Title(n)}).First(&weather).Error; err != nil {
		return nil, err
	}
	if weather.ID > 0 {
		return &weather, err
	}
	return &weather, nil
}

// GetWeathers ...
func GetWeathers() ([]Weather, error) {
	var getDB = db.GetDB()
	var weathers []Weather
	var err error

	if err = getDB.Find(&weathers).Error; err != nil {
		return nil, err
	}
	return weathers, nil
}

// TransformedWeather ...
type TransformedWeather struct {
	ID       uint    `json:"ID"`
	CityName string  `json:"CityName"`
	TempMax  float64 `json:"TempMax"`
	Temp     float64 `json:"Temp"`
	TempMin  float64 `json:"TempMin"`
	Desc     string  `json:"Description"`
}

// UpdateWeatherTable weather table
func UpdateWeatherTable(id uint, w *Weather) error {
	var getDB = db.GetDB()
	var err error

	if err = getDB.Table("weathers").Where("id IN (?)", id).Updates(&w).Error; err != nil {
		return err
	}
	return nil
}

// IsEqual ...
func (w *Weather) IsEqual(wtr *Weather) bool {
	if w.Temp != wtr.Temp || w.Desc != wtr.Desc {
		return false
	}
	return true
}
