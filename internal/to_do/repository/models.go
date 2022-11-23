package repository

import (
	"tobialbertino/portfolio-be/internal/to_do/models/domain"
	"tobialbertino/portfolio-be/internal/to_do/models/entity"
)

func ToDomain(et *entity.ToDo) *domain.ResponseToDo {
	return &domain.ResponseToDo{
		Id:         et.Id,
		Title:      et.Title,
		Status:     et.Status,
		Created_at: et.Created_at,
		Updated_at: et.Updated_at,
	}
}
