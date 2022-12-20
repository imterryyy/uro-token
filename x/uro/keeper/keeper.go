package keeper

import (
  paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
  "github.com/cosmos/cosmos-sdk/codec"
  sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
  // external keeper
  // store key
  storeKey sdk.StoreKey
  memKey sdk.StoreKey
  paramSpace paramTypes.Subspace

  // codec
  cdc codec.BinaryCodec
}
