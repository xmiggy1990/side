package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitBlockHeaders{}, "btcbridge/MsgSubmitBlockHeaders", nil)
	cdc.RegisterConcrete(&MsgSubmitDepositTransaction{}, "btcbridge/MsgSubmitDepositTransaction", nil)
	cdc.RegisterConcrete(&MsgSubmitWithdrawTransaction{}, "btcbridge/MsgSubmitWithdrawTransaction", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "btcbridge/MsgUpdateParams", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSubmitBlockHeaders{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSubmitDepositTransaction{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSubmitWithdrawTransaction{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgUpdateParams{})
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
