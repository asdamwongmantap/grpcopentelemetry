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

	// "go.opentelemetry.io/otel/label"

	"google.golang.org/grpc"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct{}

func (UsersServer) Register(ctx context.Context, param *model.User) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, param)

	log.Println("Registering user", param.String())

	return new(empty.Empty), nil
}

func (UsersServer) List(ctx context.Context, void *empty.Empty) (*model.UserList, error) {
	return localStorage, nil
}

// func initTracer() func() {
// 	// Create and install Jaeger export pipeline
// 	flush, err := jaeger.InstallNewPipeline(
// 		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
// 		jaeger.WithProcess(jaeger.Process{
// 			ServiceName: "gotracer",
// 			Tags: []label.KeyValue{
// 				label.String("exporter", "jaeger"),
// 				label.Float64("float", 312.23),
// 			},
// 		}),

// 		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return func() {
// 		flush()
// 	}
// }
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
	// config.Init()
	fn := config.InitTraceProvider("srvuser")
	defer fn()
	// fn := initTracer()

	// defer fn()

	// srv := grpc.NewServer(servOpts...)
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(grpctrace.UnaryServerInterceptor(global.Tracer("srvuser"))),
		grpc.StreamInterceptor(grpctrace.StreamServerInterceptor(global.Tracer("srvuserstream"))))
	var userSrv UsersServer
	model.RegisterUsersServer(srv, userSrv)

	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	l, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}
	// go func() {
	// 	time.Sleep(time.Second)

	// 	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	// 	tracer, _, err := serviceusr.NewJaegerTracer("testCli1", "127.0.0.1:6831")
	// 	if err != nil {
	// 		fmt.Printf("new tracer err: %+v\n", err)
	// 		os.Exit(-1)
	// 	}

	// 	if tracer != nil {
	// 		dialOpts = append(dialOpts, serviceusr.DialOption(tracer))
	// 	}

	// 	conn, err := grpc.Dial("127.0.0.1:8020", dialOpts...)
	// 	if err != nil {
	// 		fmt.Printf("grpc connect failed, err:%+v\n", err)
	// 		os.Exit(-1)
	// 	}
	// 	defer conn.Close()

	// 	// user1 := model.User{
	// 	// 	Id:       "n001",
	// 	// 	Name:     "Noval Agung",
	// 	// 	Password: "kw8d hl12/3m,a",
	// 	// 	Gender:   model.UserGender(model.UserGender_value["MALE"]),
	// 	// }
	// 	// resp, err := user.Register(context.Background(), &user1)
	// 	// if err != nil {
	// 	// 	fmt.Printf("call sayhello failed, err:%+v\n", err)
	// 	// 	// os.Exit(-1)
	// 	// } else {
	// 	// 	fmt.Printf("call sayhello suc, res:%+v\n", resp)
	// 	// }

	// }()
	log.Fatal(srv.Serve(l))
	// time.Sleep(time.Second * 3)
}
