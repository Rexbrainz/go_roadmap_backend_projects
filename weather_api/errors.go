package weather

import (
	"errors"
	"net/http"
)

var (
	ErrCityNotFound = errors.New("city not found")
	ErrProviderDown = errors.New("weather provider unavailable")
)

func (app *Application) writeHTTPError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrCityNotFound):
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errors.Is(err, ErrProviderDown):
		http.Error(w, err.Error(), http.StatusBadGateway)
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}