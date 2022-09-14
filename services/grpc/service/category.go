package service

import (
	"context"

	"github.com/jaloldinov/Udevs_task/author_service/config"
	ab "github.com/jaloldinov/Udevs_task/author_service/genproto/author_service"
	"github.com/jaloldinov/Udevs_task/author_service/pkg/logger"
	"github.com/jaloldinov/Udevs_task/author_service/storage"
)

type categoryService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	ab.UnimplementedCategoryServiceServer
}

func NewCategoryService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *categoryService {
	return &categoryService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *categoryService) Create(ctx context.Context, req *ab.CreateCategoryRequest) (*ab.Category, error) {
	id, err := s.strg.Category().Create(ctx, req)
	if err != nil {
		s.log.Error("CreateCategory", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &ab.Category{
		Id:           id,
		CategoryName: req.CategoryName,
	}, nil
}

func (s *categoryService) GetAll(ctx context.Context, req *ab.GetAllCategoryRequest) (*ab.GetAllCategoryResponse, error) {
	resp, err := s.strg.Category().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllCategory", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *categoryService) Get(ctx context.Context, req *ab.CategoryId) (*ab.Category, error) {
	resp, err := s.strg.Category().Get(req.Id)

	if err != nil {
		s.log.Error("GetCategoryID", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *categoryService) Update(ctx context.Context, req *ab.Category) (*ab.Result, error) {
	result, err := s.strg.Category().Update(req)
	if err != nil {
		s.log.Error("UpdateCategory", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &ab.Result{
		Result:  result.Result,
		Message: result.Message,
	}, nil
}

func (s *categoryService) Delete(ctx context.Context, req *ab.CategoryId) (*ab.Result, error) {

	result, err := s.strg.Category().Delete(req.Id)
	if err != nil {
		s.log.Error("DeleteCategory", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &ab.Result{
		Result:  result.Result,
		Message: result.Message,
	}, nil
}
