package postgres

import (
	"context"
	"tobialbertino/portfolio-be/internal/notes/models/entity"

	"github.com/jackc/pgx/v5"
)

type AuthRepository interface {
	AddRefreshToken(ctx context.Context, db *pgx.Conn, notes *entity.Token) (int64, error)
	DeleteRefreshToken(ctx context.Context, db *pgx.Conn, notes *entity.Token) (int64, error)
}
