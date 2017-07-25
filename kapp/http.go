package kapp

import (
	"github.com/kubex/proto-go/application"
	"golang.org/x/net/context"
	"errors"
)

// HandleHTTPRequest handles requests from HTTP sources
func (s *ApplicationServer) HandleHTTPRequest(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error) {
	for _, route := range s.routes {
		if in.Path == route.path && route.page != nil {
			if in.RequestType == application.HTTPRequest_PAGE_DEFINITION {
				if route.page.definition != nil {
					return route.page.definition(ctx, in)
				}
			} else if in.RequestType == application.HTTPRequest_REQUEST {
				return route.page.defaultPage(ctx, in)
			}
		} else if route.handler != nil && route.Match(in) {
			return route.handler(ctx, in)
		}
	}
	return nil, errors.New("Unsupported request type")
}
