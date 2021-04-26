package main

import (
	application "fiber_api/application"
	database "fiber_api/infra"
	"fiber_api/models"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "database.db")

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection opened")
	database.DBConn.AutoMigrate(&models.Book{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()

	initDatabase()

	application.SetupRoutes(app)
	app.Listen(8000)

	defer database.DBConn.Close()
}
