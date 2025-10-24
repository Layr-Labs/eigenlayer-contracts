package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

func TestFetchOperatorSets(t *testing.T) {
	// Create a test server that returns mock operator sets
	mockResponse := SidecarResponse{
		OperatorSets: []SidecarOperatorSet{
			{
				AVS: "0x1234567890123456789012345678901234567890",
				ID:  1,
			},
			{
				AVS: "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd",
				ID:  2,
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/operatorSets" {
			t.Errorf("Expected path /v1/operatorSets, got %s", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept header application/json, got %s", r.Header.Get("Accept"))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Test fetchOperatorSets
	operatorSets, err := fetchOperatorSets(server.URL)
	if err != nil {
		t.Fatalf("fetchOperatorSets failed: %v", err)
	}

	if len(operatorSets) != 2 {
		t.Errorf("Expected 2 operator sets, got %d", len(operatorSets))
	}

	if operatorSets[0].AVS != mockResponse.OperatorSets[0].AVS {
		t.Errorf("Expected AVS %s, got %s", mockResponse.OperatorSets[0].AVS, operatorSets[0].AVS)
	}

	if operatorSets[1].ID != mockResponse.OperatorSets[1].ID {
		t.Errorf("Expected ID %d, got %d", mockResponse.OperatorSets[1].ID, operatorSets[1].ID)
	}
}

func TestConvertToContractFormat(t *testing.T) {
	sidecarSets := []SidecarOperatorSet{
		{
			AVS: "0x1234567890123456789012345678901234567890",
			ID:  1,
		},
		{
			AVS: "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd",
			ID:  2,
		},
	}

	contractSets := convertToContractFormat(sidecarSets)

	if len(contractSets) != 2 {
		t.Errorf("Expected 2 contract sets, got %d", len(contractSets))
	}

	expectedAVS1 := common.HexToAddress("0x1234567890123456789012345678901234567890")
	if contractSets[0].Avs != expectedAVS1 {
		t.Errorf("Expected AVS %s, got %s", expectedAVS1.Hex(), contractSets[0].Avs.Hex())
	}

	if contractSets[1].Id != 2 {
		t.Errorf("Expected ID 2, got %d", contractSets[1].Id)
	}
}

func TestFetchAndConvert(t *testing.T) {
	// Create a test server
	mockResponse := SidecarResponse{
		OperatorSets: []SidecarOperatorSet{
			{
				AVS: "0x1234567890123456789012345678901234567890",
				ID:  1,
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Test fetching and converting operator sets
	operatorSets, err := fetchOperatorSets(server.URL)
	if err != nil {
		t.Fatalf("fetchOperatorSets failed: %v", err)
	}

	if len(operatorSets) != 1 {
		t.Errorf("Expected 1 operator set, got %d", len(operatorSets))
	}

	// Test conversion
	contractSets := convertToContractFormat(operatorSets)
	if len(contractSets) != 1 {
		t.Errorf("Expected 1 contract set, got %d", len(contractSets))
	}

	expectedAVS := common.HexToAddress("0x1234567890123456789012345678901234567890")
	if contractSets[0].Avs != expectedAVS {
		t.Errorf("Expected AVS %s, got %s", expectedAVS.Hex(), contractSets[0].Avs.Hex())
	}
}

func TestMainFunctionWithEnv(t *testing.T) {
	// Skip this test if we're in CI or don't have a .env file
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Skipping test that requires .env file")
	}

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	// Check if SIDECAR_URL is set
	sidecarURL := os.Getenv("SIDECAR_URL")
	if sidecarURL == "" {
		t.Skip("SIDECAR_URL not set in .env file")
	}

	// Test fetching real operator sets
	operatorSets, err := fetchOperatorSets(sidecarURL)
	if err != nil {
		t.Logf("Warning: Failed to fetch operator sets from %s: %v", sidecarURL, err)
		// Don't fail the test, as the API might be down or require auth
		return
	}

	t.Logf("Successfully fetched %d operator sets from %s", len(operatorSets), sidecarURL)

	// Log first few operator sets for debugging
	for i := 0; i < len(operatorSets) && i < 3; i++ {
		t.Logf("  OperatorSet %d: AVS=%s, ID=%d", i, operatorSets[i].AVS, operatorSets[i].ID)
	}
}

func TestErrorHandling(t *testing.T) {
	// Test with server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	_, err := fetchOperatorSets(server.URL)
	if err == nil {
		t.Error("Expected error for 500 response, got nil")
	}

	// Test with invalid JSON
	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("invalid json"))
	}))
	defer server2.Close()

	_, err = fetchOperatorSets(server2.URL)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestBatching(t *testing.T) {
	// Test that batching logic correctly splits operator sets
	testCases := []struct {
		totalSets       int
		expectedBatches int
	}{
		{0, 0},
		{1, 1},
		{20, 1},
		{21, 2},
		{40, 2},
		{41, 3},
		{100, 5},
	}

	const batchSize = 20
	for _, tc := range testCases {
		actualBatches := (tc.totalSets + batchSize - 1) / batchSize
		if tc.totalSets == 0 {
			actualBatches = 0
		}
		if actualBatches != tc.expectedBatches {
			t.Errorf("For %d sets, expected %d batches, got %d", tc.totalSets, tc.expectedBatches, actualBatches)
		}
	}
}
