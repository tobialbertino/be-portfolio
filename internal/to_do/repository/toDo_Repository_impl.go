package repository

import (
	"context"
	"tobialbertino/portfolio-be/internal/to_do/models/entity"
	"tobialbertino/portfolio-be/pkg/helper"

	"github.com/jackc/pgx/v5"
)

type ToDoRepositoryImpl struct {
}

func NewToDoRepository() ToDoRepository {
	return &ToDoRepositoryImpl{}
}

// Create implements ToDoRepository
func (repository *ToDoRepositoryImpl) Create(ctx context.Context, db *pgx.Conn, toDo *entity.ToDo) (int64, error) {
	SQL := `INSERT INTO to_do (title, status, created_at, updated_at) VALUES ($1, $2, $3, $4)`
	varArgs := []interface{}{
		toDo.Title,
		toDo.Status,
		toDo.Created_at,
		toDo.Updated_at,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	result, err := tx.Exec(ctx, SQL, varArgs...)
	if err != nil {
		return -1, err
	}

	i := result.RowsAffected()
	return i, nil
}

// Delete implements ToDoRepository
func (repository *ToDoRepositoryImpl) Delete(ctx context.Context, db *pgx.Conn, id int64) (int64, error) {
	SQL := `DELETE FROM to_do WHERE id = $1`
	varArgs := []interface{}{
		id,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	result, err := tx.Exec(ctx, SQL, varArgs...)
	if err != nil {
		return -1, err
	}

	i := result.RowsAffected()
	return i, nil
}

func (repository *ToDoRepositoryImpl) DeleteAll(ctx context.Context, db *pgx.Conn) (int64, error) {
	SQL := `DELETE FROM to_do`

	tx, err := db.Begin(ctx)
	if err != nil {
		return -1, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	result, err := tx.Exec(ctx, SQL)
	if err != nil {
		return -1, err
	}

	i := result.RowsAffected()
	return i, nil
}

// GetAll implements ToDoRepository
func (repository *ToDoRepositoryImpl) GetAll(ctx context.Context, db *pgx.Conn) (*entity.ListToDo, error) {
	var (
		ListResult *entity.ListToDo = new(entity.ListToDo)
		result     *entity.ToDo     = new(entity.ToDo)
	)

	SQL := `SELECT id, title, status, created_at, updated_at FROM to_do ORDER BY id ASC`

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	rows, err := tx.Query(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Title, &result.Status, &result.Created_at, &result.Updated_at)
		if err != nil {
			return nil, err
		}
		*ListResult = append(*ListResult, *result)
	}

	return ListResult, nil
}

// Update implements ToDoRepository
func (repository *ToDoRepositoryImpl) Update(ctx context.Context, db *pgx.Conn, toDo *entity.ToDo) (int64, error) {
	SQL := `UPDATE to_do SET title = $1, status = $2, updated_at = $3 WHERE id = $4`
	varArgs := []interface{}{
		toDo.Title,
		toDo.Status,
		toDo.Updated_at,
		toDo.Id,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return -1, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	result, err := tx.Exec(ctx, SQL, varArgs...)
	if err != nil {
		return -1, err
	}

	isTrue := result.RowsAffected()
	return isTrue, nil
}
