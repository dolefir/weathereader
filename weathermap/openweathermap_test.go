package weathermap

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"io/ioutil"
	"log"
	"testing"
)

func TestWeathermap_GetWeatherByCityName(t *testing.T) {
	defer gock.Off()
	b, err := ioutil.ReadFile("../fixtures/by_city_name.json")
	if err != nil {
		log.Println(err)
		panic("cannot read fixtures/campaigns.json")
	}
	gock.New("http://api.openweathermap.org").
		Get("/data/2.5/weather").
		MatchParam("q", "London").
		Reply(200).
		BodyString(string(b))

	w := &WeatherData{}
	w, err = GetWeatherByCityName("London")
	assert.Nil(t, err)
	assert.Equal(t, w.Name, "London")
	assert.Equal(t, w.Cod, int64(200)) // statusCode
	assert.True(t, gock.IsDone())
}
