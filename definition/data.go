package definition

import "github.com/kubex/potens-go/i18n"

type DataType string

const (
	DataTypeString    DataType = "string"
	DataTypeInteger   DataType = "integer"
	DataTypeLong      DataType = "long"
	DataTypeFloat     DataType = "float"
	DataTypeDouble    DataType = "double"
	DataTypeBoolean   DataType = "boolean"
	DataTypeBytes     DataType = "bytes"
	DataTypeDate      DataType = "date"
	DataTypeDateTime  DataType = "datetime"
	DataTypeTimestamp DataType = "timestamp"
	DataTypeList      DataType = "list"
	DataTypePassword  DataType = "password"
	DataTypeJson      DataType = "json"
)

type DataLocation string

const (
	DataLocationHeader      DataLocation = "header"
	DataLocationQueryString DataLocation = "querystring"
	DataLocationBody        DataLocation = "body"
	DataLocationPost        DataLocation = "post"
	DataLocationPath        DataLocation = "path"
)

type DataItem struct {
	//ID common computer friendly name e.g. api_token
	ID string
	//Key when transferring data, e.g. in Headers, you may want to set a specific key e.g. X-Auth-Token
	Key         string
	Type        DataType
	Required    bool
	Name        i18n.Translations
	Description i18n.Translations
	Options     map[string]string
	Default     string
	Location    DataLocation
	//LocationPattern e.g. /users/{userid}/members/.* (path) | user.company.name (body)
	LocationPattern string
}
