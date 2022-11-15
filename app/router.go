package app

import (
	simpleHandler "tobialbertino/portfolio-be/internal/simple/delivery/http"
	simpleUseCase "tobialbertino/portfolio-be/internal/simple/useCase"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	simpleUc := simpleUseCase.NewSimpleUseCase()
	simpleHnadler := simpleHandler.NewHandler(simpleUc)

	simpleHnadler.Route(app)
}
