package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `gorm:"title" json:"title"`
	Publication string `gorm:"publication" json:"publication"`
	Author      string `gorm:"author" json:"author"`
}

type BookDto struct {
	Title       string `json:"title"`
	Publication string `json:"publication"`
	Author      string `json:"author"`
}

func (b *Book) CreateBook() *Book {
	database.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var books []Book
	database.Find(&books)
	return books
}

func GetBookById(id int) (*Book, *gorm.DB) {
	var book Book
	db := database.Where("id = ?", id).Find(&book)
	return &book, db
}

func (b *Book) UpdateBook(book BookDto){
	updatedBook := &Book{
		Title:       book.Title,
		Publication: book.Publication,
		Author:      book.Author,
	}
	database.Model(&Book{}).Where("id = ?", b.ID).Updates(updatedBook)
}

func DeleteBook(id int) Book {
	var book Book
	database.Where("id = ?", id).Delete(&book)
	return book
}