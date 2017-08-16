package definition

import "github.com/kubex/potens-go/i18n"

type Action struct {
	ID          string
	Title       i18n.Translations
	Path        string
	Icon        string
	Color       AppColor
	Roles       []AppScope
	Permissions []AppScope
}
