package definition

import (
	"io/ioutil"
	"github.com/kubex/potens-go/i18n"
	yaml "gopkg.in/yaml.v2"
)

// AppDefinition Application Definition
type AppDefinition struct {
	Type                      AppType
	UIMode                    UIMode `yaml:"ui_mode"`
	Installable               bool
	Auth                      AuthConfig
	ConfigVersion             float32 `yaml:"config_version"`
	Version                   float32
	VendorID                  string `yaml:"vendor_id"`
	TrustedVendor             bool
	AppID                     string `yaml:"app_id"`
	GroupID                   string `yaml:"group_id"`
	Category                  string
	Priority                  int32
	Name                      i18n.Translations
	Description               i18n.Translations
	Icon                      string
	Color                     AppColor
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
	Rpcs                      []Rpc
	Permissions               []AppPermission

	Actions []AppAle
	Lookups []AppAle
	Events  []AppAle

	Integrations AppIntegrations
	//Dependencies GAIDs that this app requires
	Dependencies []string
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
	return err
}
