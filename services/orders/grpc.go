package main

import (
	"log"
	"net"

	handler "github.com/ThuraMinThein/go_microservices/services/orders/handler/orders"
	"github.com/ThuraMinThein/go_microservices/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// register the service
	orderService := service.NewOrderService()
	handler.NewGRPCOrdersService(grpcServer, orderService)

	log.Println("gRPC server is running on", s.addr)

	return grpcServer.Serve(lis)
}
