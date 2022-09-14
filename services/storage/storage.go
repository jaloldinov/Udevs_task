package storage

import (
	"context"
	"errors"

	ab "github.com/jaloldinov/Udevs_task/author_service/genproto/author_service"
)

var ErrorTheSameId = errors.New("cannot use the same uuid for 'id' and 'parent_id' fields")
var ErrorProjectId = errors.New("not valid 'project_id'")

type StorageI interface {
	Author() AuthorI
}

type AuthorI interface {
	Create(ctx context.Context, entity *ab.CreateAuthorRequest) (id string, err error)
	GetAll(ctx context.Context, req *ab.GetAllAuthorRequest) (*ab.GetAllAuthorResponse, error)
	Get(id string) (*ab.Author, error)
	Update(req *ab.Author) (string, error)
	Delete(id string) (*ab.Result, error)
}
