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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	GetUserEmail(ctx context.Context, in *GetUserEmailRequest, opts ...grpc.CallOption) (*GetUserEmailResponse, error)
	GetUsername(ctx context.Context, in *GetUsernameRequest, opts ...grpc.CallOption) (*GetUsernameResponse, error)
	CheckIfPostIsInFavorites(ctx context.Context, in *CheckFavoritesRequest, opts ...grpc.CallOption) (*CheckFavoritesResponse, error)
	CheckIfUserIsTaggable(ctx context.Context, in *CheckTaggableRequest, opts ...grpc.CallOption) (*CheckTaggableResponse, error)
	GetFollowingUsers(ctx context.Context, in *GetFollowingUsersRequest, opts ...grpc.CallOption) (UserService_GetFollowingUsersClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserEmail(ctx context.Context, in *GetUserEmailRequest, opts ...grpc.CallOption) (*GetUserEmailResponse, error) {
	out := new(GetUserEmailResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetUserEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsername(ctx context.Context, in *GetUsernameRequest, opts ...grpc.CallOption) (*GetUsernameResponse, error) {
	out := new(GetUsernameResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckIfPostIsInFavorites(ctx context.Context, in *CheckFavoritesRequest, opts ...grpc.CallOption) (*CheckFavoritesResponse, error) {
	out := new(CheckFavoritesResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/CheckIfPostIsInFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckIfUserIsTaggable(ctx context.Context, in *CheckTaggableRequest, opts ...grpc.CallOption) (*CheckTaggableResponse, error) {
	out := new(CheckTaggableResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/CheckIfUserIsTaggable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetFollowingUsers(ctx context.Context, in *GetFollowingUsersRequest, opts ...grpc.CallOption) (UserService_GetFollowingUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/proto.UserService/GetFollowingUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetFollowingUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetFollowingUsersClient interface {
	Recv() (*GetFollowingUsersResponse, error)
	grpc.ClientStream
}

type userServiceGetFollowingUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceGetFollowingUsersClient) Recv() (*GetFollowingUsersResponse, error) {
	m := new(GetFollowingUsersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetUserEmail(context.Context, *GetUserEmailRequest) (*GetUserEmailResponse, error)
	GetUsername(context.Context, *GetUsernameRequest) (*GetUsernameResponse, error)
	CheckIfPostIsInFavorites(context.Context, *CheckFavoritesRequest) (*CheckFavoritesResponse, error)
	CheckIfUserIsTaggable(context.Context, *CheckTaggableRequest) (*CheckTaggableResponse, error)
	GetFollowingUsers(*GetFollowingUsersRequest, UserService_GetFollowingUsersServer) error
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserEmail(context.Context, *GetUserEmailRequest) (*GetUserEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEmail not implemented")
}
func (UnimplementedUserServiceServer) GetUsername(context.Context, *GetUsernameRequest) (*GetUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsername not implemented")
}
func (UnimplementedUserServiceServer) CheckIfPostIsInFavorites(context.Context, *CheckFavoritesRequest) (*CheckFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfPostIsInFavorites not implemented")
}
func (UnimplementedUserServiceServer) CheckIfUserIsTaggable(context.Context, *CheckTaggableRequest) (*CheckTaggableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfUserIsTaggable not implemented")
}
func (UnimplementedUserServiceServer) GetFollowingUsers(*GetFollowingUsersRequest, UserService_GetFollowingUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFollowingUsers not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetUserEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserEmail(ctx, req.(*GetUserEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsername(ctx, req.(*GetUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckIfPostIsInFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckFavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckIfPostIsInFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/CheckIfPostIsInFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckIfPostIsInFavorites(ctx, req.(*CheckFavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckIfUserIsTaggable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTaggableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckIfUserIsTaggable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/CheckIfUserIsTaggable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckIfUserIsTaggable(ctx, req.(*CheckTaggableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetFollowingUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFollowingUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetFollowingUsers(m, &userServiceGetFollowingUsersServer{stream})
}

type UserService_GetFollowingUsersServer interface {
	Send(*GetFollowingUsersResponse) error
	grpc.ServerStream
}

type userServiceGetFollowingUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceGetFollowingUsersServer) Send(m *GetFollowingUsersResponse) error {
	return x.ServerStream.SendMsg(m)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "GetUserEmail",
			Handler:    _UserService_GetUserEmail_Handler,
		},
		{
			MethodName: "GetUsername",
			Handler:    _UserService_GetUsername_Handler,
		},
		{
			MethodName: "CheckIfPostIsInFavorites",
			Handler:    _UserService_CheckIfPostIsInFavorites_Handler,
		},
		{
			MethodName: "CheckIfUserIsTaggable",
			Handler:    _UserService_CheckIfUserIsTaggable_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFollowingUsers",
			Handler:       _UserService_GetFollowingUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user_service.proto",
}
