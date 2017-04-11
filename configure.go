package potens

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/kubex/potens-go/definition"
	"github.com/kubex/potens-go/identity"
)

func (app *Application) SetIdentity(ident *identity.AppIdentity) error {
	if ident == nil {
		ident = &identity.AppIdentity{}
		err := ident.FromJSONFile(app.relPath("app-identity.json"))
		if err != nil {
			return err
		}
	}

	block, _ := pem.Decode([]byte(ident.PrivateKey))
	if block == nil {
		return errors.New("No RSA private key found")
	}

	var key *rsa.PrivateKey
	if block.Type == "RSA PRIVATE KEY" {
		rsapk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return errors.New("Unable to read RSA private key")
		}
		key = rsapk
	}

	if !app.canBecomeGlobalAppID(ident.AppID) {
		return errors.New("The App ID specified in your identity does not match your definition")
	}

	app.identity = ident
	app.pk = key
	app.kh = ident.KeyHandle

	return nil
}

func (app *Application) SetDefinition(def *definition.AppDefinition) error {
	if def == nil {
		def = &definition.AppDefinition{}
		err := def.FromConfig(app.relPath("app-definition.yaml"))
		if err != nil {
			return err
		}
	}

	if len(def.Vendor) < 2 {
		return errors.New("The Vendor ID specified in your definition file is invalid")
	}

	if len(def.AppID) < 2 {
		return errors.New("The App ID specified in your definition file is invalid")
	}

	if !app.canBecomeGlobalAppID(def.GlobalAppID) {
		return errors.New("The App ID specified in your definition does not match your identity")
	}

	app.definition = def
	return nil
}

func (app *Application) canBecomeGlobalAppID(globalAppID string) bool {
	if app.identity != nil {
		return app.identity.AppID == globalAppID
	}

	if app.definition != nil {
		return app.definition.GlobalAppID == globalAppID
	}

	return true
}

// Identity retrieves your identity
func (app *Application) Identity() *identity.AppIdentity {
	return app.identity
}

// Definition retrieves your definition
func (app *Application) Definition() *definition.AppDefinition {
	return app.definition
}
