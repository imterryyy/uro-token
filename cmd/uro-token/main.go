package main

import (
    "github.com/tendermint/starport/starport/pkg/cosmoscmd"
    "github.com/bui-duc-huy/uro-token/app"
)

func main() {
    rootCmd, _ := cosmoscmd.NewRootCmd(
        app.Name,
        app.AccountAddressPrefix,
        app.DefaultNodeHome,
        app.Name,
        app.ModuleBasics,
        app.New,
        // this line is used by starport scaffolding # root/arguments
        // cmdOptions...,
	  )
}
