package entity

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type Collab struct {
	Id     string
	NoteId string
	UserId string
}

func (dt *Collab) ToDomain() domain.Collab {
	return domain.Collab{
		Id:     dt.Id,
		NoteId: dt.NoteId,
		UserId: dt.UserId,
	}
}
