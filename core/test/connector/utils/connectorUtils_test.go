package utils

import (
	"reflect"
	"testing"

	"github.com/ditrit/shoset/msg"

	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/gandalf/core/connector/utils"
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

func GetConnectorTypeTest(t *testing.T) {
	const (
		connectorTypeName1 string = "toto"
		connectorTypeName2 string = "tutu"
	)
	ConnectorConfig1 := &models.ConnectorConfig{Name: "toto", ConnectorTypeID: 0, ConnectorType: models.ConnectorType{Name: "toto"}}
	ConnectorConfig2 := &models.ConnectorConfig{Name: "tata", ConnectorTypeID: 1, ConnectorType: models.ConnectorType{Name: "tata"}}
	list := []*models.ConnectorConfig{ConnectorConfig1, ConnectorConfig2}
	result1 := utils.GetConnectorType(connectorTypeName1, list)

	if reflect.DeepEqual(result1, ConnectorConfig1) {
		t.Errorf("Should be equal")
	}

	result2 := utils.GetConnectorType(connectorTypeName2, list)

	if reflect.DeepEqual(result2, (models.ConnectorTypeCommand{})) {
		t.Errorf("Should not be equal")
	}
}

// GetConnectorTypeConfigByVersion : GetConnectorTypeConfigByVersion
func GetConnectorTypeConfigByVersionTest(t *testing.T) {
	const (
		version1 int64 = 0
		version2 int64 = 1
		version3 int64 = 3
	)
	ConnectorConfig1 := &models.ConnectorConfig{Name: "toto", Version: 1}
	ConnectorConfig2 := &models.ConnectorConfig{Name: "tata", Version: 2}
	list := []*models.ConnectorConfig{ConnectorConfig1, ConnectorConfig2}
	result1 := utils.GetConnectorTypeConfigByVersion(version1, list)

	if result1 != nil {
		t.Errorf("Should be equal")
	}

	result2 := utils.GetConnectorTypeConfigByVersion(version2, list)

	if reflect.DeepEqual(result2, ConnectorConfig1) {
		t.Errorf("Should be equal")
	}

	result3 := utils.GetConnectorTypeConfigByVersion(version3, list)

	if reflect.DeepEqual(result3, (models.ConnectorTypeCommand{})) {
		t.Errorf("Should be equal")
	}

}

func GetConnectorTypeCommandTest(t *testing.T) {
	const (
		commandName1 string = "toto"
		commandName2 string = "tutu"
	)
	connectorTypeCommandTest1 := models.ConnectorTypeCommand{Name: "toto", Schema: "test"}
	connectorTypeCommandTest2 := models.ConnectorTypeCommand{Name: "tata", Schema: "test"}
	list := []models.ConnectorTypeCommand{connectorTypeCommandTest1, connectorTypeCommandTest2}
	result1 := utils.GetConnectorTypeCommand(commandName1, list)

	if reflect.DeepEqual(result1, connectorTypeCommandTest1) {
		t.Errorf("Should be equal")
	}

	result2 := utils.GetConnectorTypeCommand(commandName2, list)

	if reflect.DeepEqual(result2, (models.ConnectorTypeCommand{})) {
		t.Errorf("Should not be equal")
	}
}

func GetConnectorTypeEventTest(t *testing.T) {
	const (
		eventName1 string = "toto"
		eventName2 string = "tutu"
	)
	connectorTypeEventTest1 := models.ConnectorTypeEvent{Name: "toto", Schema: "test"}
	connectorTypeEventTest2 := models.ConnectorTypeEvent{Name: "tata", Schema: "test"}
	list := []models.ConnectorTypeEvent{connectorTypeEventTest1, connectorTypeEventTest2}
	result1 := utils.GetConnectorTypeEvent(eventName1, list)

	if reflect.DeepEqual(result1, connectorTypeEventTest1) {
		t.Errorf("Should be equal")
	}

	result2 := utils.GetConnectorTypeEvent(eventName2, list)

	if reflect.DeepEqual(result2, (models.ConnectorTypeEvent{})) {
		t.Errorf("Should not be equal")
	}
}

func ValidatePayloadTest(t *testing.T) {
	const (
		payload            string = "test"
		payloadSchema      string = `{"type":"string"}`
		payloadSchemaFalse string = ""
	)
	validatetrue := utils.ValidatePayload(payload, payloadSchema)
	if !validatetrue {
		t.Errorf("Should be true")
	}

	validatefalse := utils.ValidatePayload(payload, payloadSchemaFalse)
	if validatefalse {
		t.Errorf("Should be false")
	}

}
