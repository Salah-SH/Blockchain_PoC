package cagnotte

import (
	"github.com/cagnotteApp/x/cagnotte/keeper"
	"github.com/cagnotteApp/x/cagnotte/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data types.GenesisState) {
	for _, record := range data.CagnotteRecords {
		keeper.SetCagnotte(ctx, record.Name, record)
	}
	adminAccount := data.AdminAccount
	keeper.SetAdmin(ctx, adminAccount.Address)

}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {

	var records []types.Cagnotte
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		cagnotte, _ := k.GetCagnotte(ctx, name)
		records = append(records, cagnotte)

	}
	adminAccount, _ := k.GetAdmin(ctx)

	return types.GenesisState{
		CagnotteRecords: records,
		AdminAccount:    types.AdminAddr{Address: adminAccount},
	}
}
