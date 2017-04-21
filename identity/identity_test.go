package identity_test

import (
	"testing"

	"github.com/kubex/potens-go/identity"
)

func TestReadYaml(t *testing.T) {
	ident := identity.AppIdentity{}
	err := ident.FromJSONFile("app-identity.dist.json")
	if err != nil {
		t.Fatal(err)
	}
}
