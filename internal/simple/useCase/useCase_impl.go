package usecase

import "tobialbertino/portfolio-be/internal/simple/models/web"

type SimpleUseCaseImpl struct {
}

func NewSimpleUseCase() SimpleUseCase {
	return &SimpleUseCaseImpl{}
}

// AddTwoNumber implements SimpleUseCase
func (useCase *SimpleUseCaseImpl) AddTwoNumber(webReq *web.AddRequest) (float64, error) {
	return webReq.Number1 + webReq.Number2, nil
}
