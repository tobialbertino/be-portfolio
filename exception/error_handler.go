package exception

import (
	"errors"
	"fmt"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// Custom wrap error
type WrappedError struct {
	Code    int
	Context string
	Err     error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf(`%s: %v, %v `, w.Context, w.Code, w.Err)
}

func Wrap(contextInfo string, code int, err error) *WrappedError {
	return &WrappedError{
		Context: contextInfo,
		Code:    code,
		Err:     err,
	}
}

// var CustomErrorHandler = func(ctx *fiber.Ctx, err error) error {
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	var e *fiber.Error
	var wrapErr *WrappedError
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	status := utils.StatusMessage(code)

	if errors.As(err, &wrapErr) {
		code = wrapErr.Code
		status = utils.StatusMessage(code)
		goto LABEL_RETURN
	}

	// Return from handler
LABEL_RETURN:
	return ctx.Status(code).JSON(models.WebResponseError{
		Status:  status,
		Data:    nil,
		Message: err.Error(),
	})
}
