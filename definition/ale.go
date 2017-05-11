package definition

import "github.com/kubex/potens-go/i18n"

type AppAleType string

const (
	//AppAleTypeAction Actions that can be performed against an endpoint
	AppAleTypeAction AppAleType = "action"
	//AppAleTypeLookup Data Retrieve from an endpoint
	AppAleTypeLookup AppAleType = "lookup"
	//AppAleTypeEvent Event Triggered
	AppAleTypeEvent AppAleType = "event"
)

type AppAle struct {
	ID          string
	Type        AppAleType
	Name        i18n.Translations
	Description i18n.Translations
	Inputs      map[string]DataItem
	Outputs     map[string]DataItem
	Endpoint    EndpointDefinition
}
