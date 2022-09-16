package services

import (
	"fmt"

	"github.com/jaloldinov/Udevs_task/api_gateway/config"
	"github.com/jaloldinov/Udevs_task/api_gateway/genproto/book_service"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	AuthorService() book_service.AuthorServiceClient
}

type grpcClients struct {
	authorService book_service.AuthorServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connAuthorService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.BookServiceHost, conf.BookServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		authorService: book_service.NewAuthorServiceClient(connAuthorService),
	}, nil
}

func (g *grpcClients) AuthorService() book_service.AuthorServiceClient {
	return g.authorService
}
