package potens

import (
	"log"
	"net"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

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
	app.ServiceName = strings.Replace(strings.Replace(app.ServiceName, "-", "_", -1), " ", "_", -1)

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
			log.Print("Unable to use ", osPort, " as the listen port for ", name)
		} else {
			app.Port = int(intPort)
		}
	}

	return app
}

func (app *Application) relPath(file string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		app.Log().Fatal(err.Error())
	}
	return path.Join(dir, file)
}
