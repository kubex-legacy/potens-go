package definition

import "github.com/kubex/potens-go/i18n"

type Rpc struct {
	Method string
	Name   i18n.Translations
	Roles  []AppRole
	Input  []DataItem
	Output []DataItem
}
