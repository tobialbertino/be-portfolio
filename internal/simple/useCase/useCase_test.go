package usecase_test

import (
	"log"
	"testing"
	"tobialbertino/portfolio-be/internal/simple/models/web"
	simpleUseCase "tobialbertino/portfolio-be/internal/simple/useCase"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	webReq := &web.AddRequest{
		Number1: 1.5,
		Number2: 1.6,
	}
	simpleUc := simpleUseCase.NewSimpleUseCase()
	result, err := simpleUc.AddTwoNumber(webReq)
	if err != nil {
		log.Printf("Error func: %s", err.Error())
	}

	assert.Equal(t, 3.1, result)
}
