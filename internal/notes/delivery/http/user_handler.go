package http

import (
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	usecase "tobialbertino/portfolio-be/internal/notes/useCase"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(userUC usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: userUC,
	}
}

func (h *UserHandler) Route(app *fiber.App) {
	// notes
	g := app.Group("/notes")
	users := g.Group("/users")

	// user notes
	users.Get("", h.GetUsersByUsername)
	users.Get("/:id", h.GetUserById)
	users.Post("", h.AddUser)
}

func (h *UserHandler) AddUser(c *fiber.Ctx) error {
	var request *domain.ReqAddUser = new(domain.ReqAddUser)

	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	result, err := h.UserUseCase.AddUser(request)
	if err != nil {
		return err
	}

	c.Status(201).JSON(&models.WebResponse{
		Status:  "success",
		Message: "User berhasil ditambahkan",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.UserUseCase.GetUserById(id)
	if err != nil {
		return err
	}

	c.JSON(&models.WebResponse{
		Status: "success",
		Data: domain.UserData{
			User: *result,
		},
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *UserHandler) GetUsersByUsername(c *fiber.Ctx) error {
	username := utils.CopyString(c.Query("username"))

	result, err := h.UserUseCase.GetUsersByUsername(username)
	if err != nil {
		return err
	}

	c.JSON(&models.WebResponse{
		Status: "success",
		Data: domain.UsersData{
			User: *result,
		},
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}
