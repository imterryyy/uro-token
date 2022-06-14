package app

import (
    "github.com/cosmos/cosmos-sdk/baseapp"
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

type App struct {
    *baseapp.BaseApp
}

func (app *App) Name() string {
    return app.BaseApp.Name()
}

func New() cosmoscmd.App {

}
