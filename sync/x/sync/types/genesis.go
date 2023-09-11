package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		HeaderList: []Header{},
		Admin:      DefaultAdmin,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in header
	headerIdMap := make(map[uint64]bool)
	headerCount := gs.GetHeaderCount()
	for _, elem := range gs.HeaderList {
		if _, ok := headerIdMap[elem.BlockID]; ok {
			return fmt.Errorf("duplicated id for header")
		}
		if elem.BlockID >= headerCount {
			return fmt.Errorf("header id should be lower or equal than the last id")
		}
		headerIdMap[elem.BlockID] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
