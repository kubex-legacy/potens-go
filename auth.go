package potens

import (
	"github.com/kubex/potens-go/auth"
	"golang.org/x/net/context"
)

// UserFromContext retrieves user info from given request context
func UserFromContext(ctx context.Context) auth.UserData {
	return auth.FromContext(ctx)
}
