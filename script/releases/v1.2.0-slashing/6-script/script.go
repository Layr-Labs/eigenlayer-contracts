package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/attestantio/go-eth2-client/api"
	attestantio "github.com/attestantio/go-eth2-client/http"
)

type SignedBeaconBlockResponse struct {
	Data struct {
		Message struct {
			Body struct {
				ExecutionPayload struct {
					Timestamp string `json:"timestamp"`
				} `json:"execution_payload"`
			} `json:"body"`
		} `json:"message"`
	} `json:"data"`
}

type TArgs struct {
	ChainId    string
	BeaconNode string
}

type ChainForkData struct {
	ForkSlot uint64
}

var chainForkData = map[string]ChainForkData{
	"1": { // mainnet
		ForkSlot: 0,
	},
	"17000": { // holesky
		ForkSlot: 3710976,
	},
}

func main() {
	err := runScript(TArgs{
		ChainId:    os.Getenv("CHAIN_ID"),
		BeaconNode: os.Getenv("BEACON_URL"),
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func panicOnError(msg string, err error) {
	if err != nil {
		fmt.Printf("Error: %s", msg)
		panic(err)
	}
}

func runScript(args TArgs) error {
	ctx := context.Background()

	// Validate chain ID
	if args.ChainId != "1" && args.ChainId != "17000" {
		return fmt.Errorf("invalid chain ID: %s. Must be either '1' (mainnet) or '17000' (holesky)", args.ChainId)
	}

	forkData := chainForkData[args.ChainId]

	beaconClient, err := attestantio.New(ctx,
		attestantio.WithAddress(args.BeaconNode),
	)
	panicOnError("failed to reach beacon node", err)

	httpClient := beaconClient.(*attestantio.Service)
	// Start checking from fork slot
	slotNum := forkData.ForkSlot

	for {
		opts := &api.BeaconBlockHeaderOpts{Block: strconv.FormatUint(slotNum, 10)}
		_, err := httpClient.BeaconBlockHeader(ctx, opts)
		if err != nil {
			fmt.Printf("Slot %d was missed, checking next slot...\n", slotNum)
			slotNum++
			continue
		}

		fmt.Printf("Found first non-missed slot at slot %d\n", slotNum)
		break
	}

	// Get the slot timestamp. We don't need to use the beacon api because it returns a versioned state. Just use the vanilla endpoint
	url := fmt.Sprintf("%s/eth/v2/beacon/blocks/%d", args.BeaconNode, slotNum)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	var blockResponse SignedBeaconBlockResponse
	if err := json.NewDecoder(resp.Body).Decode(&blockResponse); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	timestamp := blockResponse.Data.Message.Body.ExecutionPayload.Timestamp
	fmt.Printf("Slot timestamp: %s\n", timestamp)

	return nil
}
