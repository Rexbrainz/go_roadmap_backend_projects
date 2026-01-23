package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)


type Application struct {
	fetcher	WeatherFetcher
	store	Storer
}

// Weather Fetcher is an interface that speaks to OpenWeatherMap server
type WeatherFetcher interface {
	Fetch(context.Context, string) (Weather, error)
}

type Client struct {
	BaseURL		string
	ApiKey		string
	HTTPClient	*http.Client
}

// Contstructs and returns a new Client
func NewClient(apikey string) *Client {
	return &Client{
		BaseURL:	"https://api.openweathermap.org",
		ApiKey:		apikey,
		HTTPClient:	&http.Client{
			Timeout:	10 * time.Second,
		},
	}
}

// Contstructs and returns a Weather app
func NewWeatherApp(fetcher WeatherFetcher, store Storer) *Application {
	return &Application {
		fetcher:	fetcher,
		store:	store,
	}
}

// Formats and returns URL to to be called by http.Client.Get()
func (c *Client) FormatURL(city string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s",
		 c.BaseURL, url.QueryEscape(city), c.ApiKey)
}

// GetWeather() gets and caches response from OpenWeatherMap server
func (app *Application) GetWeather(ctx context.Context, city string) (Weather, error) {
	if w, ok := app.store.Get(city); ok {
		return w, nil
	}

	w, err := app.fetcher.Fetch(ctx, city)
	if err != nil {
		return Weather{}, err
	}

	app.store.Set(city, w, 15 * time.Minute)

	return w, nil
}

// Fetch() talks directly to OpenWeatherMap Server
func (f *Client) Fetch(ctx context.Context, city string) (Weather, error) {
	URL := f.FormatURL(city)

	resp, err := f.HTTPClient.Get(URL)
	if err != nil {
		return Weather{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("unexpected error, %v", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Weather{}, err
	}
	
	return ParseResponse(data)
}

// Unmarshal response from open weather mpa api, returns it as Weather
func ParseResponse(data []byte) (Weather, error) {
	var weather WeatherReport

	if err := json.Unmarshal(data, &weather); err != nil {
		return Weather{}, err
	}

	if len(weather.Weather) < 1 {
		return Weather{}, ErrInvalidProviderResponse
	}
	report := Weather{
		City:			weather.Name,
		Country: 		weather.Country.Country,
		Temp:			weather.Temp.Temp,
		Temp_min:		weather.Temp.Temp_min,
		Temp_max:		weather.Temp.Temp_max,
		Condition:		weather.Weather[0].Main,
		Description:	weather.Weather[0].Description,
	}
	return report, nil
}

// Write Response to the ResponseWriter
func (app *Application) writeJSON(w http.ResponseWriter, status int, report Weather) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(report)
}

// Starts the server which runs the application.
func StartServer(app *Application) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /weather", app.getWeather)
	mux.HandleFunc("GET /ping", app.ping)

	server := &http.Server{
		Addr:		":4000",
		Handler:	mux,
	}

	fmt.Println("Listening on port 4000")
	err := server.ListenAndServe()
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
