package http

import (
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	usecase "tobialbertino/portfolio-be/internal/notes/useCase"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	NotesoUseCase usecase.NotesUseCase
	UserUseCase   usecase.UserUseCase
}

func NewHandler(notesoUseCase usecase.NotesUseCase, userUC usecase.UserUseCase) *Handler {
	return &Handler{
		NotesoUseCase: notesoUseCase,
		UserUseCase:   userUC,
	}
}

func (h *Handler) Route(app *fiber.App) {
	// notes
	g := app.Group("/notes")
	g.Post("", h.Add)
	g.Get("", h.GetAll)
	g.Get("/:id", h.GetById)
	g.Put("/:id", h.UpdateById)
	g.Delete("/:id", h.DeleteById)

	// user notes
	user := g.Group("/user")
	user.Post("", h.AddUser)
	user.Get("/:id", h.GetUserById)
}

func (h *Handler) Add(c *fiber.Ctx) error {
	var request *domain.ReqAddNote = new(domain.ReqAddNote)

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.NotesoUseCase.Add(request)
	if err != nil {
		return err
	}

	return c.JSON(&models.WebResponse{
		Status: "Ok",
		Data:   result,
	})
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	var result *[]domain.Notes = new([]domain.Notes)

	result, err := h.NotesoUseCase.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(&models.WebResponse{
		Status: "Ok",
		Data:   result,
	})
}

func (h *Handler) GetById(c *fiber.Ctx) error {
	var result *domain.Notes = new(domain.Notes)
	id := c.Params("id")

	result, err := h.NotesoUseCase.GetById(id)
	if err != nil {
		return err
	}

	return c.JSON(&models.WebResponse{
		Status: "Ok",
		Data:   result,
	})
}

func (h *Handler) UpdateById(c *fiber.Ctx) error {
	var request *domain.ReqAddNote = new(domain.ReqAddNote)
	id := c.Params("id")

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.NotesoUseCase.Update(request, id)
	if err != nil {
		return err
	}

	return c.JSON(&models.WebResponse{
		Status: "Ok",
		Data:   result,
	})
}

func (h *Handler) DeleteById(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.NotesoUseCase.Delete(id)
	if err != nil {
		return err
	}

	return c.JSON(&models.WebResponse{
		Status: "Ok",
		Data:   result,
	})
}

func (h *Handler) AddUser(c *fiber.Ctx) error {
	var request *domain.ReqAddUser = new(domain.ReqAddUser)

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	result, err := h.UserUseCase.AddUser(request)
	if err != nil {
		return err
	}

	return c.JSON(&models.WebResponse{
		Status: "Ok",
		Data:   result,
	})
}

func (h *Handler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.UserUseCase.GetUserById(id)
	if err != nil {
		return err
	}

	return c.JSON(&models.WebResponse{
		Status: "Ok",
		Data:   result,
	})
}
