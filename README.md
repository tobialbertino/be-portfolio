# portfolio-be

on this page:
- [dependencies](#dependencies)
- [Framework](#framework)
- [Architecture](#architecture-folder)
- [modules / internal description](#modules--internal-description)
- [ToDo](#todo-rewrite-notesapp-from-dicoding-to-golang-in-modulesinternal-notes)

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

### ToDo: Rewrite notesApp from dicoding to golang in modules/internal: notes
- change error handler
- Try Unit Test, mainly use mocking data
- Notes: I was mistaken about using pointers to make efficient passing by reference (by default pass by value), in short using stack memory is faster than heap memory, or minimizing memory allocation