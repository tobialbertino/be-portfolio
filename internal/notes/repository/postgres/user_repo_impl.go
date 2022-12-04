package postgres

import (
	"context"
	"errors"
	"tobialbertino/portfolio-be/exception"
	"tobialbertino/portfolio-be/internal/notes/models/entity"
	"tobialbertino/portfolio-be/pkg/helper"

	"github.com/jackc/pgx/v5"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// CheckUsername implements UserRepository
func (repo *UserRepositoryImpl) CheckUsername(ctx context.Context, db *pgx.Conn, user *entity.User) (bool, error) {
	SQL := `SELECT username FROM users WHERE username = $1`
	varArgs := []interface{}{
		user.Username,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return false, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	rows, err := tx.Query(ctx, SQL, varArgs...)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	counter := 0
	for rows.Next() {
		counter++
	}

	if counter > 0 {
		return false, exception.Wrap("username exist", 403, errors.New("username exist"))
	}
	return true, nil
}

func (repo *UserRepositoryImpl) AddUser(ctx context.Context, db *pgx.Conn, user *entity.User) (int64, error) {
	SQL := `INSERT INTO users VALUES($1, $2, $3, $4) RETURNING id`
	varArgs := []interface{}{
		user.Id,
		user.Username,
		user.Passwword,
		user.FullName,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return -1, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	result, err := tx.Exec(ctx, SQL, varArgs...)
	if err != nil {
		return -1, err
	}

	i := result.RowsAffected()
	if i == 0 {
		return -1, exception.Wrap("repository", 400, errors.New("add user fail"))
	}

	return i, nil
}

func (repo *UserRepositoryImpl) GetUserById(ctx context.Context, db *pgx.Conn, user *entity.User) (*entity.User, error) {
	var res *entity.User = new(entity.User)

	SQL := `SELECT id, username, fullname FROM users WHERE id = $1`
	varArgs := []interface{}{
		user.Id,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	row := tx.QueryRow(ctx, SQL, varArgs...)
	row.Scan(&res.Id, &res.Username, &res.FullName)

	return res, nil
}
