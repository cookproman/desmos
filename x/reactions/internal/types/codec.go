package types

import "github.com/cosmos/cosmos-sdk/codec"

// ModuleCdc is the codec
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAddPostReaction{}, "desmos/MsgAddPostReaction", nil)
	cdc.RegisterConcrete(MsgRemovePostReaction{}, "desmos/MsgRemovePostReaction", nil)
}
