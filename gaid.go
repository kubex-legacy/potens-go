package potens

import (
	"errors"
	"strings"
)

var (
	//ErrInvalidGlobalAppID invalid Global App ID
	ErrInvalidGlobalAppID = errors.New("The Global App ID specified is invalid")
)

//SplitGaID Split a Global App ID into Vendor and App ID
func SplitGaID(gaid string) (string, string, error) {
	parts := strings.SplitN(gaid, "/", 3)
	if len(parts) != 2 {
		return "", "", errors.New("Invalid GAID Provided")
	}
	return parts[0], parts[1], nil
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
