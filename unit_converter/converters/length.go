package converters

import (
	"fmt"
)

var lengthFactors = map[string]float64 {
	"millimeter": 	0.001, 
	"centimeter": 	0.01,
	"meter":		1,
	"kilometer":	1000,
	"inch":			0.0254,
	"foot":			0.3048,
	"yard":			0.9144,
	"mile":			1609.34,
}

func ConvertLength(value float64, from, to string) (float64, error) {
	fromFactor, ok := lengthFactors[from]
	if !ok {
		return 0, fmt.Errorf("unknown unit: %s", from)
	}

	toFactor, ok := lengthFactors[to]
	if !ok {
		return 0, fmt.Errorf("unknown unit: %s", to)
	}

	meters := value * fromFactor
	result := meters / toFactor
	
	return result, nil
}

