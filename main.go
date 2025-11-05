package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// --- Data Structures ---

// Algorand Indexer API Response Structures
type IndexerResponse struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID                  string              `json:"id"`
	RoundTime           int64               `json:"round-time"`
	AssetTransferTxn    *AssetTransferTxn   `json:"asset-transfer-transaction,omitempty"`
	ApplicationCallTxn  *ApplicationCallTxn `json:"application-transaction,omitempty"`
}

type AssetTransferTxn struct {
	AssetID  int64  `json:"asset-id"`
	Amount   int64  `json:"amount"`
	Receiver string `json:"receiver"`
}

type ApplicationCallTxn struct {
	ApplicationID int64 `json:"application-id"`
}

// API Response Structures
type Stats struct {
	TotalTransfers int    `json:"totalTransfers"`
	LastDonation   string `json:"lastDonation"`
}

type Donation struct {
	TxnId string `json:"txnId"`
	Date  string `json:"date"`
}

// --- Global Variables ---
const (
	AMINA_ASA_ID       = 1107424865 // Amina Coin ASA ID
	TINYMAN_V2_APP_ID  = 1002541853 // Tinyman V2 AMM Application ID
	INDEXER_API        = "https://mainnet-idx.algonode.cloud/v2"
	CHECK_INTERVAL     = 30 * time.Second
)

var (
	mu               sync.Mutex
	stats            = Stats{TotalTransfers: 0, LastDonation: ""}
	lastDonation     = Donation{TxnId: "", Date: ""}
	lastCheckedRound int64 = 0
)

// --- Blockchain Monitor ---

// Monitor Algorand blockchain for Amina token swaps
func monitorAminaSwaps() {
	log.Println("🔍 Starting Amina swap monitor...")
	
	// Get current round to start from
	if lastCheckedRound == 0 {
		// Start from recent blocks (last ~5 minutes worth of blocks, ~12 blocks)
		lastCheckedRound = getCurrentRound() - 12
	}
	
	ticker := time.NewTicker(CHECK_INTERVAL)
	defer ticker.Stop()
	
	for range ticker.C {
		checkForNewSwaps()
	}
}

// Get current Algorand round
func getCurrentRound() int64 {
	resp, err := http.Get(INDEXER_API + "/health")
	if err != nil {
		log.Printf("Error checking current round: %v", err)
		return 0
	}
	defer resp.Body.Close()
	
	var health struct {
		Round int64 `json:"round"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&health); err != nil {
		log.Printf("Error decoding health response: %v", err)
		return 0
	}
	
	return health.Round
}

// Check for new Amina swaps since last checked round
func checkForNewSwaps() {
	currentRound := getCurrentRound()
	if currentRound == 0 || currentRound <= lastCheckedRound {
		return
	}
	
	// Query indexer for Amina ASA transactions
	url := fmt.Sprintf("%s/transactions?asset-id=%d&min-round=%d&max-round=%d",
		INDEXER_API, AMINA_ASA_ID, lastCheckedRound+1, currentRound)
	
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error querying indexer: %v", err)
		return
	}
	defer resp.Body.Close()
	
	var indexerResp IndexerResponse
	if err := json.NewDecoder(resp.Body).Decode(&indexerResp); err != nil {
		log.Printf("Error decoding indexer response: %v", err)
		return
	}
	
	// Count swap transactions
	swapCount := 0
	for _, txn := range indexerResp.Transactions {
		// Look for asset transfers of Amina token
		if txn.AssetTransferTxn != nil && txn.AssetTransferTxn.AssetID == AMINA_ASA_ID {
			swapCount++
		}
	}
	
	if swapCount > 0 {
		mu.Lock()
		stats.TotalTransfers += swapCount
		log.Printf("✅ Detected %d Amina transfers. Total: %d", swapCount, stats.TotalTransfers)
		
		// Check if we hit 1000 milestone
		if stats.TotalTransfers >= 1000 {
			stats.LastDonation = time.Now().Format("2006-01-02 15:04:05")
			lastDonation.TxnId = "PENDING_DONATION_TXN"
			lastDonation.Date = stats.LastDonation
			log.Println("🎉💧 MILESTONE REACHED! $100 donation triggered!")
			stats.TotalTransfers = stats.TotalTransfers % 1000 // Reset but keep overflow
		}
		mu.Unlock()
	}
	
	lastCheckedRound = currentRound
}

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
	log.Println("🌍 Amina Water - Real Blockchain Monitoring")
	
	// Start blockchain monitor in background
	go monitorAminaSwaps()
	
	// Serve static files (HTML, CSS, JS, images)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// API endpoints
		if r.URL.Path == "/stats" {
			handleStats(w, r)
			return
		}
		if r.URL.Path == "/last-donation" {
			handleLastDonation(w, r)
			return
		}
		if r.URL.Path == "/simulate" {
			handleSim(w, r)
			return
		}
		
		// Serve static files
		path := r.URL.Path
		if path == "/" {
			path = "/index.html"
		}
		
		// Set content type based on file extension
		contentType := "text/html"
		if path == "/style.css" {
			contentType = "text/css"
		} else if path == "/scripts.js" {
			contentType = "application/javascript"
		} else if path == "/coin.PNG" {
			contentType = "image/png"
		}
		
		w.Header().Set("Content-Type", contentType)
		http.ServeFile(w, r, "."+path)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port " + port)
	log.Println("📊 Serving website and API at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}