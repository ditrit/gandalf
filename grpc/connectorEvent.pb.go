// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connectorEvent.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EventMessage struct {
	Tenant               string   `protobuf:"bytes,1,opt,name=Tenant,json=tenant,proto3" json:"Tenant,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,json=token,proto3" json:"Token,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=Topic,json=topic,proto3" json:"Topic,omitempty"`
	Timeout              string   `protobuf:"bytes,4,opt,name=Timeout,json=timeout,proto3" json:"Timeout,omitempty"`
	Timestamp            string   `protobuf:"bytes,5,opt,name=Timestamp,json=timestamp,proto3" json:"Timestamp,omitempty"`
	Uuid                 string   `protobuf:"bytes,6,opt,name=Uuid,json=uuid,proto3" json:"Uuid,omitempty"`
	Event                string   `protobuf:"bytes,7,opt,name=Event,json=event,proto3" json:"Event,omitempty"`
	Payload              string   `protobuf:"bytes,8,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventMessage) Reset()         { *m = EventMessage{} }
func (m *EventMessage) String() string { return proto.CompactTextString(m) }
func (*EventMessage) ProtoMessage()    {}
func (*EventMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15ac99c650ec24e, []int{0}
}

func (m *EventMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventMessage.Unmarshal(m, b)
}
func (m *EventMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventMessage.Marshal(b, m, deterministic)
}
func (m *EventMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMessage.Merge(m, src)
}
func (m *EventMessage) XXX_Size() int {
	return xxx_messageInfo_EventMessage.Size(m)
}
func (m *EventMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EventMessage proto.InternalMessageInfo

func (m *EventMessage) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *EventMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *EventMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *EventMessage) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

func (m *EventMessage) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *EventMessage) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *EventMessage) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EventMessage) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type EventMessageWait struct {
	WorkerSource         string   `protobuf:"bytes,1,opt,name=WorkerSource,json=workerSource,proto3" json:"WorkerSource,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=Event,json=event,proto3" json:"Event,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=Topic,json=topic,proto3" json:"Topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventMessageWait) Reset()         { *m = EventMessageWait{} }
func (m *EventMessageWait) String() string { return proto.CompactTextString(m) }
func (*EventMessageWait) ProtoMessage()    {}
func (*EventMessageWait) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15ac99c650ec24e, []int{1}
}

func (m *EventMessageWait) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventMessageWait.Unmarshal(m, b)
}
func (m *EventMessageWait) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventMessageWait.Marshal(b, m, deterministic)
}
func (m *EventMessageWait) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMessageWait.Merge(m, src)
}
func (m *EventMessageWait) XXX_Size() int {
	return xxx_messageInfo_EventMessageWait.Size(m)
}
func (m *EventMessageWait) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMessageWait.DiscardUnknown(m)
}

var xxx_messageInfo_EventMessageWait proto.InternalMessageInfo

func (m *EventMessageWait) GetWorkerSource() string {
	if m != nil {
		return m.WorkerSource
	}
	return ""
}

func (m *EventMessageWait) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EventMessageWait) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*EventMessage)(nil), "grpc.EventMessage")
	proto.RegisterType((*EventMessageWait)(nil), "grpc.EventMessageWait")
}

func init() { proto.RegisterFile("connectorEvent.proto", fileDescriptor_a15ac99c650ec24e) }

