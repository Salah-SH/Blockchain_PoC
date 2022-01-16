package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cagnotteApp/x/cagnotte/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper of cagnotte store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

//This function allows the creation of cagnotte keeper.
//Ps: each module has its own keeper, which is responsible for th interaction with
//    the store and has the major part of the core functionality of a module.
//    It is also responsible for intermodule (cross-module) communications with external modules.

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
