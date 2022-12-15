package usecase

import (
	"context"
	"tobialbertino/portfolio-be/exception"
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	"tobialbertino/portfolio-be/internal/notes/models/entity"
	"tobialbertino/portfolio-be/internal/notes/repository/postgres"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CollabUseCaseImpl struct {
	CollabRepository postgres.CollabRepository
	DB               *pgxpool.Pool
	Validate         *validator.Validate
}

func NewCollabUseCase(collabRepo postgres.CollabRepository, DB *pgxpool.Pool, validate *validator.Validate) CollabUseCase {
	return &CollabUseCaseImpl{
		CollabRepository: collabRepo,
		DB:               DB,
		Validate:         validate,
	}
}

// AddCollaboration implements CollabUseCase
func (uc *CollabUseCaseImpl) AddCollaboration(req *domain.Collab) (domain.Collab, error) {
	err := uc.Validate.Struct(req)
	if err != nil {
		return domain.Collab{}, err
	}

	var request = entity.Collab{
		Id:     uuid.New().String(),
		NoteId: req.NoteId,
		UserId: req.UserId,
	}
	c, err := uc.CollabRepository.AddCollaboration(context.Background(), uc.DB, &request)
	if err != nil {
		return domain.Collab{}, err
	}
	if c.Id == "" {
		return domain.Collab{}, exception.NewClientError("Kolaborasi gagal ditambahkan", 500)
	}
	res := c.ToDomain()

	return res, nil

}

func (uc *CollabUseCaseImpl) DeleteCollaboration(req *domain.Collab) (domain.Collab, error) {
	err := uc.Validate.Struct(req)
	if err != nil {
		return domain.Collab{}, err
	}

	var request = entity.Collab{
		Id:     req.Id,
		NoteId: req.NoteId,
		UserId: req.UserId,
	}
	c, err := uc.CollabRepository.DeleteCollaboration(context.Background(), uc.DB, &request)
	if err != nil {
		return domain.Collab{}, err
	}
	if c.Id == "" {
		return domain.Collab{}, exception.NewClientError("Kolaborasi gagal dihapus", 500)
	}

	res := c.ToDomain()

	return res, nil

}

func (uc *CollabUseCaseImpl) VerifyCollaborator(req *domain.Collab) (domain.Collab, error) {
	err := uc.Validate.Struct(req)
	if err != nil {
		return domain.Collab{}, err
	}

	var request = entity.Collab{
		Id:     req.Id,
		NoteId: req.NoteId,
		UserId: req.UserId,
	}
	c, err := uc.CollabRepository.VerifyCollaborator(context.Background(), uc.DB, &request)
	if err != nil {
		return domain.Collab{}, err
	}
	if c.Id == "" {
		return domain.Collab{}, exception.NewClientError("Kolaborasi gagal diverifikasi", 500)
	}

	res := c.ToDomain()

	return res, nil

}
