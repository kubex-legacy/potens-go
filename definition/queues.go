package definition

import (
	"github.com/kubex/potens-go/i18n"
)

// AppQueue Queue provided by your app
type AppQueue struct {
	ID    string
	Name  i18n.Translations
	Icon  string
	Path  string
	Roles []AppRole
}
