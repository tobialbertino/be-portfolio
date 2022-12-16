package usecase

import (
	"tobialbertino/portfolio-be/internal/simple/models/web"

	"github.com/go-playground/validator/v10"
)

type SimpleUseCaseImpl struct {
	Validate *validator.Validate
}

func NewSimpleUseCase(validate *validator.Validate) SimpleUseCase {
	return &SimpleUseCaseImpl{
		Validate: validate,
	}
}

// AddTwoNumber implements SimpleUseCase
func (useCase *SimpleUseCaseImpl) AddTwoNumber(webReq *web.AddRequest) (float64, error) {
	err := useCase.Validate.Struct(webReq)
	if err != nil {
		return 0, err
	}

	return webReq.Number1 + webReq.Number2, nil
}
