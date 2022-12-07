package repository

import (
	"context"
	"tobialbertino/portfolio-be/internal/to_do/models/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ToDoRepository interface {
	Create(ctx context.Context, db *pgxpool.Pool, toDo *entity.ToDo) (int64, error)
	Update(ctx context.Context, db *pgxpool.Pool, toDo *entity.ToDo) (int64, error)
	Delete(ctx context.Context, db *pgxpool.Pool, id *int64) (int64, error)
	DeleteAll(ctx context.Context, db *pgxpool.Pool) (int64, error)
	GetAll(ctx context.Context, db *pgxpool.Pool) (*entity.ListToDo, error)
}
