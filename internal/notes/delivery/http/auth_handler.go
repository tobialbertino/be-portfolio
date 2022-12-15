package http

import (
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	usecase "tobialbertino/portfolio-be/internal/notes/useCase"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthUseCase usecase.AuthUseCase
}

func NewAuthHandler(authUC usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		AuthUseCase: authUC,
	}
}

func (h *AuthHandler) Route(app *fiber.App) {
	// notes
	g := app.Group("/notes")
	auth := g.Group("/authentications")

	// auth user
	auth.Post("", h.postAuthenticationHandler)
	auth.Put("", h.putAuthenticationHandler)
	auth.Delete("", h.deleteAuthenticationHandler)
}

func (h *AuthHandler) postAuthenticationHandler(c *fiber.Ctx) error {
	var req *domain.ReqLoginUser = new(domain.ReqLoginUser)

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	result, err := h.AuthUseCase.AddRefreshToken(req)
	if err != nil {
		return err
	}

	c.Status(201).JSON(&models.WebResponse{
		Status:  "success",
		Message: "Authentication berhasil ditambahkan",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *AuthHandler) putAuthenticationHandler(c *fiber.Ctx) error {
	var req *domain.ReqRefreshToken = new(domain.ReqRefreshToken)

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	result, err := h.AuthUseCase.VerifyRefreshToken(req)
	if err != nil {
		return err
	}

	c.Status(200).JSON(&models.WebResponse{
		Status:  "success",
		Message: "Access Token berhasil diperbarui",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}

func (h *AuthHandler) deleteAuthenticationHandler(c *fiber.Ctx) error {
	var req *domain.ReqRefreshToken = new(domain.ReqRefreshToken)

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	result, err := h.AuthUseCase.DeleteRefreshToken(req)
	if err != nil {
		return err
	}

	c.Status(200).JSON(&models.WebResponse{
		Status:  "success",
		Message: "Refresh token berhasil dihapus",
		Data:    result,
	})
	c.Set("content-type", "application/json; charset=utf-8")

	return nil
}
