package connector

import (
	"reflect"
	"testing"

	"github.com/ditrit/gandalf/core/connector"
)

func TestNewConnectorMember(t *testing.T) {
	const (
		logicalName   string = "test"
		tenant        string = "test"
		connectorType string = "test"
		logPath       string = "test"
		shosetType    string = "c"
	)
	versions := []int64{1, 2}

	connectorMember := connector.NewConnectorMember(logicalName, tenant, connectorType, logPath, versions)

	if connectorMember == nil {
		t.Errorf("Should not be nil")
	}

	if connectorMember.GetChaussette().GetName() != logicalName {
		t.Errorf("Should be equal")
	}

	if connectorMember.GetChaussette().GetShosetType() != shosetType {
		t.Errorf("Should be equal")
	}

	if connectorMember.GetChaussette().Context["tenant"] != tenant {
		t.Errorf("Should be equal")
	}

	if connectorMember.GetChaussette().Context["connectorType"] != connectorType {
		t.Errorf("Should be equal")
	}

	if !reflect.DeepEqual(connectorMember.GetChaussette().Context["versions"].([]int64), versions) {
		t.Errorf("Should be equal")
	}

	if connectorMember.GetChaussette().Handle["cfgjoin"] == nil {
		t.Errorf("Should not be nil")
	}

	if connectorMember.GetChaussette().Handle["cmd"] == nil {
		t.Errorf("Should not be nil")
	}

	if connectorMember.GetChaussette().Handle["evt"] == nil {
		t.Errorf("Should not be nil")
	}
}

func TestConnectorMemberInit(t *testing.T) {
	const (
		logicalName     string = "test"
		tenant          string = "test"
		shosetType      string = "c"
		bindAddress     string = "127.0.0.1:8001"
		grpcBindAddress string = "127.0.0.1:8002"
		linkAddress     string = "127.0.0.1:8003"
		connectorType   string = "test"
		product         string = "test"
		workerUrl       string = "test"
		workerPath      string = "test"
		logPath         string = "test"
		timeoutMax      int64  = 1000
	)
	versions := []int64{1, 2}

	connectorMember := connector.ConnectorMemberInit(logicalName, tenant, bindAddress, grpcBindAddress, linkAddress, connectorType, product, workerUrl, workerPath, logPath, timeoutMax, versions)

	if connectorMember == nil {
		t.Errorf("Should not be nil")
	}

	if connectorMember.GetChaussette().GetBindAddr() != bindAddress {
		t.Errorf("Should be equal")
	}

	//TODO GRPC

	//TODO LINK

	if aggregatorMember.GetTimeoutMax() != timeoutMax {
		t.Errorf("Should be equal")
	}

}
