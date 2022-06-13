package main

import (
	"context"
	"fmt"
	"github.com/Jadepypy/ticketing-system/cmd/timer"
	timer_proto "github.com/Jadepypy/ticketing-system/proto/timer"

	"google.golang.org/grpc"
	"log"
	"net"
)

//type eventServer struct {x
//}

//type timerServer struct {
//	timer_proto.UnimplementedTimerServiceServer
//}

//type ticketServer struct {
//}

//func (e eventServer) GetEvents(ctx context.Context, req *event_proto.GetEventsRequest) (*GetEventsResponse, error){
//	return nil, nil
//}

func main() {
	fmt.Println("Starting server...")
	ctx := context.WithValue(context.Background(), "MY-CLIENT-ID", "timer")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3000))
	if err != nil {
		log.Fatalf("failed to listed: %v", err)
	}

	// STEP 2-2：使用 gRPC 的 NewServer 方法來建立 gRPC Server 的實例
	service := timer.NewService()
	timerSrv := timer.MakeGRPCServer(ctx, service)
	grpcSrv := grpc.NewServer()

	// STEP 2-3：在 gRPC Server 中註冊 service 的實作
	// 使用 proto 提供的 RegisterTimerServiceServer 方法，並將 timerServer 作為參數傳入
	timer_proto.RegisterTimerServiceServer(grpcSrv, timerSrv)

	// STEP 2-4：啟動 grpcServer，並阻塞在這裡直到該程序被 kill 或 stop
	err = grpcSrv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
