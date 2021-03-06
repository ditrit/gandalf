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

// ConnectorClient is the client API for Connector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorClient interface {
	SendCommandList(ctx context.Context, in *CommandList, opts ...grpc.CallOption) (*Validate, error)
	SendStop(ctx context.Context, in *Stop, opts ...grpc.CallOption) (*Validate, error)
}

type connectorClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorClient(cc grpc.ClientConnInterface) ConnectorClient {
	return &connectorClient{cc}
}

func (c *connectorClient) SendCommandList(ctx context.Context, in *CommandList, opts ...grpc.CallOption) (*Validate, error) {
	out := new(Validate)
	err := c.cc.Invoke(ctx, "/gogrpc.Connector/SendCommandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) SendStop(ctx context.Context, in *Stop, opts ...grpc.CallOption) (*Validate, error) {
	out := new(Validate)
	err := c.cc.Invoke(ctx, "/gogrpc.Connector/SendStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorServer is the server API for Connector service.
// All implementations must embed UnimplementedConnectorServer
// for forward compatibility
type ConnectorServer interface {
	SendCommandList(context.Context, *CommandList) (*Validate, error)
	SendStop(context.Context, *Stop) (*Validate, error)
	mustEmbedUnimplementedConnectorServer()
}

// UnimplementedConnectorServer must be embedded to have forward compatible implementations.
type UnimplementedConnectorServer struct {
}

func (UnimplementedConnectorServer) SendCommandList(context.Context, *CommandList) (*Validate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommandList not implemented")
}
func (UnimplementedConnectorServer) SendStop(context.Context, *Stop) (*Validate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendStop not implemented")
}
func (UnimplementedConnectorServer) mustEmbedUnimplementedConnectorServer() {}

// UnsafeConnectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorServer will
// result in compilation errors.
type UnsafeConnectorServer interface {
	mustEmbedUnimplementedConnectorServer()
}

func RegisterConnectorServer(s *grpc.Server, srv ConnectorServer) {
	s.RegisterService(&_Connector_serviceDesc, srv)
}

func _Connector_SendCommandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).SendCommandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogrpc.Connector/SendCommandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).SendCommandList(ctx, req.(*CommandList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_SendStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).SendStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogrpc.Connector/SendStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).SendStop(ctx, req.(*Stop))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gogrpc.Connector",
	HandlerType: (*ConnectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCommandList",
			Handler:    _Connector_SendCommandList_Handler,
		},
		{
			MethodName: "SendStop",
			Handler:    _Connector_SendStop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connector.proto",
}
