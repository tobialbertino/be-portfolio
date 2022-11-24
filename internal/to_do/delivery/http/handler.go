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
	g.Put("/:id", h.UpdateTodo)
}

func (h *Handler) CreateoDo(c *fiber.Ctx) error {
	var request *domain.RequestToDo = new(domain.RequestToDo)

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

func (h *Handler) UpdateTodo(c *fiber.Ctx) error {
	var request *domain.RequestUpdateToDo = new(domain.RequestUpdateToDo)

	id, _ := c.ParamsInt("id")
	request.Id = int64(id)
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.ToDoUseCase.Update(request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(&models.WebResponse{
		Status: "Ok",
		Data:   result,
	})
}
