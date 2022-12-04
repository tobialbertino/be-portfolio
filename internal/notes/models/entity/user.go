package entity

import "tobialbertino/portfolio-be/internal/notes/models/domain"

type User struct {
	Id        string
	Username  string
	Passwword string
	FullName  string
}

func (dt *User) ToDomain() *domain.ResponseUser {
	return &domain.ResponseUser{
		Id:       dt.Id,
		Username: dt.Username,
		FullName: dt.FullName,
	}
}
