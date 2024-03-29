package uro

import (
  types "github.com/cosmos/cosmos-sdk/codec/types"
  sdk "github.com/cosmos/cosmos-sdk/types"
  "github.com/cosmos/cosmos-sdk/types/msgservice"
  "github.com/cosmos/cosmos-sdk/codec/legacy"
)

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
  registry.RegisterImplementations((*sdk.Msg)(nil),
    &MsgSend{},
  )
  msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterCodec(cdc *cdc.LegacyAmino) {
}

func RegisterLegacyAminoCodec(cdc *cdc.LegacyAmino) {
}
