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
	Client	*Client
	Store	Storer
}

type Client struct {
	BaseURL		string
	ApiKey		string
	HTTPClient	*http.Client
}

func NewWeatherApp(apikey string) *Application {
	client :=	&Client{
		BaseURL:	"https://api.openweathermap.org",
		ApiKey:		apikey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	var store Storer

	redisURL := os.Getenv("REDIS_URL")
	if redisURL != "" {
		redisStore, err := NewRedisStore(redisURL)
		if err == nil {
			store = redisStore
		}
	}

	// Fallback if Redis is not configured or failed
	if store == nil {
		fmt.Println("Using in-memory store")
		store = &MemStore{
			Cache: map[string]Weather{},
		}
	}

	app := &Application {
		Client:	client,
		Store:	store,
	}
	return app
}

func (c *Client) FormatURL(city string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s",
		 c.BaseURL, url.QueryEscape(city), c.ApiKey)
}

func (app *Application) GetWeather(ctx context.Context, city string) (Weather, error) {
	URL := app.Client.FormatURL(city)

	report, ok := app.Store.Get(city)
	if ok {
		return report, nil
	}

	resp, err := app.Client.HTTPClient.Get(URL)
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
	
	report, err = ParseResponse(data)
	app.Store.Set(city, report, 15 * time.Minute)
	return report, nil
}

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


func (app *Application) writeJSON(w http.ResponseWriter, status int, report Weather) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(report)
}

func Main() {
	apikey := os.Getenv("OpenWeatherApiKey")
	if apikey == "" {
		fmt.Fprintf(os.Stderr, "%v Internal Server Error", http.StatusInternalServerError)
		return
	}

	app := NewWeatherApp(apikey)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /weather", app.getWeather)
	mux.HandleFunc("GET /ping", app.ping)

	server := &http.Server{
		Addr:		":4000",
		Handler:	mux,
	}

	fmt.Println(Usage, "\nListening on port 4000")
	err := server.ListenAndServe()
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
