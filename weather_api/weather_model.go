package weather

type Weather struct {
	City		string	`json:"city"`
	Country		string	`json:"country"`
	Temp		float64	`json:"temp"`
	Temp_min	float64	`json:"temp_min"`
	Temp_max	float64	`json:"temp_max"`
	Condition	string	`json:"condition"`
	Description	string	`json:"description"`
}

type WeatherReport struct {
	Weather 	[]condition `json:"weather"`
	Temp		main		`json:"main"`
	Name		string		`json:"name"`
	Country		sys			`json:"sys"`
}

type condition struct {
	Main		string	`json:"main"`
	Description	string	`json:"description"`
}

type main struct {
	Temp		float64	`json:"temp"`
	Temp_min	float64	`json:"temp_min"`
	Temp_max	float64	`json:"temp_max"`
}

type sys struct {
	Country	string	`json:"country"`
}
