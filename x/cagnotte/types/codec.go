package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {

	cdc.RegisterConcrete(MsgAddCagnotte{}, "cagnotte/Addcagnotte", nil)
	cdc.RegisterConcrete(MsgCreateCagnotte{}, "cagnotte/Createcagnotte", nil)
	cdc.RegisterConcrete(MsgCloseCagnotte{}, "cagnotte/Closecagnotte", nil)
	cdc.RegisterConcrete(MsgConfirmPendingTx{}, "cagnotte/ConfirmPendingTx", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
