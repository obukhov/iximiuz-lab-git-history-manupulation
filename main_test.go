package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStringArray(t *testing.T) {
	req, err := http.NewRequest("GET", "/signs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	signsHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expectedBody := `["Aries","Taurus","Gemini","Cancer","Leo","Virgo","Libra","Scorpio","Sagittarius","Capricorn","Aquarius","Pisces"]`
	if body := rr.Body.String(); body != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedBody)
	}
}
