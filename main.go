package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/", dog)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}
