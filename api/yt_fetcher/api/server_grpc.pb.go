// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// YoutubeFetcherClient is the client API for YoutubeFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YoutubeFetcherClient interface {
	// Get videoIds from a youtube video list page
	GetVideoIds(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*Channel, error)
	// Get videos info by page url set in channel
	GetVideos(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*Videos, error)
	// Get video info by videoId
	GetVideo(ctx context.Context, in *Video, opts ...grpc.CallOption) (*Video, error)
	// Get Channel name by cid
	GetCname(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*Channel, error)
	// Get Channel info by cid
	GetChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*Channel, error)
}

type youtubeFetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewYoutubeFetcherClient(cc grpc.ClientConnInterface) YoutubeFetcherClient {
	return &youtubeFetcherClient{cc}
}

func (c *youtubeFetcherClient) GetVideoIds(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/yt_fetcher.api.YoutubeFetcher/GetVideoIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *youtubeFetcherClient) GetVideos(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*Videos, error) {
	out := new(Videos)
	err := c.cc.Invoke(ctx, "/yt_fetcher.api.YoutubeFetcher/GetVideos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *youtubeFetcherClient) GetVideo(ctx context.Context, in *Video, opts ...grpc.CallOption) (*Video, error) {
	out := new(Video)
	err := c.cc.Invoke(ctx, "/yt_fetcher.api.YoutubeFetcher/GetVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *youtubeFetcherClient) GetCname(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/yt_fetcher.api.YoutubeFetcher/GetCname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *youtubeFetcherClient) GetChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/yt_fetcher.api.YoutubeFetcher/GetChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YoutubeFetcherServer is the server API for YoutubeFetcher service.
// All implementations must embed UnimplementedYoutubeFetcherServer
// for forward compatibility
type YoutubeFetcherServer interface {
	// Get videoIds from a youtube video list page
	GetVideoIds(context.Context, *Channel) (*Channel, error)
	// Get videos info by page url set in channel
	GetVideos(context.Context, *Channel) (*Videos, error)
	// Get video info by videoId
	GetVideo(context.Context, *Video) (*Video, error)
	// Get Channel name by cid
	GetCname(context.Context, *Channel) (*Channel, error)
	// Get Channel info by cid
	GetChannel(context.Context, *Channel) (*Channel, error)
	mustEmbedUnimplementedYoutubeFetcherServer()
}

// UnimplementedYoutubeFetcherServer must be embedded to have forward compatible implementations.
type UnimplementedYoutubeFetcherServer struct {
}

func (UnimplementedYoutubeFetcherServer) GetVideoIds(context.Context, *Channel) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoIds not implemented")
}
func (UnimplementedYoutubeFetcherServer) GetVideos(context.Context, *Channel) (*Videos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideos not implemented")
}
func (UnimplementedYoutubeFetcherServer) GetVideo(context.Context, *Video) (*Video, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}
func (UnimplementedYoutubeFetcherServer) GetCname(context.Context, *Channel) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCname not implemented")
}
func (UnimplementedYoutubeFetcherServer) GetChannel(context.Context, *Channel) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannel not implemented")
}
func (UnimplementedYoutubeFetcherServer) mustEmbedUnimplementedYoutubeFetcherServer() {}

// UnsafeYoutubeFetcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YoutubeFetcherServer will
// result in compilation errors.
type UnsafeYoutubeFetcherServer interface {
	mustEmbedUnimplementedYoutubeFetcherServer()
}

func RegisterYoutubeFetcherServer(s grpc.ServiceRegistrar, srv YoutubeFetcherServer) {
	s.RegisterService(&YoutubeFetcher_ServiceDesc, srv)
}

func _YoutubeFetcher_GetVideoIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeFetcherServer).GetVideoIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yt_fetcher.api.YoutubeFetcher/GetVideoIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeFetcherServer).GetVideoIds(ctx, req.(*Channel))
	}
	return interceptor(ctx, in, info, handler)
}

func _YoutubeFetcher_GetVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeFetcherServer).GetVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yt_fetcher.api.YoutubeFetcher/GetVideos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeFetcherServer).GetVideos(ctx, req.(*Channel))
	}
	return interceptor(ctx, in, info, handler)
}

func _YoutubeFetcher_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Video)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeFetcherServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yt_fetcher.api.YoutubeFetcher/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeFetcherServer).GetVideo(ctx, req.(*Video))
	}
	return interceptor(ctx, in, info, handler)
}

func _YoutubeFetcher_GetCname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeFetcherServer).GetCname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yt_fetcher.api.YoutubeFetcher/GetCname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeFetcherServer).GetCname(ctx, req.(*Channel))
	}
	return interceptor(ctx, in, info, handler)
}

func _YoutubeFetcher_GetChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeFetcherServer).GetChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yt_fetcher.api.YoutubeFetcher/GetChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeFetcherServer).GetChannel(ctx, req.(*Channel))
	}
	return interceptor(ctx, in, info, handler)
}

// YoutubeFetcher_ServiceDesc is the grpc.ServiceDesc for YoutubeFetcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YoutubeFetcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yt_fetcher.api.YoutubeFetcher",
	HandlerType: (*YoutubeFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideoIds",
			Handler:    _YoutubeFetcher_GetVideoIds_Handler,
		},
		{
			MethodName: "GetVideos",
			Handler:    _YoutubeFetcher_GetVideos_Handler,
		},
		{
			MethodName: "GetVideo",
			Handler:    _YoutubeFetcher_GetVideo_Handler,
		},
		{
			MethodName: "GetCname",
			Handler:    _YoutubeFetcher_GetCname_Handler,
		},
		{
			MethodName: "GetChannel",
			Handler:    _YoutubeFetcher_GetChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/yt_fetcher/api/server.proto",
}
