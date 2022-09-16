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

type bookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) storage.BookI {
	return &bookRepo{
		db: db,
	}
}

func (r *bookRepo) Create(ctx context.Context, entity *ab.CreateBookRequest) (id string, err error) {
	query := `
		INSERT INTO books (
			id,
			book_name,
			author_id,
			category_id
		) 
		 VALUES ($1, $2, $3, $4)
	`

	id = uuid.NewString()

	_, err = r.db.Exec(
		ctx,
		query,
		id,
		entity.Name,
		entity.AuthorId,
		entity.CategoryId,
	)

	if err != nil {
		return "", fmt.Errorf("error while creating Book. err: %w", err)
	}

	return id, nil
}

func (r *bookRepo) GetAll(ctx context.Context, req *ab.GetAllBookRequest) (*ab.GetAllBookResponse, error) {
	var (
		resp   ab.GetAllBookResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND book_name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM books WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT
				id,
				book_name,
				author_id,
				category_id
			FROM books
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
		var book ab.Book

		err = rows.Scan(
			&book.Id,
			&book.Name,
			&book.AuthorId,
			&book.CategoryId,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning bookList err: %w", err)
		}

		resp.Books = append(resp.Books, &book)
	}

	return &resp, nil
}

func (r *bookRepo) Get(id string) (*ab.Book, error) {
	var book ab.Book

	query := `
		SELECT
		id,
		book_name,
		author_id,
		category_id
		FROM
			books
		WHERE id = $1
	`

	row := r.db.QueryRow(context.Background(), query, id)
	err := row.Scan(
		&book.Id,
		&book.Name,
		&book.AuthorId,
		&book.CategoryId,
	)

	if err != nil {
		return nil, fmt.Errorf("error while Getting book err: %w", err)
	}

	return &book, nil
}

func (r *bookRepo) Update(req *ab.Book) (*ab.Result, error) {

	query := `
		UPDATE books SET
		book_name = $1,
		author_id = $2,
		category_id = $3,
		updated_at = now()
		WHERE id = $4
	`
	result, err := r.db.Exec(
		context.Background(),
		query,
		req.Name,
		req.AuthorId,
		req.CategoryId,
		req.Id,
	)

	if err != nil {
		return nil, fmt.Errorf("error while updating book err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return nil, errors.New("not found id")
	}

	return &ab.Result{
		Result:  "OK",
		Message: "Updated",
	}, nil
}

func (r *bookRepo) Delete(id string) (*ab.Result, error) {

	query := `
	   		DELETE FROM books
	   		WHERE id = $1
	   	`

	result, err := r.db.Exec(context.Background(), query, id)

	if err != nil {
		return nil, fmt.Errorf("error while deleting book err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return nil, errors.New("not found id")
	}
	return &ab.Result{
		Result:  "OK",
		Message: "Deleted",
	}, nil
}
