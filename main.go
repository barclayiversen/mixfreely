package main

import (
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
	http.HandleFunc("/channel2", dog2)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func dog2(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	tpl.ExecuteTemplate(w, "channel2.html", nil)
}
