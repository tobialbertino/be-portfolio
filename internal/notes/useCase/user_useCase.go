package usecase

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type UserUseCase interface {
	AddUser(req *domain.ReqAddUser) (*domain.UserId, error)
	GetUserById(id string) (*domain.ResponseUser, error)
	VerifyUserCredential(req *domain.ReqLoginUser) (*domain.ResponseUser, error)
	GetUsersByUsername(username string) (*[]domain.ResponseUser, error)
}
