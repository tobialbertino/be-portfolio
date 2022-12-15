package postgres

import (
	"context"
	"errors"
	"tobialbertino/portfolio-be/exception"
	"tobialbertino/portfolio-be/internal/notes/models/entity"
	"tobialbertino/portfolio-be/pkg/helper"

	"github.com/jackc/pgx/v5/pgxpool"
)

type NotesRepositoryImpl struct {
}

func NewNotesRepository() NotesRepository {
	return &NotesRepositoryImpl{}
}

// Update implements NotesRepository
func (repository *NotesRepositoryImpl) Add(ctx context.Context, db *pgxpool.Pool, notes *entity.Notes) (int64, error) {
	SQL := `INSERT INTO notes VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	varArgs := []interface{}{
		notes.Id,
		notes.Title,
		notes.Body,
		notes.Tags,
		notes.CreatedAt,
		notes.UpdatedAt,
		notes.Owner,
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

func (repository *NotesRepositoryImpl) GetAll(ctx context.Context, db *pgxpool.Pool, notes *entity.Notes) (*entity.ListNotes, error) {
	var (
		ListResult *entity.ListNotes = new(entity.ListNotes)
		result     *entity.Notes     = new(entity.Notes)
	)

	SQL := `SELECT notes.* FROM notes
    LEFT JOIN collaborations ON collaborations.note_id = notes.id
    WHERE notes.owner = $1 OR collaborations.user_id = $1
    GROUP BY notes.id`
	varArgs := []interface{}{
		notes.Owner,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	rows, err := tx.Query(ctx, SQL, varArgs...)
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

func (repository *NotesRepositoryImpl) GetById(ctx context.Context, db *pgxpool.Pool, id string) (*entity.Notes, error) {
	var (
		result *entity.Notes = new(entity.Notes)
	)

	SQL := `SELECT notes.*, users.username
    FROM notes
    LEFT JOIN users ON users.id = notes.owner
    WHERE notes.id = $1`

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	row := tx.QueryRow(ctx, SQL, id)
	row.Scan(&result.Id, &result.Title, &result.Body, &result.Tags, &result.CreatedAt, &result.UpdatedAt, &result.Owner, &result.Username)

	return result, nil
}

func (repository *NotesRepositoryImpl) Update(ctx context.Context, db *pgxpool.Pool, notes *entity.Notes) (int64, error) {
	SQL := `UPDATE notes SET title = $1, body = $2, tags = $3, updated_at = $4 WHERE id = $5 RETURNING id`
	varArgs := []interface{}{
		notes.Title,
		notes.Body,
		notes.Tags,
		notes.UpdatedAt,
		notes.Id,
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
	if i <= 0 {
		return -1, exception.Wrap("repository not found", 404, errors.New("error not found"))
	}

	isTrue := result.RowsAffected()
	return isTrue, nil
}

func (repository *NotesRepositoryImpl) Delete(ctx context.Context, db *pgxpool.Pool, notes *entity.Notes) (int64, error) {
	SQL := `DELETE FROM notes WHERE id = $1 RETURNING id`
	varArgs := []interface{}{
		notes.Id,
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
	if i <= 0 {
		return -1, exception.Wrap("Catatan tidak ditemukan", 404, errors.New("fail"))
	}

	isTrue := result.RowsAffected()
	return isTrue, nil
}

func (repository *NotesRepositoryImpl) VerifyNoteOwner(ctx context.Context, db *pgxpool.Pool, notes *entity.Notes) (entity.Notes, error) {
	var (
		result entity.Notes = entity.Notes{}
	)

	SQL := `SELECT * FROM notes WHERE id = $1`

	tx, err := db.Begin(ctx)
	if err != nil {
		return entity.Notes{}, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	row := tx.QueryRow(ctx, SQL, notes.Id)
	row.Scan(&result.Id, &result.Title, &result.Body, &result.Tags, &result.CreatedAt, &result.UpdatedAt, &result.Owner)

	return result, nil
}
