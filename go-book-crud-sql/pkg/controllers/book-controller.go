package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zennon-sml/go-book-crud-sql/pkg/models"
	"github.com/zennon-sml/go-book-crud-sql/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	NewBooks := models.GetAllBooks() //TODO print later
	res, _ := json.Marshal(NewBooks) //TODO print later
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //TODO print later
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while it was parsing")
	}
	bookDetails, _ := models.GetBookById(ID) //TODO print later
	res, _ := json.Marshal(bookDetails)      //TODO print later
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook) //TODO print later
	b := CreateBook.CreateBook()   //TODO print later
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //TODO print later
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing to int")
	}
	book := models.DeleteBook(ID) //TODO print later
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{} //empty book variable
	utils.ParseBody(r, UpdateBook)  //gets the request and parse it to a Booktype variable that is the new book
	vars := mux.Vars(r)             //TODO print these all later and see how the work
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0) //gets the id of the book that will be updated
	if err != nil {
		fmt.Println("error while parsingInt in the update function")
	}
	bookDetails, db := models.GetBookById(ID) //TODO print later, gets the book that will be updated
	//so UpdateBook is the new book tha will have the same id has the old book bookDetails

	if updateBook.Name != "" { //if the new book paramethers have been passed right
		bookDetails.Name = updateBook.Name //then set the atributes from bookdetails to the new ones from updateBook
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
