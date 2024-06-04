package handlers

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Get all books")
}

func GetBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Get book by id, " + id)
}

func CreateBook(c *fiber.Ctx) error {
	return c.SendString("Create new book")
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Update book by id, " + id)
}
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Delete book by id, " + id)
}