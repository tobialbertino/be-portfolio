package postgres

import (
	"context"
	"tobialbertino/portfolio-be/internal/notes/models/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	AddRefreshToken(ctx context.Context, db *pgxpool.Pool, token *entity.Token) (int64, error)
	VerifyRefreshToken(ctx context.Context, db *pgxpool.Pool, token *entity.Token) (string, error)
	DeleteRefreshToken(ctx context.Context, db *pgxpool.Pool, token *entity.Token) (int64, error)
}
