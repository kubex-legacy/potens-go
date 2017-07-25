package kapp

import (
	"github.com/kubex/proto-go/application"
	"golang.org/x/net/context"
)

type ApplicationServer struct {
	routes      []*AppRoute
	namedRoutes map[string]*AppRoute
}

func New() *ApplicationServer {
	return &ApplicationServer{namedRoutes: make(map[string]*AppRoute), routes: make([]*AppRoute, 0)}
}

func (s *ApplicationServer) NewRoute(path string) *AppRoute {
	r := &AppRoute{parent: s}
	r.Path(path)
	s.routes = append(s.routes, r)
	return r
}
func (s *ApplicationServer) NewPage(path string) *AppPage {
	r := s.NewRoute(path)
	p := &AppPage{
		route: r,
	}
	return p
}

func (s *ApplicationServer) HandleHttp(path string, f func(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error)) *AppRoute {
	r := s.NewRoute(path)
	r.httpHandler = f
	return r
}

func (s *ApplicationServer) HandleSocket(path string, f func(ctx context.Context, in *application.SocketRequest) (*application.HTTPResponse, error)) *AppRoute {
	r := s.NewRoute(path)
	r.socketHandler = f
	return r
}

func (s *ApplicationServer) HandlePage(path string,
	definition func(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error),
	defaultPage func(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error)) *AppPage {
	r := s.NewPage(path)
	r.definition = definition
	r.defaultPage = defaultPage
	return r
}
