package usecase

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type AuthUseCase interface {
	AddRefreshToken(req *domain.ReqLoginUser) (*domain.ResToken, error)
	DeleteRefreshToken(req *domain.ReqLoginUser) (*domain.ResToken, error)
}
