package repository

import (
	"context"
	"tobialbertino/portfolio-be/internal/to_do/models/entity"

	"github.com/jackc/pgx/v5"
)

type ToDoRepository interface {
	Create(ctx context.Context, db *pgx.Conn, toDo *entity.ToDo) (int64, error)
	Update(ctx context.Context, db *pgx.Conn, toDo *entity.ToDo) (bool, error)
	Delete(ctx context.Context, db *pgx.Conn, id int64) (int64, error)
	DeleteAll(ctx context.Context, db *pgx.Conn) (int64, error)
	GetAll(ctx context.Context, db *pgx.Conn) (*entity.ListToDo, error)
}
