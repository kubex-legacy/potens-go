package main

import (
	"github.com/kubex/potens-go"
)

func quick() {
	app := potens.QuickStartApp("Quick Service", nil, nil)
	app.Log().Info("Started Application")
}

func main() {
	var err error
	app := potens.NewApplication("Example Service")

	err = app.SetIdentity(nil)
	app.FatalErr(err)

	err = app.SetDefinition(nil)
	app.FatalErr(err)
}
