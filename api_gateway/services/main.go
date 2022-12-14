package services

import (
	"fmt"

	"github.com/jaloldinov/Udevs_task/api_gateway/config"
	"github.com/jaloldinov/Udevs_task/api_gateway/genproto/book_service"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	AuthorService() book_service.AuthorServiceClient
	CategoryService() book_service.CategoryServiceClient
	BookService() book_service.BookServiceClient
}

type grpcClients struct {
	authorService   book_service.AuthorServiceClient
	categoryService book_service.CategoryServiceClient
	bookService     book_service.BookServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connAuthorService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.BookServiceHost, conf.BookServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connCategoryService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.BookServiceHost, conf.BookServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connBookService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.BookServiceHost, conf.BookServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		authorService:   book_service.NewAuthorServiceClient(connAuthorService),
		categoryService: book_service.NewCategoryServiceClient(connCategoryService),
		bookService:     book_service.NewBookServiceClient(connBookService),
	}, nil
}

func (g *grpcClients) AuthorService() book_service.AuthorServiceClient {
	return g.authorService
}

func (g *grpcClients) CategoryService() book_service.CategoryServiceClient {
	return g.categoryService
}

func (g *grpcClients) BookService() book_service.BookServiceClient {
	return g.bookService
}
