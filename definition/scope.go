package definition

import (
	"github.com/kubex/potens-go/i18n"
	"strings"
)

// AppScope scopes provided by your application
type AppScope struct {
	ID          string
	Name        i18n.Translations `json:",omitempty"`
	Description i18n.Translations `json:",omitempty"`
}

func NewScope(scopeID string) AppScope {
	return AppScope{ID: scopeID}
}

// VendorID Retrieves the vendor ID for this scope, empty for a global scope
func (scope *AppScope) VendorID(appDef *AppDefinition) string {
	if scope.IsBuiltIn() {
		return ""
	}
	scopeSplit := strings.SplitN(scope.ID, "/", 3)
	if len(scopeSplit) == 3 && len(scopeSplit[0]) > 0 {
		return scopeSplit[0]
	}

	return appDef.VendorID
}

// AppID Retrieves the application ID for this scope, empty for a global scope
func (scope *AppScope) AppID(appDef *AppDefinition) string {
	if scope.IsBuiltIn() {
		return ""
	}
	scopeSplit := strings.SplitN(scope.ID, "/", 3)
	if len(scopeSplit) == 3 && len(scopeSplit[1]) > 0 {
		return scopeSplit[1]
	}

	return appDef.AppID
}

// ScopeID Retrieves the Scope ID for this scope
func (scope *AppScope) ScopeID() string {
	scopeSplit := strings.SplitN(scope.ID, "/", 3)
	return scopeSplit[len(scopeSplit)-1]
}

// IsBuiltIn returns true for global scope, e.g. owner
func (scope *AppScope) IsBuiltIn() bool {
	return !strings.Contains(scope.ID, "/")
}

// IsSameVendor returns true if the vendor for the scope matches the vendor in the provided definition
func (scope *AppScope) IsSameVendor(appDef *AppDefinition) bool {
	return scope.VendorID(appDef) == appDef.VendorID
}
