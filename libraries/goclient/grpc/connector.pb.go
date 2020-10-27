// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: connector.proto

package grpc

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

type CommandList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major    int64    `protobuf:"varint,1,opt,name=Major,proto3" json:"Major,omitempty"`
	Commands []string `protobuf:"bytes,2,rep,name=Commands,proto3" json:"Commands,omitempty"`
}

func (x *CommandList) Reset() {
	*x = CommandList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandList) ProtoMessage() {}

func (x *CommandList) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandList.ProtoReflect.Descriptor instead.
func (*CommandList) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{0}
}

func (x *CommandList) GetMajor() int64 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *CommandList) GetCommands() []string {
	if x != nil {
		return x.Commands
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{1}
}

type IteratorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *IteratorMessage) Reset() {
	*x = IteratorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IteratorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IteratorMessage) ProtoMessage() {}

func (x *IteratorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IteratorMessage.ProtoReflect.Descriptor instead.
func (*IteratorMessage) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{2}
}

func (x *IteratorMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_connector_proto protoreflect.FileDescriptor

var file_connector_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x3f, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x21, 0x0a, 0x0f, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x64, 0x32, 0x40, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x33, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x30, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x69,
	0x74, 0x72, 0x69, 0x74, 0x2e, 0x67, 0x61, 0x6e, 0x64, 0x61, 0x6c, 0x66, 0x2e, 0x6a, 0x61, 0x76,
	0x61, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x42, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connector_proto_rawDescOnce sync.Once
	file_connector_proto_rawDescData = file_connector_proto_rawDesc
)

func file_connector_proto_rawDescGZIP() []byte {
	file_connector_proto_rawDescOnce.Do(func() {
		file_connector_proto_rawDescData = protoimpl.X.CompressGZIP(file_connector_proto_rawDescData)
	})
	return file_connector_proto_rawDescData
}

var file_connector_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_connector_proto_goTypes = []interface{}{
	(*CommandList)(nil),     // 0: grpc.CommandList
	(*Empty)(nil),           // 1: grpc.Empty
	(*IteratorMessage)(nil), // 2: grpc.IteratorMessage
}
var file_connector_proto_depIdxs = []int32{
	0, // 0: grpc.Connector.SendCommandList:input_type -> grpc.CommandList
	1, // 1: grpc.Connector.SendCommandList:output_type -> grpc.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_connector_proto_init() }
func file_connector_proto_init() {
	if File_connector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandList); i {
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
		file_connector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_connector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IteratorMessage); i {
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
			RawDescriptor: file_connector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connector_proto_goTypes,
		DependencyIndexes: file_connector_proto_depIdxs,
		MessageInfos:      file_connector_proto_msgTypes,
	}.Build()
	File_connector_proto = out.File
	file_connector_proto_rawDesc = nil
	file_connector_proto_goTypes = nil
	file_connector_proto_depIdxs = nil
}
