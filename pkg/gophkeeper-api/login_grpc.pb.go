// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: login.proto

package api

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

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServiceClient interface {
	Create(ctx context.Context, in *CreateLoginRequest, opts ...grpc.CallOption) (*Login, error)
	Update(ctx context.Context, in *UpdateLoginRequest, opts ...grpc.CallOption) (*Login, error)
	Get(ctx context.Context, in *GetLoginRequest, opts ...grpc.CallOption) (*Login, error)
	ListAvailableLogins(ctx context.Context, in *ListAvailableLoginsRequest, opts ...grpc.CallOption) (*ListAvailableLoginsResponse, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Create(ctx context.Context, in *CreateLoginRequest, opts ...grpc.CallOption) (*Login, error) {
	out := new(Login)
	err := c.cc.Invoke(ctx, "/gophkeeper.v1.LoginService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) Update(ctx context.Context, in *UpdateLoginRequest, opts ...grpc.CallOption) (*Login, error) {
	out := new(Login)
	err := c.cc.Invoke(ctx, "/gophkeeper.v1.LoginService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) Get(ctx context.Context, in *GetLoginRequest, opts ...grpc.CallOption) (*Login, error) {
	out := new(Login)
	err := c.cc.Invoke(ctx, "/gophkeeper.v1.LoginService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) ListAvailableLogins(ctx context.Context, in *ListAvailableLoginsRequest, opts ...grpc.CallOption) (*ListAvailableLoginsResponse, error) {
	out := new(ListAvailableLoginsResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.v1.LoginService/ListAvailableLogins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations must embed UnimplementedLoginServiceServer
// for forward compatibility
type LoginServiceServer interface {
	Create(context.Context, *CreateLoginRequest) (*Login, error)
	Update(context.Context, *UpdateLoginRequest) (*Login, error)
	Get(context.Context, *GetLoginRequest) (*Login, error)
	ListAvailableLogins(context.Context, *ListAvailableLoginsRequest) (*ListAvailableLoginsResponse, error)
	mustEmbedUnimplementedLoginServiceServer()
}

// UnimplementedLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (UnimplementedLoginServiceServer) Create(context.Context, *CreateLoginRequest) (*Login, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLoginServiceServer) Update(context.Context, *UpdateLoginRequest) (*Login, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLoginServiceServer) Get(context.Context, *GetLoginRequest) (*Login, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLoginServiceServer) ListAvailableLogins(context.Context, *ListAvailableLoginsRequest) (*ListAvailableLoginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAvailableLogins not implemented")
}
func (UnimplementedLoginServiceServer) mustEmbedUnimplementedLoginServiceServer() {}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.v1.LoginService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Create(ctx, req.(*CreateLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.v1.LoginService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Update(ctx, req.(*UpdateLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.v1.LoginService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Get(ctx, req.(*GetLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_ListAvailableLogins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAvailableLoginsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).ListAvailableLogins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.v1.LoginService/ListAvailableLogins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).ListAvailableLogins(ctx, req.(*ListAvailableLoginsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.v1.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LoginService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LoginService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LoginService_Get_Handler,
		},
		{
			MethodName: "ListAvailableLogins",
			Handler:    _LoginService_ListAvailableLogins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login.proto",
}
