package weather_test

import (
	"slices"
	"os"
	"testing"

	"github.com/Rexbrainz/weather"
)

func TestGetWeather (t *testing.T) {
	t.Parallel()

	location := "Abia,NG"
	apikey := os.Getenv("OpenWeatherApiKey")
	
	app := weather.New(apikey)
	got, err := app.GetWeather(location)
	want, ok := app.Store.Cache[location]
	if !ok {
		t.Fatal("Should not be empty")
	}

	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(want, got) {
		t.Errorf("wanted %q, got %q", want, got)
	}
}