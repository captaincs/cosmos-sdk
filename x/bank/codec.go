package bank

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSend{}, "cosmos-sdk/MsgSend", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "cosmos-sdk/MsgMultiSend", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
