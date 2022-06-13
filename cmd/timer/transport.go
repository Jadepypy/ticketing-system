package timer

import (
	"context"
	"github.com/Jadepypy/ticketing-system/proto/timer"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	setTimer grpctransport.Handler
	timer_proto.UnimplementedTimerServiceServer
}

func (g *grpcServer) SetTimer(ctx context.Context, req *timer_proto.SetTimerRequest) (*timer_proto.SetTimerResponse, error) {
	_, resp, err := g.setTimer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*timer_proto.SetTimerResponse), nil
}
func MakeGRPCServer(ctx context.Context, s Service) timer_proto.TimerServiceServer {
	return &grpcServer{
		setTimer: grpctransport.NewServer(
			makeSetTimerEndpoint(s),
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
