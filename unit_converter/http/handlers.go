package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rexbrainz/converters/converters"
)

type ConverterFunc func(float64, string, string) (float64, error)


func LengthHandler(w http.ResponseWriter, r *http.Request) {
	handleConvert(w, r, converters.ConvertLength)
}

func WeightHandler(w http.ResponseWriter, r *http.Request) {
	handleConvert(w, r, converters.ConvertWeight)
}

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	handleConvert(w, r, converters.ConvertTemperature)
}

func handleConvert(w http.ResponseWriter, r *http.Request, convert ConverterFunc) {
	// get a map of the query section of the url
	q := r.URL.Query()

	valueStr := q.Get("value")
	if valueStr == "" {
		http.Error(w, "missing value parameter", http.StatusBadRequest)
		return
	}

	from := q.Get("from")
	to := q.Get("to")
	if from == ""  || to == "" {
		http.Error(w, "missing unit parameter", http.StatusBadRequest)
		return
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		http.Error(w, "invalid value parameter", http.StatusBadRequest)
		return
	}

	result, err := convert(value, from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%.2f", result)
}