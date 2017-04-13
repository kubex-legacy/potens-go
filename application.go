package potens

import (
	"crypto/rsa"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/kubex/potens-go/definition"
	"github.com/kubex/potens-go/identity"
	"github.com/kubex/proto-go/discovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Application struct {
	/** RunTime **/
	Port          int
	IP            net.IP
	instanceID    string
	kubexDomain   string
	currentStatus discovery.ServiceStatus

	/** Definition **/
	Name       string
	definition *definition.AppDefinition
	appVersion discovery.AppVersion

	/** Identity **/
	identity *identity.AppIdentity
	pk       *rsa.PrivateKey
	kh       string

	/** Utility **/
	logger *zap.Logger

	/** Security **/
	imperiumCertificate []byte
	imperiumKey         []byte
	hostname            string

	/** gRPC **/
	server   *grpc.Server
	services *services
}

func (app *Application) FatalErr(err error) {
	if err != nil {
		app.Log().Fatal("Oops", zap.Error(err))
	}
}

//GlobalAppID returns the global app ID for the application
func (app *Application) GlobalAppID() string {
	return app.definition.GlobalAppID()
}

//VendorID returns the vendor ID for this application
func (app *Application) VendorID() string {
	return app.definition.VendorID
}

//AppID returns the app ID for this application
func (app *Application) AppID() string {
	return app.definition.AppID
}

//InstanceID returns the instance ID for this session
func (app *Application) InstanceID() string {
	return app.instanceID
}

//PortString returns the port as a string
func (app *Application) PortString() string {
	return strconv.FormatInt(int64(app.Port), 10)
}

//KubexDomain returns the kubex domain being used
func (app *Application) KubexDomain() string {
	return app.kubexDomain
}

//IsProduction returns if you are in production environment
func (app *Application) IsProduction() bool {
	return app.kubexDomain == KubexProductionDomain
}

//Log returns a logger to log against
func (app *Application) Log() *zap.Logger {
	return app.logger
}

//ServiceKey returns a key that can be used to pre-fix environment variables
func (app *Application) ServiceKey() string {
	return strings.ToUpper(strings.Replace(regexp.MustCompile("[^A-Za-z0-9\\-_]").ReplaceAllString(app.Name, ""), "-", "_", -1))
}
