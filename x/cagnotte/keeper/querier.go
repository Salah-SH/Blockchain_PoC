package keeper

import (
	// this line is used by starport scaffolding

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cagnotteApp/x/cagnotte/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for cagnotte clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2

		case types.QueryGetCagnotte:

			return getCagnotte(ctx, path[1:], k)
		case types.QueryListCagnotte:
			return listCagnotteByUser(ctx, k, path[1:])
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown cagnotte query endpoint")
		}
	}
}
