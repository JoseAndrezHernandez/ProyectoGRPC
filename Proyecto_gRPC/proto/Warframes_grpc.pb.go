// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: proto/Warframes.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WarframeService_GetWaframeInfo_FullMethodName     = "/Warframe.WarframeService/GetWaframeInfo"
	WarframeService_GetWarameList_FullMethodName      = "/Warframe.WarframeService/getWarameList"
	WarframeService_AddWarframes_FullMethodName       = "/Warframe.WarframeService/addWarframes"
	WarframeService_GetWarframesByType_FullMethodName = "/Warframe.WarframeService/GetWarframesByType"
)

// WarframeServiceClient is the client API for WarframeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WarframeServiceClient interface {
	GetWaframeInfo(ctx context.Context, in *WarframeRequest, opts ...grpc.CallOption) (*WarframeResponse, error)
	GetWarameList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WarframeResponse], error)
	AddWarframes(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[NewWarframeRequest, AddWarframesResponse], error)
	GetWarframesByType(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[WarframeTimeRequest, WarframeResponse], error)
}

type warframeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWarframeServiceClient(cc grpc.ClientConnInterface) WarframeServiceClient {
	return &warframeServiceClient{cc}
}

func (c *warframeServiceClient) GetWaframeInfo(ctx context.Context, in *WarframeRequest, opts ...grpc.CallOption) (*WarframeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WarframeResponse)
	err := c.cc.Invoke(ctx, WarframeService_GetWaframeInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warframeServiceClient) GetWarameList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WarframeResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WarframeService_ServiceDesc.Streams[0], WarframeService_GetWarameList_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Empty, WarframeResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WarframeService_GetWarameListClient = grpc.ServerStreamingClient[WarframeResponse]

func (c *warframeServiceClient) AddWarframes(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[NewWarframeRequest, AddWarframesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WarframeService_ServiceDesc.Streams[1], WarframeService_AddWarframes_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[NewWarframeRequest, AddWarframesResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WarframeService_AddWarframesClient = grpc.ClientStreamingClient[NewWarframeRequest, AddWarframesResponse]

func (c *warframeServiceClient) GetWarframesByType(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[WarframeTimeRequest, WarframeResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WarframeService_ServiceDesc.Streams[2], WarframeService_GetWarframesByType_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WarframeTimeRequest, WarframeResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WarframeService_GetWarframesByTypeClient = grpc.BidiStreamingClient[WarframeTimeRequest, WarframeResponse]

// WarframeServiceServer is the server API for WarframeService service.
// All implementations must embed UnimplementedWarframeServiceServer
// for forward compatibility.
type WarframeServiceServer interface {
	GetWaframeInfo(context.Context, *WarframeRequest) (*WarframeResponse, error)
	GetWarameList(*Empty, grpc.ServerStreamingServer[WarframeResponse]) error
	AddWarframes(grpc.ClientStreamingServer[NewWarframeRequest, AddWarframesResponse]) error
	GetWarframesByType(grpc.BidiStreamingServer[WarframeTimeRequest, WarframeResponse]) error
	mustEmbedUnimplementedWarframeServiceServer()
}

// UnimplementedWarframeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWarframeServiceServer struct{}

func (UnimplementedWarframeServiceServer) GetWaframeInfo(context.Context, *WarframeRequest) (*WarframeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWaframeInfo not implemented")
}
func (UnimplementedWarframeServiceServer) GetWarameList(*Empty, grpc.ServerStreamingServer[WarframeResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetWarameList not implemented")
}
func (UnimplementedWarframeServiceServer) AddWarframes(grpc.ClientStreamingServer[NewWarframeRequest, AddWarframesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method AddWarframes not implemented")
}
func (UnimplementedWarframeServiceServer) GetWarframesByType(grpc.BidiStreamingServer[WarframeTimeRequest, WarframeResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetWarframesByType not implemented")
}
func (UnimplementedWarframeServiceServer) mustEmbedUnimplementedWarframeServiceServer() {}
func (UnimplementedWarframeServiceServer) testEmbeddedByValue()                         {}

// UnsafeWarframeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WarframeServiceServer will
// result in compilation errors.
type UnsafeWarframeServiceServer interface {
	mustEmbedUnimplementedWarframeServiceServer()
}

func RegisterWarframeServiceServer(s grpc.ServiceRegistrar, srv WarframeServiceServer) {
	// If the following call pancis, it indicates UnimplementedWarframeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WarframeService_ServiceDesc, srv)
}

func _WarframeService_GetWaframeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WarframeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarframeServiceServer).GetWaframeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WarframeService_GetWaframeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarframeServiceServer).GetWaframeInfo(ctx, req.(*WarframeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WarframeService_GetWarameList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WarframeServiceServer).GetWarameList(m, &grpc.GenericServerStream[Empty, WarframeResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WarframeService_GetWarameListServer = grpc.ServerStreamingServer[WarframeResponse]

func _WarframeService_AddWarframes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WarframeServiceServer).AddWarframes(&grpc.GenericServerStream[NewWarframeRequest, AddWarframesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WarframeService_AddWarframesServer = grpc.ClientStreamingServer[NewWarframeRequest, AddWarframesResponse]

func _WarframeService_GetWarframesByType_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WarframeServiceServer).GetWarframesByType(&grpc.GenericServerStream[WarframeTimeRequest, WarframeResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WarframeService_GetWarframesByTypeServer = grpc.BidiStreamingServer[WarframeTimeRequest, WarframeResponse]

// WarframeService_ServiceDesc is the grpc.ServiceDesc for WarframeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WarframeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Warframe.WarframeService",
	HandlerType: (*WarframeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWaframeInfo",
			Handler:    _WarframeService_GetWaframeInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getWarameList",
			Handler:       _WarframeService_GetWarameList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "addWarframes",
			Handler:       _WarframeService_AddWarframes_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetWarframesByType",
			Handler:       _WarframeService_GetWarframesByType_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/Warframes.proto",
}