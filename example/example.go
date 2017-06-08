package main

import (
	"github.com/kubex/potens-go"
)

func quick() {
	app := potens.QuickStartApp("Quick Service", nil, nil)
	run(app)
}

func main() {
	app := potens.NewApplication("Example Service")

	app.FatalErr(app.SetIdentity(nil))
	app.FatalErr(app.SetDefinition(nil))
	app.FatalErr(app.CreateServer())
	run(app)
}

func run(app *potens.Application) {
	app.Log().Info("Ready to serve")
	app.FatalErr(app.RegisterWithDiscovery("host", "50051"))
	app.FatalErr(app.DiscoveryOnline())
	err := app.Serve()
	app.FatalErr(app.DiscoveryOffline())
	app.FatalErr(err)
}
