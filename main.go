package main

import (
	"github.com/dolefir/weathereader/controllers"
	"github.com/dolefir/weathereader/db"
	"github.com/dolefir/weathereader/middlewares"
	"github.com/dolefir/weathereader/models"
	"github.com/dolefir/weathereader/services"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.ConnectDB(); err != nil {
		panic(err)
	}
	models.Migrate() // AutoMigrate
	defer db.CloseDB()

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	v1 := r.Group("/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/signin", controllers.SignInHandler)
		auth.POST("/signup", controllers.SignUpHandler)
	}

	user := v1.Group("/users")
	user.Use(middlewares.AuthHandler())
	{
		user.GET("/:id", controllers.UserHandler)
	}

	weather := v1.Group("/weather")
	weather.Use(middlewares.AuthHandler())
	{
		weather.GET("", controllers.WeathersHandler)             // all list of weathers
		weather.GET("/monitor", controllers.CitiesByUserHandler) // user's favorite weather list
		weather.PUT("/monitor/:city", controllers.AddCityHandler)
		weather.DELETE("/monitor/cities/:city", controllers.DeleteCityHandler)
	}

	go services.MonitorWeatherChanges()
	r.Run(":5050")
}
