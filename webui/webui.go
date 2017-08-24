package webui

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/kubex/potens-go/webui/breadcrumb"
	"github.com/kubex/proto-go/application"
)

//CreateResponse creates a new initialised response
func CreateResponse() *application.HTTPResponse {
	response := &application.HTTPResponse{}
	response.Headers = make(map[string]*application.HTTPResponse_HTTPHeaderParameter)
	return response
}

//CreateJsonResponse Creates a new response
func CreateJsonResponse(content interface{}) *application.HTTPResponse {
	response := CreateResponse()
	response.ContentType = "application/json"
	jsonContent, _ := json.Marshal(content)
	response.Body = string(jsonContent)
	return response
}

//SetBreadcrumb set the breadcrumb on the response
func SetBreadcrumb(response *application.HTTPResponse, breadcrumb breadcrumb.Breadcrumb) {
	response.Headers["x-cube-breadcrumb"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{breadcrumb.Json()}}
}

//SetPageTitle set the page title on the response
func SetPageTitle(response *application.HTTPResponse, PageTitle string) {
	response.Headers["x-cube-title"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{PageTitle}}
}

//SetBackPath set the back path on the response, relative to app route
func SetBackPath(response *application.HTTPResponse, BackPath string) {
	response.Headers["x-cube-back-path"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{BackPath}}
}

//SetPageIcon set the icon url/code on the response
func SetPageIcon(response *application.HTTPResponse, Icon string) {
	response.Headers["x-cube-icon"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{Icon}}
}

//SetPageFID set the FID for the entity being shown on the page
func SetPageFID(response *application.HTTPResponse, FID string) {
	response.Headers["x-cube-page-fid"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{FID}}
}

//SetCacheSeconds set the response to cache for X seconds
func SetCacheSeconds(response *application.HTTPResponse, Seconds int64) {
	response.Headers["x-cache-seconds"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{strconv.FormatInt(Seconds, 10)}}
}

//SetCachePublic set the response to cache in public caches
func SetCachePublic(response *application.HTTPResponse) {
	response.Headers["x-cache-scope"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{"public"}}
}

// PageIntergrationType
type PageIntergrationType string

//Page Intergration Types
const (
	// PageIntergrationTypeDefault Default
	PageIntergrationTypeDefault PageIntergrationType = "default"
	// PageIntergrationTypeNone None
	PageIntergrationTypeNone PageIntergrationType = "none"
)

//SetPageIcon set the icon url/code on the response
func SetPageIntegrations(response *application.HTTPResponse, IntegrationType PageIntergrationType) {
	response.Headers["x-cube-integrations"] = &application.HTTPResponse_HTTPHeaderParameter{Values: []string{string(IntegrationType)}}
}

func GetUrl(request *application.HTTPRequest) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "apps.cubex.io",
		Path:       request.Path,
		RawPath:    request.Path,
		ForceQuery: false,
		RawQuery:   request.QueryString,
	}
}

func GetParameter(params map[string]*application.HTTPRequest_HTTPParameter, key string) string {
	if val, exists := params[key]; exists {
		return strings.Join(val.Values, ", ")
	}
	return ""
}
