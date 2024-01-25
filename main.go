package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	initRoutes()
	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func initRoutes() {
	http.HandleFunc("/signs", signsHandler)
	http.HandleFunc("/signByDate", signByDateHandler)
}

func signsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal([]string{
		"Aries",
		"Taurus",
		"Gemini",
		"Cancer",
		"Leo",
		"Virgo",
		"Libra",
		"Scorpio",
		"Sagittarius",
		"Capricorn",
		"Aquarius",
		"Pisces",
	})

	w.Write(b)
}

func signByDateHandler(w http.ResponseWriter, r *http.Request) {
	dateString := r.URL.Query().Get("date")
	if dateString == "" {
		http.Error(w, "Date parameter is required", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	fDate := date.Format("01-02")
	sign := "Unknown"

	for _, s := range []struct{ till, sign string }{
		{"01-19", "Capricorn"},
		{"02-18", "Aquarius"},
		{"03-21", "Pisces"},
		{"04-19", "Aries"},
		{"05-20", "Taurus"},
		{"06-20", "Gemini"},
		{"07-22", "Cancer"},
		{"08-22", "Leo"},
		{"09-22", "Virgo"},
		{"10-22", "Libra"},
		{"11-21", "Scorpio"},
		{"12-21", "Sagittarius"},
		{"12-31", "Capricorn"},
	} {
		if fDate <= s.till {
			sign = s.sign
			break
		}
	}

	jsonData, err := json.Marshal(sign)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
