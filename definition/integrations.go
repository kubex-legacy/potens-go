package definition

import (
	"github.com/kubex/potens-go/i18n"
	"strings"
)

type Hook struct {
	VendorID string `yaml:"vendor"`
	AppID    string `yaml:"app"`
	Path     string
}

//ToString Convert the hook to a full path
func (h *Hook) ToString() string {
	return h.VendorID + "/" + h.AppID + "/" + strings.Trim(h.Path, "/")
}

type AppIntegrationPanel struct {
	ID          string
	Hook        Hook
	Path        string
	Roles       []AppScope
	Permissions []AppScope
}

type AppIntegrationMenuItem struct {
	ID          string
	Hook        Hook
	Path        string
	Icon        string
	Title       i18n.Translations
	Description i18n.Translations
	Roles       []AppScope
	Permissions []AppScope
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
	Roles       []AppScope
	Permissions []AppScope
}

type AppIntegrations struct {
	Panels          []AppIntegrationPanel
	HeaderMenuItems []AppIntegrationMenuItem `yaml:"header_menu_items"`
	PageMenuItems   []AppIntegrationMenuItem `yaml:"page_menu_items"`
	HeaderActions   []AppIntegrationAction   `yaml:"header_actions"`
	PageActions     []AppIntegrationAction   `yaml:"page_actions"`
}
