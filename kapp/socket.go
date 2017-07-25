package kapp

import (
	"github.com/kubex/proto-go/application"
	"github.com/kubex/potens-go/webui"
	"golang.org/x/net/context"
)

func (s *ApplicationServer) HandleSocketAction(ctx context.Context, in *application.SocketRequest) (*application.HTTPResponse, error) {
	if method, ok := s.socketActions[in.Action]; ok {
		return method(ctx, in)
	}
	resp := webui.CreateJsonResponse("")
	return resp, nil
}

func (s *ApplicationServer) HandleSocketSubscribe(f func(ctx context.Context, in *application.SocketRequest) (*application.HTTPResponse, error)) {
	s.socketActions[application.SocketAction_SUBSCRIBE] = f
}

func (s *ApplicationServer) HandleSocketUnsubscribe(f func(ctx context.Context, in *application.SocketRequest) (*application.HTTPResponse, error)) {
	s.socketActions[application.SocketAction_UNSUBSCRIBE] = f
}
