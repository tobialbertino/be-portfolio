package domain

import (
	"time"
	"tobialbertino/portfolio-be/internal/to_do/models/entity"
)

type ResponseToDo struct {
	Id         int64  `json:"id" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Status     bool   `json:"status" validate:"required"`
	Created_at int64  `json:"created_at" validate:"required"`
	Updated_at int64  `json:"updated_at" validate:"required"`
}

type RequestToDo struct {
	Title string `json:"title" validate:"required"`
}

type RowsAffected struct {
	RowsAffected int64 `json:"rows_affected" validate:"required"`
}

func (req *RequestToDo) ToEntity() *entity.ToDo {
	return &entity.ToDo{
		Title:      req.Title,
		Status:     false,
		Created_at: time.Now().Unix(),
		Updated_at: time.Now().Unix(),
	}
}
