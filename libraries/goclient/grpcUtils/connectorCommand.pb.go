// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: connectorCommand.proto

package grpcUtils

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CommandMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceAggregator      string `protobuf:"bytes,1,opt,name=SourceAggregator,proto3" json:"SourceAggregator,omitempty"`
	SourceConnector       string `protobuf:"bytes,2,opt,name=SourceConnector,proto3" json:"SourceConnector,omitempty"`
	SourceWorker          string `protobuf:"bytes,3,opt,name=SourceWorker,proto3" json:"SourceWorker,omitempty"`
	DestinationAggregator string `protobuf:"bytes,4,opt,name=DestinationAggregator,proto3" json:"DestinationAggregator,omitempty"`
	DestinationConnector  string `protobuf:"bytes,5,opt,name=DestinationConnector,proto3" json:"DestinationConnector,omitempty"`
	DestinationWorker     string `protobuf:"bytes,6,opt,name=DestinationWorker,proto3" json:"DestinationWorker,omitempty"`
	Tenant                string `protobuf:"bytes,7,opt,name=Tenant,proto3" json:"Tenant,omitempty"`
	Token                 string `protobuf:"bytes,8,opt,name=Token,proto3" json:"Token,omitempty"`
	Context               string `protobuf:"bytes,9,opt,name=Context,proto3" json:"Context,omitempty"`
	Timeout               string `protobuf:"bytes,10,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Timestamp             string `protobuf:"bytes,11,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Major                 int64  `protobuf:"varint,12,opt,name=Major,proto3" json:"Major,omitempty"`
	Minor                 int64  `protobuf:"varint,13,opt,name=Minor,proto3" json:"Minor,omitempty"`
	UUID                  string `protobuf:"bytes,14,opt,name=UUID,proto3" json:"UUID,omitempty"`
	ConnectorType         string `protobuf:"bytes,15,opt,name=ConnectorType,proto3" json:"ConnectorType,omitempty"`
	CommandType           string `protobuf:"bytes,16,opt,name=CommandType,proto3" json:"CommandType,omitempty"`
	Command               string `protobuf:"bytes,17,opt,name=Command,proto3" json:"Command,omitempty"`
	Payload               string `protobuf:"bytes,18,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (x *CommandMessage) Reset() {
	*x = CommandMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectorCommand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandMessage) ProtoMessage() {}

func (x *CommandMessage) ProtoReflect() protoreflect.Message {
	mi := &file_connectorCommand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandMessage.ProtoReflect.Descriptor instead.
func (*CommandMessage) Descriptor() ([]byte, []int) {
	return file_connectorCommand_proto_rawDescGZIP(), []int{0}
}

func (x *CommandMessage) GetSourceAggregator() string {
	if x != nil {
		return x.SourceAggregator
	}
	return ""
}

func (x *CommandMessage) GetSourceConnector() string {
	if x != nil {
		return x.SourceConnector
	}
	return ""
}

func (x *CommandMessage) GetSourceWorker() string {
	if x != nil {
		return x.SourceWorker
	}
	return ""
}

func (x *CommandMessage) GetDestinationAggregator() string {
	if x != nil {
		return x.DestinationAggregator
	}
	return ""
}

func (x *CommandMessage) GetDestinationConnector() string {
	if x != nil {
		return x.DestinationConnector
	}
	return ""
}

func (x *CommandMessage) GetDestinationWorker() string {
	if x != nil {
		return x.DestinationWorker
	}
	return ""
}

func (x *CommandMessage) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CommandMessage) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CommandMessage) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *CommandMessage) GetTimeout() string {
	if x != nil {
		return x.Timeout
	}
	return ""
}

func (x *CommandMessage) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *CommandMessage) GetMajor() int64 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *CommandMessage) GetMinor() int64 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *CommandMessage) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *CommandMessage) GetConnectorType() string {
	if x != nil {
		return x.ConnectorType
	}
	return ""
}

func (x *CommandMessage) GetCommandType() string {
	if x != nil {
		return x.CommandType
	}
	return ""
}

func (x *CommandMessage) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CommandMessage) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type CommandMessageUUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *CommandMessageUUID) Reset() {
	*x = CommandMessageUUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectorCommand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandMessageUUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandMessageUUID) ProtoMessage() {}

func (x *CommandMessageUUID) ProtoReflect() protoreflect.Message {
	mi := &file_connectorCommand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandMessageUUID.ProtoReflect.Descriptor instead.
func (*CommandMessageUUID) Descriptor() ([]byte, []int) {
	return file_connectorCommand_proto_rawDescGZIP(), []int{1}
}

func (x *CommandMessageUUID) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type CommandMessageWait struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerSource string `protobuf:"bytes,1,opt,name=WorkerSource,proto3" json:"WorkerSource,omitempty"`
	Value        string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	IteratorId   string `protobuf:"bytes,3,opt,name=IteratorId,proto3" json:"IteratorId,omitempty"`
	Major        int64  `protobuf:"varint,4,opt,name=Major,proto3" json:"Major,omitempty"`
	Minor        int64  `protobuf:"varint,5,opt,name=Minor,proto3" json:"Minor,omitempty"`
}

func (x *CommandMessageWait) Reset() {
	*x = CommandMessageWait{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectorCommand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandMessageWait) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandMessageWait) ProtoMessage() {}

func (x *CommandMessageWait) ProtoReflect() protoreflect.Message {
	mi := &file_connectorCommand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandMessageWait.ProtoReflect.Descriptor instead.
func (*CommandMessageWait) Descriptor() ([]byte, []int) {
	return file_connectorCommand_proto_rawDescGZIP(), []int{2}
}

func (x *CommandMessageWait) GetWorkerSource() string {
	if x != nil {
		return x.WorkerSource
	}
	return ""
}

func (x *CommandMessageWait) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *CommandMessageWait) GetIteratorId() string {
	if x != nil {
		return x.IteratorId
	}
	return ""
}

func (x *CommandMessageWait) GetMajor() int64 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *CommandMessageWait) GetMinor() int64 {
	if x != nil {
		return x.Minor
	}
	return 0
}

var File_connectorCommand_proto protoreflect.FileDescriptor

var file_connectorCommand_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x72, 0x70, 0x63, 0x55, 0x74,
	0x69, 0x6c, 0x73, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x04, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x22, 0x0a,
	0x0c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x12, 0x34, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x15, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x6a,
	0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22,
	0x9a, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x57, 0x61, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x32, 0xff, 0x01, 0x0a,
	0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x50, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x55, 0x74,
	0x69, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x55, 0x49,
	0x44, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x12, 0x57, 0x61, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x55, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x57, 0x61, 0x69, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x55,
	0x74, 0x69, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x49, 0x74, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x37,
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69, 0x74, 0x72, 0x69, 0x74, 0x2e, 0x67, 0x61, 0x6e,
	0x64, 0x61, 0x6c, 0x66, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x42, 0x15,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connectorCommand_proto_rawDescOnce sync.Once
	file_connectorCommand_proto_rawDescData = file_connectorCommand_proto_rawDesc
)

func file_connectorCommand_proto_rawDescGZIP() []byte {
	file_connectorCommand_proto_rawDescOnce.Do(func() {
		file_connectorCommand_proto_rawDescData = protoimpl.X.CompressGZIP(file_connectorCommand_proto_rawDescData)
	})
	return file_connectorCommand_proto_rawDescData
}

var file_connectorCommand_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_connectorCommand_proto_goTypes = []interface{}{
	(*CommandMessage)(nil),     // 0: grpcUtils.CommandMessage
	(*CommandMessageUUID)(nil), // 1: grpcUtils.CommandMessageUUID
	(*CommandMessageWait)(nil), // 2: grpcUtils.CommandMessageWait
	(*Empty)(nil),              // 3: grpcUtils.Empty
	(*IteratorMessage)(nil),    // 4: grpcUtils.IteratorMessage
}
var file_connectorCommand_proto_depIdxs = []int32{
	0, // 0: grpcUtils.ConnectorCommand.SendCommandMessage:input_type -> grpcUtils.CommandMessage
	2, // 1: grpcUtils.ConnectorCommand.WaitCommandMessage:input_type -> grpcUtils.CommandMessageWait
	3, // 2: grpcUtils.ConnectorCommand.CreateIteratorCommand:input_type -> grpcUtils.Empty
	1, // 3: grpcUtils.ConnectorCommand.SendCommandMessage:output_type -> grpcUtils.CommandMessageUUID
	0, // 4: grpcUtils.ConnectorCommand.WaitCommandMessage:output_type -> grpcUtils.CommandMessage
	4, // 5: grpcUtils.ConnectorCommand.CreateIteratorCommand:output_type -> grpcUtils.IteratorMessage
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_connectorCommand_proto_init() }
func file_connectorCommand_proto_init() {
	if File_connectorCommand_proto != nil {
		return
	}
	file_connector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_connectorCommand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_connectorCommand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandMessageUUID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_connectorCommand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandMessageWait); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_connectorCommand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connectorCommand_proto_goTypes,
		DependencyIndexes: file_connectorCommand_proto_depIdxs,
		MessageInfos:      file_connectorCommand_proto_msgTypes,
	}.Build()
	File_connectorCommand_proto = out.File
	file_connectorCommand_proto_rawDesc = nil
	file_connectorCommand_proto_goTypes = nil
	file_connectorCommand_proto_depIdxs = nil
}
