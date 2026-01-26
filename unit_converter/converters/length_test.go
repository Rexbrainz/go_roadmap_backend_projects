package converters

import (
	"testing"
)

func TestConvertLength_MetersToKilometers(t *testing.T) {
	t.Parallel()

	want := 1.0
	got, err := ConvertLength(1000, "meter", "kilometer")
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestConvertLength_KilometerToMeters(t *testing.T) {
	t.Parallel()

	want := 1000.0
	got, err := ConvertLength(1, "kilometer", "meter")
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestConvertLength_SameUnitRetursEnteredValue(t *testing.T) {
	t.Parallel()

	want := 42.0
	got, err := ConvertLength(42, "meter", "meter")
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestConvertLength_UnknownFromAndToUnit(t *testing.T) {
	t.Parallel()

	if _, err := ConvertLength(10, "Lightyear", "meter"); err == nil {
		t.Error(err)
	}

	if _, err := ConvertLength(10, "meter", "parsec"); err == nil {
		t.Fatal(err)
	}
}

func TestConvertLength_ZeroValueReturnsZero(t *testing.T) {
	t.Parallel()

	got, err := ConvertLength(0, "kilometer", "meter") 
	if err != nil {
		t.Fatal(err)
	}

	if got != 0 {
		t.Fatalf("want 0, got %v", got)
	}
}

func TestConvertLength_NegativeValues(t *testing.T) {
	t.Parallel()

	want := -500.0
	got, err := ConvertLength(-5, "meter", "centimeter")
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}