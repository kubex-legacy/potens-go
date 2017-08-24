package potens

import (
	"github.com/kubex/potens-go/auth"
	"golang.org/x/net/context"
	"github.com/kubex/potens-go/definition"
)

// UserFromContext retrieves user info from given request context
func UserFromContext(ctx context.Context) auth.UserData {
	return auth.FromContext(ctx)
}

//IsPermitted Check a users roles and permissions to see if they should have access
func (app *Application) IsPermitted(user auth.UserData, roles, permissions []definition.AppScope) bool {
	return app.Definition().IsPermitted(user, roles, permissions)
}

func (app *Application) MakeScope(scopeID string) string {
	scope := definition.NewScope(scopeID)
	return scope.GenID(app.Definition())
}
