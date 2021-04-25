package repositories

import (
	"github.com/gofiber/fiber"
	db "github.com/johnatasr/fiber_api/infra/database"
)

func GetBooks(c *fiber.Ctx) {
	db := db.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)

}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := db.DBConn
	var book Book

	data := c.Body()

	book.Title = data.title
	book.Author = data.author
	book.Rating = data.rating

	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DBConn

	var book Book

	db.First(&book, id)

	if book.Title == "" {
		c.Status(500).Send("No book found with ID")
		return
	}

	db.Delete(&book)
	c.Send("Book successfully deleted")
}
