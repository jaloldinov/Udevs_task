package postgres

import (
	"context"
	"errors"
	"fmt"

	ab "github.com/jaloldinov/Udevs_task/author_service/genproto/author_service"
	"github.com/jaloldinov/Udevs_task/author_service/pkg/helper"
	"github.com/jaloldinov/Udevs_task/author_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type authorRepo struct {
	db *pgxpool.Pool
}

func NewAuthorRepo(db *pgxpool.Pool) storage.AuthorI {
	return &authorRepo{
		db: db,
	}
}

func (r *authorRepo) Create(ctx context.Context, entity *ab.CreateAuthorRequest) (id string, err error) {
	query := `
		INSERT INTO authors (
			id,
			firstname,
			lastname
		) 
		 VALUES ($1, $2, $3)
	`

	id = uuid.NewString()

	_, err = r.db.Exec(
		ctx,
		query,
		id,
		entity.FirstName,
		entity.LastName,
	)

	if err != nil {
		return "", fmt.Errorf("error while creating Author. err: %w", err)
	}

	return id, nil
}

func (r *authorRepo) GetAll(ctx context.Context, req *ab.GetAllAuthorRequest) (*ab.GetAllAuthorResponse, error) {
	var (
		resp   ab.GetAllAuthorResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND firstname || lastname ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM authors WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT
				id,
				firstname,
				lastname
			FROM authors
			WHERE true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = req.Limit
	params["offset"] = req.Offset

	q, arr = helper.ReplaceQueryParams(query, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}

	for rows.Next() {
		var author ab.Author

		err = rows.Scan(
			&author.Id,
			&author.FirstName,
			&author.LastName,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning authorList err: %w", err)
		}

		resp.Authors = append(resp.Authors, &author)
	}

	return &resp, nil
}

func (r *authorRepo) Get(id string) (*ab.Author, error) {
	var author ab.Author

	query := `
		SELECT
		id,
		firstname,
		lastname
		FROM
			authors
		WHERE id = $1
	`

	row := r.db.QueryRow(context.Background(), query, id)
	err := row.Scan(
		&author.Id,
		&author.FirstName,
		&author.LastName,
	)

	if err != nil {
		return nil, fmt.Errorf("error while Getting author err: %w", err)
	}

	return &author, nil
}

func (r *authorRepo) Update(req *ab.Author) (string, error) {

	query := `
		UPDATE authors SET
		firstname = $1,
		lastname = $2,
		updated_at = now()
		WHERE id = $3
	`
	result, err := r.db.Exec(
		context.Background(),
		query,
		req.FirstName,
		req.LastName,
		req.Id,
	)

	if err != nil {
		return "", fmt.Errorf("error while updating author err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return "not found id", nil
	}

	return "Updated", nil
}

func (r *authorRepo) Delete(id string) (*ab.Result, error) {

	query := `
	   		DELETE FROM authors
	   		WHERE id = $1
	   	`

	result, err := r.db.Exec(context.Background(), query, id)

	if err != nil {
		return nil, fmt.Errorf("error while deleting author err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return nil, errors.New("not found id")
	}
	return &ab.Result{
		Result:  "OK",
		Message: "Deleted",
	}, nil
}
