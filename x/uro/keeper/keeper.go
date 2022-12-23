package keeper

import (
  paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
  "github.com/cosmos/cosmos-sdk/codec"
  sdk "github.com/cosmos/cosmos-sdk/types"
  "github.com/bui-duc-huy/uro-token/x/uro"
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

func NewKeeper(storeKey sdk.StoreKey, memKey sdk.StoreKey, paramSpace sdk.Subspace, cdc codec.BinaryCodec) *Keeper {
  if !paramSpace.HasKeyTable() {
    paramSpace = paramSpace.WithKeyTable(types)
  }

  return &Keeper{
    storeKey: storeKey,
    memKey: memKey,
    paramSpace: paramSpace,
    cdc: cdc
  }
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
  return ctx.Logger().With("module", fmt.Sprintf("x/%s", ModuleName))
}

func (k Keeper) GetMessage(ctx sdk.Context, owner sdk.AccAddress) uro.Message {
  store = ctx.KVStore(k.storeKey)
  bz := store.Get(MessageKey(owner))

  if bz == nil {
  }
}
