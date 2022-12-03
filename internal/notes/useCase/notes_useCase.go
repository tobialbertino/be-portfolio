package usecase

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type NotesUseCase interface {
	Add(req *domain.ReqAddNote) (*domain.RowsAffected, error)
	GetAll() (*[]domain.Notes, error)
	GetById(id string) (*domain.Notes, error)
	Update(req *domain.ReqAddNote, id string) (*domain.RowsAffected, error)
	Delete(id string) (*domain.RowsAffected, error)
}
