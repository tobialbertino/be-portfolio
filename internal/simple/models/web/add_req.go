package web

type AddRequest struct {
	Number1 float64 `json:"number_1" validate:"required,number"`
	Number2 float64 `json:"number_2" validate:"required,number"`
}
