package usecase

import (
	"tobialbertino/portfolio-be/internal/to_do/models/domain"
)

type ToDoUseCase interface {
	Create(req *domain.RequestToDo) (*domain.RowsAffected, error)
	Update(req *domain.RequestUpdateToDo) (*domain.SuccessReturn, error)
	Delete(id int64) (*domain.RowsAffected, error)
	DeleteAll() (*domain.RowsAffected, error)
	GetAll() (*[]domain.ResponseToDo, error)
}
