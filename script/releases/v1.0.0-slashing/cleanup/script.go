package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/api"
	v1 "github.com/attestantio/go-eth2-client/api/v1"
	attestantio "github.com/attestantio/go-eth2-client/http"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	multicall "github.com/jbrower95/multicall-go"
	"github.com/samber/lo"
)

type TQueryAllEigenpodsOnNetworkArgs struct {
	Ctx               context.Context
	AllValidators     []ValidatorWithIndex
	Eth               *ethclient.Client
	EigenpodAbi       abi.ABI
	PodManagerAbi     abi.ABI
	PodManagerAddress string
	Mc                *multicall.MulticallClient
}

//go:embed EigenPod.abi.json
var EigenPodAbi string

//go:embed EigenPodManager.abi.json
var EigenPodManagerAbi string

type ValidatorWithIndex struct {
	Validator *v1.Validator
	Index     phase0.ValidatorIndex
}

type TArgs struct {
	Node       string
	BeaconNode string
}

func main() {
	err := runScript(TArgs{
		Node:       os.Getenv("RPC_URL"),
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

	eigenpodAbi, err := abi.JSON(strings.NewReader(EigenPodAbi))
	panicOnError("failed to load eigenpod abi", err)

	podManagerAbi, err := abi.JSON(strings.NewReader(EigenPodManagerAbi))
	panicOnError("failed to load eigenpod manager abi", err)

	eth, err := ethclient.Dial(args.Node)
	panicOnError("failed to reach eth node", err)

	beaconClient, err := attestantio.New(ctx,
		attestantio.WithAddress(args.BeaconNode),
	)
	panicOnError("failed to reach beacon node", err)

	panicOnError("failed to reach ethereum clients", err)

	mc, err := multicall.NewMulticallClient(ctx, eth, &multicall.TMulticallClientOptions{
		MaxBatchSizeBytes: 8192,
	})
	panicOnError("error initializing mc", err)

	podManagerAddress := os.Getenv("ZEUS_DEPLOYED_EigenPodManager_Proxy")

	// fetch latest beacon state.
	_validators := (func() *map[phase0.ValidatorIndex]*v1.Validator {
		if provider, isProvider := beaconClient.(eth2client.ValidatorsProvider); isProvider {
			validators, err := provider.Validators(ctx, &api.ValidatorsOpts{
				State: "head",
				Common: api.CommonOpts{
					Timeout: 60 * time.Second,
				},
			})
			panicOnError("failed to load validator set", err)
			return &validators.Data
		}
		return nil
	})()
	if _validators == nil {
		panic("failed to load validators")
	}
	validators := *_validators

	panicOnError("failed to load beacon state", err)

	panicOnError("failed to fetch validators", err)
	allValidators := lo.Map(lo.Keys(validators), func(idx phase0.ValidatorIndex, i int) ValidatorWithIndex {
		return ValidatorWithIndex{
			Validator: validators[idx],
			Index:     idx,
		}
	})

	allEigenpods, err := queryAllEigenpodsOnNetwork(ctx, allValidators, eth, &eigenpodAbi, &podManagerAbi, podManagerAddress, mc)
	panicOnError("queryAllEigenpodsOnNetwork", err)

	fmt.Printf("Discovered %d eigenpods on the network.\n", len(allEigenpods))

	// Now for each eigenpod, we want to fetch currentCheckpointTimestamp.
	// We'll do a multicall to get currentCheckpointTimestamp from each eigenpod.
	checkpointTimestamps, err := fetchCurrentCheckpointTimestamps(allEigenpods, &eigenpodAbi, mc)
	panicOnError("failed to fetch currentCheckpointTimestamps", err)

	// Build a result set
	type EigenpodInfo struct {
		Address                    string `json:"address"`
		CurrentCheckpointTimestamp uint64 `json:"currentCheckpointTimestamp"`
	}
	results := []EigenpodInfo{}

	for i, ep := range allEigenpods {
		if checkpointTimestamps[i] > 0 {
			results = append(results, EigenpodInfo{
				Address:                    ep,
				CurrentCheckpointTimestamp: checkpointTimestamps[i],
			})
		}
	}

	fmt.Printf("%d EigenPods had active checkpoints\n\n", len(results))

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	fmt.Printf("Completing these checkpoints....")

	// TODO: complete the checkpoints.
	// TODO: pass in a funded wallet.

	return enc.Encode(results)
}

// This is a simplified version of the queryAllEigenpodsOnNetwork function inline.
// It uses the logic from the provided code snippet in the commands package.
func queryAllEigenpodsOnNetwork(
	ctx context.Context,
	allValidators []ValidatorWithIndex,
	eth *ethclient.Client,
	eigenpodAbi, podManagerAbi *abi.ABI,
	podManagerAddress string,
	mc *multicall.MulticallClient,
) ([]string, error) {
	args := TQueryAllEigenpodsOnNetworkArgs{
		Ctx:               ctx,
		AllValidators:     allValidators,
		Eth:               eth,
		EigenpodAbi:       *eigenpodAbi,
		PodManagerAbi:     *podManagerAbi,
		PodManagerAddress: podManagerAddress,
		Mc:                mc,
	}
	return internalQueryAllEigenpodsOnNetwork(args)
}

// internalQueryAllEigenpodsOnNetwork is lifted from the provided snippet.
func internalQueryAllEigenpodsOnNetwork(args TQueryAllEigenpodsOnNetworkArgs) ([]string, error) {
	// Filter out validators that are withdrawing to execution layer addresses
	executionLayerWithdrawalCredentialValidators := lo.Filter(args.AllValidators, func(validator ValidatorWithIndex, i int) bool {
		return validator.Validator.Validator.WithdrawalCredentials[0] == 1
	})

	interestingWithdrawalAddresses := lo.Keys(lo.Reduce(executionLayerWithdrawalCredentialValidators, func(accum map[string]int, next ValidatorWithIndex, index int) map[string]int {
		accum[common.Bytes2Hex(next.Validator.Validator.WithdrawalCredentials[12:])] = 1
		return accum
	}, map[string]int{}))

	fmt.Printf("Querying %d addresses to see if they may be eigenpods\n", len(interestingWithdrawalAddresses))

	podOwners, err := multicall.DoManyAllowFailures[common.Address](args.Mc, lo.Map(interestingWithdrawalAddresses, func(address string, index int) *multicall.MultiCallMetaData[common.Address] {
		callMeta, err := multicall.Describe[common.Address](
			common.HexToAddress(address),
			args.EigenpodAbi,
			"podOwner",
		)
		panicOnError("failed to form mc", err)
		return callMeta
	})...)

	if podOwners == nil || err != nil || len(*podOwners) == 0 {
		panicOnError("failed to fetch podOwners", err)
		panic("loaded no pod owners")
	}

	podToPodOwner := map[string]*common.Address{}
	addressesWithPodOwners := lo.Filter(interestingWithdrawalAddresses, func(address string, i int) bool {
		success := (*podOwners)[i].Success
		if success {
			podToPodOwner[address] = (*podOwners)[i].Value
		}
		return success
	})

	fmt.Printf("Querying %d addresses on (EigenPodManager=%s) to see if it knows about these eigenpods\n", len(addressesWithPodOwners), args.PodManagerAddress)

	eigenpodForOwner, err := multicall.DoManyAllowFailures(
		args.Mc,
		lo.Map(addressesWithPodOwners, func(address string, i int) *multicall.MultiCallMetaData[common.Address] {
			claimedOwner := *podToPodOwner[address]
			call, err := multicall.Describe[common.Address](
				common.HexToAddress(args.PodManagerAddress),
				args.PodManagerAbi,
				"ownerToPod",
				claimedOwner,
			)
			panicOnError("failed to form multicall", err)
			return call
		})...,
	)
	panicOnError("failed to query", err)

	// now, see which are properly eigenpods
	return lo.Filter(addressesWithPodOwners, func(address string, i int) bool {
		return (*eigenpodForOwner)[i].Success && (*eigenpodForOwner)[i].Value.Cmp(common.HexToAddress(addressesWithPodOwners[i])) == 0
	}), nil
}

func fetchCurrentCheckpointTimestamps(
	allEigenpods []string,
	eigenpodAbi *abi.ABI,
	mc *multicall.MulticallClient,
) ([]uint64, error) {
	calls := lo.Map(allEigenpods, func(eigenpod string, i int) *multicall.MultiCallMetaData[uint64] {
		call, err := multicall.Describe[uint64](
			common.HexToAddress(eigenpod),
			*eigenpodAbi,
			"currentCheckpointTimestamp",
		)
		panicOnError("failed to form multicall", err)
		return call
	})

	results, err := multicall.DoMany(mc, calls...)
	if err != nil {
		return nil, err
	}

	out := make([]uint64, len(*results))
	for i, r := range *results {
		out[i] = *r
	}
	return out, nil
}
