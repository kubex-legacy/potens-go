package definition

// AuthType - Type of Auth for the app
type AuthType string

const (
	// AuthTypeNoAuth No Auth - Webhook project hook url
	AuthTypeNoAuth AuthType = "none"
	//AuthTypeBasic username/password auth
	AuthTypeBasic AuthType = "basic"
	//AuthTypeOAuth Oauth Flow
	AuthTypeOAuth AuthType = "oauth"
	//AuthTypeApiKey API Key
	AuthTypeApiKey AuthType = "apiKey"
)

type AuthConfig struct {
	Type       AuthType
	Parameters []DataItem `json:",omitempty"`
}
