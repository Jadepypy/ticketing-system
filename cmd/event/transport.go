package event

import (
	"context"
	event_proto "github.com/Jadepypy/ticketing-system/proto/event"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	getEvents         grpctransport.Handler
	getEvent          grpctransport.Handler
	createEvent       grpctransport.Handler
	deleteEvent       grpctransport.Handler
	updateEvent       grpctransport.Handler
	updateEventStatus grpctransport.Handler
	event_proto.UnimplementedEventServiceServer
}

func (g *grpcServer) GetEvents(ctx context.Context, req *event_proto.GetEventsRequest) (*event_proto.GetEventsResponse, error) {
	_, resp, err := g.getEvents.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*event_proto.GetEventsResponse), nil
}

func (g *grpcServer) GetEvent(ctx context.Context, req *event_proto.GetEventRequest) (*event_proto.GetEventResponse, error) {
	_, resp, err := g.getEvent.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*event_proto.GetEventResponse), nil
}

func (g *grpcServer) CreateEvent(ctx context.Context, req *event_proto.CreateEventRequest) (*event_proto.CreateEventResponse, error) {
	_, resp, err := g.createEvent.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*event_proto.CreateEventResponse), nil
}

func (g *grpcServer) DeleteEvent(ctx context.Context, req *event_proto.DeleteEventRequest) (*event_proto.DeleteEventResponse, error) {
	_, resp, err := g.deleteEvent.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*event_proto.DeleteEventResponse), nil
}

func (g *grpcServer) UpdateEvent(ctx context.Context, req *event_proto.UpdateEventRequest) (*event_proto.UpdateEventResponse, error) {
	_, resp, err := g.updateEvent.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*event_proto.UpdateEventResponse), nil
}

func (g *grpcServer) UpdateEventStatus(ctx context.Context, req *event_proto.UpdateEventStatusRequest) (*event_proto.UpdateEventStatusResponse, error) {
	_, resp, err := g.updateEventStatus.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*event_proto.UpdateEventStatusResponse), nil
}

func MakeGRPCServer(ctx context.Context, s Service) event_proto.EventServiceServer {
	return &grpcServer{
		createEvent: grpctransport.NewServer(
			makeCreateEventEndpoint(s),
			decodeRequest,
			encodeResponse,
		),
	}
}

func decodeRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}
func encodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
