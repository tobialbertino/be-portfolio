package usecase

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type NotesUseCase interface {
	Add(req *domain.ReqAddNote) (*domain.RowsAffected, error)
}
