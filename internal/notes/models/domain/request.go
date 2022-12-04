package domain

type ReqAddNote struct {
	Title string   `json:"title" validate:"required"`
	Tags  []string `json:"tags" validate:"required"`
	Body  string   `json:"body" validate:"required"`
}

type ReqAddUser struct {
	Username  string `json:"username" validate:"required"`
	Passwword string `json:"password" validate:"required,min=6"`
	FullName  string `json:"fullname" validate:"required"`
}
