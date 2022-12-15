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

type ReqLoginUser struct {
	Username  string `json:"username" validate:"required"`
	Passwword string `json:"password" validate:"required,min=6"`
}

type Token struct {
	Token string `json:"token" validate:"required"`
}

type ReqAuthPayload struct {
	Username  string `json:"username" validate:"required"`
	Passwword string `json:"password" validate:"required"`
}

type ReqRefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type Collab struct {
	Id     string `json:"collaborationId,omitempty"`
	NoteId string `json:"noteId,omitempty" validate:"required"`
	UserId string `json:"userId,omitempty" validate:"required"`
}
