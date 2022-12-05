package postgres

import (
	"context"
	"tobialbertino/portfolio-be/internal/notes/models/entity"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	CheckUsername(ctx context.Context, db *pgx.Conn, user *entity.User) (bool, error)
	AddUser(ctx context.Context, db *pgx.Conn, user *entity.User) (string, error)
	GetUserById(ctx context.Context, db *pgx.Conn, user *entity.User) (*entity.User, error)
	VerifyUserCredential(ctx context.Context, db *pgx.Conn, user *entity.User) (*entity.User, error)
	GetUsersByUsername(ctx context.Context, db *pgx.Conn, user *entity.User) (*entity.ListUser, error)
}
