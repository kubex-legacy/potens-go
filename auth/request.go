package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	keyprefix = "portc-"

	projectKey  = "project"
	usernameKey = "username"
	userIDKey   = "userid"
	appIDkey    = "appid"
	vendorKey   = "appvendor"
	sigKey      = "signature"

	firstNameKey = "first-name"
	lastNameKey  = "last-name"

	rolesKey       = "roles"
	permissionsKey = "permissions"
)

// GetKeyPrefix returns portcullis key prefix
func GetKeyPrefix() string {
	return keyprefix
}

// GetSignatureKey retrieves the key used for portcullis verification signature
func GetSignatureKey() string {
	return keyprefix + sigKey
}

// GetAppIDKey retrieves the key used for App ID
func GetAppIDKey() string {
	return keyprefix + appIDkey
}

// GetAppVendorKey retrieves the key used for app vendor
func GetAppVendorKey() string {
	return keyprefix + vendorKey
}

// GetProjectKey retrieves the key used for project
func GetProjectKey() string {
	return keyprefix + projectKey
}

// GetUsernameKey retrieves the key used for username
func GetUsernameKey() string {
	return keyprefix + usernameKey
}

// GetUserIDKey retrieves the key used for user ID
func GetUserIDKey() string {
	return keyprefix + userIDKey
}

// GetFirstNameKey retrieves the first name of the user make the request
func GetFirstNameKey() string {
	return keyprefix + firstNameKey
}

// GetLastNameKey retrieves the last name of the user making the request
func GetLastNameKey() string {
	return keyprefix + lastNameKey
}

// GetRolesKey key for retrieving roles from the request
func GetRolesKey() string {
	return keyprefix + rolesKey
}

// GetPermissionsKey key for retrieving permissions from the request
func GetPermissionsKey() string {
	return keyprefix + permissionsKey
}

// GetGenericKeyForString retrieves key for given generic value
func GetGenericKeyForString(in string) string {
	key := strings.Replace(in, " ", "-", -1)
	key = strings.ToLower(key)
	return keyprefix + key
}

// UserData is the structure for deserialised request information
type UserData struct {
	ProjectID   string
	UserID      string
	AppID       string
	VendorID    string
	Username    string
	FirstName   string
	LastName    string
	signature   string
	meta        metadata.MD
	Roles       []string
	Permissions []string
}

// Verify checks that the request signature matches using signature key
func (r *UserData) Verify(sigKey string) bool {
	mac := hmac.New(sha256.New, []byte(sigKey))
	mk := make([]string, len(r.meta))
	i := 0
	for k := range r.meta {
		mk[i] = k
		i++
	}
	sort.Strings(mk)

	m := ""
	for _, v := range mk {
		if strings.HasPrefix(v, GetKeyPrefix()) {
			m = m + v
			b := r.meta[v]
			sort.Strings(b)
			for _, a := range b {
				m = m + a
			}
		}
	}

	mac.Write([]byte(m))
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(r.signature), []byte(expectedMAC))
}

// GlobalAppID is getter for requesting app's Global ID
func (r *UserData) GlobalAppID() string {
	return fmt.Sprintf("%s/%s", r.VendorID, r.AppID)
}

// HasRole check if the user has a specific role
func (r *UserData) HasRole(checkRole string) bool {
	for _, role := range r.Roles {
		if role == r.ProjectID+"|"+checkRole {
			return true
		}
	}
	return false
}

// HasPermission check if the user has a specific permission
func (r *UserData) HasPermission(checkPermission string) bool {
	return *r.HasPermissionStrict(checkPermission, true)
}

func (r *UserData) HasPermissionStrict(checkPermission string, strict bool) *bool {
	var hasPermission bool
	if strict {
		hasPermission = false
	}
	for _, permission := range r.Permissions {
		if permission == r.ProjectID+"|"+checkPermission {
			hasPermission = true
		} else if permission == "!"+r.ProjectID+"|"+checkPermission {
			hasPermission = false
			return &hasPermission
		}
	}
	return &hasPermission
}

// FromContext retrieves user info from given request context
func FromContext(ctx context.Context) UserData {
	md, _ := metadata.FromIncomingContext(ctx)
	res := UserData{
		ProjectID:   safeGetMetaValString(GetProjectKey(), md),
		UserID:      safeGetMetaValString(GetUserIDKey(), md),
		Username:    safeGetMetaValString(GetUsernameKey(), md),
		FirstName:   safeGetMetaValString(GetFirstNameKey(), md),
		LastName:    safeGetMetaValString(GetLastNameKey(), md),
		AppID:       safeGetMetaValString(GetAppIDKey(), md),
		VendorID:    safeGetMetaValString(GetAppVendorKey(), md),
		signature:   safeGetMetaValString(GetSignatureKey(), md),
		Roles:       safeGetMetaValStringSlice(GetRolesKey(), md),
		Permissions: safeGetMetaValStringSlice(GetPermissionsKey(), md),
		meta:        md,
	}
	return res
}

func safeGetMetaValString(key string, md metadata.MD) string {
	result := ""
	if md != nil {
		if len(md[key]) != 0 {
			result = md[key][0]
		}
	}
	return result
}

func safeGetMetaValStringSlice(key string, md metadata.MD) []string {
	result := []string{}
	if md != nil {
		if sliceKeys, hasKey := md[key]; hasKey {
			for _, sliceValue := range sliceKeys {
				result = append(result, sliceValue)
			}
		}
	}
	return result
}
