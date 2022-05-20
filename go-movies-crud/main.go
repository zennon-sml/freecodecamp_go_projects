package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct { //movies table class
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` //refers to a director type
}

type Director struct { //director table class
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie //dumb database

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(movies)
	params := mux.Vars(r)
	fmt.Println(params)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")
	params := mux.Vars(r)
	fmt.Println(params)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	fmt.Println(params)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000)) // generates new id
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r) //gets the movie to be updated
	for index, item := range movies {
		if item.ID == params["id"] { //find the specifc movie
			movies = append(movies[:index], movies[index+1:]...) //deletes that movie from the slice
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) //decode from what we sent into the var movie
			movie.ID = params["id"]                    //define its id by the old one
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie) //return it to the page i guess (yes)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	//append movies to dumb database
	movies = append(movies, Movie{ID: "1", Isbn: "10", Title: "Fantastic Animals", Director: /*refers director to a type director*/ &Director{Firstname: "harry", Lastname: "Potter"}})
	movies = append(movies, Movie{ID: "2", Isbn: "11", Title: "Morbius", Director: &Director{Firstname: "Steven", Lastname: "Strange"}})
	movies = append(movies, Movie{ID: "3", Isbn: "12", Title: "Avatar", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
	movies = append(movies, Movie{ID: "4", Isbn: "13", Title: "Eragon", Director: &Director{Firstname: "Peter", Lastname: "Buchman"}})
	//all routes calling out theyrs respective functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("the server is in port: 8000")
	log.Fatal(http.ListenAndServe(":8000", r)) // start the server
}
