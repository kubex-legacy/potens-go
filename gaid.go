package potens

import (
	"errors"
	"strings"
	"github.com/cubex/cubeutil-go"
	"github.com/kubex/potens-go/definition"
)

var (
	//ErrInvalidGlobalAppID invalid Global App ID
	ErrInvalidGlobalAppID = errors.New("The Global App ID specified is invalid")
)

type GlobalAppID struct {
	VendorID string
	AppID    string
	AppType  definition.AppType
}

func (gaid *GlobalAppID) String() string {
	return gaid.VendorID + "/" + gaid.AppID
}

func NewGlobalAppID(vendorID string, applicationID string) GlobalAppID {
	gaid := GlobalAppID{
		VendorID: vendorID,
		AppID:    applicationID,
	}
	return gaid
}

func BuildGlobalAppID(gaid string) (GlobalAppID, string) {
	parts := strings.SplitN(gaid, "/", 3)
	return NewGlobalAppID(parts[0], parts[1]), parts[2]
}

//MakeGaID Create a Global App ID from your vendor and application IDs
func MakeGaID(vendor string, application string) (string, error) {
	err := cubeutil.ValidateID(vendor)
	if err != nil {
		return "", err
	}

	err = cubeutil.ValidateID(application)
	if err != nil {
		return "", err
	}

	return vendor + "/" + application, nil
}

//SplitGaID Split a Global App ID into Vendor and App ID
func SplitGaID(gaid string) (string, string, error) {
	parts := strings.SplitN(gaid, "/", 3)
	if len(parts) != 2 {
		return "", "", errors.New("Invalid GAID Provided")
	}
	return parts[0], parts[1], nil
}

//SplitGaIDEx Split a Global App ID into Vendor and App ID, and remaining parts
func SplitGaIDEx(gaid string) (string, string, string, error) {
	parts := strings.SplitN(gaid, "/", 3)
	if len(parts) < 2 {
		return "", "", "", errors.New("Invalid GAID Provided")
	} else if len(parts) < 3 {
		return parts[0], parts[1], "", nil
	}
	return parts[0], parts[1], parts[2], nil
}

//ValidateGlobalAppID Validates a Global App ID
func ValidateGlobalAppID(gapid string) error {
	if strings.Trim(gapid, "/") != gapid {
		return ErrInvalidGlobalAppID
	}
	if !strings.Contains(gapid, "/") {
		return ErrInvalidGlobalAppID
	}
	if len(gapid) < 3 {
		return ErrInvalidGlobalAppID
	}
	return nil
}
