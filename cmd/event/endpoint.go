package event

import (
	"context"
	"errors"
	event_proto "github.com/Jadepypy/ticketing-system/proto/event"
	"github.com/go-kit/kit/endpoint"
)

func makeCreateEventEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*event_proto.CreateEventRequest)
		if !ok {
			return nil, errors.New("can't cast request")
		}
		v, err := s.CreateEvent(ctx, req)
		if err != nil {
			return nil, errors.New("internal server error")
		}
		return v, nil
	}
}
