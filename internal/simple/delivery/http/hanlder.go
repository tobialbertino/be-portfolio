package http

import (
	"tobialbertino/portfolio-be/internal/simple/models/web"
	usecase "tobialbertino/portfolio-be/internal/simple/useCase"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	SimpleUseCase usecase.SimpleUseCase
}

func NewHandler(SimpleUseCase usecase.SimpleUseCase) *Handler {
	return &Handler{
		SimpleUseCase: SimpleUseCase,
	}
}

func (h *Handler) Route(app *fiber.App) {
	g := app.Group("/simple")
	g.Post("/add-two-number", h.AddTwoNumber)
}

func (h *Handler) AddTwoNumber(c *fiber.Ctx) error {
	var request *web.AddRequest

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}

	response, _ := h.SimpleUseCase.AddTwoNumber(request)

	return c.JSON(models.WebResponse{
		Status: "Ok",
		Data:   response,
	})
}
