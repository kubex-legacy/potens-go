package potens

import (
	"crypto/tls"
	"errors"
	"net"
	"os"
	"strings"

	"github.com/cubex/portcullis-go/keys"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// GetGrpcContext context to use when communicating with other services
func (app *Application) GetGrpcContext() context.Context {
	md := metadata.Pairs(
		keys.GetAppIDKey(), app.Definition().AppID,
		keys.GetAppVendorKey(), app.Definition().VendorID,
	)
	return metadata.NewContext(context.Background(), md)
}

// CreateServer creates a gRPC server with your tls certificates
func (app *Application) CreateServer() error {

	if app.imperiumKey == nil || app.imperiumCertificate == nil || app.hostname == "" {
		return errors.New("CreateServer called before GetCertificate, or GetCertificate call failed")
	}

	cert, err := tls.X509KeyPair(app.imperiumCertificate, app.imperiumKey)
	if err != nil {
		return err
	}

	app.server = grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&cert)))

	return nil
}

func (app *Application) Serve() error {

	lis, err := net.Listen("tcp", app.hostname+":"+app.PortString())
	if err != nil {
		return err
	}

	return app.server.Serve(lis)
}

func (app *Application) GetServiceConnection(service string) (*grpc.ClientConn, error) {
	location := os.Getenv(app.ServiceKey() + envServiceLocationSuffix)

	kubexServiceDomain := os.Getenv(envKubexServiceDomain)
	if kubexServiceDomain == "" {
		kubexServiceDomain = KubexProductionServicesDomain
	}

	if location == "" {
		location = strings.ToLower(service) + "." + kubexServiceDomain
	}

	return grpc.Dial(location, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
}
