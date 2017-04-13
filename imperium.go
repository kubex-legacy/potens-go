package potens

import (
	"github.com/kubex/proto-go/imperium"
	"go.uber.org/zap"
)

func (app *Application) GetCertificate() error {

	imperiumConnection, err := app.GetServiceConnection(imperiumServiceName)
	if err != nil {
		return err
	}
	c := imperium.NewImperiumClient(imperiumConnection)
	response, err := c.Request(app.GetGrpcContext(), &imperium.CertificateRequest{
		AppId: app.GlobalAppID(),
	})

	if err != nil {
		return err
	}

	app.hostname = response.Hostname
	app.imperiumCertificate = []byte(response.Certificate)
	app.imperiumKey = []byte(response.PrivateKey)

	app.Log().Debug("Received TLS Certificates", zap.String("hostname", app.hostname))

	return nil
}
