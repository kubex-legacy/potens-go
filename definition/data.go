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
)

type DataItem struct {
	ID          string
	Type        DataType
	Required    bool
	Name        i18n.Translations
	Description i18n.Translations
	Options     map[string]string
	Default     string
}
