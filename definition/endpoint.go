package definition

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
	//EndpointRequestTypeImpart data pushed in via impart service (for events only)
	EndpointRequestTypeImpart EndpointRequestType = "impart"
)

type EndpointDefinition struct {
	Type            EndpointType
	Endpoint        string
	RequestType     EndpointRequestType
	RequestHeaders  map[string]DataItem
	ResponseHeaders map[string]DataItem
}
