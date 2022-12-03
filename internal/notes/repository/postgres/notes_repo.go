package postgres

import (
	"context"
	"tobialbertino/portfolio-be/internal/notes/models/entity"

	"github.com/jackc/pgx/v5"
)

type NotesRepository interface {
	Add(ctx context.Context, db *pgx.Conn, notes *entity.Notes) (int64, error)
}
