package usecase

import (
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	"tobialbertino/portfolio-be/internal/notes/repository/postgres"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
)

type AuthUseCaseImpl struct {
	AuthRepository postgres.AuthRepository
	DB             *pgx.Conn
	Validate       *validator.Validate
}

func NewAuthUseCase(authRepo postgres.AuthRepository, DB *pgx.Conn, validate *validator.Validate) AuthUseCase {
	return &AuthUseCaseImpl{
		AuthRepository: authRepo,
		DB:             DB,
		Validate:       validate,
	}
}

// AddRefreshToken implements AuthUseCase
func (*AuthUseCaseImpl) AddRefreshToken(req *domain.Token) (*domain.RowsAffected, error) {
	panic("unimplemented")
}

// DeleteRefreshToken implements AuthUseCase
func (*AuthUseCaseImpl) DeleteRefreshToken(req *domain.Token) (*domain.RowsAffected, error) {
	panic("unimplemented")
}
