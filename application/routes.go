package application

import (
	bookRepo "fiber_api/repositories"

	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/books", bookRepo.GetBooks)
	app.Get("/api/book/:id", bookRepo.GetBook)
	app.Post("/api/book-insert", bookRepo.NewBook)
	app.Delete("/api/book-delete/:id", bookRepo.DeleteBook)
}
