package aggregator

import (
	"github.com/mathieucaroff/gandalf-core/aggregator"
	"testing"
)

func TestNewAggregatorMember(t *testing.T) {
	const (
		logicalName string = "test"
		tenant      string = "test"
		shosetType  string = "a"
	)

	aggregatorMember := aggregator.NewAggregatorMember(logicalName, tenant)

	if aggregatorMember == nil {
		t.Errorf("Should not be nil")
	}

	if aggregatorMember.GetChaussette().GetName() != logicalName {
		t.Errorf("Should be equal")
	}

	if aggregatorMember.GetChaussette().GetShosetType() != shosetType {
		t.Errorf("Should be equal")
	}

	if aggregatorMember.GetChaussette().Context["tenant"] != tenant {
		t.Errorf("Should be equal")
	}

	if aggregatorMember.GetChaussette().Handle["cfgjoin"] == nil {
		t.Errorf("Should not be nil")
	}

	if aggregatorMember.GetChaussette().Handle["cmd"] == nil {
		t.Errorf("Should not be nil")
	}

	if aggregatorMember.GetChaussette().Handle["evt"] == nil {
		t.Errorf("Should not be nil")
	}
}

func TestAggregatorMemberInit(t *testing.T) {
	const (
		logicalName string = "test"
		tenant      string = "test"
		shosetType  string = "a"
		bindAddress string = "127.0.0.1:8001"
		linkAddress string = "127.0.0.1:8002"
	)

	aggregatorMember := aggregator.AggregatorMemberInit(logicalName, tenant, bindAddress, linkAddress)

	if aggregatorMember == nil {
		t.Errorf("Should not be nil")
	}

	if aggregatorMember.GetChaussette().GetBindAddr() != bindAddress {
		t.Errorf("Should be equal")
	}

	//TODO LINK
}
