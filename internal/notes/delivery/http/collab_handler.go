package http

import (
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	usecase "tobialbertino/portfolio-be/internal/notes/useCase"
	"tobialbertino/portfolio-be/pkg/helper"
	"tobialbertino/portfolio-be/pkg/middleware"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type CollabHandler struct {
	CollabUseCase usecase.CollabUseCase
	NotesUseCase  usecase.NotesUseCase
}

func NewCollabHandler(CollabUC usecase.CollabUseCase, notesUseCase usecase.NotesUseCase) *CollabHandler {
	return &CollabHandler{
		CollabUseCase: CollabUC,
		NotesUseCase:  notesUseCase,
	}
}

func (h *CollabHandler) Route(app *fiber.App) {
	// notes
	g := app.Group("/notes")
	collab := g.Group("/collaborations", middleware.ProtectedJWT())

	// Collab Notes
	collab.Post("", h.postCollaborationHandler)
	collab.Delete("", h.deleteCollaborationHandler)
}

func (h *CollabHandler) postCollaborationHandler(c *fiber.Ctx) error {
	var req domain.Collab

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	userId := helper.GetIDUserFromToken(c)
	IsTrue, err := h.NotesUseCase.VerifyNoteOwner(req.NoteId, userId)
	if err != nil && !IsTrue {
		return err
	}

	result, err := h.CollabUseCase.AddCollaboration(&req)
	if err != nil {
		return err
	}

	c.Status(201).JSON(&models.WebResponse{
		Status:  "success",
		Message: "Kolaborasi berhasil ditambahkan",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *CollabHandler) deleteCollaborationHandler(c *fiber.Ctx) error {
	var req domain.Collab

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	userId := helper.GetIDUserFromToken(c)
	IsTrue, err := h.NotesUseCase.VerifyNoteOwner(req.NoteId, userId)
	if err != nil && !IsTrue {
		return err
	}

	_, err = h.CollabUseCase.DeleteCollaboration(&req)
	if err != nil {
		return err
	}

	c.JSON(&models.WebResponse{
		Status:  "success",
		Message: "Kolaborasi berhasil dihapus",
		Data:    nil,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}
