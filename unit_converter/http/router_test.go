package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter_LengthRoute(t *testing.T) {
	t.Parallel()

	router := NewRouter()
	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/length?value=1000&from=meter&to=kilometer",
		nil,
	)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Result().StatusCode)
	}

	if w.Body.String() != "1.00" {
		t.Fatalf("expected 1.00, got %q", w.Body.String())
	}
}

func TestRouter_WeightRoute(t *testing.T) {
	t.Parallel()

	router := NewRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/weight?value=1000&from=gram&to=kilogram",
		nil,
	)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Result().StatusCode)
	}

	if w.Body.String() != "1.00" {
		t.Fatalf("expected 1.00, got %q", w.Body.String())
	}
}

func TestRouter_TemperatureRoute(t *testing.T) {
	t.Parallel()

	router := NewRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/convert/temperature?value=212&from=fahrenheit&to=celsius",
		nil,
	)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Result().StatusCode)
	}

	if w.Body.String() != "100.00" {
		t.Fatalf("expected 100.00, got %q", w.Body.String())
	}
}

func TestRouter_UnknownRoute(t *testing.T) {
	t.Parallel()
	
	router := NewRouter()

	req := httptest.NewRequest(http.MethodGet, "/does-not-exist", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Result().StatusCode)
	}
}
