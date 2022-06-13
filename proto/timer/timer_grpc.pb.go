// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package timer_proto

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

// TimerServiceClient is the client API for TimerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimerServiceClient interface {
	SetTimer(ctx context.Context, in *SetTimerRequest, opts ...grpc.CallOption) (*SetTimerResponse, error)
}

type timerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimerServiceClient(cc grpc.ClientConnInterface) TimerServiceClient {
	return &timerServiceClient{cc}
}

func (c *timerServiceClient) SetTimer(ctx context.Context, in *SetTimerRequest, opts ...grpc.CallOption) (*SetTimerResponse, error) {
	out := new(SetTimerResponse)
	err := c.cc.Invoke(ctx, "/timer_proto.TimerService/SetTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimerServiceServer is the server API for TimerService service.
// All implementations must embed UnimplementedTimerServiceServer
// for forward compatibility
type TimerServiceServer interface {
	SetTimer(context.Context, *SetTimerRequest) (*SetTimerResponse, error)
	mustEmbedUnimplementedTimerServiceServer()
}

// UnimplementedTimerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimerServiceServer struct {
}

func (UnimplementedTimerServiceServer) SetTimer(context.Context, *SetTimerRequest) (*SetTimerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTimer not implemented")
}
func (UnimplementedTimerServiceServer) mustEmbedUnimplementedTimerServiceServer() {}

// UnsafeTimerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimerServiceServer will
// result in compilation errors.
type UnsafeTimerServiceServer interface {
	mustEmbedUnimplementedTimerServiceServer()
}

func RegisterTimerServiceServer(s grpc.ServiceRegistrar, srv TimerServiceServer) {
	s.RegisterService(&TimerService_ServiceDesc, srv)
}

func _TimerService_SetTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimerServiceServer).SetTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timer_proto.TimerService/SetTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimerServiceServer).SetTimer(ctx, req.(*SetTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimerService_ServiceDesc is the grpc.ServiceDesc for TimerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "timer_proto.TimerService",
	HandlerType: (*TimerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetTimer",
			Handler:    _TimerService_SetTimer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timer.proto",
}