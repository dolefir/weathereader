package weathermap

import (
	"encoding/json"
	"fmt"
	"github.com/simplewayUA/weathereader/models"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const baseURL string = "http://api.openweathermap.org/data/2.5/weather?%s"

// APIKEY for example available, you can hide .evn
// const APIKEY = os.Getenv("APIKEY")

// Clouds ...
type Clouds struct {
	All int64 `json:"all"`
}

// Coord ...
type Coord struct {
	Lon float64 `json:"lon"` // City geo location, longitude
	Lat float64 `json:"lat"` // City geo location, latitude
}

// Main group of weather parameters (Rain, Snow, Extreme etc.)
type Main struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity int64   `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

// Sys internal parameter
type Sys struct {
	Type    int64   `json:"type"`
	ID      int64   `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int64   `json:"sunrise"`
	Sunset  int64   `json:"sunset"`
}

// WeatherElement ...
type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Wind ...
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

// WeatherData model
type WeatherData struct {
	Coord       Coord            `json:"coord"`      // Geo location
	WeatherDesc []WeatherElement `json:"weather"`    // More info Weather condition codes
	Base        string           `json:"base"`       // Internal parameter
	Main        Main             `json:"main"`       // Main(etc.)
	Visibility  int64            `json:"visibility"` // Visibility, meter
	Wind        Wind             `json:"wind"`       // Wind
	Clouds      Clouds           `json:"clouds"`     // Cloudiness
	Dt          int64            `json:"dt"`         // Time of data calculation, unix, UTC
	Sys         Sys              `json:"sys"`        // Sys internal parameter
	ID          int64            `json:"id"`         // City ID
	Name        string           `json:"name"`       // City name
	Cod         int64            `json:"cod"`        // Internal parameter
}

// WeatherModel ...
func (wd *WeatherData) WeatherModel() *models.Weather {
	var desc string

	for _, v := range wd.WeatherDesc {
		if v.Description != "" {
			desc = v.Description
		} else {
			desc = "no description, sorry"
		}
	}

	return &models.Weather{
		CityName: wd.Name,
		TempMax:  wd.Main.TempMax,
		Temp:     wd.Main.Temp,
		TempMin:  wd.Main.TempMin,
		Desc:     desc,
	}
}

// GetWeatherByCityName ...
func GetWeatherByCityName(name string) (*WeatherData, error) {
	response, err := http.Get(fmt.Sprintf(fmt.Sprintf(baseURL, "q=%s&APPID=%s&units=metric"), strings.Title(name), os.Getenv("APIKEY")))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	weatherResult := WeatherData{}
	if response.StatusCode == 404 {
		return &weatherResult, fmt.Errorf("city not found")
	}

	err = json.Unmarshal(result, &weatherResult)
	if err != nil {
		return nil, err
	}

	return &weatherResult, nil
}
