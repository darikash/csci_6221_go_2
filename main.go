package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", homePage)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Starting web server on port 2111")
	http.ListenAndServe(":2111", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "/template.html")
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
