package forms_test

import (
	"encoding/json"
	"testing"

	"github.com/kubex/potens-go/webui/forms"
)

func TestError(t *testing.T) {
	frmErr := forms.NewError()
	frmErr.AddError("field_1", "unable to do something")
	frmErr.AddError("field_1", "unable to do another thing")
	frmErr.AddError("field_2", "generic")

	jsonErr, _ := json.Marshal(frmErr)
	if string(jsonErr) != "{\"Errors\":{\"field_1\":[\"unable to do something\",\"unable to do another thing\"],\"field_2\":[\"generic\"]}}" {
		t.Fatal("Invalid json generated")
	}
}
