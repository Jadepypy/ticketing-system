package timer

import (
	"context"
	"errors"
	timer_proto "github.com/Jadepypy/ticketing-system/proto/timer"
	"github.com/go-kit/kit/endpoint"
)

//type Endpoints struct {
//	SetTimer endpoint.Endpoint
//}

func makeSetTimerEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*timer_proto.SetTimerRequest)
		if !ok {
			return nil, errors.New("can't cast request")
		}
		v, err := s.SetTimer(ctx, req)
		if err != nil {
			return nil, errors.New("internal server error")
		}
		return v, nil
	}
}
