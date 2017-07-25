package kapp

import (
	"github.com/kubex/proto-go/application"
	"golang.org/x/net/context"
)

type ApplicationServer struct {
	routes        []*AppRoute
	namedRoutes   map[string]*AppRoute
	socketActions map[application.SocketAction]func(ctx context.Context, in *application.SocketRequest) (*application.HTTPResponse, error)
	modifyActions map[application.ProjectModifyRequest_ModifyMode]func(ctx context.Context, distributorID, projectID string, data map[string]string) (*application.ProjectModifyResponse, error)
}

func New() *ApplicationServer {
	return &ApplicationServer{namedRoutes: make(map[string]*AppRoute), routes: make([]*AppRoute, 0)}
}

func (s *ApplicationServer) newRoute(path string) *AppRoute {
	r := &AppRoute{parent: s}
	r.Path(path)
	s.routes = append(s.routes, r)
	return r
}
func (s *ApplicationServer) newPage(path string) *AppPage {
	r := s.newRoute(path)
	p := &AppPage{
		route: r,
	}
	r.page = p
	return p
}

func (s *ApplicationServer) HandleHttp(path string, f func(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error)) *AppRoute {
	r := s.newRoute(path)
	r.handler = f
	return r
}

func (s *ApplicationServer) HandlePage(path string,
	definition func(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error),
	defaultPage func(ctx context.Context, in *application.HTTPRequest) (*application.HTTPResponse, error)) *AppPage {
	r := s.newPage(path)
	r.definition = definition
	r.defaultPage = defaultPage
	return r
}
