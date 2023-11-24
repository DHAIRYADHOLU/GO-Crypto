package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CryptoData represents the JSON response structure from the external API.
type CryptoData struct {
	Bitcoin  map[string]float64 `json:"bitcoin"`
	Ethereum map[string]float64 `json:"ethereum"`
	Tether   map[string]float64 `json:"tether"`
}

// fetchData fetches the current data for Bitcoin, Ethereum, and Tether from the external API.
func fetchData() (CryptoData, error) {
	// You can replace the API URL with the one you selected.
	apiURL := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum,tether&vs_currencies=cad"

	// Make an HTTP request to the external API.
	response, err := http.Get(apiURL)
	if err != nil {
		return CryptoData{}, err
	}
	defer response.Body.Close()

	// Decode the JSON response.
	var cryptoData CryptoData
	err = json.NewDecoder(response.Body).Decode(&cryptoData)
	if err != nil {
		return CryptoData{}, err
	}

	return cryptoData, nil
}

// dataHandler is the HTTP handler for the "data" endpoint.
func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch the current data for Bitcoin, Ethereum, and Tether.
	cryptoData, err := fetchData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with individual JSON payloads for each coin.
	w.Header().Set("Content-Type", "application/json")

	// Bitcoin JSON response
	bitcoinJSON, err := json.Marshal(map[string]interface{}{
		"bitcoin_data": cryptoData.Bitcoin["cad"],
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s\n", bitcoinJSON)

	// Ethereum JSON response
	ethereumJSON, err := json.Marshal(map[string]interface{}{
		"ethereum_data": cryptoData.Ethereum["cad"],
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s\n", ethereumJSON)

	// Tether JSON response
	tetherJSON, err := json.Marshal(map[string]interface{}{
		"tether_data": cryptoData.Tether["cad"],
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s\n", tetherJSON)
}

func main() {
	// Register the dataHandler function for the "/data" endpoint.
	http.HandleFunc("/data", dataHandler)

	// Start the HTTP server
	fmt.Println("Server listening on :9090")
	http.ListenAndServe(":9090", nil)
}
