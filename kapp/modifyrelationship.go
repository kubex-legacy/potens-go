package kapp

import (
	"github.com/kubex/proto-go/application"
	"golang.org/x/net/context"
)

func (s *ApplicationServer) ModifyRelationship(ctx context.Context, in *application.ProjectModifyRequest) (*application.ProjectModifyResponse, error) {

	if method, ok := s.modifyActions[in.Mode]; ok {
		return method(ctx, in.DistributorId, in.ProjectId, in.Data)
	}

	return &application.ProjectModifyResponse{Success: true}, nil
}

func (s *ApplicationServer) OnInstall(f func(ctx context.Context, distributorID, projectID string, data map[string]string) (*application.ProjectModifyResponse, error)) {
	s.modifyActions[application.ProjectModifyRequest_INSTALL] = f
}

func (s *ApplicationServer) OnUninstall(f func(ctx context.Context, distributorID, projectID string, data map[string]string) (*application.ProjectModifyResponse, error)) {
	s.modifyActions[application.ProjectModifyRequest_UNINSTALL] = f
}

func (s *ApplicationServer) OnSuspend(f func(ctx context.Context, distributorID, projectID string, data map[string]string) (*application.ProjectModifyResponse, error)) {
	s.modifyActions[application.ProjectModifyRequest_SUSPEND] = f
}

func (s *ApplicationServer) OnReactivate(f func(ctx context.Context, distributorID, projectID string, data map[string]string) (*application.ProjectModifyResponse, error)) {
	s.modifyActions[application.ProjectModifyRequest_REACTIVATE] = f
}

func (s *ApplicationServer) OnCancel(f func(ctx context.Context, distributorID, projectID string, data map[string]string) (*application.ProjectModifyResponse, error)) {
	s.modifyActions[application.ProjectModifyRequest_CANCEL] = f
}

func (s *ApplicationServer) OnBilling(f func(ctx context.Context, distributorID, projectID string, data map[string]string) (*application.ProjectModifyResponse, error)) {
	s.modifyActions[application.ProjectModifyRequest_BILLING] = f
}
