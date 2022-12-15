package http

import (
	usecase "tobialbertino/portfolio-be/internal/notes/useCase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	NotesoUseCase usecase.NotesUseCase
	UserUseCase   usecase.UserUseCase
	AuthUseCase   usecase.AuthUseCase
	CollabUseCase usecase.CollabUseCase
}

func NewHandler(notesoUseCase usecase.NotesUseCase, userUC usecase.UserUseCase, authUC usecase.AuthUseCase, collabUseCase usecase.CollabUseCase) *Handler {
	return &Handler{
		NotesoUseCase: notesoUseCase,
		UserUseCase:   userUC,
		AuthUseCase:   authUC,
		CollabUseCase: collabUseCase,
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

	// collab notes
	ch := NewCollabHandler(h.CollabUseCase, h.NotesoUseCase)
	ch.Route(app)
}
