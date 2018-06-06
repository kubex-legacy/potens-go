package definition_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/kubex/potens-go/definition"
	"github.com/kubex/potens-go/i18n"
)

func TestReadYaml(t *testing.T) {
	def := definition.AppDefinition{}
	err := def.FromConfig("app-definition.dist.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if def.Type != definition.AppTypeProject {
		t.Error("Incorrect application type")
	}

	if def.UIMode != definition.UIModeFull {
		t.Error("Incorrect UI mode")
	}

	if !def.Installable {
		t.Error("Incorrect installable flag")
	}

	if def.ConfigVersion != 1.0 {
		t.Error("Incorrect config version")
	}

	if def.VendorID != "vendor-id" {
		t.Error("Incorrect vendor ID")
	}

	if def.AppID != "app-id" {
		t.Error("Incorrect App ID")
	}

	if def.GroupID != "" {
		t.Error("Incorrect Group ID")
	}

	if def.Priority != 500 {
		t.Error("Incorrect Priority")
	}

	if i18n.NewTranslatable(def.Name).Get("fr") != "Les clients" {
		t.Error("Failed to read translation")
	}

	if i18n.NewTranslatable(def.Name).Get("en") != "Customers" {
		t.Error("Failed to read translation")
	}

	if i18n.NewTranslatable(def.Name).Get("eeewf") != "Customers" {
		t.Error("Failed to read default")
	}

	if def.Icon != "social:group" {
		t.Error("Incorrect Icon")
	}

	if def.Color != definition.AppColorPurple {
		t.Error("Incorrect Color")
	}

	if len(def.Navigation) != 1 {
		t.Error("Incorrect Navigation Items")
	}

	if def.Navigation[0].ID != "list-customers" {
		t.Error("Incorrect Navigation ID property")
	}

	if i18n.NewTranslatable(def.Navigation[0].Name).Get("en") != "List Customers" {
		t.Error("Incorrect Navigation Name property")
	}

	if i18n.NewTranslatable(def.Navigation[0].Description).Get("fr") != "Montrer tous vos clients" {
		t.Error("Incorrect Navigation Description property")
	}

	if def.Navigation[0].Icon != "social:group" {
		t.Error("Incorrect Navigation Icon property")
	}

	if def.Navigation[0].Path != "customers/" {
		t.Error("Incorrect Navigation Path property")
	}

	if len(def.Navigation[0].Roles) != 2 {
		t.Error("Incorrect Navigation Roles count")
	}
	if len(def.Navigation[0].Permissions) != 1 {
		t.Error("Incorrect Navigation Permissions count")
	}

	jsonBytes, err := json.Marshal(def)
	log.Println(string(jsonBytes))
}
