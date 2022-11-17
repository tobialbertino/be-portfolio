package exception

import (
	"errors"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

var CustomErrorHandler = func(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	status := utils.StatusMessage(code)

	// Return from handler
	return ctx.Status(code).JSON(models.WebResponseError{
		Status:  status,
		Data:    nil,
		Message: err.Error(),
	})
}
