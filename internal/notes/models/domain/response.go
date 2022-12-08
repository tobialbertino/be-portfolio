package domain

type UserId struct {
	UserId string `json:"userId"`
}

type NoteId struct {
	NoteId string `json:"noteId"`
}

type ListNotes struct {
	Notes []Notes `json:"notes"`
}

type NoteRes struct {
	Note Notes `json:"note"`
}

type Notes struct {
	Id        string   `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Tags      []string `json:"tags"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
	Owner     *string  `json:"owner"`
}

type RowsAffected struct {
	RowsAffected int64 `json:"rows_affected"`
}

type UserData struct {
	User ResponseUser `json:"user"`
}

type UsersData struct {
	User []ResponseUser `json:"users"`
}

type ResponseUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
}

type ResToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
