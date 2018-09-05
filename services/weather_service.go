package services

import (
	"fmt"
	"github.com/simplewayUA/weathereader/models"
	"github.com/simplewayUA/weathereader/weathermap"
	"gopkg.in/gomail.v2"
	"time"
)

// MonitorWeatherChanges ...
func MonitorWeatherChanges() {
	for {
		weathers, err := models.GetWeathers()
		if err != nil {
			continue
		}
		var weather *models.Weather
		var weatherData *weathermap.WeatherData
		for _, v := range weathers {
			weatherData, err = weathermap.GetWeatherByCityName(v.CityName)
			if err != nil {
				continue
			}
			weather = weatherData.WeatherModel()
			if !weather.IsEqual(&v) {
				usercity, err := models.GetUserCities(v.ID)
				if err != nil {
					continue
				}
				for _, val := range usercity {
					user, err := models.GetUserWithID(val.UserID)
					if err != nil {
						continue
					}
					if err = sendEmail(user.Email, weather); err != nil {
						continue
					}
				}
			}
			err = models.UpdateWeatherTable(v.ID, weather)
			if err != nil {
				continue
			}
		}
		time.Sleep(time.Minute * 15)
	}
}

func sendEmail(name string, w *models.Weather) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "exm.no.reply@gmail.com")
	m.SetHeader("To", name)
	m.SetHeader("Subject", "Hello, that you weather :)")
	m.SetBody("text/html", fmt.Sprintf("Weather for %s did changed!<br> Temperature: %.2f, characteristics: %s", w.CityName, w.Temp, w.Desc))

	d := gomail.NewDialer("smtp.gmail.com", 465, "exm.no.reply", "67Hu81HG7n12")
	// Send email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
