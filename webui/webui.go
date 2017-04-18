package webui

import (
	"net/url"

	"github.com/kubex/potens-go/webui/breadcrumb"
	"github.com/kubex/proto-go/application"
)

//CreateResponse creates a new initialised response
func CreateResponse() *application.HTTPResponse {
	response := &application.HTTPResponse{}
	response.Headers = make(map[string]*application.HTTPResponse_HTTPHeaderParameter)
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