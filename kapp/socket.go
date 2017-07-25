package kapp

import (
	"github.com/kubex/proto-go/application"
	"github.com/kubex/potens-go/webui"
	"golang.org/x/net/context"
)

func (s *ApplicationServer) HandleSocketAction(ctx context.Context, in *application.SocketRequest) (*application.HTTPResponse, error) {
	resp := webui.CreateJsonResponse("")
	return resp, nil
}
