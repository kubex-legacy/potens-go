package potens

import (
	"github.com/kubex/potens-go/definition"
	"github.com/kubex/potens-go/identity"
)

func QuickStartApp(name string) *Application {
	return StartApp(name, nil, nil)
}

func StartApp(name string, ident *identity.AppIdentity, def *definition.AppDefinition) *Application {
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

func QuickStartService(name string) *Application {
	return StartService(name, nil, nil)
}

func StartService(name string, ident *identity.AppIdentity, def *definition.AppDefinition) *Application {
	app := NewService(name)
	app.Log().Debug("Assuming service name " + app.ServiceKey())
	app.Log().Info("Starting Service " + name)
	app.FatalErr(app.SetIdentity(ident))
	app.Log().Info("Processed Identity")
	app.FatalErr(app.SetDefinition(def))
	app.Log().Info("Processed Definition")
	app.FatalErr(app.CreateServer())
	app.Log().Info("Created gRPC Server")
	return app
}
