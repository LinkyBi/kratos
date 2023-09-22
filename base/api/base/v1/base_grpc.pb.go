// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: base/v1/base.proto

package v1

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
	Base_CreateBase_FullMethodName = "/api.base.v1.Base/CreateBase"
	Base_UpdateBase_FullMethodName = "/api.base.v1.Base/UpdateBase"
	Base_DeleteBase_FullMethodName = "/api.base.v1.Base/DeleteBase"
	Base_GetBase_FullMethodName    = "/api.base.v1.Base/GetBase"
	Base_ListBase_FullMethodName   = "/api.base.v1.Base/ListBase"
)

// BaseClient is the client API for Base service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaseClient interface {
	CreateBase(ctx context.Context, in *CreateBaseRequest, opts ...grpc.CallOption) (*CreateBaseReply, error)
	UpdateBase(ctx context.Context, in *UpdateBaseRequest, opts ...grpc.CallOption) (*UpdateBaseReply, error)
	DeleteBase(ctx context.Context, in *DeleteBaseRequest, opts ...grpc.CallOption) (*DeleteBaseReply, error)
	GetBase(ctx context.Context, in *GetBaseRequest, opts ...grpc.CallOption) (*GetBaseReply, error)
	ListBase(ctx context.Context, in *ListBaseRequest, opts ...grpc.CallOption) (*ListBaseReply, error)
}

type baseClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseClient(cc grpc.ClientConnInterface) BaseClient {
	return &baseClient{cc}
}

func (c *baseClient) CreateBase(ctx context.Context, in *CreateBaseRequest, opts ...grpc.CallOption) (*CreateBaseReply, error) {
	out := new(CreateBaseReply)
	err := c.cc.Invoke(ctx, Base_CreateBase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) UpdateBase(ctx context.Context, in *UpdateBaseRequest, opts ...grpc.CallOption) (*UpdateBaseReply, error) {
	out := new(UpdateBaseReply)
	err := c.cc.Invoke(ctx, Base_UpdateBase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) DeleteBase(ctx context.Context, in *DeleteBaseRequest, opts ...grpc.CallOption) (*DeleteBaseReply, error) {
	out := new(DeleteBaseReply)
	err := c.cc.Invoke(ctx, Base_DeleteBase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) GetBase(ctx context.Context, in *GetBaseRequest, opts ...grpc.CallOption) (*GetBaseReply, error) {
	out := new(GetBaseReply)
	err := c.cc.Invoke(ctx, Base_GetBase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseClient) ListBase(ctx context.Context, in *ListBaseRequest, opts ...grpc.CallOption) (*ListBaseReply, error) {
	out := new(ListBaseReply)
	err := c.cc.Invoke(ctx, Base_ListBase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseServer is the server API for Base service.
// All implementations must embed UnimplementedBaseServer
// for forward compatibility
type BaseServer interface {
	CreateBase(context.Context, *CreateBaseRequest) (*CreateBaseReply, error)
	UpdateBase(context.Context, *UpdateBaseRequest) (*UpdateBaseReply, error)
	DeleteBase(context.Context, *DeleteBaseRequest) (*DeleteBaseReply, error)
	GetBase(context.Context, *GetBaseRequest) (*GetBaseReply, error)
	ListBase(context.Context, *ListBaseRequest) (*ListBaseReply, error)
	mustEmbedUnimplementedBaseServer()
}

// UnimplementedBaseServer must be embedded to have forward compatible implementations.
type UnimplementedBaseServer struct {
}

func (UnimplementedBaseServer) CreateBase(context.Context, *CreateBaseRequest) (*CreateBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBase not implemented")
}
func (UnimplementedBaseServer) UpdateBase(context.Context, *UpdateBaseRequest) (*UpdateBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBase not implemented")
}
func (UnimplementedBaseServer) DeleteBase(context.Context, *DeleteBaseRequest) (*DeleteBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBase not implemented")
}
func (UnimplementedBaseServer) GetBase(context.Context, *GetBaseRequest) (*GetBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBase not implemented")
}
func (UnimplementedBaseServer) ListBase(context.Context, *ListBaseRequest) (*ListBaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBase not implemented")
}
func (UnimplementedBaseServer) mustEmbedUnimplementedBaseServer() {}

// UnsafeBaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseServer will
// result in compilation errors.
type UnsafeBaseServer interface {
	mustEmbedUnimplementedBaseServer()
}

func RegisterBaseServer(s grpc.ServiceRegistrar, srv BaseServer) {
	s.RegisterService(&Base_ServiceDesc, srv)
}

func _Base_CreateBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).CreateBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Base_CreateBase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).CreateBase(ctx, req.(*CreateBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_UpdateBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).UpdateBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Base_UpdateBase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).UpdateBase(ctx, req.(*UpdateBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_DeleteBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).DeleteBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Base_DeleteBase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).DeleteBase(ctx, req.(*DeleteBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_GetBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).GetBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Base_GetBase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).GetBase(ctx, req.(*GetBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Base_ListBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseServer).ListBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Base_ListBase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseServer).ListBase(ctx, req.(*ListBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Base_ServiceDesc is the grpc.ServiceDesc for Base service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Base_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.base.v1.Base",
	HandlerType: (*BaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBase",
			Handler:    _Base_CreateBase_Handler,
		},
		{
			MethodName: "UpdateBase",
			Handler:    _Base_UpdateBase_Handler,
		},
		{
			MethodName: "DeleteBase",
			Handler:    _Base_DeleteBase_Handler,
		},
		{
			MethodName: "GetBase",
			Handler:    _Base_GetBase_Handler,
		},
		{
			MethodName: "ListBase",
			Handler:    _Base_ListBase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/v1/base.proto",
}
