package weather

import (
	"net/http"
)

func (app *Application) route() http.Handler {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /weather", app.getWeather)
	return mux
}

func (app *Application) getWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	report, err := app.GetWeather(r.Context(), city)
	if err != nil {
		writeHTTPError(w, err)
	}

}

func (app *Application) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}