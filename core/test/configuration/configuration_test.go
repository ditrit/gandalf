package configuration

import (
	"github.com/ditrit/gandalf/core/configuration"
	"testing"
)

func TestSetStringKeyConfig(t *testing.T) {

	err := configuration.SetStringKeyConfig("test", "testStrKey", "", "toto", "test string usage", false)
	if err != nil {
		t.Errorf("SetStringConfig shouldn't fail")
	}
}

func TestSetIntegerKeyConfig(t *testing.T) {
	err := configuration.SetIntegerKeyConfig("test", "testIntKey", "", 10, "test integer usage", false)
	if err != nil {
		t.Errorf("SetIntegerConfig shouldn't fail")
	}
}

func TestParseConfig(t *testing.T) {
	configuration.InitMainConfigKeys()
	err := configuration.ParseConfig()
	if err != nil {
		t.Errorf("impossible to parse config")
	}
}

func TestIsConfigValid(t *testing.T) {
	err := configuration.IsConfigValid()
	if err == nil {
		t.Errorf("Configuration should be valid")
	}
}

func TestGetStringConfig(t *testing.T) {
	strVal, err := configuration.GetStringConfig("testStrKey")

	if err != nil {
		t.Errorf("GetStringConfig shouldn't have failed because the key exists")
	}

	if strVal == "" {
		t.Errorf("the value of testKey shouldn't be an empty string")
	}
}

func TestGetIntegerConfig(t *testing.T) {
	intVal, err := configuration.GetIntegerConfig("testIntKey")

	if err != nil {
		t.Errorf("GetIntegerConfig shouldn't have falied because the key exists")
	}

	if intVal == -1 {
		t.Errorf("The value of intVal shouldn't be -1")
	}
}
