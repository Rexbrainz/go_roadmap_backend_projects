package http

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/convert/length", LengthHandler)
	mux.HandleFunc("/convert/weight", WeightHandler)
	mux.HandleFunc("/convert/temperature", TemperatureHandler)

	// Frontend api for template handlers
	mux.HandleFunc("/length", LengthPage)
	mux.HandleFunc("/weight", WeightPage)
	mux.HandleFunc("/temperature", TemperaturePage)
	mux.HandleFunc("/", IndexPage)

	return mux
}