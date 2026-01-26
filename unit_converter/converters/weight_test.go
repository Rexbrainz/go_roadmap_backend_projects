package converters

import (
	"testing"
)

func TestConvertWeight_GramsToKilograms(t *testing.T) {
	t.Parallel()

	want := 1.0
	got, err := ConvertWeight(1000, "gram", "kilogram")
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want , got)
	}
}

func TestConvertWeight_kilogramsToGrams(t *testing.T) {
	t.Parallel()

	want := 1000.0
	got, err := ConvertWeight(1, "kilogram", "gram")
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestConvertWeight_SameUnitRetursEnteredValue(t *testing.T) {
	t.Parallel()

	want := 42.0
	got, err := ConvertWeight(42, "gram", "gram")
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestConvertWeight_UnknownFromAndToUnit(t *testing.T) {
	t.Parallel()

	if _, err := ConvertWeight(10, "meter", "pound"); err == nil {
		t.Error(err)
	}

	if _, err := ConvertWeight(10, "milligram", "fahrenheit"); err == nil {
		t.Fatal(err)
	}
}

func TestConvertWeight_ZeroValueReturnsZero(t *testing.T) {
	t.Parallel()

	got, err := ConvertWeight(0, "gram", "kilogram") 
	if err != nil {
		t.Fatal(err)
	}

	if got != 0 {
		t.Fatalf("want 0, got %v", got)
	}
}

func TestConvertWeight_NegativeValues(t *testing.T) {
	t.Parallel()

	want := -50000.0
	got, err := ConvertWeight(-5, "gram", "milligram")
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}