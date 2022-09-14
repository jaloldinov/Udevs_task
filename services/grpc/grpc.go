package grpc

import (
	"github.com/jaloldinov/Udevs_task/author_service/config"
	"github.com/jaloldinov/Udevs_task/author_service/genproto/author_service"
	"github.com/jaloldinov/Udevs_task/author_service/grpc/service"
	"github.com/jaloldinov/Udevs_task/author_service/pkg/logger"
	"github.com/jaloldinov/Udevs_task/author_service/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	author_service.RegisterAuthorServiceServer(grpcServer, service.NewAuthorService(cfg, log, strg))
	author_service.RegisterCategoryServiceServer(grpcServer, service.NewCategoryService(cfg, log, strg))
	// author_service.RegisterBookServiceServer(grpcServer, service.NewBookService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
