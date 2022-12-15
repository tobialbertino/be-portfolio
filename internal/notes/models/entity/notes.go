package entity

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type ListNotes []Notes
type Notes struct {
	Id        string
	Title     string
	Body      string
	Tags      []string
	CreatedAt int64
	UpdatedAt int64
	Owner     *string
	Username  *string
}

func (dt *Notes) ToDomain() *domain.Notes {
	return &domain.Notes{
		Id:        dt.Id,
		Title:     dt.Title,
		Body:      dt.Body,
		Tags:      dt.Tags,
		CreatedAt: dt.CreatedAt,
		UpdatedAt: dt.UpdatedAt,
		Owner:     dt.Username,
	}
}

func (ldt *ListNotes) ToDomain() *[]domain.Notes {
	var result []domain.Notes = make([]domain.Notes, 0, 10)
	for _, td := range *ldt {
		result = append(result, *td.ToDomain())
	}
	return &result
}
