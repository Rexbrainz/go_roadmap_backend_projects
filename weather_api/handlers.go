package weather

import (
	"net/http"
)

func (app *Application) getWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		app.writeHTTPError(w, ErrCityNotFound)
		return
	}
	report, err := app.GetWeather(r.Context(), city)
	if err != nil {
		app.writeHTTPError(w, err)
		return
	}
	if err := app.writeJSON(w, http.StatusOK, report); err != nil {
		app.writeHTTPError(w, err)
		return
	}
}


func (app *Application) ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
