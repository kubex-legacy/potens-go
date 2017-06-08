package potens

import (
	"github.com/kubex/potens-go/definition"
	"github.com/kubex/potens-go/identity"
)

func QuickStartApp(name string, ident *identity.AppIdentity, def *definition.AppDefinition) (*Application) {
	app := NewApplication(name)
	app.FatalErr(app.SetIdentity(ident))
	app.FatalErr(app.SetDefinition(def))
	app.FatalErr(app.CreateServer())
	return app
}
