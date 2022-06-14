package event

import (
	"context"
	"fmt"
	event_proto "github.com/Jadepypy/ticketing-system/proto/event"
)

type Service interface {
	GetEvents(ctx context.Context, req *event_proto.GetEventsRequest) (*event_proto.GetEventsResponse, error)
	GetEvent(ctx context.Context, req *event_proto.GetEventRequest) (*event_proto.GetEventResponse, error)
	CreateEvent(ctx context.Context, req *event_proto.CreateEventRequest) (*event_proto.CreateEventResponse, error)
	DeleteEvent(ctx context.Context, req *event_proto.DeleteEventRequest) (*event_proto.DeleteEventResponse, error)
	UpdateEvent(ctx context.Context, req *event_proto.UpdateEventRequest) (*event_proto.UpdateEventResponse, error)
	UpdateEventStatus(ctx context.Context, req *event_proto.UpdateEventStatusRequest) (*event_proto.UpdateEventStatusResponse, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) GetEvents(ctx context.Context, req *event_proto.GetEventsRequest) (*event_proto.GetEventsResponse, error) {
	fmt.Println("set timer...")
	return &event_proto.GetEventsResponse{}, nil
}

func (s *service) GetEvent(ctx context.Context, req *event_proto.GetEventRequest) (*event_proto.GetEventResponse, error) {
	fmt.Println("set timer...")
	return &event_proto.GetEventResponse{}, nil
}

func (s *service) CreateEvent(ctx context.Context, req *event_proto.CreateEventRequest) (*event_proto.CreateEventResponse, error) {
	fmt.Println("create event...")
	return &event_proto.CreateEventResponse{}, nil
}

func (s *service) DeleteEvent(ctx context.Context, req *event_proto.DeleteEventRequest) (*event_proto.DeleteEventResponse, error) {
	fmt.Println("set timer...")
	return &event_proto.DeleteEventResponse{}, nil
}

func (s *service) UpdateEvent(ctx context.Context, req *event_proto.UpdateEventRequest) (*event_proto.UpdateEventResponse, error) {
	fmt.Println("set timer...")
	return &event_proto.UpdateEventResponse{}, nil
}

func (s *service) UpdateEventStatus(ctx context.Context, req *event_proto.UpdateEventStatusRequest) (*event_proto.UpdateEventStatusResponse, error) {
	fmt.Println("set timer...")
	return &event_proto.UpdateEventStatusResponse{}, nil
}
