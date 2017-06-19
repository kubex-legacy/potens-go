package definition

import "github.com/kubex/potens-go/i18n"

type DashboardPanel struct {
	ID          string
	Title       i18n.Translations
	Description i18n.Translations
	Path        string
	Icon        string
	Roles       []AppRole
	Actions     []DashboardPanelAction
}

type DashboardPanelAction struct {
	ID    string
	Title i18n.Translations
	Path  string
	Icon  string
	Roles []AppRole
}
