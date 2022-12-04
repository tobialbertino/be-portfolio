package usecase

import (
	"context"
	"fmt"
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	"tobialbertino/portfolio-be/internal/notes/models/entity"
	"tobialbertino/portfolio-be/internal/notes/repository/postgres"
	"tobialbertino/portfolio-be/pkg/security"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserUseCaseImpl struct {
	UserRepository postgres.UserRepository
	DB             *pgx.Conn
	Validate       *validator.Validate
}

func NewUserUseCase(userRepo postgres.UserRepository, DB *pgx.Conn, validate *validator.Validate) UserUseCase {
	return &UserUseCaseImpl{
		UserRepository: userRepo,
		DB:             DB,
		Validate:       validate,
	}
}

// AddUser implements UserUseCase
func (useCase *UserUseCaseImpl) AddUser(req *domain.ReqAddUser) (*domain.RowsAffected, error) {
	// Verifikasi username, pastikan belum terdaftar.
	var request *entity.User = new(entity.User)
	request = &entity.User{
		Username: req.Username,
	}

	bool, err := useCase.UserRepository.CheckUsername(context.Background(), useCase.DB, request)
	if err != nil && !bool {
		return nil, err
	}

	// Bila verifikasi lolos, maka masukkan user baru ke database.
	// Hash Passwword
	hashedPassword, err := security.HashPassword(req.Passwword)
	if err != nil {
		return nil, err
	}
	request = &entity.User{
		Id:        fmt.Sprintf("user-%v", uuid.New().String()),
		Username:  req.Username,
		Passwword: hashedPassword,
		FullName:  req.FullName,
	}

	i, err := useCase.UserRepository.AddUser(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	response := &domain.RowsAffected{
		RowsAffected: i,
	}
	return response, err
}

func (useCase *UserUseCaseImpl) GetUserById(id string) (*domain.ResponseUser, error) {
	var request *entity.User = new(entity.User)

	request = &entity.User{
		Id: id,
	}
	result, err := useCase.UserRepository.GetUserById(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	res := result.ToDomain()

	return res, nil
}
