package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jaloldinov/Udevs_task/author_service/config"
	"github.com/jaloldinov/Udevs_task/author_service/storage"
)

type Store struct {
	db       *pgxpool.Pool
	author   storage.AuthorI
	category storage.CategoryI
}

func NewPostgres(psqlConnString string, cfg config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(psqlConnString)
	if err != nil {
		return nil, err
	}

	config.AfterConnect = nil
	config.MaxConns = int32(cfg.PostgresMaxConnections)

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	return &Store{
		db: pool,
	}, err
}

func (s *Store) Author() storage.AuthorI {
	if s.author == nil {
		s.author = NewAuthorRepo(s.db)
	}
	return s.author
}

func (s *Store) Category() storage.CategoryI {
	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}
	return s.category
}
