package http

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Rexbrainz/converters/converters"
)

type HTMLConverterFunc func(float64, string, string) (float64, error)

var (
	lengthTemplate = template.Must(template.ParseFiles(
		"templates/layout.html",
		 "templates/length.html",
		))
	weightTemplate = template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/weight.html",
		))
	temperatureTemplate = template.Must(template.ParseFiles(
		"templates/layout.html", 
		"templates/temperature.html",
		))
	indexTemplate = template.Must(
		template.ParseFiles(
			"templates/layout.html",
			"templates/index.html",
		))
)

type PageData struct {
	Result	string
	Error	string
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, nil)
}

func LengthPage(w http.ResponseWriter, r *http.Request) {
	handleHTMLConvert(w, r, lengthTemplate, converters.ConvertLength)
}

func WeightPage(w http.ResponseWriter, r *http.Request) {
	handleHTMLConvert(w, r, weightTemplate, converters.ConvertWeight)
}

func TemperaturePage(w http.ResponseWriter, r *http.Request) {
	handleHTMLConvert(w, r, temperatureTemplate, converters.ConvertTemperature)
}

func handleHTMLConvert(w http.ResponseWriter, r *http.Request,
	 tmpl *template.Template, convert HTMLConverterFunc) {

	data := PageData{}

	if r.Method == http.MethodPost {
		valueStr := r.FormValue("value")
		from := r.FormValue("from")
		to := r.FormValue("to")

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			data.Error = "Invalid value"
		} else {
			result, err := convert(value, from, to)
			if err != nil {
				data.Error = err.Error()
			} else {
				data.Result = strconv.FormatFloat(result, 'f', 2, 64)
			}
		}
	}
	tmpl.Execute(w, data)
}