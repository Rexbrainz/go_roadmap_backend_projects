package converters

import (
	"testing"
)

func TestConvertTemperature_CelsiusToFahrenheit(t *testing.T) {
	t.Parallel()

	want := 32.0

	got, err := ConvertTemperature(0, "celsius", "fahrenheit")
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestConvertTemperature_FahrenheitToCelsius(t *testing.T) {
	t.Parallel()

	want := 100.0
	got, err := ConvertTemperature(212, "fahrenheit", "celsius")
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestConvertTemperature_SameUnit(t *testing.T) {
	t.Parallel()

	got, err := ConvertTemperature(25, "celsius", "celsius")
	if err != nil {
		t.Fatal(err)
	}

	if got != 25 {
		t.Fatalf("want 25, got %v", got)
	}
}

func TestConvertTemperature_CelsiusToKelvin(t *testing.T) {
	t.Parallel()

	want := 273.15
	got, err := ConvertTemperature(0, "celsius", "kelvin")
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestConvertTemperature_KelvinToCelsius(t *testing.T) {
	t.Parallel()

	got, err := ConvertTemperature(273.15, "kelvin", "celsius")
	if err != nil {
		t.Fatal(err)
	}

	if got != 0 {
		t.Fatalf("want 0, got %v", got)
	}
}

func TestConvertTemperature_FahrenheitToKelvin(t *testing.T) {
	got, err := ConvertTemperature(32, "fahrenheit", "kelvin")
	if err != nil {
		t.Fatal(err)
	}

	want := 273.15
	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}


func TestConvertTemperature_UnknownFromUnit(t *testing.T) {
	_, err := ConvertTemperature(10, "rankine", "celsius")
	if err == nil {
		t.Fatal("expected error for unknown from unit")
	}
}


func TestConvertTemperature_UnknownToUnit(t *testing.T) {
	_, err := ConvertTemperature(10, "celsius", "rankine")
	if err == nil {
		t.Fatal("expected error for unknown to unit")
	}
}


func TestConvertTemperature_AbsoluteZero(t *testing.T) {
	got, err := ConvertTemperature(0, "kelvin", "celsius")
	if err != nil {
		t.Fatal(err)
	}

	want := -273.15
	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}
