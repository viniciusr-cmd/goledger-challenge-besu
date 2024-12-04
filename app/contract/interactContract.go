package contract

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SetContract(value uint) (*types.Receipt, error) {
	dataValueBigInt := new(big.Int).SetUint64(uint64(value))

	abiFile, err := os.ReadFile("../besu/artifacts/contracts/SimpleStorage.sol/SimpleStorage.json") // found under besu/artifacts/contracts/SimpleStorage.sol/SimpleStorage.json
	if err != nil {
		log.Fatalf("error reading ABI file: %v", err)
	}

	var abiJSON map[string]interface{}
	err = json.Unmarshal(abiFile, &abiJSON)
	if err != nil {
		log.Fatalf("error unmarshalling ABI file: %v", err)
	}

	abiString, err := json.Marshal(abiJSON["abi"])
	if err != nil {
		log.Fatalf("error marshalling ABI: %v", err)
	}

	abi, err := abi.JSON(strings.NewReader(string(abiString)))
	if err != nil {
		log.Fatalf("error parsing ABI: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, "http://localhost:8545") // e.g., http://localhost:8545
	if err != nil {
		log.Fatalf("error dialing node: %v", err)
	}

	slog.Info("querying chain id")

	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("error querying chain id: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")) // will be returned during startDev.sh execution

	boundContract := bind.NewBoundContract(
		contractAddress,
		abi,
		client,
		client,
		client,
	)

	priv, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY")) // this can be found in the genesis.json file
	if err != nil {
		log.Fatalf("error loading private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(priv, chainId)
	if err != nil {
		log.Fatalf("error creating transactor: %v", err)
	}

	tx, err := boundContract.Transact(auth, "set", &dataValueBigInt)
	if err != nil {
		log.Fatalf("error transacting: %v", err)
	}

	fmt.Println("waiting until transaction is mined",
		"tx", tx.Hash().Hex(),
	)

	receipt, err := bind.WaitMined(
		context.Background(),
		client,
		tx,
	)
	if err != nil {
		log.Fatalf("error waiting for transaction to be mined: %v", err)
	}

	fmt.Printf("transaction mined: %v\n", receipt)

	return receipt, err
}

func GetContract() (interface{}, error) {
	var result interface{}

	abiFile, err := os.ReadFile("../besu/artifacts/contracts/SimpleStorage.sol/SimpleStorage.json") // found under besu/artifacts/contracts/SimpleStorage.sol/SimpleStorage.json
	if err != nil {
		log.Fatalf("error reading ABI file: %v", err)
	}

	var abiJSON map[string]interface{}
	err = json.Unmarshal(abiFile, &abiJSON)
	if err != nil {
		log.Fatalf("error unmarshalling ABI file: %v", err)
	}

	abiString, err := json.Marshal(abiJSON["abi"])
	if err != nil {
		log.Fatalf("error marshalling ABI: %v", err)
	}

	abi, err := abi.JSON(strings.NewReader(string(abiString)))
	if err != nil {
		log.Fatalf("error parsing ABI: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, "http://localhost:8545") // e.g., http://localhost:8545
	if err != nil {
		log.Fatalf("error connecting to eth client: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")) // will be returned during startDev.sh execution
	caller := bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	boundContract := bind.NewBoundContract(
		contractAddress,
		abi,
		client,
		client,
		client,
	)

	var output []interface{}
	err = boundContract.Call(&caller, &output, "get")
	if err != nil {
		log.Fatalf("error calling contract: %v", err)
	}
	result = output

	fmt.Println("Successfully called contract!", result)

	return result, err
}
