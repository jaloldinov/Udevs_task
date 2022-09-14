package service

import (
	"context"

	ab "github.com/jaloldinov/Udevs_task/author_service/genproto/author_service"

	"github.com/jaloldinov/Udevs_task/author_service/config"
	"github.com/jaloldinov/Udevs_task/author_service/pkg/logger"
	"github.com/jaloldinov/Udevs_task/author_service/storage"
)

type authorService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	ab.UnimplementedAuthorServiceServer
}

func NewAuthorService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *authorService {
	return &authorService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *authorService) Create(ctx context.Context, req *ab.CreateAuthorRequest) (*ab.Author, error) {
	id, err := s.strg.Author().Create(ctx, req)
	if err != nil {
		s.log.Error("CreateAuthor", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &ab.Author{
		Id:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}, nil
}

func (s *authorService) GetAll(ctx context.Context, req *ab.GetAllAuthorRequest) (*ab.GetAllAuthorResponse, error) {
	resp, err := s.strg.Author().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllAuthor", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *authorService) Get(ctx context.Context, req *ab.AuthorId) (*ab.Author, error) {
	resp, err := s.strg.Author().Get(req.Id)

	if err != nil {
		s.log.Error("GetAuthorID", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *authorService) Update(ctx context.Context, req *ab.Author) (*ab.Result, error) {
	result, err := s.strg.Author().Update(req)
	if err != nil {
		s.log.Error("UpdateAuthor", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &ab.Result{
		Result:  result,
		Message: "Author updated",
	}, nil
}

func (s *authorService) Delete(ctx context.Context, req *ab.AuthorId) (*ab.Result, error) {

	result, err := s.strg.Author().Delete(req.Id)
	if err != nil {
		s.log.Error("DeleteAuthor", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &ab.Result{
		Result:  result.Result,
		Message: result.Message,
	}, nil
}
