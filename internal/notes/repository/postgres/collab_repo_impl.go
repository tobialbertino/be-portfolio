package postgres

import (
	"context"
	"tobialbertino/portfolio-be/internal/notes/models/entity"
	"tobialbertino/portfolio-be/pkg/helper"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CollabRepositoryImpl struct {
}

func NewCollabRepository() CollabRepository {
	return &CollabRepositoryImpl{}
}

// AddCollaboration implements CollabRepository
func (*CollabRepositoryImpl) AddCollaboration(ctx context.Context, db *pgxpool.Pool, collab *entity.Collab) (entity.Collab, error) {
	var res entity.Collab

	SQL := `INSERT INTO collaborations VALUES ($1, $2, $3) RETURNING id`
	varArgs := []interface{}{
		collab.Id,
		collab.NoteId,
		collab.UserId,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return entity.Collab{}, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	row := tx.QueryRow(ctx, SQL, varArgs...)
	row.Scan(&res.Id)

	return res, nil
}

func (*CollabRepositoryImpl) DeleteCollaboration(ctx context.Context, db *pgxpool.Pool, collab *entity.Collab) (entity.Collab, error) {
	var res entity.Collab

	SQL := `DELETE FROM collaborations WHERE note_id = $1 AND user_id = $2 RETURNING id`
	varArgs := []interface{}{
		collab.NoteId,
		collab.UserId,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return entity.Collab{}, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	row := tx.QueryRow(ctx, SQL, varArgs...)
	row.Scan(&res.Id)

	return res, nil
}

func (*CollabRepositoryImpl) VerifyCollaborator(ctx context.Context, db *pgxpool.Pool, collab *entity.Collab) (entity.Collab, error) {
	var res entity.Collab

	SQL := `SELECT * FROM collaborations WHERE note_id = $1 AND user_id = $2`
	varArgs := []interface{}{
		collab.NoteId,
		collab.UserId,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return entity.Collab{}, err
	}
	defer helper.CommitOrRollback(err, ctx, tx)

	row := tx.QueryRow(ctx, SQL, varArgs...)
	row.Scan(&res.Id, &res.NoteId, &res.UserId)

	return res, nil
}
