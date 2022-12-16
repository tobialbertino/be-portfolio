package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"
	"tobialbertino/portfolio-be/exception"
	"tobialbertino/portfolio-be/internal/simple/models/web"
	simpleUseCase "tobialbertino/portfolio-be/internal/simple/useCase"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupAppRoute() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.CustomErrorHandler,
	})
	validate := validator.New()
	simpleUc := simpleUseCase.NewSimpleUseCase(validate)
	simpleHandler := NewHandler(simpleUc)

	simpleHandler.Route(app)
	return app
}

type TableTest struct {
	Name   string
	Input  web.AddRequest
	Output float64
}

func TestTableSimpleSuccess(t *testing.T) {
	table := []TableTest{
		{
			Name: "add->10+10",
			Input: web.AddRequest{
				Number1: 10,
				Number2: 10,
			},
			Output: 20,
		},
		{
			Name: "add->1.0+1.0",
			Input: web.AddRequest{
				Number1: 1.0,
				Number2: 1.0,
			},
			Output: 2,
		},
	}

	app := setupAppRoute()
	for _, test := range table {
		t.Run(test.Name, func(t *testing.T) {
			requestBody, _ := json.Marshal(test.Input)
			b := bytes.NewBuffer(requestBody)

			request := httptest.NewRequest("POST", "/simple/add-two-number", b)
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("Accept", "application/json")

			response, err := app.Test(request, -1)
			assert.Equal(t, 200, response.StatusCode, err)

			responseBody, err := io.ReadAll(response.Body)
			webResponse := models.WebResponse{}
			json.Unmarshal(responseBody, &webResponse)
			assert.Equal(t, "Ok", webResponse.Status, err)

			responseData := webResponse.Data.(float64)
			assert.NotNil(t, responseData)
			assert.Equal(t, test.Output, responseData)
		})
	}

}

func TestSimple(t *testing.T) {
	app := setupAppRoute()

	createRequest := web.AddRequest{
		Number1: 15,
		Number2: 15.5,
	}
	requestBody, _ := json.Marshal(createRequest)
	b := bytes.NewBuffer(requestBody)

	request := httptest.NewRequest("POST", "/simple/add-two-number", b)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request, -1)
	assert.Equal(t, 200, response.StatusCode, err)

	responseBody, err := io.ReadAll(response.Body)
	webResponse := models.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, "Ok", webResponse.Status, err)

	responseData := webResponse.Data.(float64)
	assert.NotNil(t, responseData)
	assert.Equal(t, 30.5, responseData)
}
