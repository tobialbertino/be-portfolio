package domain

type RowsAffected struct {
	RowsAffected int64 `json:"rows_affected"`
}

type Notes struct {
	Id        string   `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
	Owner     *string  `json:"owner"`
}

type ResponseUser struct {
	Id       string `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	FullName string `json:"fullname" validate:"required"`
}
