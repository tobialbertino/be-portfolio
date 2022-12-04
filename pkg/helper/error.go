package helper

import "github.com/gofiber/fiber/v2"

// fail, goFiber / fasthttp PanicHandler function not available
func PanicIfError(err *error) {
	panic("error")
}

// fail, should directly return error in the main function handler, or costum errors following interface error
func BadRequestIfError(err *error) *fiber.Error {
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	return nil
}
