package potens

import (
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kubex/potens-go/auth"

	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// GrpcBackgroundContext context to use when communicating with other services
func (app *Application) GrpcBackgroundContext() context.Context {
	return app.GrpcContext(context.Background())
}

// GrpcTimeoutContext context to use when communicating with other services
func (app *Application) GrpcTimeoutContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(app.GrpcContext(context.Background()), timeout)
}

// GrpcContext context to use when communicating with other services
func (app *Application) GrpcContext(parent context.Context) context.Context {
	md := metadata.Pairs(
		auth.GetAppIDKey(), app.Definition().AppID,
		auth.GetAppVendorKey(), app.Definition().VendorID,
	)

	if parentMd, hasParentMd := metadata.FromIncomingContext(parent); hasParentMd {
		return metadata.NewOutgoingContext(parent, metadata.Join(md, parentMd))
	}
	return metadata.NewOutgoingContext(parent, md)
}

// CreateServer creates a gRPC server with your tls certificates
func (app *Application) CreateServer() error {

	app.server = grpc.NewServer()

	return nil
}

//GetGrpcServer returns the grpc server
func (app *Application) GetGrpcServer() *grpc.Server {
	return app.server
}

func (app *Application) Serve() error {
	lis, err := net.Listen("tcp", ":"+app.PortString())
	if err != nil {
		return err
	}

	app.Log().Debug("Serving", zap.Int("port", app.Port))

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

func (app *Application) GetServiceConnection(service string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithDialer(app.grpcServiceDialer))
	return grpc.Dial(service, opts...)
}

func (app *Application) getEnvLocation(prefix string, service string) string {
	service = strings.Replace(service, "-", "_", -1)
	serviceHost := os.Getenv(strings.ToUpper(prefix+service) + EnvServiceHostSuffix)
	servicePort := os.Getenv(strings.ToUpper(prefix+service) + EnvServicePortSuffix + "_GRPC")
	if servicePort == "" {
		servicePort = os.Getenv(strings.ToUpper(prefix+service) + EnvServicePortSuffix + "_DEFAULTPORT")
	}
	if servicePort == "" {
		servicePort = os.Getenv(strings.ToUpper(prefix+service) + EnvServicePortSuffix)
	}
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
