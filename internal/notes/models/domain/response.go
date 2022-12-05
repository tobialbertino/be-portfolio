package domain

type RowsAffected struct {
	RowsAffected int64 `json:"rows_affected"`
}

type UserId struct {
	UserId string `json:"userId"`
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
