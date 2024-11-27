package main

import (
	"fmt"
	groupie "groupie/func"
	"html/template"
	"log"
	"net/http"
)

func main() {
	//declare the handlers
	http.HandleFunc("/", homeHandle)
	http.HandleFunc("/artist", artistHandle)

	//serve the html and css files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server started at http://localhost:8080")

	//declare which port and start the server
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		groupie.ErrorDisplay(w, http.StatusNotFound, "Page not found")
		return
	}

	//get the data from API
	data, err := groupie.getData()
	if err != nil {
		groupie.ErrorDisplay(w, http.StatusInternalServerError, err.Error())
		return
	}

	//Parse html file
	templ, err := template.ParseFiles("static/index.html")
	if err != nil {
		groupie.ErrorDisplay(w, http.StatusInternalServerError, err.Error())
		return
	}

	//Execute if everything is fine
	//rather than nill it is gonna be data
	err = templ.Execute(w, data)
	if err != nil {
		groupie.ErrorDisplay(w, http.StatusInternalServerError, err.Error())
		return
	}

}

func artistHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		groupie.ErrorDisplay(w, http.StatusNotFound, "Page not found")
		return
	}

}
