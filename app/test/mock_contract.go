package test

import (
	"github.com/ethereum/go-ethereum/core/types"
)

var MockSetContract func(value uint) (*types.Receipt, error)
var MockGetContract func() (*types.Receipt, error)
var MockSyncContract func() (*types.Receipt, error)
var MockCheckContract func() (*types.Receipt, error)

func SetContract(value uint) (*types.Receipt, error) {
	if MockSetContract != nil {
		return MockSetContract(value)
	}
	return nil, nil
}

func GetContract() (*types.Receipt, error) {
	if MockGetContract != nil {
		return MockGetContract()
	}
	return nil, nil
}

func SyncContract() (*types.Receipt, error) {
	if MockSyncContract != nil {
		return MockSyncContract()
	}
	return nil, nil
}

func CheckContract() (*types.Receipt, error) {
	if MockCheckContract != nil {
		return MockCheckContract()
	}
	return nil, nil
}
