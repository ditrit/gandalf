package test

import (
	"testing"

	"github.com/ditrit/gandalf-core/configuration"
)

func TestSetStringKeyConfig(t *testing.T) {

	err := configuration.SetStringKeyConfig("test", "testStrKey", "", "toto", "test string usage", true)
	if err != nil {
		t.Errorf("SetStringConfig shouldn't fail")
	}
}

func TestSetIntegerKeyConfig(t *testing.T) {

	err := configuration.SetIntegerKeyConfig("test", "testIntKey", "", 10, "test integer usage", true)
	if err != nil {
		t.Errorf("SetIntegerConfig shouldn't fail")
	}
}

func TestGetStringConfig(t *testing.T) {
	strVal, err := configuration.GetStringConfig("testStrKey")

	if err != nil {
		t.Errorf("GetStringConfig shouldn't have failed because the key exists")
	}

	if strVal != "" {
		t.Errorf("the value of testKey shoud be an empty string")
	}
}

func TestGetIntegerConfig(t *testing.T) {
	_, err := configuration.GetIntegerConfig("testIntKey")

	if err == nil {
		t.Errorf("GetIntegerConfig should fail because testIntKey is an empty string")
	}
}
