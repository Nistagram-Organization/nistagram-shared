// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// MediaServiceClient is the client API for MediaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaServiceClient interface {
	SaveMedia(ctx context.Context, in *SaveMediaRequest, opts ...grpc.CallOption) (*SaveMediaResponse, error)
	GetMedia(ctx context.Context, in *GetMediaRequest, opts ...grpc.CallOption) (*GetMediaResponse, error)
}

type mediaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaServiceClient(cc grpc.ClientConnInterface) MediaServiceClient {
	return &mediaServiceClient{cc}
}

func (c *mediaServiceClient) SaveMedia(ctx context.Context, in *SaveMediaRequest, opts ...grpc.CallOption) (*SaveMediaResponse, error) {
	out := new(SaveMediaResponse)
	err := c.cc.Invoke(ctx, "/proto.MediaService/SaveMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) GetMedia(ctx context.Context, in *GetMediaRequest, opts ...grpc.CallOption) (*GetMediaResponse, error) {
	out := new(GetMediaResponse)
	err := c.cc.Invoke(ctx, "/proto.MediaService/GetMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaServiceServer is the server API for MediaService service.
// All implementations must embed UnimplementedMediaServiceServer
// for forward compatibility
type MediaServiceServer interface {
	SaveMedia(context.Context, *SaveMediaRequest) (*SaveMediaResponse, error)
	GetMedia(context.Context, *GetMediaRequest) (*GetMediaResponse, error)
	mustEmbedUnimplementedMediaServiceServer()
}

// UnimplementedMediaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMediaServiceServer struct {
}

func (UnimplementedMediaServiceServer) SaveMedia(context.Context, *SaveMediaRequest) (*SaveMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMedia not implemented")
}
func (UnimplementedMediaServiceServer) GetMedia(context.Context, *GetMediaRequest) (*GetMediaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedia not implemented")
}
func (UnimplementedMediaServiceServer) mustEmbedUnimplementedMediaServiceServer() {}

// UnsafeMediaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaServiceServer will
// result in compilation errors.
type UnsafeMediaServiceServer interface {
	mustEmbedUnimplementedMediaServiceServer()
}

func RegisterMediaServiceServer(s grpc.ServiceRegistrar, srv MediaServiceServer) {
	s.RegisterService(&MediaService_ServiceDesc, srv)
}

func _MediaService_SaveMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).SaveMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MediaService/SaveMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).SaveMedia(ctx, req.(*SaveMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_GetMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMediaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).GetMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MediaService/GetMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).GetMedia(ctx, req.(*GetMediaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MediaService_ServiceDesc is the grpc.ServiceDesc for MediaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MediaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MediaService",
	HandlerType: (*MediaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveMedia",
			Handler:    _MediaService_SaveMedia_Handler,
		},
		{
			MethodName: "GetMedia",
			Handler:    _MediaService_GetMedia_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "media_service.proto",
}
