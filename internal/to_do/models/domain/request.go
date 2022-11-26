package domain

type RequestUpdateToDo struct {
	Id     int64  `json:"id" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Status bool   `json:"status" validate:"required"`
}

type RequestToDo struct {
	Title string `json:"title" validate:"required"`
}
