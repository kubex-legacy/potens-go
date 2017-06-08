package definition

import (
	"github.com/kubex/potens-go/i18n"
)

type AppPermission struct {
	VendorID string `yaml:"vendor"`
	AppID    string `yaml:"app"`
	Rpc      string
	Required bool
	Reason   i18n.Translations
}