var fileDescriptor_a15ac99c650ec24e = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x29, 0x96, 0x22, 0x23, 0x1a, 0xb2, 0x21, 0x66, 0xd3, 0x78, 0x30, 0x3d, 0x79, 0xea,
	0x41, 0xe3, 0xdd, 0x60, 0x3c, 0x9a, 0x10, 0xc1, 0x70, 0x5e, 0xda, 0x09, 0xd9, 0x40, 0x77, 0x37,
	0xdb, 0xa9, 0xca, 0x4f, 0xf0, 0xcf, 0xf9, 0x9b, 0xcc, 0xee, 0x82, 0x94, 0xc8, 0xf1, 0x7d, 0x6f,
	0xdb, 0x37, 0x33, 0x0f, 0xc6, 0x85, 0x56, 0x0a, 0x0b, 0xd2, 0xf6, 0xe5, 0x03, 0x15, 0xe5, 0xc6,
	0x6a, 0xd2, 0x2c, 0x5e, 0x59, 0x53, 0xa4, 0x97, 0x15, 0xd6, 0xb5, 0x58, 0x61, 0x80, 0xd9, 0x4f,
	0x04, 0x43, 0xff, 0xe8, 0x35, 0x60, 0x76, 0x0d, 0xc9, 0x1c, 0x95, 0x50, 0xc4, 0xa3, 0xdb, 0xe8,
	0x6e, 0xf0, 0x96, 0x90, 0x57, 0x6c, 0x0c, 0xbd, 0xb9, 0x5e, 0xa3, 0xe2, 0x5d, 0x8f, 0x7b, 0xe4,
	0x44, 0xa0, 0x46, 0x16, 0xfc, 0x6c, 0x4f, 0x8d, 0x2c, 0x18, 0x87, 0xfe, 0x5c, 0x56, 0xa8, 0x1b,
	0xe2, 0xb1, 0xe7, 0x7d, 0x0a, 0x92, 0xdd, 0xc0, 0xc0, 0x39, 0x35, 0x89, 0xca, 0xf0, 0x9e, 0xf7,
	0x06, 0xb4, 0x07, 0x8c, 0x41, 0xfc, 0xde, 0xc8, 0x92, 0x27, 0xde, 0x88, 0x9b, 0x46, 0x96, 0x2e,
	0xc1, 0xcf, 0xc7, 0xfb, 0x21, 0x01, 0x9d, 0x70, 0x09, 0x53, 0xb1, 0xdd, 0x68, 0x51, 0xf2, 0xf3,
	0x90, 0x60, 0x82, 0xcc, 0x96, 0x30, 0x6a, 0xef, 0xb3, 0x10, 0x92, 0x58, 0x06, 0xc3, 0x85, 0xb6,
	0x6b, 0xb4, 0x33, 0xdd, 0xd8, 0x02, 0x77, 0x9b, 0x0d, 0x3f, 0x5b, 0xec, 0x90, 0xd3, 0x6d, 0xe7,
	0x9c, 0xdc, 0xef, 0xfe, 0x3b, 0x82, 0xab, 0xe7, 0xa3, 0x13, 0xb3, 0x47, 0x18, 0xcd, 0x50, 0x95,
	0x47, 0xa7, 0x64, 0xb9, 0xbb, 0x78, 0xde, 0x66, 0xe9, 0xc5, 0x8e, 0x55, 0x86, 0xb6, 0x59, 0x87,
	0x3d, 0xc1, 0xc8, 0x4d, 0x78, 0xdc, 0xc0, 0xff, 0xcf, 0xdc, 0x9b, 0xf4, 0xc4, 0xef, 0xb2, 0xce,
	0x24, 0x87, 0x54, 0xea, 0xe0, 0xe0, 0x97, 0xa8, 0xcc, 0x06, 0xeb, 0xfc, 0xaf, 0xfe, 0xc9, 0x61,
	0xcc, 0xa9, 0xab, 0x7b, 0x1a, 0x2d, 0x13, 0xdf, 0xfb, 0xc3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x58, 0x08, 0x48, 0x87, 0x24, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConnectorEventClient is the client API for ConnectorEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectorEventClient interface {
	SendEventMessage(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (*Empty, error)
	WaitEventMessage(ctx context.Context, in *EventMessageWait, opts ...grpc.CallOption) (*EventMessage, error)
}

type connectorEventClient struct {
	cc *grpc.ClientConn
}

func NewConnectorEventClient(cc *grpc.ClientConn) ConnectorEventClient {
	return &connectorEventClient{cc}
}

func (c *connectorEventClient) SendEventMessage(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorEvent/SendEventMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorEventClient) WaitEventMessage(ctx context.Context, in *EventMessageWait, opts ...grpc.CallOption) (*EventMessage, error) {
	out := new(EventMessage)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorEvent/WaitEventMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorEventServer is the server API for ConnectorEvent service.
type ConnectorEventServer interface {
	SendEventMessage(context.Context, *EventMessage) (*Empty, error)
	WaitEventMessage(context.Context, *EventMessageWait) (*EventMessage, error)
}

// UnimplementedConnectorEventServer can be embedded to have forward compatible implementations.
type UnimplementedConnectorEventServer struct {
}

func (*UnimplementedConnectorEventServer) SendEventMessage(ctx context.Context, req *EventMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEventMessage not implemented")
}
func (*UnimplementedConnectorEventServer) WaitEventMessage(ctx context.Context, req *EventMessageWait) (*EventMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitEventMessage not implemented")
}

func RegisterConnectorEventServer(s *grpc.Server, srv ConnectorEventServer) {
	s.RegisterService(&_ConnectorEvent_serviceDesc, srv)
}

func _ConnectorEvent_SendEventMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).SendEventMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorEvent/SendEventMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).SendEventMessage(ctx, req.(*EventMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorEvent_WaitEventMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventMessageWait)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).WaitEventMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorEvent/WaitEventMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).WaitEventMessage(ctx, req.(*EventMessageWait))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConnectorEvent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ConnectorEvent",
	HandlerType: (*ConnectorEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEventMessage",
			Handler:    _ConnectorEvent_SendEventMessage_Handler,
		},
		{
			MethodName: "WaitEventMessage",
			Handler:    _ConnectorEvent_WaitEventMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connectorEvent.proto",
}
