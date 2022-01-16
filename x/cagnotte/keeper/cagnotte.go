package keeper

import (
	"errors"

	"github.com/cagnotteApp/x/cagnotte/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateCagnotte(ctx sdk.Context, name string, creator sdk.AccAddress) error {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.CagnottePrefix + name)

	cagnotte := types.Cagnotte{
		Owner:  creator,
		Status: true,
		Name:   name,
	}
	exist := k.Exists(ctx, name)
	if exist {
		return errors.New("this  name is already used")
	}
	value := k.cdc.MustMarshalBinaryLengthPrefixed(cagnotte)
	store.Set(key, value)

	return nil
}

func (k Keeper) GetCagnotte(ctx sdk.Context, key string) (types.Cagnotte, error) {
	store := ctx.KVStore(k.storeKey)
	var cagnotte types.Cagnotte
	byteKey := []byte(types.CagnottePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &cagnotte)
	if err != nil {
		return cagnotte, err
	}

	return cagnotte, nil
}

// Check if the key exists in the store
func (k Keeper) Exists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)

	return store.Has([]byte(types.CagnottePrefix + key))
}

// SetCagnotte sets a cagnotte.
func (k Keeper) SetCagnotte(ctx sdk.Context, name string, cagnotte types.Cagnotte) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(cagnotte)
	key := []byte(types.CagnottePrefix + name)
	store.Set(key, bz)
}

func (k Keeper) CloseCagnotte(ctx sdk.Context, name string) error {
	cagnotte, _ := k.GetCagnotte(ctx, name)
	if len(cagnotte.PendingTransactions) > 0 {
		return errors.New("The cagnotte cannot be closed")
	}
	cagnotte.Status = false
	k.SetCagnotte(ctx, name, cagnotte)
	return nil
}

func (k Keeper) AddPendingAmount(ctx sdk.Context, name string, client sdk.AccAddress, pendingamount int) {
	cagnotte, _ := k.GetCagnotte(ctx, name)
	user := types.Participator{
		User:   client,
		Amount: pendingamount,
	}
	cagnotte.PendingTransactions = append(cagnotte.PendingTransactions, user)
	k.SetCagnotte(ctx, name, cagnotte)
}

func (k Keeper) findPending(ctx sdk.Context, name string, amount int, sender sdk.AccAddress) (bool, int) {
	cagnotte, _ := k.GetCagnotte(ctx, name)
	var rank int
	found := false
	for index, participation := range cagnotte.PendingTransactions {
		if participation.User.Equals(sender) {
			if participation.Amount == amount {
				found = true
				rank = index
				break
			}
		}
	}
	return found, rank
	// if found {
	// 	cagnotte.PendingTransactions = append(cagnotte.PendingTransactions[1:], cagnotte.PendingTransactions[1+1:]...)
	// }

}
func (k Keeper) AddAmount(ctx sdk.Context, name string, amount int, sender sdk.AccAddress, results bool, user sdk.AccAddress) error {

	cagnotte, _ := k.GetCagnotte(ctx, name)

	exist, rank := k.findPending(ctx, cagnotte.Name, amount, user)
	if exist {
		if results {
			cagnotte.Amount = cagnotte.Amount + cagnotte.PendingTransactions[rank].Amount
			cagnotte.ValidTransactions = append(cagnotte.ValidTransactions, cagnotte.PendingTransactions[rank])
			cagnotte.PendingTransactions = append(cagnotte.PendingTransactions[:rank], cagnotte.PendingTransactions[rank+1:]...)
		} else {
			cagnotte.InvalidTransactions = append(cagnotte.InvalidTransactions, cagnotte.PendingTransactions[rank])
			cagnotte.PendingTransactions = append(cagnotte.PendingTransactions[:rank], cagnotte.PendingTransactions[rank+1:]...)
		}
	} else {
		return errors.New("Cannot find the participation")
	}


	k.SetCagnotte(ctx, name, cagnotte)
	return nil
}
func (k Keeper) IsActive(ctx sdk.Context, name string) bool {

	cagnotte, _ := k.GetCagnotte(ctx, name)
	return cagnotte.Status
}

func (k Keeper) Allowed(ctx sdk.Context, executedAddr sdk.AccAddress) (bool, error) {
	adminAddr, err := k.GetAdmin(ctx)
	if err != nil {
		return false, err
	}
	return executedAddr.Equals(adminAddr), nil
}

func (k Keeper) AddParticipator(ctx sdk.Context, name string, participator sdk.AccAddress) {

	cagnotte, _ := k.GetCagnotte(ctx, name)
	cagnotte.Participators = append(cagnotte.Participators, participator)
	k.SetCagnotte(ctx, name, cagnotte)
}

// Get owner of the item
func (k Keeper) GetCreator(ctx sdk.Context, key string) sdk.AccAddress {
	cagnotte, _ := k.GetCagnotte(ctx, key)
	return cagnotte.Owner
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.CagnottePrefix))
}

// SetAdmin sets the admin.
func (k Keeper) SetAdmin(ctx sdk.Context, adminAddr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(adminAddr)
	key := []byte(types.AdminPrefix + "adminAccount")
	store.Set(key, bz)
}

func (k Keeper) GetAdmin(ctx sdk.Context) (sdk.AccAddress, error) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.AdminPrefix + "adminAccount")
	var adminAddr sdk.AccAddress
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &adminAddr)
	if err != nil {
		return nil, err
	}

	return adminAddr, nil
}
func getAdmin(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	admin, err := k.GetAdmin(ctx)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func listCagnotte(ctx sdk.Context, k Keeper) ([]types.Cagnotte, error) {
	var whoisList []types.Cagnotte
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.CagnottePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var whois types.Cagnotte
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &whois)
		whoisList = append(whoisList, whois)
	}

	return whoisList, nil
}
func getCagnotte(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	cagnotte, err := k.GetCagnotte(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, cagnotte)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
func listCagnotteByUser(ctx sdk.Context, k Keeper, path []string) (res []byte, sdkError error) {
	data, err := listCagnotte(ctx, k)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	userAddr := path[0]
	listcagnotte := []types.Cagnotte{}
	for _, value := range data {
		if value.Owner.String() == userAddr {
			listcagnotte = append(listcagnotte, value)

		}
	}
	res = codec.MustMarshalJSONIndent(k.cdc, listcagnotte)

	return res, nil
}
