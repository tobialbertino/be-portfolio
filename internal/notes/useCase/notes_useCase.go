package usecase

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type NotesUseCase interface {
	Add(req *domain.ReqAddNote, id string) (*domain.NoteId, error)
	GetAll(owner string) (*[]domain.Notes, error)
	GetById(id string) (*domain.Notes, error)
	Update(req *domain.ReqAddNote, id string) (*domain.Notes, error)
	Delete(id string) (*domain.RowsAffected, error)

	VerifyNoteOwner(id, owner string) (bool, error)
	VerifyNoteAccess(noteId, userId string) (bool, error)
}
