package main

import (
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func projectpage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/projects.html")
}

func aboutpage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func contactpage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	http.HandleFunc("/home", homepage)
	http.HandleFunc("/projects", projectpage)
	http.HandleFunc("/about", aboutpage)
	http.HandleFunc("/contact", contactpage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
