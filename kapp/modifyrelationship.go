package kapp

import (
	"github.com/kubex/proto-go/application"
	"golang.org/x/net/context"
)

func (s *ApplicationServer) ModifyRelationship(ctx context.Context, in *application.ProjectModifyRequest) (*application.ProjectModifyResponse, error) {
	return &application.ProjectModifyResponse{Success: true}, nil
}
