package potens

import (
	"log"
	"net"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"strings"
)

//NewApplication creates an instance of your application, name will be converted to upper and _
func NewApplication(name string) (*Application) {
	app := makeApp(name)

	osPort := os.Getenv("APP_" + app.ServiceKey() + EnvListenPortSuffix)
	if osPort != "" {
		intPort, err := strconv.ParseInt(osPort, 10, 32)
		if err != nil {
			log.Print("Unable to use ", osPort, " as the listen port for ", name)
		} else {
			app.Port = int(intPort)
		}
	}

	return app
}

//NewService create a new instance of a service, this is usually not required
func NewService(name string) (*Application) {
	app := makeApp(name)
	return app
}

func makeApp(name string) (*Application) {
	var err error
	app := &Application{
		services:       &serviceCache{},
		Name:           name,
		IP:             net.IPv4(127, 0, 0, 1),
		Port:           KubexDefaultGRPCPort,
		consoleDomain:  KubexProductionConsoleDomain,
		logDiscoveryHB: false,
	}

	app.instanceID = uuid.NewV4().String()

	kubexDomain := os.Getenv(EnvKubexConsoleDomain)
	if kubexDomain != "" {
		app.consoleDomain = kubexDomain
	}

	if app.consoleDomain != KubexProductionConsoleDomain {
		app.logger, err = zap.NewDevelopment()
	} else {
		app.logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatal(err)
	}

	osPort := os.Getenv(app.ServiceKey() + EnvListenPortSuffix)
	if osPort != "" {
		intPort, err := strconv.ParseInt(osPort, 10, 32)
		if err != nil {
			log.Print("Unable to use ", osPort, " as the listen port for ", name)
		} else {
			app.Port = int(intPort)
		}
	}

	return app
}

func (app *Application) relPath(file string) string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	app.FatalErr(err)

	// if executable is in TempDir use getwd instead (go run)
	if strings.HasPrefix(dir, os.TempDir()) {
		dir, err = os.Getwd()
		app.FatalErr(err)
		dir, err = filepath.Abs(dir)
		app.FatalErr(err)
	}

	return path.Join(dir, file)
}
