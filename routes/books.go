package routes

import (
	"github.com/Tirivashe/book_management/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	booksApi := v1.Group("/books")
	registerBooksRoutes(booksApi)
}

func registerBooksRoutes(api fiber.Router) {
	api.Get("/", handlers.GetBooks)
	api.Get("/:id", handlers.GetBookById)
	api.Post("/", handlers.CreateBook)
	api.Patch("/:id", handlers.UpdateBook)
	api.Delete("/:id", handlers.DeleteBook)
}