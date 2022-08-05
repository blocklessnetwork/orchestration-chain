package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		OrderList:            []Order{},
		CompletedOrderList:   []CompletedOrder{},
		ClaimedOrderList:     []ClaimedOrder{},
		NodeRegistrationList: []NodeRegistration{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in order
	orderIndexMap := make(map[string]struct{})

	for _, elem := range gs.OrderList {
		index := string(OrderKey(elem.Index))
		if _, ok := orderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for order")
		}
		orderIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in completedOrder
	completedOrderIndexMap := make(map[string]struct{})

	for _, elem := range gs.CompletedOrderList {
		index := string(CompletedOrderKey(elem.Index))
		if _, ok := completedOrderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for completedOrder")
		}
		completedOrderIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in claimedOrder
	claimedOrderIndexMap := make(map[string]struct{})

	for _, elem := range gs.ClaimedOrderList {
		index := string(ClaimedOrderKey(elem.Index))
		if _, ok := claimedOrderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for claimedOrder")
		}
		claimedOrderIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in nodeRegistration
	nodeRegistrationIndexMap := make(map[string]struct{})

	for _, elem := range gs.NodeRegistrationList {
		index := string(NodeRegistrationKey(elem.Index))
		if _, ok := nodeRegistrationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nodeRegistration")
		}
		nodeRegistrationIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
