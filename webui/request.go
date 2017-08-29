package webui

import (
	"github.com/kubex/proto-go/application"
	"net/url"
	"strings"
)

//GetURL convert an application http reqest to a URL
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

//GetParameter Retrieve a parameter from http request
func GetParameter(params map[string]*application.HTTPRequest_HTTPParameter, key string) string {
	if val, exists := params[key]; exists {
		return strings.Join(val.Values, ", ")
	}
	return ""
}
