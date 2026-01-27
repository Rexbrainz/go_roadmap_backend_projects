package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLengthHandler_MetersToKilometers(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/length?value=1000&from=meter&to=kilometer",
		nil,
	)
	w := httptest.NewRecorder()

	LengthHandler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}

	body := w.Body.String()
	if body != "1.00" {
		t.Fatalf("expected body '1', got %q", body)
	}
}

func TestLengthHandler_MissingValue(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/length?from=meter&to=kilometer",
		nil,
	)
	w := httptest.NewRecorder()

	LengthHandler(w, req)

	res := w.Result()
	
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", res.StatusCode)
	}
}

func TestLengthHandler_InvalidValue(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/length?value=abc&from=meter&to=kilometer",
		nil,
	)
	w := httptest.NewRecorder()

	LengthHandler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", res.StatusCode)
	}
}

func TestLengthHandler_MissingUnit(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/length?value=1000",
		nil,
	)
	w := httptest.NewRecorder()

	LengthHandler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", res.StatusCode)
	}
}

func TestLengthHandler_UnknownUnit(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/length?value=10&from=lightyear&to=meter",
		nil,
	)

	w := httptest.NewRecorder()

	LengthHandler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", res.StatusCode)
	}
}

func TestWeightHandler_GramsToKilograms(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/weight?value=1000&from=gram&to=kilogram",
		nil,
	)
	w := httptest.NewRecorder()

	WeightHandler(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %v", resp.StatusCode)
	}
	
	body := w.Body.String()
	if body != "1.00" {
		t.Fatalf("expected 1, got %q", body)
	}
}

func TestWeightHandler_MissingValueParameter(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/weight?from=gram&to=kilogram",
		nil,
	)
	w := httptest.NewRecorder()

	WeightHandler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestWeightHandler_MissingUnitParameters(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/weight?value=1000",
		nil,
	)
	w := httptest.NewRecorder()

	WeightHandler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestWeightHandler_InvalidValueParameter(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/weight?value=abc&from=gram&to=kilogram",
		nil,
	)
	w := httptest.NewRecorder()

	WeightHandler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestWeightHandler_InvalidUnitParameter(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/weight?value=100&from=gram&to=fahrenheit",
		nil,
	)
	w := httptest.NewRecorder()

	WeightHandler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestTemperatureHandler_CelsiusToFahrenheit(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/temperature?value=0&from=celsius&to=fahrenheit",
		nil,
	)
	w := httptest.NewRecorder()

	TemperatureHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected %v, got %v", http.StatusOK, res.StatusCode)
	}
	
	body := w.Body.String()
	if body != "32.00" {
		t.Fatalf("expected 32, got %q", body)
	}
}

func TestTemperatureHandler_MissingValueParameter(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/temperature?from=celsius&to=fahrenheit",
		nil,
	)
	w := httptest.NewRecorder()

	TemperatureHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestTemperatureHandler_MissingUnitParameter(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/temperature?value=32.0",
		nil,
	)
	w := httptest.NewRecorder()

	TemperatureHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestTemperatureHandler_InvalidValueParameter(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/temperature?value=abc&from=celsius&to=fahrenheit",
		nil,
	)
	w := httptest.NewRecorder()

	TemperatureHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestTemperatureHandler_InvalidUnitParameter(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/temperature?value=0&from=persec&to=fahrenheit",
		nil,
	)
	w := httptest.NewRecorder()

	TemperatureHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected %v, got %v", http.StatusBadRequest, res.StatusCode)
	}
}

func TestTemperatureHandler_FahrenheitToCelsius(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/temperature?value=212&from=fahrenheit&to=celsius",
		nil,
	)
	w := httptest.NewRecorder()

	TemperatureHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected %v, got %v", http.StatusOK, res.StatusCode)
	}

	body := w.Body.String()
	if body != "100.00" {
		t.Fatalf("expected 100.00, got %q", body)
	}
}
