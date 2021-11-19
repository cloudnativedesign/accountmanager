package main

import (
	"fmt"
	"log"
	"net"

	"github.com/cloudnativedesign/accountmanager/internal/gRPC/impl"
	"github.com/cloudnativedesign/accountmanager/internal/gRPC/service"
	"google.golang.org/grpc"
)

func main() {
	netListener := getNetListener(50051)

	accountServiceImpl := impl.NewAccountServiceGrpcImpl()
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	service.RegisterAccountServiceServer(grpcServer, accountServiceImpl)

	// Start the server
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	log.Printf("Server started. Listening on %d", 50051)
	return lis
}
