package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func runTestWithForkSlot(t *testing.T, forkSlot uint64, expectedOutput string) {
	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	// Override chainForkData for testing
	chainForkData["17000"] = ChainForkData{
		ForkSlot: forkSlot,
	}

	// Get environment variables
	chainId := "17000" // Force Holesky testnet
	beaconNode := os.Getenv("BEACON_URL")

	if beaconNode == "" {
		t.Fatal("BEACON_URL must be set in .env file")
	}

	// Run the script
	err = runScript(TArgs{
		ChainId:    chainId,
		BeaconNode: beaconNode,
	})

	// Restore stdout
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = oldStdout
	output := buf.String()

	if err != nil {
		t.Errorf("runScript failed: %v\nOutput:\n%s", err, output)
		return
	}

	if expectedOutput != "" && !strings.Contains(output, expectedOutput) {
		t.Errorf("Expected output to contain:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

// Test on a missed slot, should print out that it missed the first slot
func TestRunScript_ForkSlot3667156(t *testing.T) {
	expectedOutputs := []string{
		"Slot 3667156 was missed, checking next slot...",
		"Found first non-missed slot at slot 3667157",
		"Slot timestamp: 173990828",
	}
	runTestWithForkSlot(t, 3667156, strings.Join(expectedOutputs, "\n"))
}

// Test on a non-missed slot. Should immediately find the slot
func TestRunScript_ForkSlot3667157(t *testing.T) {
	expectedOutputs := []string{
		"Found first non-missed slot at slot 3667157",
		"Slot timestamp: 173990828",
	}
	runTestWithForkSlot(t, 3667157, strings.Join(expectedOutputs, "\n"))
}
