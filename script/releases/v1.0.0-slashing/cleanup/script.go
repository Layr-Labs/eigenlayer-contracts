package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"
	"time"

	proofgen "github.com/Layr-Labs/eigenpod-proofs-generation/cli/core"
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

type EigenpodInfo struct {
	Address                    string `json:"address"`
	CurrentCheckpointTimestamp uint64 `json:"currentCheckpointTimestamp"`
}

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
	Sender     string
}

func main() {
	err := runScript(TArgs{
		Node:       os.Getenv("RPC_URL"),
		BeaconNode: os.Getenv("BEACON_URL"),
		Sender:     os.Getenv("SENDER_PK"),
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

	chainId, err := eth.ChainID(ctx)
	panicOnError("failed to read chainId", err)

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

	fmt.Printf("Found %d validators\n", len(validators))

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

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	fmt.Printf("Discovered %d eigenpods on the network.\n", len(allEigenpods))

	pods := lo.Map(allEigenpods, func(pod string, i int) string {
		return fmt.Sprintf("0x%s", pod)
	})
	sort.Strings(pods)
	fmt.Printf("%s\n", enc.Encode(pods))

	// Now for each eigenpod, we want to fetch currentCheckpointTimestamp.
	// We'll do a multicall to get currentCheckpointTimestamp from each eigenpod.
	checkpointTimestamps, err := fetchCurrentCheckpointTimestamps(allEigenpods, &eigenpodAbi, mc)
	panicOnError("failed to fetch currentCheckpointTimestamps", err)

	results := []EigenpodInfo{}

	for i, ep := range allEigenpods {
		if checkpointTimestamps[i] > 0 {
			results = append(results, EigenpodInfo{
				Address:                    fmt.Sprintf("0x%s", ep),
				CurrentCheckpointTimestamp: checkpointTimestamps[i],
			})
		}
	}

	if len(results) == 0 {
		fmt.Printf("No eigenpods had active checkpoints. OK.")
		return nil
	}

	fmt.Printf("%d EigenPods had active checkpoints\n\n", len(results))
	fmt.Printf("%s\n", enc.Encode(results))

	fmt.Printf("Completing %d checkpoints....", len(results))
	coreBeaconClient, _, err := proofgen.NewBeaconClient(args.BeaconNode, true /* verbose */)
	panicOnError("failed to instantiate beaconClient", err)

	for i := 0; i < len(results); i++ {
		fmt.Printf("Completing [%d/%d]...", i+1, len(results))
		fmt.Printf("NOTE: this is expensive, and may take several minutes.")
		completeCheckpointForEigenpod(ctx, results[i].Address, eth, chainId, coreBeaconClient, args.Sender)
	}

	checkpointTimestamps, err = fetchCurrentCheckpointTimestamps(allEigenpods, &eigenpodAbi, mc)
	panicOnError("failed to fetch currentCheckpointTimestamps", err)

	// require that all eigenpods have a checkpoint timestamp of 0
	for i, timestamp := range checkpointTimestamps {
		if timestamp != 0 {
			panic(fmt.Sprintf("expected all eigenpods to have a checkpoint timestamp of 0, but found %d on %s", timestamp, allEigenpods[i]))
		}
	}

	return nil
}

func completeCheckpointForEigenpod(ctx context.Context, eigenpodAddress string, eth *ethclient.Client, chainId *big.Int, coreBeaconClient proofgen.BeaconClient, sender string) {
	res, err := proofgen.GenerateCheckpointProof(ctx, eigenpodAddress, eth, chainId, coreBeaconClient, true)
	panicOnError(fmt.Sprintf("failed to generate checkpoint proof for eigenpod:%s", eigenpodAddress), err)

	txns, err := proofgen.SubmitCheckpointProof(ctx, sender, eigenpodAddress, chainId, res, eth, 80 /* ideal checkpoint proof batch size */, true /* noPrompt */, false /* noSend */, true /* verbose */)
	panicOnError(fmt.Sprintf("failed to submit checkpoint proof for eigenpod:%s", eigenpodAddress), err)
	if txns == nil {
		panic("submitting checkpoint proof generated no transactions. this is a bug.")
	}

	for i, txn := range txns {
		fmt.Printf("[%d/%d] %s\n", i+1, len(txns), txn.Hash())
	}
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

	fmt.Printf("Querying %d beacon-chain withdrawal addresses to see if they may be eigenpods\n", len(interestingWithdrawalAddresses))

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

	eigenpodForOwner, err := multicall.DoMany(
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
		return (*eigenpodForOwner)[i].Cmp(common.HexToAddress(addressesWithPodOwners[i])) == 0
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
