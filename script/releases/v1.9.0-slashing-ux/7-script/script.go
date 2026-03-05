package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	allocationmanager "github.com/Layr-Labs/eigenlayer-contracts/pkg/bindings/AllocationManager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SidecarOperatorSet represents the JSON response from the sidecar API
type SidecarOperatorSet struct {
	AVS string `json:"avs"`
	ID  uint32 `json:"operatorSetId"`
}

// SidecarResponse represents the full response from the sidecar API
type SidecarResponse struct {
	OperatorSets []SidecarOperatorSet `json:"operatorSets"`
}

type ScriptArgs struct {
	SidecarURL            string
	AllocationManagerAddr string
	RPCEndpoint           string
	PrivateKey            string
}

func main() {
	// Get environment variables
	args := ScriptArgs{
		SidecarURL:            os.Getenv("SIDECAR_URL"),
		AllocationManagerAddr: os.Getenv("ZEUS_DEPLOYED_AllocationManager_Proxy"),
		RPCEndpoint:           os.Getenv("RPC_URL"),
		PrivateKey:            os.Getenv("PRIVATE_KEY"),
	}

	// Validate required environment variables
	if args.SidecarURL == "" {
		fmt.Println("Error: SIDECAR_URL environment variable is required")
		os.Exit(1)
	}
	if args.AllocationManagerAddr == "" {
		fmt.Println("Error: ZEUS_DEPLOYED_AllocationManager_Proxy environment variable is required")
		os.Exit(1)
	}
	if args.RPCEndpoint == "" {
		fmt.Println("Error: RPC_URL environment variable is required")
		os.Exit(1)
	}
	if args.PrivateKey == "" {
		fmt.Println("Error: PRIVATE_KEY environment variable is required")
		os.Exit(1)
	}

	err := runScript(args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func runScript(args ScriptArgs) error {
	// 1. Fetch operator sets from sidecar API
	operatorSets, err := fetchOperatorSets(args.SidecarURL)
	if err != nil {
		return fmt.Errorf("failed to fetch operator sets: %v", err)
	}

	fmt.Printf("Fetched %d operator sets from sidecar API\n", len(operatorSets))

	// If no operator sets, nothing to do
	if len(operatorSets) == 0 {
		fmt.Println("No operator sets to migrate")
		return nil
	}

	// 2. Convert to contract format
	operatorSetsToCall := convertToContractFormat(operatorSets)

	// 3. Process in batches of 20
	const batchSize = 20
	totalBatches := (len(operatorSetsToCall) + batchSize - 1) / batchSize

	fmt.Printf("Processing %d operator sets in %d batches of up to %d each\n", len(operatorSetsToCall), totalBatches, batchSize)

	for i := 0; i < len(operatorSetsToCall); i += batchSize {
		end := i + batchSize
		if end > len(operatorSetsToCall) {
			end = len(operatorSetsToCall)
		}

		batch := operatorSetsToCall[i:end]
		batchNum := i/batchSize + 1

		fmt.Printf("\n--- Processing batch %d/%d (%d operator sets) ---\n", batchNum, totalBatches, len(batch))

		err = callMigrateSlashers(args, batch)
		if err != nil {
			return fmt.Errorf("failed to process batch %d: %v", batchNum, err)
		}

		// Add a delay
		if end < len(operatorSetsToCall) {
			fmt.Println("Waiting before next batch...")
			time.Sleep(2 * time.Second)
		}
	}

	fmt.Printf("\n✅ Successfully migrated all %d operator sets in %d batches\n", len(operatorSetsToCall), totalBatches)

	return nil
}

func fetchOperatorSets(sidecarURL string) ([]SidecarOperatorSet, error) {
	// Make HTTP request to sidecar API
	url := fmt.Sprintf("%s/v1/operatorSets", strings.TrimRight(sidecarURL, "/"))

	fmt.Printf("Fetching operator sets from %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var response SidecarResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return response.OperatorSets, nil
}

func convertToContractFormat(sidecarSets []SidecarOperatorSet) []allocationmanager.OperatorSet {
	contractSets := make([]allocationmanager.OperatorSet, len(sidecarSets))

	for i, set := range sidecarSets {
		contractSets[i] = allocationmanager.OperatorSet{
			Avs: common.HexToAddress(set.AVS),
			Id:  set.ID,
		}
	}

	return contractSets
}

func callMigrateSlashers(args ScriptArgs, operatorSets []allocationmanager.OperatorSet) error {
	// Connect to Ethereum client
	client, err := ethclient.Dial(args.RPCEndpoint)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(args.PrivateKey, "0x"))
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Create auth transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	// Create AllocationManager instance
	allocationManager, err := allocationmanager.NewAllocationManager(common.HexToAddress(args.AllocationManagerAddr), client)
	if err != nil {
		return fmt.Errorf("failed to create AllocationManager instance: %v", err)
	}

	// Call migrateSlashers
	fmt.Printf("Calling migrateSlashers with %d operator sets:\n", len(operatorSets))
	for idx, opSet := range operatorSets {
		fmt.Printf("  [%d] AVS: %s, ID: %d\n", idx+1, opSet.Avs.Hex(), opSet.Id)
	}

	tx, err := allocationManager.MigrateSlashers(auth, operatorSets)
	if err != nil {
		return fmt.Errorf("failed to call migrateSlashers: %v", err)
	}

	fmt.Printf("Transaction sent! Hash: %s\n", tx.Hash().Hex())

	// Wait for transaction receipt
	fmt.Println("Waiting for transaction to be mined...")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %v", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}

	fmt.Printf("Transaction mined! Block: %d, Gas used: %d\n", receipt.BlockNumber.Uint64(), receipt.GasUsed)

	return nil
}
