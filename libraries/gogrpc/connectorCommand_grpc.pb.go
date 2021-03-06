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

// ConnectorCommandClient is the client API for ConnectorCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorCommandClient interface {
	SendCommandMessage(ctx context.Context, in *CommandMessage, opts ...grpc.CallOption) (*CommandMessageUUID, error)
	WaitCommandMessage(ctx context.Context, in *CommandMessageWait, opts ...grpc.CallOption) (*CommandMessage, error)
	CreateIteratorCommand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IteratorMessage, error)
}

type connectorCommandClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorCommandClient(cc grpc.ClientConnInterface) ConnectorCommandClient {
	return &connectorCommandClient{cc}
}

func (c *connectorCommandClient) SendCommandMessage(ctx context.Context, in *CommandMessage, opts ...grpc.CallOption) (*CommandMessageUUID, error) {
	out := new(CommandMessageUUID)
	err := c.cc.Invoke(ctx, "/gogrpc.ConnectorCommand/SendCommandMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorCommandClient) WaitCommandMessage(ctx context.Context, in *CommandMessageWait, opts ...grpc.CallOption) (*CommandMessage, error) {
	out := new(CommandMessage)
	err := c.cc.Invoke(ctx, "/gogrpc.ConnectorCommand/WaitCommandMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorCommandClient) CreateIteratorCommand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IteratorMessage, error) {
	out := new(IteratorMessage)
	err := c.cc.Invoke(ctx, "/gogrpc.ConnectorCommand/CreateIteratorCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorCommandServer is the server API for ConnectorCommand service.
// All implementations must embed UnimplementedConnectorCommandServer
// for forward compatibility
type ConnectorCommandServer interface {
	SendCommandMessage(context.Context, *CommandMessage) (*CommandMessageUUID, error)
	WaitCommandMessage(context.Context, *CommandMessageWait) (*CommandMessage, error)
	CreateIteratorCommand(context.Context, *Empty) (*IteratorMessage, error)
	mustEmbedUnimplementedConnectorCommandServer()
}

// UnimplementedConnectorCommandServer must be embedded to have forward compatible implementations.
type UnimplementedConnectorCommandServer struct {
}

func (UnimplementedConnectorCommandServer) SendCommandMessage(context.Context, *CommandMessage) (*CommandMessageUUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommandMessage not implemented")
}
func (UnimplementedConnectorCommandServer) WaitCommandMessage(context.Context, *CommandMessageWait) (*CommandMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitCommandMessage not implemented")
}
func (UnimplementedConnectorCommandServer) CreateIteratorCommand(context.Context, *Empty) (*IteratorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIteratorCommand not implemented")
}
func (UnimplementedConnectorCommandServer) mustEmbedUnimplementedConnectorCommandServer() {}

// UnsafeConnectorCommandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorCommandServer will
// result in compilation errors.
type UnsafeConnectorCommandServer interface {
	mustEmbedUnimplementedConnectorCommandServer()
}

func RegisterConnectorCommandServer(s *grpc.Server, srv ConnectorCommandServer) {
	s.RegisterService(&_ConnectorCommand_serviceDesc, srv)
}

func _ConnectorCommand_SendCommandMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorCommandServer).SendCommandMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogrpc.ConnectorCommand/SendCommandMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorCommandServer).SendCommandMessage(ctx, req.(*CommandMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorCommand_WaitCommandMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandMessageWait)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorCommandServer).WaitCommandMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogrpc.ConnectorCommand/WaitCommandMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorCommandServer).WaitCommandMessage(ctx, req.(*CommandMessageWait))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorCommand_CreateIteratorCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorCommandServer).CreateIteratorCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogrpc.ConnectorCommand/CreateIteratorCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorCommandServer).CreateIteratorCommand(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConnectorCommand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gogrpc.ConnectorCommand",
	HandlerType: (*ConnectorCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCommandMessage",
			Handler:    _ConnectorCommand_SendCommandMessage_Handler,
		},
		{
			MethodName: "WaitCommandMessage",
			Handler:    _ConnectorCommand_WaitCommandMessage_Handler,
		},
		{
			MethodName: "CreateIteratorCommand",
			Handler:    _ConnectorCommand_CreateIteratorCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connectorCommand.proto",
}
