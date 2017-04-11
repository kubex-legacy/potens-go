package potens

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type Application struct {
	Name        string
	ServiceName string
	Port        int
	IP          net.IP

	instanceID  string
	kubexDomain string

	logger *zap.Logger
}

//NewApplication creates an instance of your application, name will be converted to upper and _
func NewApplication(name string) (*Application) {
	var err error
	app := &Application{
		Name:        name,
		IP:          net.IPv4(127, 0, 0, 1),
		Port:        KubexDefaultGRPCPort,
		kubexDomain: KubexProductionDomain,
	}

	app.instanceID = uuid.NewV4().String()

	app.ServiceName = strings.ToUpper(name)
	app.ServiceName = strings.Replace(app.ServiceName, " ", "_", -1)

	kubexDomain := os.Getenv("KUBEX_DOMAIN")
	if kubexDomain != "" {
		app.kubexDomain = kubexDomain
	}

	if app.kubexDomain != KubexProductionDomain {
		app.logger, err = zap.NewDevelopment()
	} else {
		app.logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatal(err)
	}

	osPort := os.Getenv(app.ServiceName + "_LISTEN_PORT")
	if osPort != "" {
		intPort, err := strconv.ParseInt(osPort, 10, 32)
		if err != nil {
			app.Port = int(intPort)
		} else {
			log.Print("Unable to use ", osPort, " as the listen port for ", name)
		}
	}

	return app
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
