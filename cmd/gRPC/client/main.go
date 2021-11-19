package main

import (
	"context"
	"flag"
	"log"

	"github.com/cloudnativedesign/accountmanager/internal/gRPC/domain"
	"github.com/cloudnativedesign/accountmanager/internal/gRPC/service"
	"google.golang.org/grpc"
)

var (
	tls               = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile            = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr        = flag.String("addr", "localhost:50051", "The server address in the format host:port")
	serverHostOverrid = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS")
)

// addAccount creates a new account for the specified provider
func addAccount(client service.AccountServiceClient, account *domain.Account) {

	response, err := client.Add(context.Background(), account)
	if err != nil {
		log.Fatal("Unable to create account")
	}
	log.Printf("Successfully created account with id: %d", response.Account.Id)
}

func main() {
	flag.Parse()
	// var opts []grpc.DialOption

	conn, e := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if e != nil {
		panic(e)
	}

	defer conn.Close()
	client := service.NewAccountServiceClient(conn)

	// Creating a new account
	account := &domain.Account{
		Id:       64521,
		Name:     "frank.winkler",
		Provider: domain.Provider_LINKEDIN,
	}

	// Create a new account
	addAccount(client, account)
}
