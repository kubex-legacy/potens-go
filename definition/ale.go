package definition

import "github.com/kubex/potens-go/i18n"

type AppAle struct {
	ID          string
	Name        i18n.Translations
	Description i18n.Translations
	Inputs      map[string]DataItem
	Outputs     map[string]DataItem
	Endpoint    EndpointDefinition
}

// EndpointType - Type of API Call
type EndpointType string

const (
	// EndpointTypeGRPC gRPC
	EndpointTypeGRPC EndpointType = "grpc"
	// EndpointTypeRest Rest
	EndpointTypeRest EndpointType = "rest"
)

type EndpointRequestType string

const (
	//EndpointRequestTypeCall default for Lookups & Actions
	EndpointRequestTypeCall EndpointRequestType = "call"
	//EndpointRequestTypeWebhook data pushed in by webhook (for events only)
	EndpointRequestTypeWebhook EndpointRequestType = "webhook"
	//EndpointRequestTypePoll poll the endpoint for changes (for events only)
	EndpointRequestTypePoll EndpointRequestType = "poll"
	//EndpointRequestTypePipe data pushed in via datapipe service (for events only)
	EndpointRequestTypePipe EndpointRequestType = "pipe"
)

type EndpointDefinition struct {
	Type            EndpointType
	Endpoint        string
	RequestType     EndpointRequestType
	RequestHeaders  map[string]DataItem
	ResponseHeaders map[string]DataItem
}
