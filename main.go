package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Struktur data untuk informasi gempa dari BMKG
type BMKGResponse struct {
	Infogempa struct {
		Gempa map[string]string `json:"gempa"`
	} `json:"Infogempa"`
}

// Handler untuk endpoint gempa
func earthquakeHandler(w http.ResponseWriter, r *http.Request) {
	// Panggil API BMKG untuk mendapatkan informasi gempa terbaru
	resp, err := http.Get("https://data.bmkg.go.id/DataMKG/TEWS/autogempa.json")
	if err != nil {
		http.Error(w, "Failed to fetch earthquake data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Dekode respons JSON
	var bmkgResp BMKGResponse
	if err := json.NewDecoder(resp.Body).Decode(&bmkgResp); err != nil {
		http.Error(w, "Failed to decode earthquake data", http.StatusInternalServerError)
		return
	}

	shakeMapURL := "https://data.bmkg.go.id/DataMKG/TEWS/" + bmkgResp.Infogempa.Gempa["Shakemap"]
	bmkgResp.Infogempa.Gempa["Shakemap"] = shakeMapURL

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	// Kembalikan informasi gempa dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bmkgResp.Infogempa.Gempa)
}

func main() {
	// Routing endpoint
	http.HandleFunc("/earthquake", earthquakeHandler)

	// Mulai server di port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
