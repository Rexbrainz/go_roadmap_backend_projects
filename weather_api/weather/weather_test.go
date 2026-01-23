package weather_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Rexbrainz/weather/weather"
	"github.com/google/go-cmp/cmp"
)

type FakeFetcher struct {
	Called	int
	Result	weather.Weather
	Err		error
}

func (f *FakeFetcher) Fetch(ctx context.Context, location string) (weather.Weather, error) {
	f.Called++
	return f.Result, f.Err
}

type FakeStore struct {
	data map[string]weather.Weather
}

func NewFakeStore() *FakeStore {
	return &FakeStore{data: make(map[string]weather.Weather)}
}

func (s *FakeStore) Get(location string) (weather.Weather, bool) {
	w, ok := s.data[location]
	return w, ok
}

func (s *FakeStore) Set(location string, w weather.Weather, ttl time.Duration) {
	s.data[location] = w
}

func TestGetWeather_CacheMiss(t *testing.T) {
	want := weather.Weather{
		City:			"Abia",
		Country:		"NG",
		Temp:			311.69,
		Temp_min:		311.69,
		Temp_max:		311.69,
		Condition:		"Clouds",
		Description:	"overcast clouds",		
	}

	fetcher := &FakeFetcher{
		Result: want,
	}
	store := NewFakeStore()

	app := weather.NewWeatherApp(fetcher, store)

	got, err := app.GetWeather(context.Background(), "Abia,NG")
	if err != nil {
		t.Fatal("expected error")
	}

	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}

	if fetcher.Called != 1 {
		t.Fatalf("expected fetcher to be called once, got %d", fetcher.Called)
	}
}


func TestFetch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    
	w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{
			"name": "Abia",
			"sys": { "country": "NG" },
			"main": {
				"temp": 311.69,
				"temp_min": 311.69,
				"temp_max": 311.69
			},
			"weather": [{
				"main": "Clouds",
				"description": "overcast clouds"
			}]
		}`))
	}))
	defer ts.Close()

	client := weather.NewClient("dummy-key")
	client.BaseURL = ts.URL
	client.HTTPClient = ts.Client()

	got, err := client.Fetch(context.Background(), "Abia,NG")
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Weather{
		City:        "Abia",
		Country:     "NG",
		Temp:        311.69,
		Temp_min:    311.69,
		Temp_max:    311.69,
		Condition:   "Clouds",
		Description: "overcast clouds",
	}

	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestParseResponse_ReturnsWeatherReport(t *testing.T) {
	t.Parallel()

	want := weather.Weather{
		City:			"Abia",
		Country:		"NG",
		Temp:			311.69,
		Temp_min:		311.69,
		Temp_max:		311.69,
		Condition:		"Clouds",
		Description:	"overcast clouds",		
	}

	data := []byte(`{
  "name": "Abia",
  "sys": { "country": "NG" },
  "main": {
    "temp": 311.69,
    "temp_min": 311.69,
    "temp_max": 311.69
  },
  "weather": [{
    "main": "Clouds",
    "description": "overcast clouds"
  }]
}`)

	got, err := weather.ParseResponse([]byte(data))
	if err != nil {
		t.Fatal("expected error")
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}