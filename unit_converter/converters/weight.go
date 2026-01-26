package converters

import (
	"fmt"
)

var weightFactors = map[string]float64 {
	"milligram":	0.0001,
	"gram":			1,
	"kilogram":		1000,
	"ounce":		28.3495,
	"pound":		453.592,
}


func ConvertWeight(value float64, from, to string) (float64, error) {
	fromFactor, ok := weightFactors[from]
	if !ok {
		return 0, fmt.Errorf("unknown unit: %s", from)
	}

	toFactor, ok := weightFactors[to]
	if !ok {
		return 0, fmt.Errorf("unknown unit: %s", to)
	}

	grams := value * fromFactor
	result := grams / toFactor
	
	return result, nil
}
