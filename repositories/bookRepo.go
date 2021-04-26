package repositories

import (
	"fiber_api/models"

	db "fiber_api/infra"

	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	db := db.DBConn
	var books []models.Book

	if err := db.Find(&books).Error; err != nil {
		c.Send("No one book found")
	}

	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DBConn

	var book models.Book

	if err := db.Find(&book, id).Error; err != nil {
		c.Send("No one book with this ID")
	}

	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := db.DBConn

	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DBConn

	var book models.Book

	db.First(&book, id)

	if book.Title == "" {
		c.Status(500).Send("No book found with ID")
		return
	}

	db.Delete(&book)
	c.Send("Book successfully deleted")
}
