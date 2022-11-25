// learn & try pattern repository, structure folder
package entity

import "tobialbertino/portfolio-be/internal/to_do/models/domain"

type ListToDo []ToDo
type ToDo struct {
	Id         int64
	Title      string
	Status     bool
	Created_at int64
	Updated_at int64
}

func (dt *ToDo) ToDomain() *domain.ResponseToDo {
	return &domain.ResponseToDo{
		Id:         dt.Id,
		Title:      dt.Title,
		Status:     dt.Status,
		Created_at: dt.Created_at,
		Updated_at: dt.Updated_at,
	}
}

func (ldt *ListToDo) ToDomain() *[]domain.ResponseToDo {
	var result []domain.ResponseToDo = make([]domain.ResponseToDo, 0, 10)

	for _, td := range *ldt {
		result = append(result, *td.ToDomain())
	}

	return &result
}
