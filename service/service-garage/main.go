package main

import (
	"context"
	"log"
	"net"

	"tesprotogrpc/common/config"
	"tesprotogrpc/common/model"

	"github.com/golang/protobuf/ptypes/empty"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/instrumentation/grpctrace"
	"google.golang.org/grpc"
)

var localStorage *model.GarageListByUser

func init() {
	localStorage = new(model.GarageListByUser)
	localStorage.List = make(map[string]*model.GarageList)
}

type GaragesServer struct{}

func (GaragesServer) Add(ctx context.Context, param *model.GarageAndUserId) (*empty.Empty, error) {
	userId := param.UserId
	garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(model.GarageList)
		localStorage.List[userId].List = make([]*model.Garage, 0)
	}
	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	log.Println("Adding garage", garage.String(), "for user", userId)

	return new(empty.Empty), nil
}
func (GaragesServer) List(ctx context.Context, param *model.GarageUserId) (*model.GarageList, error) {
	userId := param.UserId

	return localStorage.List[userId], nil
}
func main() {
	// var servOpts []grpc.ServerOption
	// tracer, _, err := serviceusr.NewJaegerTracer("testSrv1", "127.0.0.1:6831")
	// if err != nil {
	// 	fmt.Printf("new tracer err: %+v\n", err)
	// 	os.Exit(-1)
	// }
	// if tracer != nil {
	// 	servOpts = append(servOpts, serviceusr.ServerOption(tracer))
	// }
	// srv := grpc.NewServer(servOpts...)
	fn := config.InitTraceProvider("srvgarage")
	defer fn()
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(grpctrace.UnaryServerInterceptor(global.Tracer("srvgarage"))),
		grpc.StreamInterceptor(grpctrace.StreamServerInterceptor(global.Tracer("srvgaragestream"))))
	var garageSrv GaragesServer
	model.RegisterGaragesServer(srv, garageSrv)

	log.Println("Starting RPC server at", config.SERVICE_GARAGE_PORT)

	l, err := net.Listen("tcp", config.SERVICE_GARAGE_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_GARAGE_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
