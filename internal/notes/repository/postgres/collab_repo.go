package postgres

import (
	"context"
	"tobialbertino/portfolio-be/internal/notes/models/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CollabRepository interface {
	AddCollaboration(ctx context.Context, db *pgxpool.Pool, notes *entity.Collab) (entity.Collab, error)
	DeleteCollaboration(ctx context.Context, db *pgxpool.Pool, notes *entity.Collab) (entity.Collab, error)
	VerifyCollaborator(ctx context.Context, db *pgxpool.Pool, notes *entity.Collab) (entity.Collab, error)
}
