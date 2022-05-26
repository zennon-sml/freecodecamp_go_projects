package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zennon-sml/go-book-crud-sql/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter() //creates my router
	routes.RegisterBookStoreRoutes(r) //define all the routes for my router
	http.Handle("/", r) //TODO what it does?
	log.Fatal(http.ListenAndServe(":8000", r)) //start the server
}
