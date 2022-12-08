package usecase

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type NotesUseCase interface {
	Add(req *domain.ReqAddNote) (*domain.NoteId, error)
	GetAll() (*[]domain.Notes, error)
	GetById(id string) (*domain.Notes, error)
	Update(req *domain.ReqAddNote, id string) (*domain.Notes, error)
	Delete(id string) (*domain.RowsAffected, error)
}
