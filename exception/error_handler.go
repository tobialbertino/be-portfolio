package exception

import (
	"errors"
	"tobialbertino/portfolio-be/pkg/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// var CustomErrorHandler = func(ctx *fiber.Ctx, err error) error {
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	var (
		e           *fiber.Error
		wrapErr     *WrappedError
		clientError *ClientError
		authError   *AuthorizationError
	)

	// defaults Status code to 500
	code := fiber.StatusInternalServerError
	status := "fail"
	message := utils.StatusMessage(code)

	// Retrieve the custom status code if it's a *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	if errors.As(err, &wrapErr) {
		code = wrapErr.Code
		status = wrapErr.Err.Error()
		message = wrapErr.Context
		goto LABEL_RETURN
	}

	if errors.As(err, &clientError) {
		code = clientError.Code
		message = clientError.Message
		goto LABEL_RETURN_CLIENT_ERROR
	}

	if errors.As(err, &authError) {
		code = authError.Code
		message = authError.Message
		goto LABEL_RETURN_CLIENT_ERROR
	}

	// Return from handler
LABEL_RETURN:
	ctx.Status(code).JSON(models.WebResponseError{
		Status:  status,
		Data:    nil,
		Message: message,
	})
	ctx.Set("content-type", "application/json; charset=utf-8")
	return nil

LABEL_RETURN_CLIENT_ERROR:
	ctx.Status(code).JSON(models.WebResponseError{
		Status:  "fail",
		Data:    nil,
		Message: message,
	})
	ctx.Set("content-type", "application/json; charset=utf-8")
	return nil
}
