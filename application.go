package potens

import (
	"crypto/rsa"
	"net"
	"strconv"
	"github.com/kubex/potens-go/definition"
	"github.com/kubex/potens-go/identity"
	"go.uber.org/zap"
)

type Application struct {
	/** RunTime **/
	Port        int
	IP          net.IP
	instanceID  string
	kubexDomain string

	/** Definition **/
	ServiceName string
	Name        string
	definition  *definition.AppDefinition

	/** Identity **/
	identity *identity.AppIdentity
	pk       *rsa.PrivateKey
	kh       string

	/** Utility **/
	logger *zap.Logger
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
