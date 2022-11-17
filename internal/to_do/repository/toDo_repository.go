package repository

import (
	"tobialbertino/portfolio-be/internal/to_do/models/domain"

	"github.com/jackc/pgx/v5"
)

type ToDoRepository interface {
	Create(db *pgx.Conn, tx pgx.Tx, toDo *domain.ToDo) *domain.ToDo
	Update(db *pgx.Conn, tx pgx.Tx, toDo *domain.ToDo) *domain.ToDo
	Delete(db *pgx.Conn, tx pgx.Tx, toDo *domain.ToDo)
	GetAll(db *pgx.Conn, tx pgx.Tx) *[]domain.ToDo
}
