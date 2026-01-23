package weather_test

import (
	"context"
	"os"
	"slices"
	"testing"

	"github.com/Rexbrainz/weather"
)

func TestGetWeather (t *testing.T) {
	t.Parallel()

	location := "Abia,NG"
	apikey := os.Getenv("OpenWeatherApiKey")
	
	app := weather.NewWeatherApp(apikey)
	got, err := app.GetWeather(context.Background(), location)
	
	want, ok := app.Store.Get(location)
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