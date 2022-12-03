package domain

type RowsAffected struct {
	RowsAffected int64 `json:"rows_affected" validate:"required"`
}
