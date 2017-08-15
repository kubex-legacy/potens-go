package definition

import (
	"github.com/kubex/potens-go/i18n"
)

// AppNavigation Application Navigation ITem
type AppNavigation struct {
	ID          string
	Name        i18n.Translations
	Description i18n.Translations
	Icon        string
	Path        string
	Roles       []AppScope
	Permissions []AppScope
}
