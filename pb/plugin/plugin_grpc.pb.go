// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: plugin/plugin.proto

package plugin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PluginServiceClient is the client API for PluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginServiceClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (PluginService_CallClient, error)
	Directory(ctx context.Context, in *DirectoryRequest, opts ...grpc.CallOption) (*DirectoryResponse, error)
}

type pluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginServiceClient(cc grpc.ClientConnInterface) PluginServiceClient {
	return &pluginServiceClient{cc}
}

func (c *pluginServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/PluginService/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (PluginService_CallClient, error) {
	stream, err := c.cc.NewStream(ctx, &PluginService_ServiceDesc.Streams[0], "/PluginService/Call", opts...)
	if err != nil {
		return nil, err
	}
	x := &pluginServiceCallClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PluginService_CallClient interface {
	Recv() (*CallResponse, error)
	grpc.ClientStream
}

type pluginServiceCallClient struct {
	grpc.ClientStream
}

func (x *pluginServiceCallClient) Recv() (*CallResponse, error) {
	m := new(CallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pluginServiceClient) Directory(ctx context.Context, in *DirectoryRequest, opts ...grpc.CallOption) (*DirectoryResponse, error) {
	out := new(DirectoryResponse)
	err := c.cc.Invoke(ctx, "/PluginService/Directory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServiceServer is the server API for PluginService service.
// All implementations must embed UnimplementedPluginServiceServer
// for forward compatibility
type PluginServiceServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Call(*CallRequest, PluginService_CallServer) error
	Directory(context.Context, *DirectoryRequest) (*DirectoryResponse, error)
	mustEmbedUnimplementedPluginServiceServer()
}

// UnimplementedPluginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPluginServiceServer struct {
}

func (UnimplementedPluginServiceServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedPluginServiceServer) Call(*CallRequest, PluginService_CallServer) error {
	return status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedPluginServiceServer) Directory(context.Context, *DirectoryRequest) (*DirectoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Directory not implemented")
}
func (UnimplementedPluginServiceServer) mustEmbedUnimplementedPluginServiceServer() {}

// UnsafePluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServiceServer will
// result in compilation errors.
type UnsafePluginServiceServer interface {
	mustEmbedUnimplementedPluginServiceServer()
}

func RegisterPluginServiceServer(s grpc.ServiceRegistrar, srv PluginServiceServer) {
	s.RegisterService(&PluginService_ServiceDesc, srv)
}

func _PluginService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PluginService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_Call_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CallRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PluginServiceServer).Call(m, &pluginServiceCallServer{stream})
}

type PluginService_CallServer interface {
	Send(*CallResponse) error
	grpc.ServerStream
}

type pluginServiceCallServer struct {
	grpc.ServerStream
}

func (x *pluginServiceCallServer) Send(m *CallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PluginService_Directory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).Directory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PluginService/Directory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).Directory(ctx, req.(*DirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginService_ServiceDesc is the grpc.ServiceDesc for PluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PluginService",
	HandlerType: (*PluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _PluginService_Connect_Handler,
		},
		{
			MethodName: "Directory",
			Handler:    _PluginService_Directory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Call",
			Handler:       _PluginService_Call_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "plugin/plugin.proto",
}
