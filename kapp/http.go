package kapp

import (
	"github.com/kubex/proto-go/application"
	"golang.org/x/net/context"
	"errors"
)

// HandleHTTPRequest handles requests from HTTP sources
func (s *ApplicationServer) HandleHTTPRequest(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error) {
	if in.RequestType == application.HTTPRequest_PAGE_DEFINITION {
	} else if in.RequestType == application.HTTPRequest_REQUEST {
	}
	return nil, errors.New("Unsupported request type")
}
