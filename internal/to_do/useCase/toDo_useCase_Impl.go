package usecase

import (
	"context"
	"errors"
	"time"
	"tobialbertino/portfolio-be/exception"
	"tobialbertino/portfolio-be/internal/to_do/models/domain"
	"tobialbertino/portfolio-be/internal/to_do/models/entity"
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

	request := &entity.ToDo{
		Title:      req.Title,
		Status:     false,
		Created_at: time.Now().Unix(),
		Updated_at: time.Now().Unix(),
	}
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
func (useCase *ToDoUseCaseImpl) Delete(id *int64) (*domain.RowsAffected, error) {
	request := id
	i, err := useCase.ToDoRepository.Delete(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	if i <= 0 {
		err = exception.Wrap("not found", 404, errors.New("rows affected: 0"))
		return nil, err
	}

	response := &domain.RowsAffected{
		RowsAffected: i,
	}
	return response, err
}

func (useCase *ToDoUseCaseImpl) DeleteAll() (*domain.RowsAffected, error) {
	i, err := useCase.ToDoRepository.DeleteAll(context.Background(), useCase.DB)
	if err != nil {
		return nil, err
	}

	response := &domain.RowsAffected{
		RowsAffected: i,
	}
	return response, err
}

// GetAll implements ToDoUseCase
func (useCase *ToDoUseCaseImpl) GetAll() (*[]domain.ResponseToDo, error) {
	var listResult *entity.ListToDo = new(entity.ListToDo)

	listResult, err := useCase.ToDoRepository.GetAll(context.Background(), useCase.DB)
	if err != nil {
		return nil, err
	}

	result := listResult.ToDomain()
	return result, nil
}

// Update implements ToDoUseCase
func (useCase *ToDoUseCaseImpl) Update(req *domain.RequestUpdateToDo) (*domain.RowsAffected, error) {
	err := useCase.Validate.Struct(req)
	if err != nil {
		return nil, err
	}

	request := &entity.ToDo{
		Id:         req.Id,
		Title:      req.Title,
		Status:     req.Status,
		Updated_at: time.Now().Unix(),
	}
	i, err := useCase.ToDoRepository.Update(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	if i <= 0 {
		err = exception.Wrap("not found", 404, errors.New("rows affected: 0"))
		return nil, err
	}

	response := &domain.RowsAffected{
		RowsAffected: i,
	}

	return response, nil
}
