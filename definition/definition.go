package definition

import (
	"io/ioutil"
	"github.com/kubex/potens-go/i18n"
	"gopkg.in/yaml.v2"
)

// AppDefinition Application Definition
type AppDefinition struct {
	Type          AppType
	UIMode        UIMode `yaml:"ui_mode"`
	Installable   bool
	ConfigVersion float32 `yaml:"config_version"`

	Release  AppRelease
	VendorID string `yaml:"vendor_id"`
	AppID    string `yaml:"app_id"`
	GroupID  string `yaml:"group_id"`
	Category string
	Priority int32

	//Dependencies GAIDs that this app requires
	Dependencies []string

	Name        i18n.Translations
	Description i18n.Translations
	Icon        string
	Color       AppColor

	Navigation []AppNavigation
	Auth       AuthConfig

	Entities []AppEntity
	//EntityVendorKey 3 Character Vendor Key
	EntityVendorKey string `yaml:"entity_vendor_key"`
	//EntityAppKey 2 Chatacter App Key
	EntityAppKey string `yaml:"entity_app_key"`

	QuickActions    []AppQuickAction  `yaml:"quick_actions"`
	SearchActions   []AppSearchAction `yaml:"search_actions"`
	Queues          []AppQueue
	DashboardPanels []DashboardPanel`yaml:"dashboard_panels"`

	AdvancedNotificationsPath string `yaml:"advanced_notifications_path"`
	Notifications             []AppNotification

	Roles              []AppScope
	Permissions        []AppScope
	UtilisePermissions []AppPermission `yaml:"utilise_permissions"`

	Integrations AppIntegrations

	AdvancedConfigPath string `yaml:"advanced_config_path"`
	Config             []AppConfig

	Actions []AppAle
	Lookups []AppAle
	Events  []AppAle

	TrustedVendor bool
}

func New() AppDefinition {
	def := AppDefinition{}
	def.Installable = true
	def.Auth.Type = AuthTypeNoAuth
	return def
}

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
	d.Parse()
	return err
}
