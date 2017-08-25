package definition

import (
	"github.com/kubex/potens-go/i18n"
	"strings"
	"github.com/kubex/potens-go/auth"
)

// ScopeType Scope Type
type ScopeType string

//Scope Types
const (
	//ScopeTypeRole Role
	ScopeTypeRole ScopeType = "role"
	//ScopeTypePermission Permission
	ScopeTypePermission ScopeType = "perm"
)

// AppScope scopes provided by your application
type AppScope struct {
	ID          string
	Type        ScopeType
	Name        i18n.Translations `json:",omitempty"`
	Description i18n.Translations `json:",omitempty"`
}

//NewScope Creates a new untyped scope
func NewScope(scopeID string) AppScope {
	return AppScope{ID: scopeID}
}

//NewPermission Creates a new permission scope
func NewPermission(scopeID string) AppScope {
	return AppScope{ID: scopeID, Type: ScopeTypePermission}
}

//NewRole Creates a new role scope
func NewRole(scopeID string) AppScope {
	return AppScope{ID: scopeID, Type: ScopeTypeRole}
}

//TypedScopes Combine scopes with types set
func TypedScopes(roles, permissions []AppScope) []AppScope {
	scopes := []AppScope{}
	for _, role := range roles {
		role.Type = ScopeTypeRole
		scopes = append(scopes, role)
	}
	for _, perm := range permissions {
		perm.Type = ScopeTypePermission
		scopes = append(scopes, perm)
	}
	return scopes
}

// VendorID Retrieves the vendor ID for this scope, empty for a global scope
func (scope *AppScope) VendorID(appDef *AppDefinition) string {
	if scope.IsBuiltIn() {
		return ""
	}
	scopeSplit := strings.SplitN(scope.ID, "/", 4)
	if len(scopeSplit) == 4 && len(scopeSplit[1]) > 0 {
		return scopeSplit[1]
	}

	return appDef.VendorID
}

// AppID Retrieves the application ID for this scope, empty for a global scope
func (scope *AppScope) AppID(appDef *AppDefinition) string {
	if scope.IsBuiltIn() {
		return ""
	}
	scopeSplit := strings.SplitN(scope.ID, "/", 4)
	if len(scopeSplit) == 4 && len(scopeSplit[2]) > 0 {
		return scopeSplit[2]
	}

	return appDef.AppID
}

// ProjectID Retrieves the project ID for this scope
func (scope *AppScope) ProjectID() string {
	scopeSplit := strings.SplitN(scope.ID, "/", 2)
	return scopeSplit[0]
}

// ScopeKey Retrieves the Scope ID for this scope
func (scope *AppScope) ScopeKey() string {
	scopeSplit := strings.SplitN(scope.ID, "/", 4)
	return scopeSplit[len(scopeSplit)-1]
}

// IsBuiltIn returns true for global scope, e.g. owner
func (scope *AppScope) IsBuiltIn() bool {
	return strings.Count(scope.ID, "/") == 1
}

// IsSameVendor returns true if the vendor for the scope matches the vendor in the provided definition
func (scope *AppScope) IsSameVendor(appDef *AppDefinition) bool {
	return scope.VendorID(appDef) == appDef.VendorID
}

func (scope *AppScope) GenID(appDef *AppDefinition) string {
	return strings.Join([]string{scope.VendorID(appDef), scope.AppID(appDef), scope.ScopeKey()}, "/")
}

//MakeScopeID creates a scope for the current vendor/app e.g. vendor/app/id
func (d *AppDefinition) MakeScopeID(ID string) string {
	return d.VendorID + "/" + d.AppID + "/" + ID
}

func (d *AppDefinition) IsPermitted(user auth.UserData, scopes ...AppScope) bool {
	//Do roles or permissions means all users
	if len(scopes) == 0 {
		return true
	}

	//Project Owners can do anything
	if user.HasRole(string(auth.RoleProjectOwner)) {
		return true
	}

	//Check Permissions
	for _, perm := range scopes {
		if perm.Type == ScopeTypePermission {
			hasPerm := user.HasPermissionStrict(perm.GenID(d), false)
			if hasPerm != nil {
				return *hasPerm
			}
		}
	}

	//Check Roles
	for _, role := range scopes {
		if role.Type == ScopeTypeRole && user.HasRole(role.GenID(d)) {
			return true
		}
	}

	return false
}
