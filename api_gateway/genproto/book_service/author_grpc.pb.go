// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: author.proto

package book_service

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

// AuthorServiceClient is the client API for AuthorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorServiceClient interface {
	Create(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	GetAll(ctx context.Context, in *GetAllAuthorRequest, opts ...grpc.CallOption) (*GetAllAuthorResponse, error)
	Get(ctx context.Context, in *AuthorId, opts ...grpc.CallOption) (*Author, error)
	Update(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Result, error)
	Delete(ctx context.Context, in *AuthorId, opts ...grpc.CallOption) (*Result, error)
}

type authorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorServiceClient(cc grpc.ClientConnInterface) AuthorServiceClient {
	return &authorServiceClient{cc}
}

func (c *authorServiceClient) Create(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/genproto.AuthorService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) GetAll(ctx context.Context, in *GetAllAuthorRequest, opts ...grpc.CallOption) (*GetAllAuthorResponse, error) {
	out := new(GetAllAuthorResponse)
	err := c.cc.Invoke(ctx, "/genproto.AuthorService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) Get(ctx context.Context, in *AuthorId, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/genproto.AuthorService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) Update(ctx context.Context, in *Author, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/genproto.AuthorService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) Delete(ctx context.Context, in *AuthorId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/genproto.AuthorService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServiceServer is the server API for AuthorService service.
// All implementations must embed UnimplementedAuthorServiceServer
// for forward compatibility
type AuthorServiceServer interface {
	Create(context.Context, *CreateAuthorRequest) (*Author, error)
	GetAll(context.Context, *GetAllAuthorRequest) (*GetAllAuthorResponse, error)
	Get(context.Context, *AuthorId) (*Author, error)
	Update(context.Context, *Author) (*Result, error)
	Delete(context.Context, *AuthorId) (*Result, error)
	mustEmbedUnimplementedAuthorServiceServer()
}

// UnimplementedAuthorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorServiceServer struct {
}

func (UnimplementedAuthorServiceServer) Create(context.Context, *CreateAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAuthorServiceServer) GetAll(context.Context, *GetAllAuthorRequest) (*GetAllAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAuthorServiceServer) Get(context.Context, *AuthorId) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAuthorServiceServer) Update(context.Context, *Author) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAuthorServiceServer) Delete(context.Context, *AuthorId) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAuthorServiceServer) mustEmbedUnimplementedAuthorServiceServer() {}

// UnsafeAuthorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorServiceServer will
// result in compilation errors.
type UnsafeAuthorServiceServer interface {
	mustEmbedUnimplementedAuthorServiceServer()
}

func RegisterAuthorServiceServer(s grpc.ServiceRegistrar, srv AuthorServiceServer) {
	s.RegisterService(&AuthorService_ServiceDesc, srv)
}

func _AuthorService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AuthorService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).Create(ctx, req.(*CreateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AuthorService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAll(ctx, req.(*GetAllAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AuthorService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).Get(ctx, req.(*AuthorId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Author)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AuthorService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).Update(ctx, req.(*Author))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AuthorService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).Delete(ctx, req.(*AuthorId))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorService_ServiceDesc is the grpc.ServiceDesc for AuthorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AuthorService_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AuthorService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AuthorService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AuthorService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AuthorService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "author.proto",
}
