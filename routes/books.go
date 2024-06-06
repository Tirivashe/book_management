package routes

import (
	"github.com/Tirivashe/book_management/db"
	"github.com/Tirivashe/book_management/handlers"
	"github.com/Tirivashe/book_management/models"
	"github.com/gofiber/fiber/v2"
)

func bookRoutes(api fiber.Router) {
	database := db.GetDB()
	database.AutoMigrate(&models.Book{})
	api.Get("/", handlers.GetBooks)
	api.Get("/:id", handlers.GetBookById)
	api.Post("/", handlers.CreateBook)
	api.Patch("/:id", handlers.UpdateBook)
	api.Delete("/:id", handlers.DeleteBook)
}