package definition

import "github.com/kubex/potens-go/i18n"

type AppEventType string

const (
	AppEventTypeWebhook AppEventType = "webhook"
	AppEventTypePoll    AppEventType = "poll"
	AppEventTypeAPI     AppEventType = "api"
)

type AppEvent struct {
	ID          string
	Name        i18n.Translations
	Description i18n.Translations
	CreateType  AppEventType
	Properties  map[string]DataItem
}
