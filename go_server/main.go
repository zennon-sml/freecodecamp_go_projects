package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func MAE() string { //return a message according with the time(Morning, Afternoon, Evening)
	now := time.Now()
	if now.Hour() > 4 && now.Hour() < 12 {
		return "its " + now.Format("15:04") + " in this beautifull Morning\n"
	} else if now.Hour() > 12 && now.Hour() < 18 {
		return "its " + now.Format("15:04") + " in this beautifull Afternoon\n"
	} else {
		return "its " + now.Format("15:04") + " in this beautifull Evening\n"
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "paserForm() err: %v\n", err) //if the form was parsed wrong then show the error
		return
	}
	fmt.Fprintf(w, "POST request sucefullllll\n")
	name := r.FormValue("name") //gets name from the requested form
	address := r.FormValue("address")
	fmt.Fprintf(w, MAE())
	fmt.Fprintf(w, "WELCOME %s now we now that your address is %s\nbe ready...\n", name, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) { //request from user, response from me
	if r.URL.Path != "/hello" { //checking if the request is really the hello
		http.Error(w, "404 not found my friend", http.StatusNotFound)
		return
	}
	if r.Method != "GET" { //checking if the method is get if not return a error
		http.Error(w, "request method invalid try GET", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "vraaau!")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static")) //points to my static folder where my files are
	http.Handle("/", fileserver)                        //my index.html will be found on fileserver and it will be my route /
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("the server is starting in port 8080 if jesus help hahaha")
	if err := http.ListenAndServe(":8080", nil); err != nil { //start the server and if any error occur return a message
		log.Fatal(err)
	}

}
