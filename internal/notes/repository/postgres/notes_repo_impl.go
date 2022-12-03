package postgres

import (
	"context"
	"tobialbertino/portfolio-be/internal/notes/models/entity"
	"tobialbertino/portfolio-be/pkg/helper"

	"github.com/jackc/pgx/v5"
)

type NotesRepositoryImpl struct {
}

func NewNotesRepository() NotesRepository {
	return &NotesRepositoryImpl{}
}

// Update implements NotesRepository
func (repository *NotesRepositoryImpl) Add(ctx context.Context, db *pgx.Conn, notes *entity.Notes) (int64, error) {
	SQL := `INSERT INTO notes VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	varArgs := []interface{}{
		notes.Id,
		notes.Title,
		notes.Body,
		notes.Tags,
		notes.CreatedAt,
		notes.UpdatedAt,
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

	isTrue := result.RowsAffected()
	return isTrue, nil
}
