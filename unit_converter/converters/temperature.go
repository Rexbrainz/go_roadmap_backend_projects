package converters

import (
	"fmt"
)

func ConvertTemperature(value float64, from, to string) (float64, error) {
	celsius, err := toCelsius(value, from)
	if err != nil {
		return 0, err
	}

	return fromCelsius(celsius, to)
}

func toCelsius(value float64, unit string) (float64, error) {
	switch unit {
	case "celsius":
		return value, nil
	case "fahrenheit":
		return (value- 32) * 5 / 9, nil
	case "kelvin":
		return value - 273.15, nil
	default:
		return 0, fmt.Errorf("unknown unit: %s", unit)
	}
}

func fromCelsius(value float64, unit string) (float64, error) {
	switch unit {
	case "celsius":
		return value, nil
	case "fahrenheit":
		return value * 9/5 + 32, nil
	case "kelvin":
		return value + 273.15, nil
	default:
		return 0, fmt.Errorf("unknown unit %s", unit)
	}
}
