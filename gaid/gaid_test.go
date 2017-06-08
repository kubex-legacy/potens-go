package gaid_test

import (
	"testing"

	"github.com/kubex/potens-go/gaid"
)

func TestValidate(t *testing.T) {
	if gaid.FromString("/asd").Validate(true) == nil {
		t.Fail()
	}

	if gaid.FromString("asd/").Validate(true) == nil {
		t.Fail()
	}

	if gaid.FromString("/adewd").Validate(true) == nil {
		t.Fail()
	}

	if gaid.FromString("ddew").Validate(true) == nil {
		t.Fail()
	}

	if gaid.FromString("dwfwe").Validate(true) == nil {
		t.Fail()
	}

	if gaid.FromString("abc/def").Validate(true) != nil {
		t.Error(gaid.FromString("abc/def").Validate(true))
	}

}

func TestFromString(t *testing.T) {
	glapid := gaid.FromString("cubex/app")
	if glapid.VendorID != "cubex" {
		t.Error("Incorrect vendor ID")
	}
	if glapid.AppID != "app" {
		t.Error("Incorrect app ID")
	}

	glapid = gaid.FromString("cubex")
	if glapid.Validate(true) == nil {
		t.Error("Failed to error on invalid GAID")
	}
	glapid = gaid.FromString("cubex/klhfw/fwejhfew")

	if glapid.Validate(true) == nil {
		t.Error("Failed to error on invalid GAID")
	}
}

func TestRemainder(t *testing.T) {
	glapid := gaid.FromString("cubex/app")
	if glapid.Validate(true) != nil {
		t.Error(glapid.Validate(true).Error())
	}
	if glapid.VendorID != "cubex" {
		t.Error("Incorrect vendor ID")
	}
	if glapid.AppID != "app" {
		t.Error("Incorrect app ID")
	}
	if glapid.Remainder != "" {
		t.Error("Incorrect remaining")
	}

	glapid = gaid.FromString("cubex")
	if glapid.Validate(true) == nil {
		t.Error("Failed to error on invalid GAID")
	}

	glapid = gaid.FromString("cubex/abc/def")
	if glapid.VendorID != "cubex" {
		t.Error("Incorrect vendor ID")
	}
	if glapid.AppID != "abc" {
		t.Error("Incorrect app ID")
	}
	if glapid.Remainder != "def" {
		t.Error("Incorrect remaining")
	}
	glapid = gaid.FromString("cubex/abc/def/xyz")
	if glapid.Remainder != "def/xyz" {
		t.Error("Incorrect remaining")
	}
}
