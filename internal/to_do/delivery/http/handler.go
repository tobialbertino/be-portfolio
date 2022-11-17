package http

import (
	"tobialbertino/portfolio-be/internal/to_do/models/domain"
	usecase "tobialbertino/portfolio-be/internal/to_do/useCase"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	ToDoUseCase usecase.ToDoUseCase
}

func NewHandler(toDoUseCase usecase.ToDoUseCase) *Handler {
	return &Handler{
		ToDoUseCase: toDoUseCase,
	}
}

func (h *Handler) Route(app *fiber.App) {
	g := app.Group("/to-do")
	g.Post("", h.CreateoDo)
}

func (h *Handler) CreateoDo(c *fiber.Ctx) error {
	var request *domain.RequestToDo

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.ToDoUseCase.Create(request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(&models.WebResponse{
		Status: "Ok",
		Data:   result,
	})
}
