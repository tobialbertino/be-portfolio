package usecase

import (
	"context"
	"tobialbertino/portfolio-be/internal/to_do/models/domain"
	"tobialbertino/portfolio-be/internal/to_do/repository"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
)

type ToDoUseCaseImpl struct {
	ToDoRepository repository.ToDoRepository
	DB             *pgx.Conn
	Validate       *validator.Validate
}

func NewToDoUseCase(toDoRepo repository.ToDoRepository, DB *pgx.Conn, validate *validator.Validate) ToDoUseCase {
	return &ToDoUseCaseImpl{
		ToDoRepository: toDoRepo,
		DB:             DB,
		Validate:       validate,
	}
}

// Create implements ToDoUseCase
func (useCase *ToDoUseCaseImpl) Create(req *domain.RequestToDo) (*domain.RowsAffected, error) {
	err := useCase.Validate.Struct(req)
	if err != nil {
		return nil, err
	}

	request := req.ToEntity()
	i, err := useCase.ToDoRepository.Create(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	response := &domain.RowsAffected{
		RowsAffected: i,
	}
	return response, err
}

// Delete implements ToDoUseCase
func (useCase *ToDoUseCaseImpl) Delete(req *domain.RequestToDo) error {
	panic("unimplemented")
}

// GetAll implements ToDoUseCase
func (useCase *ToDoUseCaseImpl) GetAll() (*[]domain.ResponseToDo, error) {
	panic("unimplemented")
}

// Update implements ToDoUseCase
func (useCase *ToDoUseCaseImpl) Update(req *domain.RequestToDo) (*domain.RequestToDo, error) {
	panic("unimplemented")
}
