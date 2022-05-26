package models

import (
	"github.com/zennon-sml/go-book-crud-sql/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct { // defines the model/struct for my book
	gorm.Model // create some atributes as: id(if the primary key hasn't been set), created at and others
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect() // conect with the db specified in config
	db = config.GetDB() //gets the db variable that is the db
	db.AutoMigrate(&Book{}) // create the schema of the database
}

func (b *Book) CreateBook() *Book { // is aplied to a book "b" get that book andcreate, and return a book
	db.Create(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book { // gets all books in the db an return them in a slice
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) { // receives the id of the chosen book and query it
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db // return the book and the db as well
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)//TODO its wrong
	return book
}

