// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/streaming.proto

package __

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
	VideoStreaming_StreamVideo_FullMethodName = "/streaming.VideoStreaming/StreamVideo"
)

// VideoStreamingClient is the client API for VideoStreaming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoStreamingClient interface {
	StreamVideo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[VideoChunk, VideoChunk], error)
}

type videoStreamingClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoStreamingClient(cc grpc.ClientConnInterface) VideoStreamingClient {
	return &videoStreamingClient{cc}
}

func (c *videoStreamingClient) StreamVideo(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[VideoChunk, VideoChunk], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoStreaming_ServiceDesc.Streams[0], VideoStreaming_StreamVideo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[VideoChunk, VideoChunk]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoStreaming_StreamVideoClient = grpc.BidiStreamingClient[VideoChunk, VideoChunk]

// VideoStreamingServer is the server API for VideoStreaming service.
// All implementations must embed UnimplementedVideoStreamingServer
// for forward compatibility.
type VideoStreamingServer interface {
	StreamVideo(grpc.BidiStreamingServer[VideoChunk, VideoChunk]) error
	mustEmbedUnimplementedVideoStreamingServer()
}

// UnimplementedVideoStreamingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVideoStreamingServer struct{}

func (UnimplementedVideoStreamingServer) StreamVideo(grpc.BidiStreamingServer[VideoChunk, VideoChunk]) error {
	return status.Errorf(codes.Unimplemented, "method StreamVideo not implemented")
}
func (UnimplementedVideoStreamingServer) mustEmbedUnimplementedVideoStreamingServer() {}
func (UnimplementedVideoStreamingServer) testEmbeddedByValue()                        {}

// UnsafeVideoStreamingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoStreamingServer will
// result in compilation errors.
type UnsafeVideoStreamingServer interface {
	mustEmbedUnimplementedVideoStreamingServer()
}

func RegisterVideoStreamingServer(s grpc.ServiceRegistrar, srv VideoStreamingServer) {
	// If the following call pancis, it indicates UnimplementedVideoStreamingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VideoStreaming_ServiceDesc, srv)
}

func _VideoStreaming_StreamVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoStreamingServer).StreamVideo(&grpc.GenericServerStream[VideoChunk, VideoChunk]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type VideoStreaming_StreamVideoServer = grpc.BidiStreamingServer[VideoChunk, VideoChunk]

// VideoStreaming_ServiceDesc is the grpc.ServiceDesc for VideoStreaming service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoStreaming_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streaming.VideoStreaming",
	HandlerType: (*VideoStreamingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamVideo",
			Handler:       _VideoStreaming_StreamVideo_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/streaming.proto",
}
