package usecase

import "tobialbertino/portfolio-be/internal/simple/models/web"

type SimpleUseCase interface {
	AddTwoNumber(webReq *web.AddRequest) (float64, error)
}
