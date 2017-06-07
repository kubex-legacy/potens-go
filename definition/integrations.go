package definition

import (
	"github.com/kubex/potens-go/i18n"
)

type Hook struct {
	VendorID string `yaml:"vendor"`
	AppID    string `yaml:"app"`
	Path     string
}

type AppIntegrationPanel struct {
	ID   string
	Hook Hook
	Path string
}

type AppIntegrationMenuItem struct {
	ID          string
	Hook        Hook
	Path        string
	Icon        string
	Title       i18n.Translations
	Description i18n.Translations
}

// AppIntegrationActionMode Launch mode for a integration action
type AppIntegrationActionMode string

const (
	// AppIntegrationActionModePage Redirect to a new page
	AppIntegrationActionModePage AppIntegrationActionMode = "page"
	// AppIntegrationActionModeDialog Open in a dialog
	AppIntegrationActionModeDialog AppIntegrationActionMode = "dialog"
	// AppIntegrationActionModeWindow Open in a new window
	AppIntegrationActionModeWindow AppIntegrationActionMode = "window"
)

type AppIntegrationAction struct {
	ID          string
	Hook        Hook
	Path        string
	Icon        string
	Mode        AppIntegrationActionMode
	Title       i18n.Translations
	Description i18n.Translations
}

type AppIntegrations struct {
	Panels          []AppIntegrationPanel
	HeaderMenuItems []AppIntegrationMenuItem `yaml:"header_menu_items"`
	PageMenuItems   []AppIntegrationMenuItem `yaml:"page_menu_items"`
	HeaderActions   []AppIntegrationAction   `yaml:"header_actions"`
	PageActions     []AppIntegrationAction   `yaml:"page_actions"`
}
