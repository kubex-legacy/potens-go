package potens

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kubex/potens-go/services"
	"github.com/kubex/proto-go/discovery"
)

var ErrDiscoveryUnableToRegister = errors.New("Unable to register with the discovery service")
var ErrDiscoveryFailedToGoOnline = errors.New("Unable to take application online with discovery service")
var ErrDiscoveryFailedToGoOffline = errors.New("Unable to take application offline with discovery service")

func (app *Application) connectToDiscovery() error {
	if app.services.discoveryClient == nil {
		discoveryConn, err := app.GetServiceConnection(services.Discovery().Key())
		if err != nil {
			return err
		}
		app.services.discoveryClient = discovery.NewDiscoveryClient(discoveryConn)
	}
	return nil
}

func (app *Application) RegisterWithDiscovery(hostname string, port string) error {
	err := app.connectToDiscovery()
	if err != nil {
		return err
	}

	version := os.Getenv(app.ServiceKey() + EnvServiceVersionSuffix)
	app.appVersion = discovery.AppVersion_STABLE
	if version != "" {
		v, ok := discovery.AppVersion_value[version]
		if ok {
			app.appVersion = discovery.AppVersion(v)
		}
	}

	portInt, _ := strconv.ParseInt(port, 10, 32)

	regResult, err := app.services.discoveryClient.Register(app.GetGrpcContext(), &discovery.RegisterRequest{
		AppId:        app.GlobalAppID(),
		InstanceUuid: app.instanceID,
		ServiceHost:  hostname,
		Version:      app.appVersion,
		ServicePort:  int32(portInt),
	})
	if err != nil {
		return err
	}

	if !regResult.Recorded {
		return ErrDiscoveryUnableToRegister
	}

	app.services.discoveryHost = hostname
	app.services.discoveryPort = port

	return nil
}

func (app *Application) discoveryHeartBeat() {
	app.connectToDiscovery()
	for {
		if app.currentStatus != discovery.ServiceStatus_ONLINE {
			return
		}
		app.Log().Debug("Sending heartbeat to discovery")
		_, err := app.services.discoveryClient.HeartBeat(app.GetGrpcContext(), &discovery.HeartBeatRequest{
			AppId:        app.GlobalAppID(),
			InstanceUuid: app.instanceID,
			Version:      app.appVersion,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "unregistered") {
				app.services.discoveryClient = nil
				app.connectToDiscovery()
			}
			app.RegisterWithDiscovery(app.services.discoveryHost, app.services.discoveryPort)
		} else {
			time.Sleep(5 * time.Second)
		}
	}
}

func (app *Application) DiscoveryOnline() error {
	err := app.connectToDiscovery()
	if err != nil {
		return err
	}

	statusResult, err := app.services.discoveryClient.Status(app.GetGrpcContext(), &discovery.StatusRequest{
		AppId:        app.GlobalAppID(),
		InstanceUuid: app.instanceID,
		Version:      app.appVersion,
		Status:       discovery.ServiceStatus_ONLINE,
		Target:       discovery.StatusTarget_BOTH,
	})

	if err != nil {
		return err
	}

	if !statusResult.Recorded {
		return ErrDiscoveryFailedToGoOnline
	}

	app.currentStatus = discovery.ServiceStatus_ONLINE

	go app.discoveryHeartBeat()
	return nil
}

func (app *Application) DiscoveryOffline() error {
	err := app.connectToDiscovery()
	if err != nil {
		return err
	}

	statusResult, err := app.services.discoveryClient.Status(app.GetGrpcContext(), &discovery.StatusRequest{
		AppId:        app.GlobalAppID(),
		InstanceUuid: app.instanceID,
		Version:      app.appVersion,
		Status:       discovery.ServiceStatus_OFFLINE,
		Target:       discovery.StatusTarget_INSTANCE,
	})

	if err != nil {
		return err
	}

	if !statusResult.Recorded {
		return ErrDiscoveryFailedToGoOffline
	}

	app.currentStatus = discovery.ServiceStatus_OFFLINE

	return nil
}
