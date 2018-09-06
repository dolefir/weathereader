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
				users, err := models.UsersWithCityID(v.ID)
				if err != nil {
					continue
				}
				for _, u := range users {
					if err = sendEmail(u.Email, weather); err != nil {
						continue
					}
				}
				err = models.UpdateWeatherTable(v.ID, weather)
				if err != nil {
					continue
				}
			}
		}
		time.Sleep(time.Minute * 15)
	}
}

func sendEmail(email string, w *models.Weather) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "exm.no.reply@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello, thats your weather :)")
	m.SetBody("text/html", fmt.Sprintf("Weather for %s did changed!<br> Temperature: %.2f °C, characteristics: %s", w.CityName, w.Temp, w.Desc))

	d := gomail.NewDialer("smtp.gmail.com", 465, "exm.no.reply", "67Hu81HG7n12")
	// Send email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
