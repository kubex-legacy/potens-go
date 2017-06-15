package definition

import (
	"github.com/kubex/potens-go/i18n"
)

// AppConfigType - Type of config value
type AppConfigType string

const (
	// AppConfigTypeString String
	AppConfigTypeString AppConfigType = "string"
	// AppConfigTypeInteger Integer
	AppConfigTypeInteger AppConfigType = "integer"
	// AppConfigTypeFloat Float
	AppConfigTypeFloat AppConfigType = "float"
	// AppConfigTypeBoolean Boolean
	AppConfigTypeBoolean AppConfigType = "boolean"
	// AppConfigTypeJSON Json
	AppConfigTypeJSON AppConfigType = "json"
	// AppConfigTypeURI Uri
	AppConfigTypeURI AppConfigType = "uri"
	// AppConfigTypeOptions Options
	AppConfigTypeOptions AppConfigType = "options"
	// AppConfigTypeArrayComma ArrayComma
	AppConfigTypeArrayComma AppConfigType = "array:comma"
	// AppConfigTypeArrayLine ArrayLine
	AppConfigTypeArrayLine AppConfigType = "array:line"
)

// AppConfig Configurable item for your app per organisation
type AppConfig struct {
	ID          string
	Name        i18n.Translations
	Description i18n.Translations
	Note        i18n.Translations
	Help        i18n.Translations
	Type        AppConfigType
	Values      map[string]i18n.Translations
	Default     string
	Required    bool
}
