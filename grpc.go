package potens

import (
	"crypto/tls"
	"errors"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cubex/portcullis-go/keys"
	"go.uber.org/zap"
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

	app.server = grpc.NewServer()

	//Do not secure with imperium for initial development
	if true {
		return nil
	}

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

//GetServer returns the grpc server
func (app *Application) GetServer() *grpc.Server {
	return app.server
}

func (app *Application) Serve() error {
	lis, err := net.Listen("tcp", ":"+app.PortString())
	if err != nil {
		return err
	}

	return app.server.Serve(lis)
}

func (app *Application) grpcServiceDialer(service string, timeout time.Duration) (net.Conn, error) {
	location := os.Getenv(strings.ToUpper(service) + EnvServiceLocationSuffix)

	kubexServiceDomain := os.Getenv(EnvKubexServiceDomain)
	if kubexServiceDomain == "" {
		kubexServiceDomain = KubexProductionServicesDomain
	}

	location = app.GetServiceEnvLocation(service)

	if location == "" {
		location = strings.ToLower(service) + "." + kubexServiceDomain
		location += ":" + strconv.FormatInt(int64(KubexDefaultGRPCPort), 10)
	}

	app.Log().Debug("Dialing GRPC", zap.String("service", service), zap.String("location", location))

	return net.DialTimeout("tcp", location, timeout)
}

func (app *Application) GetServiceConnection(service string) (*grpc.ClientConn, error) {
	return grpc.Dial(service, grpc.WithInsecure(), grpc.WithDialer(app.grpcServiceDialer))
}

func (app *Application) getEnvLocation(prefix string, service string) string {
	service = strings.Replace(service, "-", "_", -1)
	serviceHost := os.Getenv(strings.ToUpper(prefix+service) + EnvServiceHostSuffix)
	servicePort := os.Getenv(strings.ToUpper(prefix+service) + EnvServicePortSuffix)
	if serviceHost+servicePort != "" {
		return serviceHost + ":" + servicePort
	}

	return ""
}

func (app *Application) GetServiceEnvLocation(service string) string {
	return app.getEnvLocation("", service)
}

func (app *Application) GetAppEnvLocation(service string) string {
	return app.getEnvLocation("app_", service)
}
