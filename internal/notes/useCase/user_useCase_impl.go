package usecase

import (
	"context"
	"fmt"
	"tobialbertino/portfolio-be/exception"
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	"tobialbertino/portfolio-be/internal/notes/models/entity"
	"tobialbertino/portfolio-be/internal/notes/repository/postgres"
	"tobialbertino/portfolio-be/pkg/security"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserUseCaseImpl struct {
	UserRepository postgres.UserRepository
	DB             *pgxpool.Pool
	Validate       *validator.Validate
}

func NewUserUseCase(userRepo postgres.UserRepository, DB *pgxpool.Pool, validate *validator.Validate) UserUseCase {
	return &UserUseCaseImpl{
		UserRepository: userRepo,
		DB:             DB,
		Validate:       validate,
	}
}

// AddUser implements UserUseCase
func (useCase *UserUseCaseImpl) AddUser(req *domain.ReqAddUser) (*domain.UserId, error) {
	// Verifikasi username, pastikan belum terdaftar.
	err := useCase.Validate.Struct(req)
	if err != nil {
		return nil, exception.NewClientError("Error validation", 400)
	}

	var request *entity.User = new(entity.User)
	request = &entity.User{
		Username: req.Username,
	}

	count, err := useCase.UserRepository.CheckUsername(context.Background(), useCase.DB, request)
	if err != nil || count > 0 {
		return nil, exception.NewClientError("Gagal menambahkan user. Username sudah digunakan.", 400)
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

	id, err := useCase.UserRepository.AddUser(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	response := &domain.UserId{
		UserId: id,
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

	if result.Id == "" {
		return nil, exception.NewClientError("User tidak ditemukan", 404)
	}

	res := result.ToDomain()

	return res, nil
}

func (useCase *UserUseCaseImpl) VerifyUserCredential(req *entity.User) (*entity.User, error) {
	var request *entity.User = new(entity.User)

	request = &entity.User{
		Username:  req.Username,
		Passwword: req.Passwword,
	}

	// verify user
	result, err := useCase.UserRepository.VerifyUserCredential(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	// compare password
	isValid := security.CheckPasswordHash(req.Passwword, result.Passwword)
	if !isValid {
		return nil, exception.NewClientError("Kredensial yang Anda berikan salah", 401)
	}

	return result, nil
}

func (useCase *UserUseCaseImpl) GetUsersByUsername(username string) (*[]domain.ResponseUser, error) {
	var request *entity.User = new(entity.User)

	request = &entity.User{
		Username: username,
	}

	result, err := useCase.UserRepository.GetUsersByUsername(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	res := result.ToDomain()

	return res, nil
}
