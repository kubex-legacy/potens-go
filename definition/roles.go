package definition

import (
	"strings"

	"github.com/kubex/potens-go/i18n"
)

// AppRole Roles provided by your application
type AppRole struct {
	ID          string
	Name        i18n.Translations
	Description i18n.Translations
}

// VendorID Retrieves the vendor ID for this role, empty for a global role
func (role *AppRole) VendorID(appDef *AppDefinition) string {
	if role.IsBuiltIn() {
		return ""
	}
	roleSplit := strings.SplitN(role.ID, "/", 3)
	if len(roleSplit) == 3 && len(roleSplit[0]) > 0 {
		return roleSplit[0]
	}

	return appDef.VendorID
}

// AppID Retrieves the application ID for this role, empty for a global role
func (role *AppRole) AppID(appDef *AppDefinition) string {
	if role.IsBuiltIn() {
		return ""
	}
	roleSplit := strings.SplitN(role.ID, "/", 3)
	if len(roleSplit) == 3 && len(roleSplit[1]) > 0 {
		return roleSplit[1]
	}

	return appDef.AppID
}

// IsBuiltIn returns true for global roles, e.g. owner
func (role *AppRole) IsBuiltIn() bool {
	return !strings.Contains(role.ID, "/")
}

// IsSameVendor returns true if the vendor for the role matches the vendor in the provided definition
func (role *AppRole) IsSameVendor(appDef *AppDefinition) bool {
	return role.VendorID(appDef) == appDef.VendorID
}
