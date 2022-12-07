package usecase

import (
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	"tobialbertino/portfolio-be/internal/notes/models/entity"
)

type UserUseCase interface {
	AddUser(req *domain.ReqAddUser) (*domain.UserId, error)
	GetUserById(id string) (*domain.ResponseUser, error)
	GetUsersByUsername(username string) (*[]domain.ResponseUser, error)

	VerifyUserCredential(req *entity.User) (*entity.User, error)
}
