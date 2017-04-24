package definition

import (
	"github.com/kubex/potens-go/i18n"
)

// AppSearchAction Search Action provided by your app
type AppSearchAction struct {
	ID          string
	Name        i18n.Translations
	Description i18n.Translations
	Icon        string
	Path        string
	Roles       []string
	Tokens      []string
}
