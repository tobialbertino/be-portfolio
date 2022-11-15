package helper

import "github.com/gofiber/fiber/v2"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func BadRequestIfError(err error) *fiber.Error {
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}
	return nil
}
