package definition

import (
	"github.com/kubex/potens-go/i18n"
)

// QuickActionMode Launch mode for a quick action
type QuickActionMode string

const (
	// QuickActionModePage Redirect to a new page
	QuickActionModePage QuickActionMode = "page"
	// QuickActionModeDialog Open a dialog window
	QuickActionModeDialog QuickActionMode = "dialog"
	// QuickActionModeWindow Open in a new window
	QuickActionModeWindow QuickActionMode = "window"
)

// AppQuickAction Quick Action provided by your app
type AppQuickAction struct {
	ID    string
	Name  i18n.Translations
	Icon  string
	mode  QuickActionMode
	Path  string
	Roles []AppRole
}
