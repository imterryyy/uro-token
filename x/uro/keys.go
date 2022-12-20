package uro

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
