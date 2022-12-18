package usecase_test

import (
	"log"
	"testing"
	"tobialbertino/portfolio-be/internal/simple/models/web"
	simpleUseCase "tobialbertino/portfolio-be/internal/simple/useCase"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	validate *validator.Validate = validator.New()
	simpleUc                     = simpleUseCase.NewSimpleUseCase(validate)
)

type TableSimpleSuccess struct {
	Name     string
	Req      web.AddRequest
	Expected float64
}

func TestTableSimpleSuccess(t *testing.T) {
	table := []TableSimpleSuccess{
		{
			Name: "add 12+12",
			Req: web.AddRequest{
				Number1: 12,
				Number2: 12,
			},
			Expected: 24,
		},
		{
			Name: "add 12.5 + 12.5",
			Req: web.AddRequest{
				Number1: 12.5,
				Number2: 12.5,
			},
			Expected: 25,
		},
	}

	for _, test := range table {
		t.Run(test.Name, func(t *testing.T) {
			result, err := simpleUc.AddTwoNumber(&test.Req)
			require.Equal(t, test.Expected, result, err)
		})
	}
}

func TestSimple(t *testing.T) {
	webReq := web.AddRequest{
		Number1: 1.5,
		Number2: 1.6,
	}

	result, err := simpleUc.AddTwoNumber(&webReq)
	if err != nil {
		log.Printf("Error func: %s", err.Error())
	}

	assert.Equal(t, 3.1, result, "Result must be 3.1")
}

func TestSimpleFail(t *testing.T) {
	webReq := web.AddRequest{
		Number1: 1.6,
		Number2: 1.6,
	}

	result, err := simpleUc.AddTwoNumber(&webReq)
	if err != nil {
		log.Printf("Error func: %s", err.Error())
	}

	assert.NotEqual(t, 3.1, result, "Result must be 3.2")
}
