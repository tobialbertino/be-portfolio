package repository

import (
	"tobialbertino/portfolio-be/internal/to_do/models/domain"

	"github.com/jackc/pgx/v5"
)

type ToDoRepositoryImpl struct {
	Conn *pgx.Conn
}

func NewToDoRepository(conn *pgx.Conn) ToDoRepository {
	return &ToDoRepositoryImpl{
		Conn: conn,
	}
}

// Create implements ToDoRepository
func (*ToDoRepositoryImpl) Create(db *pgx.Conn, tx pgx.Tx, toDo *domain.ToDo) *domain.ToDo {
	panic("unimplemented")
}

// Delete implements ToDoRepository
func (*ToDoRepositoryImpl) Delete(db *pgx.Conn, tx pgx.Tx, toDo *domain.ToDo) {
	panic("unimplemented")
}

// GetAll implements ToDoRepository
func (*ToDoRepositoryImpl) GetAll(db *pgx.Conn, tx pgx.Tx) *[]domain.ToDo {
	panic("unimplemented")
}

// Update implements ToDoRepository
func (*ToDoRepositoryImpl) Update(db *pgx.Conn, tx pgx.Tx, toDo *domain.ToDo) *domain.ToDo {
	panic("unimplemented")
}
