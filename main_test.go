package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchDataHandler(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(dataHandler))
	defer server.Close()

	// Make a request to the test server
	response, err := http.Get(server.URL + "/data")
	if err != nil {
		t.Fatalf("Error making request to test server: %v", err)
	}
	defer response.Body.Close()

	// Check the status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	// Decode the JSON response
	var responseData map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		t.Fatalf("Error decoding JSON response: %v", err)
	}

	// Print the actual response
	fmt.Printf("Actual Response: %+v\n", responseData)

	// Check the presence of expected keys in the response
	expectedKeys := []string{"bitcoin_data", "ethereum_data", "tether_data"}
	for _, key := range expectedKeys {
		if _, exists := responseData[key]; !exists {
			t.Errorf("Expected key %s not found in the response", key)
		}
	}
}

func TestCryptoDataValues(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(dataHandler))
	defer server.Close()

	// Make a request to the test server
	response, err := http.Get(server.URL + "/data")
	if err != nil {
		t.Fatalf("Error making request to test server: %v", err)
	}
	defer response.Body.Close()

	// Check the status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	// Decode the JSON response
	var responseData map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		t.Fatalf("Error decoding JSON response: %v", err)
	}

	// Check the values for each cryptocurrency
	currencies := []string{"bitcoin_data", "ethereum_data", "tether_data"}
	for _, currency := range currencies {
		switch value := responseData[currency].(type) {
		case float64:
			fmt.Printf("%s: %f\n", currency, value)
		default:
			t.Errorf("Expected numeric value for %s, got %v", currency, responseData[currency])
		}
	}
}
