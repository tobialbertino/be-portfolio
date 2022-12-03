package http

import (
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	usecase "tobialbertino/portfolio-be/internal/notes/useCase"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	NotesoUseCase usecase.NotesUseCase
}

func NewHandler(notesoUseCase usecase.NotesUseCase) *Handler {
	return &Handler{
		NotesoUseCase: notesoUseCase,
	}
}

func (h *Handler) Route(app *fiber.App) {
	g := app.Group("/notes")
	g.Post("", h.Add)
}

func (h *Handler) Add(c *fiber.Ctx) error {
	var request *domain.ReqAddNote = new(domain.ReqAddNote)

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.NotesoUseCase.Add(request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(&models.WebResponse{
		Status: "Ok",
		Data:   result,
	})
}
