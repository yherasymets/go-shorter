// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: shorter.proto

package proto

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

const (
	Shorter_Create_FullMethodName = "/proto.Shorter/Create"
	Shorter_Get_FullMethodName    = "/proto.Shorter/Get"
)

// ShorterClient is the client API for Shorter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShorterClient interface {
	Create(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error)
	Get(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error)
}

type shorterClient struct {
	cc grpc.ClientConnInterface
}

func NewShorterClient(cc grpc.ClientConnInterface) ShorterClient {
	return &shorterClient{cc}
}

func (c *shorterClient) Create(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error) {
	out := new(UrlResponse)
	err := c.cc.Invoke(ctx, Shorter_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shorterClient) Get(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error) {
	out := new(UrlResponse)
	err := c.cc.Invoke(ctx, Shorter_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShorterServer is the server API for Shorter service.
// All implementations must embed UnimplementedShorterServer
// for forward compatibility
type ShorterServer interface {
	Create(context.Context, *UrlRequest) (*UrlResponse, error)
	Get(context.Context, *UrlRequest) (*UrlResponse, error)
	mustEmbedUnimplementedShorterServer()
}

// UnimplementedShorterServer must be embedded to have forward compatible implementations.
type UnimplementedShorterServer struct {
}

func (UnimplementedShorterServer) Create(context.Context, *UrlRequest) (*UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedShorterServer) Get(context.Context, *UrlRequest) (*UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedShorterServer) mustEmbedUnimplementedShorterServer() {}

// UnsafeShorterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShorterServer will
// result in compilation errors.
type UnsafeShorterServer interface {
	mustEmbedUnimplementedShorterServer()
}

func RegisterShorterServer(s grpc.ServiceRegistrar, srv ShorterServer) {
	s.RegisterService(&Shorter_ServiceDesc, srv)
}

func _Shorter_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShorterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shorter_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShorterServer).Create(ctx, req.(*UrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shorter_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShorterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Shorter_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShorterServer).Get(ctx, req.(*UrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shorter_ServiceDesc is the grpc.ServiceDesc for Shorter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shorter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Shorter",
	HandlerType: (*ShorterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Shorter_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Shorter_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shorter.proto",
}