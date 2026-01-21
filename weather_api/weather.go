package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)


type Application struct {
	Client	*Client
	Store	Storer
}

type Weather struct {
	City		string
	Country		string
	Temperature	float64
	Summary		string
}

type Client struct {
	BaseURL		string
	ApiKey		string
	HTTPClient	*http.Client
}

type MemoryStore struct {
	Cache	map[string]Weather
}

type Storer interface {
	Get(city string) (Weather, bool)
	Set(city string, weather Weather, ttl time.Duration)
}

func NewWeatherApp(apikey string) *Application {
	client :=	&Client{
		BaseURL:	"https://api.openweathermap.org",
		ApiKey:		apikey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	cache := &Store{
		Cache: map[string]json.RawMessage{},
	}

	app := &Application {
		Client:	client,
		Store:	cache,
	}
	return app
}

func (c *Client) FormatURL(location string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", c.BaseURL, location, c.ApiKey)
}

func (app *Application) GetWeather(ctx context.Context, location string) (Weather, error) {
	URL := app.Client.FormatURL(location)

	report, ok := app.Store.Cache[location]
	if ok {
		return report, nil
	}
	resp, err := app.Client.HTTPClient.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		var result []byte

		err := json.Unmarshal(data, &result)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%q",result)
	}
	app.Store.Cache[location] = data
	return data, nil
}

// func (app *Application) serverError(w http.ResponseWriter, r *http.Request) {
// 	var (
// 		method 	= r.Method
// 		uri		= r.URL.RequestURI()
// 	)
// 	htt
// }

func Main() {
	apikey := os.Getenv("OpenWeatherApiKey")
	if apikey == "" {
		fmt.Fprintf(os.Stderr, "%v Internal Server Error", http.StatusInternalServerError)
		return
	}
	app := NewWeatherApp(apikey)
	server := &http.Server{
		Addr:		":4000",
		Handler:	app.route(),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
