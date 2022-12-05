package usecase

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type AuthUseCase interface {
	AddRefreshToken(req *domain.Token) (*domain.RowsAffected, error)
	DeleteRefreshToken(req *domain.Token) (*domain.RowsAffected, error)
}
