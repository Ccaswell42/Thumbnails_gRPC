// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/srvc.proto

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

// GetPictureClient is the client API for GetPicture service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetPictureClient interface {
	GetThumbnail(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type getPictureClient struct {
	cc grpc.ClientConnInterface
}

func NewGetPictureClient(cc grpc.ClientConnInterface) GetPictureClient {
	return &getPictureClient{cc}
}

func (c *getPictureClient) GetThumbnail(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.GetPicture/GetThumbnail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetPictureServer is the server API for GetPicture service.
// All implementations must embed UnimplementedGetPictureServer
// for forward compatibility
type GetPictureServer interface {
	GetThumbnail(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedGetPictureServer()
}

// UnimplementedGetPictureServer must be embedded to have forward compatible implementations.
type UnimplementedGetPictureServer struct {
}

func (UnimplementedGetPictureServer) GetThumbnail(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThumbnail not implemented")
}
func (UnimplementedGetPictureServer) mustEmbedUnimplementedGetPictureServer() {}

// UnsafeGetPictureServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetPictureServer will
// result in compilation errors.
type UnsafeGetPictureServer interface {
	mustEmbedUnimplementedGetPictureServer()
}

func RegisterGetPictureServer(s grpc.ServiceRegistrar, srv GetPictureServer) {
	s.RegisterService(&GetPicture_ServiceDesc, srv)
}

func _GetPicture_GetThumbnail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetPictureServer).GetThumbnail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GetPicture/GetThumbnail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetPictureServer).GetThumbnail(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GetPicture_ServiceDesc is the grpc.ServiceDesc for GetPicture service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetPicture_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GetPicture",
	HandlerType: (*GetPictureServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetThumbnail",
			Handler:    _GetPicture_GetThumbnail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/srvc.proto",
}
