package connector

import (
	"testing"

	"github.com/ditrit/gandalf/core/connector"
)

func TestNewConnectorMember(t *testing.T) {
	const (
		logicalName string = "test"
		tenant      string = "test"
		shosetType  string = "c"
	)

	connectorMember := connector.NewConnectorMember(logicalName, tenant)

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
		timeoutMax      int64  = 1000
	)

	aggregatorMember := connector.ConnectorMemberInit(logicalName, tenant, bindAddress, grpcBindAddress, linkAddress, timeoutMax)

	if aggregatorMember == nil {
		t.Errorf("Should not be nil")
	}

	if aggregatorMember.GetChaussette().GetBindAddr() != bindAddress {
		t.Errorf("Should be equal")
	}

	//TODO GRPC

	//TODO LINK

	if aggregatorMember.GetTimeoutMax() != timeoutMax {
		t.Errorf("Should be equal")
	}

}
