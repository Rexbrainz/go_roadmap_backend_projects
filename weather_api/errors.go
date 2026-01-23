package weather

import (
	"errors"
	"fmt"
	"net/http"
)

const Usage = `Send a Get http request to http://localhost:4000/weather?city={CITY_NAME,COUNTRYCODE}
Example: http://localhost:4000/weather?city=London,Uk`

var (
	ErrCityNotFound 			= fmt.Errorf("city not found\n%s", Usage)
	ErrProviderDown 			= errors.New("weather provider unavailable")
	ErrInvalidProviderResponse	= errors.New("invalid response from weather api provider")
)

func (app *Application) writeHTTPError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrCityNotFound):
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errors.Is(err, ErrProviderDown):
		http.Error(w, err.Error(), http.StatusBadGateway)
	case errors.Is(err, ErrInvalidProviderResponse):
		http.Error(w, err.Error(), http.StatusPartialContent)
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}