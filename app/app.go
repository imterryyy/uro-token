package app

import (
    //"github.com/tendermint/starport/starport/pkg/cosmoscmd"

    "github.com/tendermint/tendermint/libs/log"
    dbm "github.com/tendermint/tm-db"

    "github.com/cosmos/cosmos-sdk/baseapp"
    "github.com/cosmos/cosmos-sdk/codec"
    simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
    "github.com/cosmos/cosmos-sdk/codec/types"

    "github.com/cosmos/cosmos-sdk/types/module"
)

const (
	  AccountAddressPrefix = "uro"
	  Name                 = "uro"
)

var (
    DefaultNodeHome string

    ModuleBasics = module.NewBasicManager(
        // this line is used by starport scaffolding # stargate/app/moduleBasic
	  )
)

// App extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct {
    *baseapp.BaseApp

    cdc               *codec.LegacyAmino
	  appCodec          codec.Codec
	  interfaceRegistry types.InterfaceRegistry

    invCheckPeriod uint
}

func (app *App) Name() string {
    return app.BaseApp.Name()
}

func New(
    logger log.Logger,
    db dbm.DB,
    encodingConfig simappparams.EncodingConfig,
    baseAppOptions ...func(*baseapp.BaseApp),
) *App {
    bApp := baseapp.NewBaseApp(
        Name,
        logger,
        db,
        encodingConfig.TxConfig.TxDecoder(),
        baseAppOptions...,
	  )
    
    app := &App{
        BaseApp: bApp,
    };

    return app;
}

