package handlers

import (
	"strconv"

	"github.com/Tirivashe/book_management/models"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	books := models.GetAllBooks()
	c.Set("Content-Type", "application/json")
	return c.JSON(books)
}

func GetBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	
	book, _ := models.GetBookById(idInt)
	if book.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Book not found"})
	}

	c.Set("Content-Type", "application/json")
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	var book models.BookDto
	if err := c.BodyParser(&book); err != nil {
		return err
	}
	bookModel := models.Book{
		Title:       book.Title,
		Publication: book.Publication,
		Author:      book.Author,
	}
	result := bookModel.CreateBook()
	return c.Status(201).JSON(fiber.Map{ "message": "Book created", "book": result })
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Update book by id, " + id)
}
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	models.DeleteBook(idInt)
	return c.Status(200).JSON(fiber.Map{"message": "Book deleted"})
}