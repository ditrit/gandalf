package grpc

import (
	"github.com/mathieucaroff/gandalf-core/grpc"
	"strconv"
	"testing"
)

func TestCommandFromGrpc(t *testing.T) {
	const (
		tenant    string = "test"
		token     string = "test"
		timeout   string = "test"
		timestamp string = "test"
		major     string = "test"
		minor     string = "test"
		uUID      string = "test"
		command   string = "test"
		payload   string = "test"
	)

	commandGrpcMessage := grpc.CommandMessage{Tenant: tenant, Token: token, Timeout: timeout, Timestamp: timestamp, Major: major, Minor: minor, UUID: uUID, Command: command, Payload: payload}
	commandMessage := grpc.CommandFromGrpc(&commandGrpcMessage)

	if commandMessage.GetTenant() != tenant {
		t.Errorf("Should be equal")
	}

	if commandMessage.GetToken() != token {
		t.Errorf("Should be equal")
	}

	testTimeout, _ := strconv.ParseInt(timeout, 10, 64)
	if commandMessage.GetTimeout() != testTimeout {
		t.Errorf("Should be equal")
	}

	testTimestamp, _ := strconv.ParseInt(timeout, 10, 64)
	if commandMessage.GetTimestamp() != testTimestamp {
		t.Errorf("Should be equal")
	}

	testMajor, _ := strconv.ParseInt(minor, 10, 8)
	if commandMessage.GetMajor() != int8(testMajor) {
		t.Errorf("Should be equal")
	}

	testMinor, _ := strconv.ParseInt(minor, 10, 8)
	if commandMessage.GetMinor() != int8(testMinor) {
		t.Errorf("Should be equal")
	}

	if commandMessage.GetUUID() != uUID {
		t.Errorf("Should be equal")
	}

	if commandMessage.GetCommand() != command {
		t.Errorf("Should be equal")
	}

	if commandMessage.GetPayload() != payload {
		t.Errorf("Should be equal")
	}
}
