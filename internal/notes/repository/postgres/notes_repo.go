package postgres

import (
	"context"
	"tobialbertino/portfolio-be/internal/notes/models/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NotesRepository interface {
	Add(ctx context.Context, db *pgxpool.Pool, notes *entity.Notes) (int64, error)
	GetAll(ctx context.Context, db *pgxpool.Pool, notes *entity.Notes) (*entity.ListNotes, error)
	GetById(ctx context.Context, db *pgxpool.Pool, id string) (*entity.Notes, error)
	Update(ctx context.Context, db *pgxpool.Pool, notes *entity.Notes) (int64, error)
	Delete(ctx context.Context, db *pgxpool.Pool, notes *entity.Notes) (int64, error)
	VerifyNoteOwner(ctx context.Context, db *pgxpool.Pool, notes *entity.Notes) (*entity.Notes, error)
}
