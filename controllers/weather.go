package controllers

import (
	"net/http"

	"github.com/dolefir/weathereader/middlewares"
	"github.com/dolefir/weathereader/models"
	"github.com/dolefir/weathereader/weathermap"
	"github.com/gin-gonic/gin"
)

// WeathersHandler ...
func WeathersHandler(c *gin.Context) {
	qParam := c.Query("city")
	if qParam == "" {
		c.JSON(http.StatusNotFound,
			gin.H{"Error: ": "Invalid weather on search filter!"})
		return
	}

	var weather *models.Weather
	var err error
	weather, err = models.FindCityByItsName(qParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	w := &models.TransformedWeather{
		ID:       weather.ID,
		CityName: weather.CityName,
		TempMax:  weather.TempMax,
		Temp:     weather.Temp,
		TempMin:  weather.TempMin,
		Desc:     weather.Desc,
	}

	c.JSON(http.StatusOK, &w)
}

// AddCityHandler ...
func AddCityHandler(c *gin.Context) {
	var user = &models.User{}
	var weather *models.Weather
	var weatherData *weathermap.WeatherData
	var err error

	cityName := c.Param("city")
	if weatherData, err = weathermap.GetWeatherByCityName(cityName); err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "entered the wrong city"})
		return
	}

	userID, err := middlewares.ReturnUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = models.GetUserWithID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	weather, err = models.FindCityByItsName(cityName)
	if err != nil {
		weather = (*models.Weather)(weatherData.WeatherModel())
	}
	user.UserСities = models.GetUserWithWeather(user.ID, weather.ID)
	// user save
	if err = user.Save(); err != nil {
		return
	}
	// weather save
	if err = weather.Save(); err != nil {
		return
	}

	w := &models.TransformedWeather{
		ID:       weather.ID,
		CityName: weather.CityName,
		TempMax:  weather.TempMax,
		Temp:     weather.Temp,
		TempMin:  weather.TempMin,
		Desc:     weather.Desc,
	}

	c.JSON(http.StatusCreated, &w)
}

// CitiesByUserHandler ...
func CitiesByUserHandler(c *gin.Context) {
	var user = &models.User{}
	var err error
	userID, err := middlewares.ReturnUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = models.GetUserWithID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var w models.Weather
	var responseJSON []models.TransformedWeather
	for i := 0; i < len(user.UserСities); i++ {
		w = user.UserСities[i].WeatherCity
		responseJSON = append(
			responseJSON, models.TransformedWeather{
				ID:       w.ID,
				CityName: w.CityName,
				Temp:     w.Temp,
				TempMax:  w.TempMax,
				TempMin:  w.TempMin,
				Desc:     w.Desc,
			})
	}

	c.JSON(http.StatusAccepted, &responseJSON)
}

// DeleteCityHandler ...
func DeleteCityHandler(c *gin.Context) {
	var weather *models.Weather
	var user = &models.User{}

	cityName := c.Param("city")
	weather, err := models.FindCityByItsName(cityName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID, err := middlewares.ReturnUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err = models.GetUserWithID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = models.DeleteUserWithCity(user.ID, weather.ID)
	if err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "record not found"})
		return
	}

	c.JSON(http.StatusAccepted,
		gin.H{"status": http.StatusAccepted, "message": "successfully deleted"})
}
