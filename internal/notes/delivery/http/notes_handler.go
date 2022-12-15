package http

import (
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	usecase "tobialbertino/portfolio-be/internal/notes/useCase"
	"tobialbertino/portfolio-be/pkg/helper"
	"tobialbertino/portfolio-be/pkg/middleware"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type NotesHandler struct {
	NotesUseCase usecase.NotesUseCase
}

func NewNotesHandler(notesoUseCase usecase.NotesUseCase) *NotesHandler {
	return &NotesHandler{
		NotesUseCase: notesoUseCase,
	}
}

func (h *NotesHandler) Route(app *fiber.App) {
	// notes
	g := app.Group("/notes")
	notes := g.Group("/notes", middleware.ProtectedJWT())

	notes.Post("", h.Add)
	notes.Get("", h.GetAll)
	notes.Get("/:id", h.GetById)
	notes.Put("/:id", h.UpdateById)
	notes.Delete("/:id", h.DeleteById)
}

func (h *NotesHandler) Add(c *fiber.Ctx) error {
	var request *domain.ReqAddNote = new(domain.ReqAddNote)

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userId := helper.GetIDUserFromToken(c)

	result, err := h.NotesUseCase.Add(request, userId)
	if err != nil {
		return err
	}

	c.Status(201).JSON(&models.WebResponse{
		Status:  "success",
		Message: "Catatan berhasil ditambahkan",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *NotesHandler) GetAll(c *fiber.Ctx) error {
	var result *[]domain.Notes = new([]domain.Notes)

	userId := helper.GetIDUserFromToken(c)

	result, err := h.NotesUseCase.GetAll(userId)
	if err != nil {
		return err
	}

	c.JSON(&models.WebResponse{
		Status: "success",
		Data: &domain.ListNotes{
			Notes: *result,
		},
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *NotesHandler) GetById(c *fiber.Ctx) error {
	var result *domain.Notes = new(domain.Notes)
	id := c.Params("id")

	// TODO: Bida dibuat middleware
	userId := helper.GetIDUserFromToken(c)
	IsTrue, err := h.NotesUseCase.VerifyNoteAccess(id, userId)
	if err != nil && !IsTrue {
		return err
	}

	result, err = h.NotesUseCase.GetById(id)
	if err != nil {
		return err
	}

	c.JSON(&models.WebResponse{
		Status: "success",
		Data: &domain.NoteRes{
			Note: *result,
		},
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *NotesHandler) UpdateById(c *fiber.Ctx) error {
	var request *domain.ReqAddNote = new(domain.ReqAddNote)
	id := c.Params("id")

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// TODO: Bida dibuat middleware
	userId := helper.GetIDUserFromToken(c)
	IsTrue, err := h.NotesUseCase.VerifyNoteAccess(id, userId)
	if err != nil && !IsTrue {
		return err
	}

	result, err := h.NotesUseCase.Update(request, id)
	if err != nil {
		return err
	}

	c.JSON(&models.WebResponse{
		Status:  "success",
		Message: "Catatan berhasil diperbarui",
		Data: &domain.NoteRes{
			Note: *result,
		},
	})
	c.Set("content-type", "application/json; charset=utf-8")
	return nil
}

func (h *NotesHandler) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")

	// TODO: Bida dibuat middleware
	userId := helper.GetIDUserFromToken(c)
	IsTrue, err := h.NotesUseCase.VerifyNoteOwner(id, userId)
	if err != nil && !IsTrue {
		return err
	}

	result, err := h.NotesUseCase.Delete(id)
	if err != nil {
		return err
	}

	c.JSON(&models.WebResponse{
		Status:  "success",
		Message: "Catatan berhasil dihapus",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")
	return nil
}
