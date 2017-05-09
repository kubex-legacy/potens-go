package definition

// EndpointType - Type of API Call
type EndpointType string

const (
	// EndpointTypeGRPC gRPC
	EndpointTypeGRPC EndpointType = "grpc"
	// EndpointTypeRest Rest
	EndpointTypeRest EndpointType = "rest"
)

type EndpointDefinition struct {
	Type            EndpointType
	Endpoint        string
	RequestHeaders  map[string]DataItem
	ResponseHeaders map[string]DataItem
}
