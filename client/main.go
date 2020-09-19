package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"tesprotogrpc/common/config"
	"tesprotogrpc/common/model"

	"github.com/golang/protobuf/ptypes/empty"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/instrumentation/grpctrace"
	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewGaragesClient(conn)
}
func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}
func main() {
	fn := config.InitTraceProvider("client")
	defer fn()
	tracer := global.Tracer("client")

	user1 := model.User{
		Id:       "n001",
		Name:     "Noval Agung",
		Password: "kw8d hl12/3m,a",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}

	garage1 := model.Garage{
		Id:   "q001",
		Name: "Quel'thalas",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.123123123,
			Longitude: 54.1231313123,
		},
	}

	user := serviceUser()

	fmt.Println("\n", "===========> user test")

	// register user1
	user.Register(context.Background(), &user1)

	// register user2
	// user.Register(context.Background(), &user2)
	// show all registered users
	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	//garage
	garage := serviceGarage()

	fmt.Println("\n", "===========> garage test A")

	// add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})
	// add garage2 to user1
	// garage.Add(context.Background(), &model.GarageAndUserId{
	// 	UserId: user1.Id,
	// 	Garage: &garage2,
	// })
	// show all garages of user1
	res2, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))
	fmt.Println("\n", "===========> garage test B")

	// add garage3 to user2
	// garage.Add(context.Background(), &model.GarageAndUserId{
	// 	UserId: user2.Id,
	// 	Garage: &garage3,
	// })
	// show all garages of user2
	// res3, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// res3String, _ := json.Marshal(res3.List)
	// log.Println(string(res3String))
	//opentracing
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
	// 	conn1, err := grpc.Dial("127.0.0.1:7000", dialOpts...)
	// 	if err != nil {
	// 		fmt.Printf("grpc connect failed, err:%+v\n", err)
	// 		os.Exit(-1)
	// 	}
	// 	defer conn1.Close()

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
	// time.Sleep(time.Second * 3)
	// end opentracing
	cc, err := grpc.Dial("localhost:8020", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpctrace.UnaryClientInterceptor(tracer)))

	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer cc.Close()
	cc1, err := grpc.Dial("localhost:7000", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpctrace.UnaryClientInterceptor(tracer)))

	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer cc1.Close()
}
