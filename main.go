package main

import (
	"fmt"
	groupie "groupie/func"
	"html/template"
	"log"
	"net/http"
	"strconv"
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
	data, err := groupie.GetData()
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
	err = templ.Execute(w, data)
	if err != nil {
		groupie.ErrorDisplay(w, http.StatusInternalServerError, err.Error())
		return
	}

}

func artistHandle(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("artist")

	if id == "" {
		groupie.ErrorDisplay(w, http.StatusBadRequest, "Missing artist ID")
		return
	}

	if r.URL.Path != "/artist" {
		groupie.ErrorDisplay(w, http.StatusNotFound, "Page not found")
		return
	}

	//convert the id to int
	aid, _ := strconv.Atoi(id)
	//get the index of the artist
	i := aid - 1

	//struct to hold both artist and relations data
	type PageData struct {
		Artist    groupie.Artist
		Relations groupie.Relations
	}

	artist, err := groupie.GetData()
	if err != nil {
		groupie.ErrorDisplay(w, http.StatusInternalServerError, err.Error())
		return
	}

	relations, err := groupie.GetRelations()
	if err != nil {
		groupie.ErrorDisplay(w, http.StatusInternalServerError, err.Error())
		return
	}

	//combine the data
	data := PageData{
		Artist:    artist[i],
		Relations: relations[i],
	}

	templ, err := template.ParseFiles("static/artist.html")
	if err != nil {
		groupie.ErrorDisplay(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = templ.Execute(w, data)
	if err != nil {
		groupie.ErrorDisplay(w, http.StatusInternalServerError, err.Error())
		return
	}
}
