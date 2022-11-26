package exception

import (
	"errors"
	"strings"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

var CustomErrorHandler = func(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	var e *fiber.Error
	code := fiber.StatusInternalServerError
	status := utils.StatusMessage(code)

	if strings.Contains(err.Error(), "not found") {
		code = fiber.StatusNotFound
		status = utils.StatusMessage(code)
		goto LABEL_RETURN
	}

	// Retrieve the custom status code if it's a *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Return from handler
LABEL_RETURN:
	return ctx.Status(code).JSON(models.WebResponseError{
		Status:  status,
		Data:    nil,
		Message: err.Error(),
	})
}
