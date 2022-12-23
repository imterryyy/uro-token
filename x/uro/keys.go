package uro

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  "github.com/cosmos/cosmos-sdk/types/address"
)

const (
  ModuleName = "uro"

  StoreKey = ModuleName

  RouterKey = ModuleName

  QuerierRoute = ModuleName

  MemStoreKey = "mem_" + ModuleName
)

func KeyPrefix(p string) []byte {
  return []byte(p)
}

func MessageKey(owner sdk.AccAddress) []byte {
  return address.MustLengthPrefix(owner.Bytes())...
}
