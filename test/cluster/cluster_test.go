package cluster

import (
	"core/cluster"
	"testing"
)

func TestNewClusterMember(t *testing.T) {
	const (
		logicalName string = "test"
		shosetType  string = "cl"
	)

	connectorMember := cluster.NewClusterMember(logicalName)

	if connectorMember == nil {
		t.Errorf("Should not be nil")
	}

	if connectorMember.GetChaussette().GetName() != logicalName {
		t.Errorf("Should be equal")
	}

	if connectorMember.GetChaussette().GetShosetType() != shosetType {
		t.Errorf("Should be equal")
	}

	if connectorMember.MapDatabaseClient == nil {
		t.Errorf("Should not be nil")
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

func TestClusterMemberInit(t *testing.T) {
	const (
		logicalName string = "test"
		shosetType  string = "cls"
		bindAddress string = "127.0.0.1:8001"
	)

	clusterMember := cluster.ClusterMemberInit(logicalName, bindAddress)

	if clusterMember == nil {
		t.Errorf("Should not be nil")
	}

	if clusterMember.GetChaussette().GetBindAddr() != bindAddress {
		t.Errorf("Should be equal")
	}
}
