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

func (repository *NotesRepositoryImpl) GetAll(ctx context.Context, db *pgx.Conn) (*entity.ListNotes, error) {
	var (
		ListResult *entity.ListNotes = new(entity.ListNotes)
		result     *entity.Notes     = new(entity.Notes)
	)

	SQL := `SELECT * FROM notes ORDER BY created_at ASC`

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	rows, err := tx.Query(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Title, &result.Body, &result.Tags, &result.CreatedAt, &result.UpdatedAt, &result.Owner)
		if err != nil {
			return nil, err
		}
		*ListResult = append(*ListResult, *result)
	}

	return ListResult, nil
}

func (repository *NotesRepositoryImpl) GetById(ctx context.Context, db *pgx.Conn, id string) (*entity.Notes, error) {
	var (
		result *entity.Notes = new(entity.Notes)
	)

	SQL := `SELECT * FROM notes WHERE id = $1`

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	row := tx.QueryRow(ctx, SQL, id)
	row.Scan(&result.Id, &result.Title, &result.Body, &result.Tags, &result.CreatedAt, &result.UpdatedAt, &result.Owner)

	// if result.Id == "" {
	// 	return result, exception.Wrap("repo", 404, errors.New("error not found"))
	// }

	return result, nil
}
