# portfolio-be

# dependencies

using gofiber v2  
```
go get -u github.com/gofiber/fiber/v2
```

## Framework

- Framework : GoFiber
- Configuration : GoDotEnv

## Architecture

per-modules / internal:  
Delivery -> UseCase -> Repository

## modules / internal description:

- simple: simple API, add 2 number from body payload
- to_do: CRUD list To do App

ToDo: Rewrite notesApp from dicoding to golang in modules/internal: notes