package definition

import (
	"io/ioutil"
	"github.com/kubex/potens-go/i18n"
	yaml "gopkg.in/yaml.v2"
)

// AppDefinition Application Definition
type AppDefinition struct {
	Type                      AppType
	UIMode                    UIMode
	ConfigVersion             float32 `yaml:"config_version"`
	Version                   float32
	VendorID                  string
	TrustedVendor             bool
	AppID                     string `yaml:"app_id"`
	GroupID                   string `yaml:"group_id"`
	Category                  string
	Priority                  int32
	Name                      i18n.Translations
	Description               i18n.Translations
	Icon                      string
	AdvancedNotificationsPath string `yaml:"advanced_notifications_path"`
	AdvancedConfigPath        string `yaml:"advanced_config_path"`
	Navigation                []AppNavigation
	Entities                  []AppEntity
	QuickActions              []AppQuickAction  `yaml:"quick_actions"`
	SearchActions             []AppSearchAction `yaml:"search_actions"`
	Queues                    []AppQueue
	Notifications             []AppNotification
	Roles                     []AppRole
	Config                    []AppConfig
	Integrations              AppIntegrations
	//Dependencies GAIDs that this app requires
	Dependencies []string
}

// AppType Application Type
type AppType string

//App Types
const (
	// AppTypePlatformProjectApplication Standard Application
	AppTypePlatformProjectApplication AppType = "kubex.platform.application.project"

	// AppTypePlatformUserApplication Standard User Application (Sidebar)
	AppTypePlatformUserApplication AppType = "kubex.platform.application.user"
)

// UIMode UI Mode Provided By The App
type UIMode string

//UI Modes
const (
	//UIModeNone - No UI Provided
	UIModeNone UIMode = "none"
	//UIModeIntegration - Integrated into pages
	UIModeIntegration UIMode = "integration"
	//UIModeFull - Full UI Application
	UIModeFull UIMode = "full"
)

// FromConfig Populates your definition based on your app-definition.dist.yaml
func (d *AppDefinition) FromConfig(yamlFile string) error {
	yamlContent, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return err
	}

	return d.FromYamlString(string(yamlContent))
}

//GlobalAppID returns the global app ID for the definition
func (d *AppDefinition) GlobalAppID() string {
	return d.VendorID + "/" + d.AppID
}

func (d *AppDefinition) FromYamlString(yamlContent string) error {
	err := yaml.Unmarshal([]byte(yamlContent), d)
	return err
}
