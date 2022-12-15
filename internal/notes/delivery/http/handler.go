package http

import (
	usecase "tobialbertino/portfolio-be/internal/notes/useCase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	NotesoUseCase usecase.NotesUseCase
	UserUseCase   usecase.UserUseCase
	AuthUseCase   usecase.AuthUseCase
}

func NewHandler(notesoUseCase usecase.NotesUseCase, userUC usecase.UserUseCase, authUC usecase.AuthUseCase) *Handler {
	return &Handler{
		NotesoUseCase: notesoUseCase,
		UserUseCase:   userUC,
		AuthUseCase:   authUC,
	}
}

func (h *Handler) Route(app *fiber.App) {
	// notes
	nh := NewNotesHandler(h.NotesoUseCase)
	nh.Route(app)

	// user notes
	uh := NewUserHandler(h.UserUseCase)
	uh.Route(app)

	// auth user
	ah := NewAuthHandler(h.AuthUseCase)
	ah.Route(app)
}
