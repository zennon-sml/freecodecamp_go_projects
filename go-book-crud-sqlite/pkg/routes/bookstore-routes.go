package routes

import (
	"github.com/gorilla/mux"
	"github.com/zennon-sml/go-book-crud-sql/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) { // receives a mux.router as parameter and set the routes for that router
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
}
