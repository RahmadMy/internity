package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	// Inisialisasi router bawaan Go
	mux := http.NewServeMux()

	// Route percobaan menggunakan Go Native
	mux.HandleFunc("GET /api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		
		response := map[string]string{
			"message": "Backend Internity berjalan dengan Go Native!",
		}
		json.NewEncoder(w).Encode(response)
	})

	// Menjalankan server di port 8080
	http.ListenAndServe(":8080", mux)
}