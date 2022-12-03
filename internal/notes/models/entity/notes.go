package entity

type Notes struct {
	Id        string
	Title     string
	Body      string
	Tags      []string
	CreatedAt int64
	UpdatedAt int64
	Owner     *string
}
