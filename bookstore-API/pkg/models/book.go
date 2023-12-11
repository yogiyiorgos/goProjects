// Define a Book model that establishes a database connection using
// the `config` package, and provides functions to create, retrieve
// and delete book records in the database using GORM.

package models

import (
	"github.com/jinzhu/gorm"                                     //GORM package for ORM functionality
	"github.com/yogiyiorgos/goProjects/bookstore-API/pkg/config" // Custom package for handling database connection
)

var db *gorm.DB // Global variable to hold the GORM db instance

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
