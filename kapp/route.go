package kapp

import (
	"github.com/kubex/proto-go/application"
	"golang.org/x/net/context"
)

type AppRoute struct {
	parent        *ApplicationServer
	httpHandler   func(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error)
	socketHandler func(ctx context.Context, in *application.SocketRequest) (*application.HTTPResponse, error)
	path          string
}

func (r *AppRoute) Path(tpl string) *AppRoute {
	r.path = tpl
	return r
}
