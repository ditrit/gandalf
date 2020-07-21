package utils

import (
	"gandalf/core/connector/utils"
	"shoset/msg"
	"testing"
)

func CreateValidationEventTest(t *testing.T) {
	const (
		tenant  string = "test"
		event   string = "ON_GOING"
		payload string = ""
		timeout int64  = 10000
	)
	command := msg.Command{Command: "Test"}

	msgevent := utils.CreateValidationEvent(command, tenant)

	if msgevent == nil {
		t.Errorf("Should not be nil")
	}

	if msgevent.GetTopic() != command.GetCommand() {
		t.Errorf("Should be equal")
	}

	if msgevent.GetEvent() != event {
		t.Errorf("Should be equal")
	}

	if msgevent.GetPayload() != payload {
		t.Errorf("Should be equal")
	}

	if msgevent.GetReferencesUUID() != command.GetUUID() {
		t.Errorf("Should be equal")
	}

	if msgevent.GetTimeout() != timeout {
		t.Errorf("Should be equal")
	}
}

func IsExecAllTest(t *testing.T) {
	//IsExecAll(mode os.FileMode)
}

func GetMaxVersionTest(t *testing.T) {
	versions := []int64{1, 2, 3}
	max := utils.GetMaxVersion(versions)

	if max != 3 {
		t.Errorf("Should be 3")
	}
}
