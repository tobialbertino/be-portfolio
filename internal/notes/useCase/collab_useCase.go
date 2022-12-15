package usecase

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type CollabUseCase interface {
	AddCollaboration(req *domain.Collab) (domain.Collab, error)
	DeleteCollaboration(req *domain.Collab) (domain.Collab, error)
	VerifyCollaborator(req *domain.Collab) (domain.Collab, error)
}
