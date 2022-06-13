// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ticket_proto

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

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketServiceClient interface {
	GetTickets(ctx context.Context, in *GetTicketsRequest, opts ...grpc.CallOption) (*GetTicketsResponse, error)
	CreateTickets(ctx context.Context, in *CreateTicketsRequest, opts ...grpc.CallOption) (*CreateTicketsResponse, error)
	DeleteTickets(ctx context.Context, in *DeleteTicketsRequest, opts ...grpc.CallOption) (*DeleteTicketsResponse, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) GetTickets(ctx context.Context, in *GetTicketsRequest, opts ...grpc.CallOption) (*GetTicketsResponse, error) {
	out := new(GetTicketsResponse)
	err := c.cc.Invoke(ctx, "/ticket_proto.TicketService/GetTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) CreateTickets(ctx context.Context, in *CreateTicketsRequest, opts ...grpc.CallOption) (*CreateTicketsResponse, error) {
	out := new(CreateTicketsResponse)
	err := c.cc.Invoke(ctx, "/ticket_proto.TicketService/CreateTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) DeleteTickets(ctx context.Context, in *DeleteTicketsRequest, opts ...grpc.CallOption) (*DeleteTicketsResponse, error) {
	out := new(DeleteTicketsResponse)
	err := c.cc.Invoke(ctx, "/ticket_proto.TicketService/DeleteTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
// All implementations must embed UnimplementedTicketServiceServer
// for forward compatibility
type TicketServiceServer interface {
	GetTickets(context.Context, *GetTicketsRequest) (*GetTicketsResponse, error)
	CreateTickets(context.Context, *CreateTicketsRequest) (*CreateTicketsResponse, error)
	DeleteTickets(context.Context, *DeleteTicketsRequest) (*DeleteTicketsResponse, error)
	mustEmbedUnimplementedTicketServiceServer()
}

// UnimplementedTicketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTicketServiceServer struct {
}

func (UnimplementedTicketServiceServer) GetTickets(context.Context, *GetTicketsRequest) (*GetTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTickets not implemented")
}
func (UnimplementedTicketServiceServer) CreateTickets(context.Context, *CreateTicketsRequest) (*CreateTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTickets not implemented")
}
func (UnimplementedTicketServiceServer) DeleteTickets(context.Context, *DeleteTicketsRequest) (*DeleteTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTickets not implemented")
}
func (UnimplementedTicketServiceServer) mustEmbedUnimplementedTicketServiceServer() {}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
	s.RegisterService(&TicketService_ServiceDesc, srv)
}

func _TicketService_GetTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_proto.TicketService/GetTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetTickets(ctx, req.(*GetTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_CreateTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).CreateTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_proto.TicketService/CreateTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).CreateTickets(ctx, req.(*CreateTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_DeleteTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).DeleteTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_proto.TicketService/DeleteTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).DeleteTickets(ctx, req.(*DeleteTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticket_proto.TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTickets",
			Handler:    _TicketService_GetTickets_Handler,
		},
		{
			MethodName: "CreateTickets",
			Handler:    _TicketService_CreateTickets_Handler,
		},
		{
			MethodName: "DeleteTickets",
			Handler:    _TicketService_DeleteTickets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticket.proto",
}