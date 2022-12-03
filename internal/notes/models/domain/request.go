package domain

type ReqAddNote struct {
	Title string   `json:"title" validate:"required"`
	Tags  []string `json:"tags" validate:"required"`
	Body  string   `json:"body" validate:"required"`
}
