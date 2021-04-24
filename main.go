package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"

	database "github.com/johnatasr/fiber-api/infra/database"
	bookRepo "github.com/johnatasr/fiber-api/repositories/bookRepo"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", bookRepo.GetBooks)
	app.Get("/api/v1/book/:id", bookRepo.GetBook)
	app.Post("/api/v1/book", bookRepo.NewBook)
	app.Delete("/api/v1/book/:id", bookRepo.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection opened")
}

func main() {
	app := fiber.New()

	initDatabase()

	setupRoutes(app)
	app.Listen(8000)

	defer database.DBConn.Close()
}
