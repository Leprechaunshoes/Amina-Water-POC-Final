package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// --- Data Structures ---
type Stats struct {
	TotalTransfers int    `json:"totalTransfers"`
	LastDonation   string `json:"lastDonation"`
}

type Donation struct {
	TxnId string `json:"txnId"`
	Date  string `json:"date"`
}

// --- Global Variables ---
var (
	mu           sync.Mutex
	stats        = Stats{TotalTransfers: 0, LastDonation: ""}
	lastDonation = Donation{TxnId: "", Date: ""}
)

// --- Handlers ---
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func handleStats(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(stats)
}

func handleLastDonation(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(lastDonation)
}

func handleSim(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	mu.Lock()
	defer mu.Unlock()

	stats.TotalTransfers++
	if stats.TotalTransfers >= 1000 {
		stats.LastDonation = time.Now().Format("2006-01-02 15:04:05")
		lastDonation.TxnId = "SAMPLE" + time.Now().Format("20060102150405")
		lastDonation.Date = stats.LastDonation
		log.Println("💧 Amina Humanity Fund: $100 donation triggered!")
		stats.TotalTransfers = 0
	}
	json.NewEncoder(w).Encode(stats)
}

// --- Main Server ---
func main() {
	// Serve static files (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	
	// API endpoints
	http.HandleFunc("/stats", handleStats)
	http.HandleFunc("/last-donation", handleLastDonation)
	http.HandleFunc("/simulate", handleSim)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🌍 Amina Water backend running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}