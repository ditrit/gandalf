// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gogrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ConnectorEventClient is the client API for ConnectorEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorEventClient interface {
	SendEventMessage(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (*Empty, error)
	WaitEventMessage(ctx context.Context, in *EventMessageWait, opts ...grpc.CallOption) (*EventMessage, error)
	WaitTopicMessage(ctx context.Context, in *TopicMessageWait, opts ...grpc.CallOption) (*EventMessage, error)
	CreateIteratorEvent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IteratorMessage, error)
}

type connectorEventClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorEventClient(cc grpc.ClientConnInterface) ConnectorEventClient {
	return &connectorEventClient{cc}
}

func (c *connectorEventClient) SendEventMessage(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/gogrpc.ConnectorEvent/SendEventMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorEventClient) WaitEventMessage(ctx context.Context, in *EventMessageWait, opts ...grpc.CallOption) (*EventMessage, error) {
	out := new(EventMessage)
	err := c.cc.Invoke(ctx, "/gogrpc.ConnectorEvent/WaitEventMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorEventClient) WaitTopicMessage(ctx context.Context, in *TopicMessageWait, opts ...grpc.CallOption) (*EventMessage, error) {
	out := new(EventMessage)
	err := c.cc.Invoke(ctx, "/gogrpc.ConnectorEvent/WaitTopicMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorEventClient) CreateIteratorEvent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IteratorMessage, error) {
	out := new(IteratorMessage)
	err := c.cc.Invoke(ctx, "/gogrpc.ConnectorEvent/CreateIteratorEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorEventServer is the server API for ConnectorEvent service.
// All implementations must embed UnimplementedConnectorEventServer
// for forward compatibility
type ConnectorEventServer interface {
	SendEventMessage(context.Context, *EventMessage) (*Empty, error)
	WaitEventMessage(context.Context, *EventMessageWait) (*EventMessage, error)
	WaitTopicMessage(context.Context, *TopicMessageWait) (*EventMessage, error)
	CreateIteratorEvent(context.Context, *Empty) (*IteratorMessage, error)
	mustEmbedUnimplementedConnectorEventServer()
}

// UnimplementedConnectorEventServer must be embedded to have forward compatible implementations.
type UnimplementedConnectorEventServer struct {
}

func (UnimplementedConnectorEventServer) SendEventMessage(context.Context, *EventMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEventMessage not implemented")
}
func (UnimplementedConnectorEventServer) WaitEventMessage(context.Context, *EventMessageWait) (*EventMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitEventMessage not implemented")
}
func (UnimplementedConnectorEventServer) WaitTopicMessage(context.Context, *TopicMessageWait) (*EventMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitTopicMessage not implemented")
}
func (UnimplementedConnectorEventServer) CreateIteratorEvent(context.Context, *Empty) (*IteratorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIteratorEvent not implemented")
}
func (UnimplementedConnectorEventServer) mustEmbedUnimplementedConnectorEventServer() {}

// UnsafeConnectorEventServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorEventServer will
// result in compilation errors.
type UnsafeConnectorEventServer interface {
	mustEmbedUnimplementedConnectorEventServer()
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
		FullMethod: "/gogrpc.ConnectorEvent/SendEventMessage",
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
		FullMethod: "/gogrpc.ConnectorEvent/WaitEventMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).WaitEventMessage(ctx, req.(*EventMessageWait))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorEvent_WaitTopicMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicMessageWait)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).WaitTopicMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogrpc.ConnectorEvent/WaitTopicMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).WaitTopicMessage(ctx, req.(*TopicMessageWait))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorEvent_CreateIteratorEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).CreateIteratorEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogrpc.ConnectorEvent/CreateIteratorEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).CreateIteratorEvent(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConnectorEvent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gogrpc.ConnectorEvent",
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
		{
			MethodName: "WaitTopicMessage",
			Handler:    _ConnectorEvent_WaitTopicMessage_Handler,
		},
		{
			MethodName: "CreateIteratorEvent",
			Handler:    _ConnectorEvent_CreateIteratorEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connectorEvent.proto",
}
