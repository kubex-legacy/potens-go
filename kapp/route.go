package kapp

import (
	"github.com/kubex/proto-go/application"
	"golang.org/x/net/context"
	"strings"
)

type AppRoute struct {
	parent  *ApplicationServer
	page    *AppPage
	handler func(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error)
	path    string
	methods []string
	//Matched params from the url
	params map[string]string
}

func (r *AppRoute) Path(tpl string) *AppRoute {
	r.path = tpl
	return r
}

func (r *AppRoute) Method(methods ...string) *AppRoute {
	r.methods = methods
	return r
}

func (r *AppRoute) Match(in *application.HTTPRequest) bool {
	return r.matchMethod(in.Method) && r.matchPath(in.Path)
}

func (r *AppRoute) matchPath(path string) bool {
	return strings.HasPrefix(path, r.path)
}

func (r *AppRoute) matchMethod(method string) bool {
	if len(r.methods) > 0 {
		for _, meth := range r.methods {
			if strings.EqualFold(meth, method) {
				return true
			}
		}
		return false
	}
	return true
}

func (r *AppRoute) GetParam(key string, failover string) string {
	if val, ok := r.params[key]; ok {
		return val
	}
	return failover
}
