package app

import (
	simpleHandler "tobialbertino/portfolio-be/internal/simple/delivery/http"
	simpleUseCase "tobialbertino/portfolio-be/internal/simple/useCase"
	toDoHandler "tobialbertino/portfolio-be/internal/to_do/delivery/http"
	toDoRepository "tobialbertino/portfolio-be/internal/to_do/repository"
	toDoUseCase "tobialbertino/portfolio-be/internal/to_do/useCase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func InitRouter(app *fiber.App, DB *pgx.Conn, validate *validator.Validate) {
	// simple app setup
	simpleUc := simpleUseCase.NewSimpleUseCase()
	simpleHnadler := simpleHandler.NewHandler(simpleUc)
	simpleHnadler.Route(app)

	// to do app setup
	toDoRepo := toDoRepository.NewToDoRepository()
	toDoUc := toDoUseCase.NewToDoUseCase(toDoRepo, DB, validate)
	toDoHandler := toDoHandler.NewHandler(toDoUc)
	toDoHandler.Route(app)

}
