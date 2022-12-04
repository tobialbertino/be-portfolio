package usecase

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type UserUseCase interface {
	AddUser(req *domain.ReqAddUser) (*domain.RowsAffected, error)
	GetUserById(id string) (*domain.ResponseUser, error)
}
