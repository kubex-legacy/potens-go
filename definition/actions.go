package definition

import "github.com/kubex/potens-go/i18n"

type AppAction struct {
	ID          string
	Name        i18n.Translations
	Description i18n.Translations
	Inputs      map[string]DataItem
	Outputs     map[string]DataItem
	Endpoint    EndpointDefinition
}
