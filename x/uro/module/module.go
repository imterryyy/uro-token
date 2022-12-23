package module

import (
  "github.com/bui-duc-huy/uro-token/x/uro"
  "github.com/cosmos/cosmos-sdk/codec"
  "github.com/bui-duc-huy/uro-token/x/uro/types"
  "github.com/bui-duc-huy/uro-token/x/uro/keeper"
)

type AppModuleBasic struct {
  cdc codec.BinaryCodec
}


func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
  return AppModuleBasic{cdc: cdc}
}

func (AppModuleBasic) Name() string {
  return types.ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
  uro.RegisterCodec(cdc)
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
  uro.RegisterLegacyAminoCodec(cdc);
}

// RegisterInterfaces registers the module's interface types
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
}

// ValidateGenesis performs genesis state validation for the capability module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
}

// RegisterRESTRoutes registers the capability module's REST service handlers.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	// this line is used by starport scaffolding # 2
}

// GetTxCmd returns the capability module's root tx command.
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
}

// GetQueryCmd returns the capability module's root query command.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
}

/// app module
type AppModule struct {
  AppModuleBasic

  keeper keeper.Keeper
}

func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
  return AppModule {
    AppModuleBasic: NewAppModuleBasic(cdc),
    keeper: keeper
  }
}

func (am AppModule) Name() string {
	return am.AppModuleBasic.Name()
}
