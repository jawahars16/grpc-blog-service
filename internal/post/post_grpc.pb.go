// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: internal/post/post.proto

package post

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
	Blog_CreatePost_FullMethodName = "/post.Blog/CreatePost"
	Blog_GetPost_FullMethodName    = "/post.Blog/GetPost"
	Blog_UpdatePost_FullMethodName = "/post.Blog/UpdatePost"
	Blog_DeletePost_FullMethodName = "/post.Blog/DeletePost"
)

// BlogClient is the client API for Blog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
}

type blogClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogClient(cc grpc.ClientConnInterface) BlogClient {
	return &blogClient{cc}
}

func (c *blogClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, Blog_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error) {
	out := new(GetPostResponse)
	err := c.cc.Invoke(ctx, Blog_GetPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error) {
	out := new(UpdatePostResponse)
	err := c.cc.Invoke(ctx, Blog_UpdatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, Blog_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServer is the server API for Blog service.
// All implementations must embed UnimplementedBlogServer
// for forward compatibility
type BlogServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	mustEmbedUnimplementedBlogServer()
}

// UnimplementedBlogServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServer struct {
}

func (UnimplementedBlogServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedBlogServer) GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedBlogServer) UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedBlogServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedBlogServer) mustEmbedUnimplementedBlogServer() {}

// UnsafeBlogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServer will
// result in compilation errors.
type UnsafeBlogServer interface {
	mustEmbedUnimplementedBlogServer()
}

func RegisterBlogServer(s grpc.ServiceRegistrar, srv BlogServer) {
	s.RegisterService(&Blog_ServiceDesc, srv)
}

func _Blog_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blog_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Blog_ServiceDesc is the grpc.ServiceDesc for Blog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post.Blog",
	HandlerType: (*BlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _Blog_CreatePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _Blog_GetPost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Blog_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Blog_DeletePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/post/post.proto",
}
