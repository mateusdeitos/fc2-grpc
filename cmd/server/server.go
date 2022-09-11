package main

import (
	"log"
	"net"

	"github.com/codeedu/fc2-grpc/pb"
	"github.com/codeedu/fc2-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// Cria o listener para o servidor na porta 50051
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	log.Default().Println("Server is running...")

	// inicia o servidor
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &services.UserService{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
