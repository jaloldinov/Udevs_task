package service

import (
	"context"

	"github.com/jaloldinov/Udevs_task/author_service/config"
	ab "github.com/jaloldinov/Udevs_task/author_service/genproto/author_service"
	"github.com/jaloldinov/Udevs_task/author_service/pkg/logger"
	"github.com/jaloldinov/Udevs_task/author_service/storage"
)

type bookService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	ab.UnimplementedBookServiceServer
}

func NewBookService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *bookService {
	return &bookService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (s *bookService) Create(ctx context.Context, req *ab.CreateBookRequest) (*ab.Book, error) {
	id, err := s.strg.Book().Create(ctx, req)
	if err != nil {
		s.log.Error("CreateBook", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &ab.Book{
		Id:         id,
		Name:       req.Name,
		AuthorId:   req.AuthorId,
		CategoryId: req.CategoryId,
	}, nil
}

func (s *bookService) GetAll(ctx context.Context, req *ab.GetAllBookRequest) (*ab.GetAllBookResponse, error) {
	resp, err := s.strg.Book().GetAll(ctx, req)
	if err != nil {
		s.log.Error("GetAllBook", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *bookService) Get(ctx context.Context, req *ab.BookId) (*ab.Book, error) {
	resp, err := s.strg.Book().Get(req.Id)

	if err != nil {
		s.log.Error("GetBookID", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (s *bookService) Update(ctx context.Context, req *ab.Book) (*ab.Result, error) {
	result, err := s.strg.Book().Update(req)
	if err != nil {
		s.log.Error("UpdateBook", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &ab.Result{
		Result:  result.Result,
		Message: result.Message,
	}, nil
}

func (s *bookService) Delete(ctx context.Context, req *ab.BookId) (*ab.Result, error) {

	result, err := s.strg.Book().Delete(req.Id)
	if err != nil {
		s.log.Error("DeleteBook", logger.Any("req", req), logger.Error(err))
		return nil, err
	}

	return &ab.Result{
		Result:  result.Result,
		Message: result.Message,
	}, nil
}
