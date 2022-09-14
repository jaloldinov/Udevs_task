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

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) storage.CategoryI {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Create(ctx context.Context, entity *ab.CreateCategoryRequest) (id string, err error) {
	query := `
		INSERT INTO book_category (
			id,
			category_name
		) 
		 VALUES ($1, $2)
	`

	id = uuid.NewString()

	_, err = r.db.Exec(
		ctx,
		query,
		id,
		entity.CategoryName,
	)

	if err != nil {
		return "", fmt.Errorf("error while creating Category. err: %w", err)
	}

	return id, nil
}

func (r *categoryRepo) GetAll(ctx context.Context, req *ab.GetAllCategoryRequest) (*ab.GetAllCategoryResponse, error) {
	var (
		resp   ab.GetAllCategoryResponse
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if req.Search != "" {
		filter += " AND category_name ILIKE '%' || :search || '%' "
		params["search"] = req.Search
	}

	countQuery := `SELECT count(1) FROM book_category WHERE true ` + filter

	q, arr := helper.ReplaceQueryParams(countQuery, params)
	err = r.db.QueryRow(ctx, q, arr...).Scan(
		&resp.Count,
	)

	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT
				id,
				category_name
			FROM book_category
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
		var category ab.Category

		err = rows.Scan(
			&category.Id,
			&category.CategoryName,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning categoryList err: %w", err)
		}

		resp.Categories = append(resp.Categories, &category)
	}

	return &resp, nil
}

func (r *categoryRepo) Get(id string) (*ab.Category, error) {
	var category ab.Category

	query := `
		SELECT
		id,
		category_name
		FROM
		book_category
		WHERE id = $1
	`

	row := r.db.QueryRow(context.Background(), query, id)
	err := row.Scan(
		&category.Id,
		&category.CategoryName,
	)

	if err != nil {
		return nil, fmt.Errorf("error while Getting category err: %w", err)
	}

	return &category, nil
}

func (r *categoryRepo) Update(req *ab.Category) (*ab.Result, error) {

	query := `
		UPDATE book_category SET
		category_name = $1,
		updated_at = now()
		WHERE id = $2
	`
	result, err := r.db.Exec(
		context.Background(),
		query,
		req.CategoryName,
		req.Id,
	)

	if err != nil {
		return nil, fmt.Errorf("error while updating category err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return nil, errors.New("not found id")
	}

	return &ab.Result{
		Result:  "OK",
		Message: "Updated",
	}, nil
}

func (r *categoryRepo) Delete(id string) (*ab.Result, error) {

	query := `
	   		DELETE FROM book_category
	   		WHERE id = $1
	   	`

	result, err := r.db.Exec(context.Background(), query, id)

	if err != nil {
		return nil, fmt.Errorf("error while deleting category err: %w", err)
	}

	if i := result.RowsAffected(); i == 0 {
		return nil, errors.New("not found id")
	}
	return &ab.Result{
		Result:  "OK",
		Message: "Deleted",
	}, nil
}
