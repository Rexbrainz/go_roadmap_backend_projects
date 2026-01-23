package main

import (
	"fmt"
	"os"
	"github.com/Rexbrainz/weather/weather"
)

func main() {
	apikey := os.Getenv("OpenWeatherApiKey")
	if  apikey == "" {
		fmt.Fprintf(os.Stderr, "Empty Api key")
		os.Exit(1)
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		fmt.Fprintf(os.Stderr, "Empty Redis URL")
		os.Exit(1)
	}

	client := weather.NewClient(apikey)
	var store weather.Storer
	store, err := weather.NewRedisStore(redisURL)
	if err != nil {
		store = weather.NewMemStore()
	}

	app := weather.NewWeatherApp(client, store)
	weather.StartServer(app)
}