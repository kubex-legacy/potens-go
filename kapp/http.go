package kapp

import (
	"github.com/kubex/proto-go/application"
	"golang.org/x/net/context"
	"errors"
)

// HandleHTTPRequest handles requests from HTTP sources
func (s *ApplicationServer) HandleHTTPRequest(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error) {
	for _, route := range s.routes {
		if route.Match(in) {
			if in.RequestType == application.HTTPRequest_PAGE_DEFINITION && route.page != nil {
				if route.page.definition != nil {
					return route.page.definition(ctx, in)
				}
			} else if in.RequestType == application.HTTPRequest_REQUEST {
				if route.page == nil {
					return route.handler(ctx, in)
				} else {
					return route.page.defaultPage(ctx, in)
				}
			}
		}
	}
	return nil, errors.New("Unsupported request type")
}
