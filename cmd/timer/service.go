package timer

import (
	"context"
	"fmt"
	"github.com/Jadepypy/ticketing-system/proto/timer"
)

type Service interface {
	SetTimer(ctx context.Context, req *timer_proto.SetTimerRequest) (*timer_proto.SetTimerResponse, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) SetTimer(ctx context.Context, req *timer_proto.SetTimerRequest) (*timer_proto.SetTimerResponse, error) {
	fmt.Println("set timer...")
	return &timer_proto.SetTimerResponse{}, nil
}
