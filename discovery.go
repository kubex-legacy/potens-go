package potens

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kubex/potens-go/services"
	"github.com/kubex/proto-go/discovery"
	"go.uber.org/zap"
	"net"
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
	app.Log().Debug("Registering with Discovery")

	if hostname == "" {
		hostname = getLocalIP()
	}

	if port == "" {
		port = app.PortString()
	}

	err := app.connectToDiscovery()
	if err != nil {
		app.Log().Info("Discovery Connect Failure", zap.Error(err))
		return err
	}

	app.Log().Debug("Opened Connection to Discovery")
	version := os.Getenv(app.ServiceKey() + EnvServiceVersionSuffix)
	if version != "" {
		v, ok := discovery.AppRelease_value[version]
		if ok {
			app.appRelease = discovery.AppRelease(v)
		}
	}

	portInt, _ := strconv.ParseInt(port, 10, 32)

	app.Log().Debug("Starting Registration")
	tCtx, _ := app.GrpcTimeoutContext(time.Second * 5)
	regResult, err := app.services.discoveryClient.Register(tCtx, &discovery.RegisterRequest{
		AppId:        app.GlobalAppID(),
		InstanceUuid: app.instanceID,
		ServiceHost:  hostname,
		Release:      app.appRelease,
		ServicePort:  int32(portInt),
	})
	if err != nil {
		app.Log().Info("Registration Failure", zap.Error(err))
		return err
	}

	if !regResult.Recorded {
		app.Log().Debug("Unable to register")
		return ErrDiscoveryUnableToRegister
	}

	app.services.discoveryHost = hostname
	app.services.discoveryPort = port
	app.services.discoveryRegistered = true

	return nil
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func (app *Application) discoveryHeartBeat() {
	app.connectToDiscovery()
	for {
		if app.currentStatus != discovery.ServiceStatus_ONLINE {
			return
		}
		app.Log().Debug("Sending heartbeat to discovery")
		tCtx, _ := app.GrpcTimeoutContext(time.Second * 1)
		_, err := app.services.discoveryClient.HeartBeat(tCtx, &discovery.HeartBeatRequest{
			AppId:        app.GlobalAppID(),
			InstanceUuid: app.instanceID,
			Release:      app.appRelease,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "unregistered") {
				app.services.discoveryClient = nil
				app.connectToDiscovery()
			}
			app.RegisterWithDiscovery(app.services.discoveryHost, app.services.discoveryPort)
		}
		time.Sleep(5 * time.Second)
	}
}

func (app *Application) DiscoveryOnline() error {

	if !app.services.discoveryRegistered {
		regErr := app.RegisterWithDiscovery("", "")
		if regErr != nil {
			return regErr
		}
	}

	err := app.connectToDiscovery()
	if err != nil {
		return err
	}

	tCtx, _ := app.GrpcTimeoutContext(time.Second * 5)
	statusResult, err := app.services.discoveryClient.Status(tCtx, &discovery.StatusRequest{
		AppId:        app.GlobalAppID(),
		InstanceUuid: app.instanceID,
		Release:      app.appRelease,
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

	tCtx, _ := app.GrpcTimeoutContext(time.Second * 5)
	statusResult, err := app.services.discoveryClient.Status(tCtx, &discovery.StatusRequest{
		AppId:        app.GlobalAppID(),
		InstanceUuid: app.instanceID,
		Release:      app.appRelease,
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
