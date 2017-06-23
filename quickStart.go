package potens

import (
	"github.com/kubex/potens-go/definition"
	"github.com/kubex/potens-go/identity"
)

func QuickStartApp(name string, ident *identity.AppIdentity, def *definition.AppDefinition) (*Application) {
	app := NewApplication(name)
	app.Log().Info("Starting Application " + name)
	app.FatalErr(app.SetIdentity(ident))
	app.Log().Info("Processed Identity")
	app.FatalErr(app.SetDefinition(def))
	app.Log().Info("Processed Definition")
	app.FatalErr(app.CreateServer())
	app.Log().Info("Created gRPC Server")
	return app
}
