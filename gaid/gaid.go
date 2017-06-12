package gaid

import (
	"errors"
	"strings"

	"github.com/cubex/cubeutil-go"
)

var (
	//ErrInvalidGlobalAppID invalid Global App ID
	ErrInvalidGlobalAppID = errors.New("The Global App ID specified is invalid")
)

//GlobalAppID
type GlobalAppID struct {
	VendorID  string
	AppID     string
	Remainder string
}

func (glapid *GlobalAppID) String() string {
	return glapid.VendorID + "/" + glapid.AppID
}

//New Create a Global App ID from your vendor and application IDs
func New(vendorID string, applicationID string) *GlobalAppID {
	resp := &GlobalAppID{
		VendorID: vendorID,
		AppID:    applicationID,
	}
	return resp
}

//FromString Take a string starting with a GlobalAppID, and extract the vendor, app and remainder
func FromString(gaidString string) *GlobalAppID {
	glapid := &GlobalAppID{}
	parts := strings.SplitN(gaidString, "/", 3)
	if len(parts) > 1 {
		glapid = New(parts[0], parts[1])
		if len(parts) > 2 {
			glapid.Remainder = parts[2]
		}
	}
	return glapid
}

//ValidateGlobalAppID Validates a Global App ID
func (glapid *GlobalAppID) Validate(strict bool) error {
	err := cubeutil.ValidateID(glapid.VendorID)
	if err != nil {
		return errors.New("Invalid Vendor ID " + glapid.VendorID)
	}
	err = cubeutil.ValidateID(glapid.AppID)
	if err != nil {
		return errors.New("Invalid App ID " + glapid.AppID)
	}
	if strict && glapid.Remainder != "" {
		return ErrInvalidGlobalAppID
	}
	return nil

}
