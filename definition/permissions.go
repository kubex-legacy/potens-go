package definition

import "github.com/kubex/potens-go/i18n"

// AppPermission permissions utilised by your application
type AppPermission struct {
	Scope    AppScope
	Required bool
	Reason   i18n.Translations `json:",omitempty"`
}
