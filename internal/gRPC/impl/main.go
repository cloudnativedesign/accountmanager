package impl

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/cloudnativedesign/accountmanager/internal/gRPC/domain"
	"github.com/cloudnativedesign/accountmanager/internal/gRPC/service"
)

type AccountServiceGrpcImpl struct {
	service.UnimplementedAccountServiceServer
}

func NewAccountServiceGrpcImpl() *AccountServiceGrpcImpl {
	return &AccountServiceGrpcImpl{}
}

// Add a new account to the system by providing a fully specified account object
func (serviceImpl *AccountServiceGrpcImpl) Add(ctx context.Context, in *domain.Account) (*service.AddAccountResponse, error) {
	log.Println("Received request for adding account with id " + strconv.FormatInt(in.Id, 10))

	// Logic to persist to DB
	log.Println("Account persisted to storage")

	return &service.AddAccountResponse{
		Provider: in.Provider,
		Status:   200,
		Account: &domain.Account{
			Id:       in.Id,
			Name:     in.Name,
			Provider: in.Provider,
		},
	}, nil

}

// Update an account in the system by providing an updated account object
// to replace the current object state
func (serviceImpl *AccountServiceGrpcImpl) Update(ctx context.Context, in *domain.Account) (*service.UpdateAccountResponse, error) {
	log.Println("Received request to update account with id " + strconv.FormatInt(in.Id, 10))

	// Logic to update account in the Database
	log.Println("Account updated in database")

	return &service.UpdateAccountResponse{
		Id:     in.Id,
		Status: 200,
		Account: &domain.Account{
			Id:       in.Id,
			Name:     in.Name,
			Provider: in.Provider,
		},
	}, nil
}

// Connect to the given provider using your account credentials
func (serviceImpl *AccountServiceGrpcImpl) Connect(ctx context.Context, in *service.ConnectionRequest) (*service.ConnectionResponse, error) {
	log.Println(fmt.Sprintf("Received request to connect to %s", in.Provider))

	// Logic to connect to the given provider using the specified account credentials

	return &service.ConnectionResponse{
		Provider: in.Provider,
		Status:   200,
		Token:    "akdsjfl2q4q35334@$lasd",
	}, nil
}

// Disconnect from a currently authenticated provider
func (serviceImpl *AccountServiceGrpcImpl) Disconnect(ctx context.Context, in *service.DisconnectionRequest) (*service.DisconnectionResponse, error) {
	log.Println(fmt.Sprintf("Received request to disconnect from %s with user_name: %s", in.Provider, in.UserName))

	// Logic to disconnect from provider

	return &service.DisconnectionResponse{
		Status: 200,
	}, nil
}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Unable to listen: %v ", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	return lis
}
