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
	NewBooks := models.GetAllBooks() //TODO print later | get all the books from the db 
	res, _ := json.Marshal(NewBooks) //TODO print later | convert the books of the db to json
	w.Header().Set("Content-Type", "pkglication/json") //I GUESS sets the header to gets json has its standart type
    w.Write(res)//write it on the page
    // i guess that the views would enter here
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //TODO print later | use mux to take the response and put it in a dictionarie
	bookId := vars["bookId"] //gets the id of the book  
	ID, err := strconv.ParseInt(bookId, 0, 0) //convert the str(id) to int
	if err != nil {
		fmt.Println("error while it was parsing")
	}
	bookDetails, _ := models.GetBookById(ID) //TODO print later | gets the book by id
	res, _ := json.Marshal(bookDetails) //TODO print later | converts the book to json
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK) //gives the header a status 200
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
	book := models.DeleteBook(ID) //TODO print later | deletes the book and return him
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{} //empty book variable
	utils.ParseBody(r, UpdateBook)  //read the response(body) and convert the book to json
	vars := mux.Vars(r)             //TODO print these all later and see how the work
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0) //gets the int id of the book that will be updated
	if err != nil {
		fmt.Println("error while parsingInt in the update function")
	}
	bookDetails, db := models.GetBookById(ID) //TODO print later | gets the book that will be updated(bookDetails)
	//so UpdateBook is the new book tha will have the same id has the old book bookDetails

	if updateBook.Name != "" { //if the new book paramethers arent null
		bookDetails.Name = updateBook.Name //then set the atributes from bookdetails to the new ones from updateBook
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails) //saves the book in the db
	res, _ := json.Marshal(bookDetails) //transform into json
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)//and write him
}
