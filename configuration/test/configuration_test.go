package test

import (
	"core/configuration"
	"testing"
)

func TestParse(t *testing.T){

err := configuration.ParseConfig()

	if err != nil {
		t.Errorf("test error")
	}
}


