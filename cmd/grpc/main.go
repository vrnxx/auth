package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
	infraConf "github.com/vrnxx/auth/internal/infrastructure/config"
	grpcDelivery "github.com/vrnxx/auth/internal/presentation/api/delivery/grpc"
	mainConf "github.com/vrnxx/auth/internal/presentation/config"
	desc "github.com/vrnxx/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = ":55031"

func main() {
	c := mainConf.Config{}
	infraConf.LoadConfig(&c, "")
	fmt.Println(color.GreenString("loaded config: %v", c))

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(color.RedString("net.Listen error: %s", err.Error()))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, grpcDelivery.NewV1GRPCServer())
	fmt.Println(color.GreenString("server started on"), color.HiGreenString(grpcPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal(
			color.RedString(fmt.Sprintf("server error:%s", err.Error())),
		)
	}
}
