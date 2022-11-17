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
func (repository *ToDoRepositoryImpl) Delete(ctx context.Context, db *pgx.Conn, toDo *entity.ToDo) error {
	panic("unimplemented")
}

// GetAll implements ToDoRepository
func (repository *ToDoRepositoryImpl) GetAll(ctx context.Context, db *pgx.Conn) (*[]entity.ToDo, error) {
	panic("unimplemented")
}

// Update implements ToDoRepository
func (repository *ToDoRepositoryImpl) Update(ctx context.Context, db *pgx.Conn, toDo *entity.ToDo) (*entity.ToDo, error) {
	panic("unimplemented")
}
