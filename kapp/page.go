package kapp

import (
	"github.com/kubex/proto-go/application"
	"golang.org/x/net/context"
)

type AppPage struct {
	definition  func(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error)
	defaultPage func(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error)
	pageRoutes  []*AppRoute
	route       *AppRoute
}

func (p *AppPage) NewRoute(path string) *AppRoute {
	return p.route.parent.newRoute(path)
}
func (p *AppPage) SubPage(tpl string, f func(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error)) *AppPage {
	spRoute := p.NewRoute(tpl)
	spRoute.handler = f
	return p
}
