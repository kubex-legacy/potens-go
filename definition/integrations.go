package definition

import (
	"github.com/kubex/potens-go/i18n"
)

type AppIntegrationPanel struct {
	ID   string
	Hook string
	Path string
}

// AppIntegrationMenuItemMode Launch mode for a integration menu item
type AppIntegrationMenuItemMode string

const (
	// AppIntegrationMenuItemModePage Redirect to a new page
	AppIntegrationMenuItemModePage AppIntegrationMenuItemMode = "page"
	// AppIntegrationMenuItemModeIntegrated Open within the content area of the entity page
	AppIntegrationMenuItemModePagelet AppIntegrationMenuItemMode = "pagelet"
)

type AppIntegrationMenuItem struct {
	ID          string
	Hook        string
	Path        string
	Icon        string
	Mode        AppIntegrationMenuItemMode
	Title       i18n.Translations
	Description i18n.Translations
}

// AppIntegrationActionMode Launch mode for a integration action
type AppIntegrationActionMode string

const (
	// AppIntegrationActionModePage Redirect to a new page
	AppIntegrationActionModePage AppIntegrationActionMode = "page"
	// AppIntegrationActionModePagelet Open within the content area of the entity page
	AppIntegrationActionModePagelet AppIntegrationActionMode = "pagelet"
	// AppIntegrationActionModeDialog Open in a dialog
	AppIntegrationActionModeDialog AppIntegrationActionMode = "dialog"
	// AppIntegrationActionModeWindow Open in a new window
	AppIntegrationActionModeWindow AppIntegrationActionMode = "window"
)

type AppIntegrationAction struct {
	ID          string
	Hook        string
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
