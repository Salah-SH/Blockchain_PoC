package types

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
)

// GenesisState - all cagnotte state that must be provided at genesis
type GenesisState struct {
	// TODO: Fill out what is needed by the module for genesis
	CagnotteRecords []Cagnotte `json:"cagnotte_records"`
	AdminAccount    AdminAddr  `json:"admin_account"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState( /* TODO: Fill out with what is needed for genesis state */ ) GenesisState {
	return GenesisState{
		// TODO: Fill out according to your genesis state
		CagnotteRecords: nil,
		AdminAccount:    AdminAddr{},
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		CagnotteRecords: []Cagnotte{},
		AdminAccount:    AdminAddr{},
	}
}

// ValidateGenesis validates the cagnotte genesis parameters
func ValidateGenesis(data GenesisState) error {
	// TODO: Create a sanity check to make sure the state conforms to the modules needs
	for _, record := range data.CagnotteRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid CagnotteisRecord: Owner: %s. Error: Missing Owner", record.Owner)
		}

	}
	if address := data.AdminAccount; address.Address.Empty() {
		return fmt.Errorf("invalid Admin account: Error: Missing Admin Address")
	}
	return nil
}

func GetGenesisStateFromAppState(cdc *codec.Codec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}
